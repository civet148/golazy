#!/bin/bash
# join-k3s-worker.sh
# 使用 Docker 部署 K3s Agent (Worker 节点)
# 支持 macOS 和 Linux，自动安装 Docker

set -e

# ==================== 配置区域 ====================
MASTER_IP="${MASTER_IP:-192.168.0.76}"
MASTER_URL="https://${MASTER_IP}:9443"
# =================================================

# 颜色输出
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

log_info() { echo -e "${GREEN}[INFO]${NC} $1"; }
log_warn() { echo -e "${YELLOW}[WARN]${NC} $1"; }
log_error() { echo -e "${RED}[ERROR]${NC} $1"; }

# 检测操作系统
detect_os() {
    OS="$(uname -s)"
    case "$OS" in
        Darwin)  echo "macOS" ;;
        Linux)   echo "Linux" ;;
        *)       echo "Unsupported"; exit 1 ;;
    esac
}

# 安装 Docker (Linux)
install_docker_linux() {
    log_info "检测到 Linux，开始安装 Docker..."
    if command -v docker &>/dev/null; then
        log_info "Docker 已安装，跳过安装步骤"
        return
    fi
    
    # 卸载旧版本
    sudo apt-get remove -y docker docker-engine docker.io containerd runc 2>/dev/null || true
    
    # 安装依赖
    sudo apt-get update
    sudo apt-get install -y ca-certificates curl gnupg lsb-release
    
    # 添加 Docker 官方 GPG 密钥
    sudo mkdir -m 0755 -p /etc/apt/keyrings
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
    
    # 添加 Docker 仓库
    echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
    
    # 安装 Docker
    sudo apt-get update
    sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
    
    # 将当前用户加入 docker 组
    sudo usermod -aG docker $USER
    
    # 启动 Docker
    sudo systemctl enable docker
    sudo systemctl start docker
    
    log_info "Docker 安装完成！"
    log_warn "请重新登录以使 docker 组权限生效，或运行: newgrp docker"
}

# 安装 Docker (macOS)
install_docker_macos() {
    log_info "检测到 macOS..."
    if command -v docker &>/dev/null && docker info &>/dev/null 2>&1; then
        log_info "Docker 已安装并正在运行"
        return
    fi
    
    if command -v docker &>/dev/null; then
        log_warn "Docker 已安装但未运行，请手动启动 Docker Desktop"
        exit 1
    fi
    
    log_error "macOS 请手动安装 Docker Desktop"
    log_error "下载地址: https://www.docker.com/products/docker-desktop"
    exit 1
}

# 安装 Docker
install_docker() {
    OS_TYPE=$(detect_os)
    log_info "操作系统: $OS_TYPE"
    
    case "$OS_TYPE" in
        macOS)  install_docker_macos ;;
        Linux)  install_docker_linux ;;
        *)      exit 1 ;;
    esac
    
    # 等待 Docker 就绪
    log_info "等待 Docker 就绪..."
    for i in {1..30}; do
        if docker info &>/dev/null 2>&1; then
            log_info "Docker 已就绪"
            return 0
        fi
        sleep 1
    done
    log_error "Docker 未能正常启动"
    exit 1
}

# 获取 Token
get_token() {
    if [ -z "$K3S_TOKEN" ]; then
        echo ""
        read -sp "请输入 Master 节点的 K3s Token: " K3S_TOKEN
        echo ""
    fi
    
    if [ -z "$K3S_TOKEN" ]; then
        log_error "未提供 Token，无法加入集群"
        exit 1
    fi
}

# 清理旧容器
cleanup_old() {
    if docker ps -a --filter "name=k3s-agent" --format "{{.Names}}" | grep -q "k3s-agent"; then
        log_info "停止并移除旧的 k3s-agent 容器..."
        docker stop k3s-agent 2>/dev/null || true
        docker rm k3s-agent 2>/dev/null || true
    fi
}

# 启动 K3s Agent
start_k3s_agent() {
    # 获取本机 IP
    WORKER_IP=""
    if command -v ip &>/dev/null; then
        WORKER_IP=$(ip route get 1 2>/dev/null | awk '{print $NF;exit}' 2>/dev/null)
    fi
    if [ -z "$WORKER_IP" ] && command -v ifconfig &>/dev/null; then
        if [[ "$(uname -s)" == "Darwin" ]]; then
            WORKER_IP=$(ifconfig | grep "inet " | grep -v 127.0.0.1 | head -1 | awk '{print $2}')
        else
            WORKER_IP=$(ifconfig | grep "inet " | grep -v 127.0.0.1 | head -1 | awk '{print $2}' | cut -d: -f2)
        fi
    fi
    if [ -z "$WORKER_IP" ]; then
        WORKER_IP=$(hostname -I 2>/dev/null | awk '{print $1}')
    fi
    if [ -z "$WORKER_IP" ]; then
        log_warn "无法自动获取本机 IP，将使用默认配置"
        WORKER_IP="0.0.0.0"
    fi
    log_info "本机 Worker IP: $WORKER_IP"
    
    log_info "启动 K3s Agent 容器加入集群..."
    log_info "Master URL: $MASTER_URL"
    
    docker run -d \
        --name k3s-agent \
        --restart unless-stopped \
        --privileged \
        -e K3S_URL="$MASTER_URL" \
        -e K3S_TOKEN="$K3S_TOKEN" \
        -e K3S_NODE_NAME="$(hostname)-worker" \
        rancher/k3s:latest agent \
        --node-ip "$WORKER_IP"
    
    if [ $? -eq 0 ]; then
        log_info "K3s Agent 容器已启动"
    else
        log_error "K3s Agent 启动失败"
        exit 1
    fi
}

# ==================== 主程序 ====================
main() {
    log_info "开始加入 K3s 集群 (Master: $MASTER_IP)"
    
    # 步骤 1: 获取 Token
    get_token
    
    # 步骤 2: 安装 Docker
    install_docker
    
    # 步骤 3: 清理旧容器
    cleanup_old
    
    # 步骤 4: 启动 K3s Agent
    start_k3s_agent
    
    echo ""
    echo "=========================================="
    echo -e "${GREEN}✅ Worker 节点正在加入集群...${NC}"
    echo "=========================================="
    echo "Master URL: $MASTER_URL"
    echo ""
    echo "📋 请在 Master 节点执行以下命令验证:"
    echo "   kubectl get nodes"
    echo ""
    echo "⏱️  等待约 10-20 秒后节点状态应为 Ready"
    echo "=========================================="
}

# 执行主程序
main
