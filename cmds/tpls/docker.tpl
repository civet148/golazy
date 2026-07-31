#!/bin/bash
# ============================================================
# 脚本名称: install-docker.sh
# 功能描述: 在 Linux 和 macOS 上自动安装 Docker 服务
# 支持系统: Ubuntu/Debian/CentOS/RHEL/Fedora/macOS
# ============================================================

set -e

# 颜色输出
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

log_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# 检测操作系统
detect_os() {
    OS="$(uname -s)"
    case "${OS}" in
        Linux*)     OS_TYPE="linux";;
        Darwin*)    OS_TYPE="macos";;
        *)          log_error "不支持的操作系统: ${OS}"; exit 1;;
    esac
    log_info "检测到操作系统: ${OS_TYPE}"
}

# 检测 Linux 发行版
detect_linux_distro() {
    if [ -f /etc/os-release ]; then
        . /etc/os-release
        DISTRO="${ID}"
        VERSION_ID="${VERSION_ID}"
    elif [ -f /etc/redhat-release ]; then
        DISTRO="rhel"
    else
        DISTRO="unknown"
    fi
    log_info "Linux 发行版: ${DISTRO}"
}

# 检查 Docker 是否已安装
check_docker_installed() {
    if command -v docker &> /dev/null; then
        log_info "Docker 已安装，版本: $(docker --version)"
        read -p "是否重新安装? (y/N): " reinstall
        if [[ ! "$reinstall" =~ ^[Yy]$ ]]; then
            log_info "退出安装"
            exit 0
        fi
    fi
}

# ============================================================
# Linux 安装函数
# ============================================================
install_linux() {
    detect_linux_distro
    
    case "${DISTRO}" in
        ubuntu|debian)
            install_linux_debian
            ;;
        centos|rhel|fedora|rocky|almalinux)
            install_linux_rhel
            ;;
        *)
            log_warn "未识别的发行版，尝试使用官方通用脚本..."
            install_linux_generic
            ;;
    esac
}

# Debian/Ubuntu 安装
install_linux_debian() {
    log_info "开始安装 Docker (Debian/Ubuntu)..."
    
    # 更新包索引
    sudo apt-get update -y
    
    # 安装依赖
    sudo apt-get install -y \
        apt-transport-https \
        ca-certificates \
        curl \
        gnupg \
        lsb-release \
        software-properties-common
    
    # 添加 Docker 官方 GPG 密钥
    sudo mkdir -p /etc/apt/keyrings
    curl -fsSL https://download.docker.com/linux/${DISTRO}/gpg | \
        sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
    sudo chmod a+r /etc/apt/keyrings/docker.gpg
    
    # 添加 Docker APT 源
    echo \
        "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] \
        https://download.docker.com/linux/${DISTRO} \
        $(lsb_release -cs) stable" | \
        sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
    
    # 安装 Docker
    sudo apt-get update -y
    sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
    
    # 启动并设置开机自启
    sudo systemctl start docker
    sudo systemctl enable docker
    
    # 将当前用户加入 docker 组
    sudo usermod -aG docker "${USER}"
    log_warn "已将用户 ${USER} 加入 docker 组，可能需要重新登录才能生效"
}

# RHEL/CentOS/Fedora 安装
install_linux_rhel() {
    log_info "开始安装 Docker (RHEL/CentOS/Fedora)..."
    
    # 卸载旧版本
    sudo yum remove -y docker docker-client docker-client-latest \
        docker-common docker-latest docker-latest-logrotate \
        docker-logrotate docker-engine podman runc 2>/dev/null || true
    
    # 安装依赖
    sudo yum install -y yum-utils device-mapper-persistent-data lvm2
    
    # 添加 Docker 官方 YUM 源
    if [[ "${DISTRO}" == "fedora" ]]; then
        sudo dnf config-manager --add-repo \
            https://download.docker.com/linux/fedora/docker-ce.repo
    else
        sudo yum-config-manager --add-repo \
            https://download.docker.com/linux/centos/docker-ce.repo
    fi
    
    # 安装 Docker
    sudo yum install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
    
    # 启动并设置开机自启
    sudo systemctl start docker
    sudo systemctl enable docker
    
    # 将当前用户加入 docker 组
    sudo usermod -aG docker "${USER}"
    log_warn "已将用户 ${USER} 加入 docker 组，可能需要重新登录才能生效"
}

# 通用脚本安装（备用方案）
install_linux_generic() {
    log_info "使用官方通用脚本安装 Docker..."
    curl -fsSL https://get.docker.com | sh
    
    # 启动服务
    sudo systemctl start docker || sudo service docker start
    sudo systemctl enable docker 2>/dev/null || sudo chkconfig docker on
    
    sudo usermod -aG docker "${USER}"
    log_warn "已将用户 ${USER} 加入 docker 组，可能需要重新登录才能生效"
}

# ============================================================
# macOS 安装函数
# ============================================================
install_macos() {
    log_info "开始安装 Docker (macOS)..."
    
    # 检查是否已安装 Homebrew
    if ! command -v brew &> /dev/null; then
        log_warn "未检测到 Homebrew，正在安装..."
        /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
    fi
    
    # 使用 Homebrew 安装 Docker
    brew install --cask docker
    
    log_info "Docker 安装完成！"
    log_info "请手动打开 /Applications/Docker.app 启动 Docker Desktop"
    log_info "启动后，可以通过 'docker --version' 验证安装"
    
    # 提示用户启动
    read -p "是否现在打开 Docker Desktop? (y/N): " open_docker
    if [[ "$open_docker" =~ ^[Yy]$ ]]; then
        open /Applications/Docker.app
    fi
}

# ============================================================
# 验证安装
# ============================================================
verify_installation() {
    log_info "验证 Docker 安装..."
    
    if command -v docker &> /dev/null; then
        log_info "✅ Docker 版本: $(docker --version)"
        log_info "✅ Docker Compose 版本: $(docker compose version 2>/dev/null || echo '未安装')"
        
        # 尝试运行 hello-world
        if [[ "${OS_TYPE}" == "linux" ]]; then
            if groups "${USER}" | grep -q docker; then
                log_info "✅ 用户 ${USER} 已在 docker 组中"
            else
                log_warn "⚠️ 用户 ${USER} 不在 docker 组中，请重新登录或执行: newgrp docker"
            fi
        fi
    else
        log_error "❌ Docker 安装验证失败"
        exit 1
    fi
}

# ============================================================
# 主函数
# ============================================================
main() {
    log_info "===== Docker 自动安装脚本 ====="
    
    detect_os
    check_docker_installed
    
    case "${OS_TYPE}" in
        linux)
            install_linux
            ;;
        macos)
            install_macos
            ;;
    esac
    
    verify_installation
    
    log_info "===== Docker 安装完成 ====="
    log_info "Linux 用户注意：如果提示权限错误，请重新登录终端后重试"
    log_info "macOS 用户注意：请确保 Docker Desktop 已启动运行"
}

# 执行主函数
main "$@"