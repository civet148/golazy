#!/bin/bash
# install-k3s-server.sh
# 使用 Docker 部署 K3s Server (Master 节点)
# 支持 macOS 和 Linux，自动安装 Docker

set -e

# ==================== 配置区域 ====================
MASTER_IP="${MASTER_IP:-192.168.0.76}"
K3S_TOKEN="${K3S_TOKEN:-12345678}"
INGRESS_EXTERNAL_PORT=443 # ingress外部端口
INGRESS_INTERNAL_PORT=443 # ingress内部端口
API_EXTERNAL_PORT=9443    # api server外部端口
API_INTERNAL_PORT=9443    # api server内部端口
WEB_EXTERNAL_PORT=80      # web服务外部端口
WEB_INTERNAL_PORT=80      # web服务内部端口
# =================================================

# 颜色输出
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

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

# 清理旧容器
cleanup_old() {
    if docker ps -a --filter "name=k3s-server" --format "{{.Names}}" | grep -q "k3s-server"; then
        log_info "停止并移除旧的 k3s-server 容器..."
        docker stop k3s-server 2>/dev/null || true
        docker rm k3s-server 2>/dev/null || true
    fi
}

# 启动 K3s Server
start_k3s_server() {
    log_info "启动 K3s Server 容器..."

	# 判断操作系统，决定 IP 绑定方式
	if [[ "$(uname -s)" == "Darwin" ]]; then
		# macOS: 使用 0.0.0.0 让容器监听所有接口
		BIND_IP="0.0.0.0"
		NODE_IP="0.0.0.0"  # 容器内节点 IP
		ADVERTISE_IP="$MASTER_IP"  # 对外公布的 IP（用宿主机 IP）
	else
		# Linux: 直接使用配置的 IP
		BIND_IP="$MASTER_IP"
		NODE_IP="$MASTER_IP"
		ADVERTISE_IP="$MASTER_IP"
	fi

	docker run -d \
		--name k3s-server \
		--restart unless-stopped \
		--privileged \
		-p ${API_EXTERNAL_PORT}:${API_INTERNAL_PORT} \
		-p ${WEB_EXTERNAL_PORT}:${WEB_INTERNAL_PORT} \
		-p ${INGRESS_EXTERNAL_PORT}:${INGRESS_INTERNAL_PORT} \
		-e K3S_TOKEN="$K3S_TOKEN" \
		-e K3S_KUBECONFIG_MODE="644" \
		rancher/k3s:latest server \
		--bind-address "$BIND_IP" \
		--advertise-address "$ADVERTISE_IP" \
		--https-listen-port ${API_INTERNAL_PORT} \
		--tls-san "$MASTER_IP"
}

# 配置 kubectl
configure_kubectl() {
    log_info "配置 kubectl..."
    
    # 等待 K3s 就绪
    sleep 10
    
    # 从容器复制 kubeconfig
    mkdir -p ~/.kube
    docker cp k3s-server:/etc/rancher/k3s/k3s.yaml ~/.kube/config 2>/dev/null || {
        log_error "无法复制 kubeconfig，请检查容器是否正常运行"
        docker logs k3s-server --tail 20
        exit 1
    }
    
    # 修改 server 地址为外部 IP
    if [[ "$(uname -s)" == "Darwin" ]]; then
        sed -i '' "s/127.0.0.1/$MASTER_IP/g" ~/.kube/config
    else
        sed -i "s/127.0.0.1/$MASTER_IP/g" ~/.kube/config
    fi
    
    log_info "kubectl 配置完成"
}

# ==================== 主程序 ====================
main() {
    log_info "开始安装 K3s Server (Master 节点: $MASTER_IP)"
    
    # 步骤 1: 安装 Docker
    install_docker
    
    # 步骤 2: 清理旧容器
    cleanup_old
    
    # 步骤 3: 启动 K3s Server
    start_k3s_server
    
    # 步骤 4: 配置 kubectl
    configure_kubectl
    
    # 步骤 5: 验证
    log_info "等待集群就绪..."
    sleep 10
    
    if command -v kubectl &>/dev/null; then
        kubectl get nodes
    else
        log_warn "kubectl 未安装，请运行: brew install kubectl (macOS) 或 apt install kubectl (Linux)"
    fi
    
    echo ""
    echo "=========================================="
    echo -e "${GREEN}✅ K3s Server 安装完成！${NC}"
    echo "=========================================="
    echo "Master IP: $MASTER_IP"
    echo "Token:     $K3S_TOKEN"
    echo ""
    echo "📋 在 Worker 节点使用此 Token 加入集群"
    echo "运行命令: ./join-k3s-worker.sh"
    echo "=========================================="
}

# 执行主程序
main
