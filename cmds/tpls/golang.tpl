#!/bin/bash
#===============================================================================
# Go 1.26 安装脚本
# 功能:
#   - 支持 Linux 和 macOS (AMD64/ARM64)
#   - 下载并解压到 /usr/local/go
#   - 自动设置 GOPATH=~/code
#   - 自动设置 GOROOT=/usr/local/go
#   - 自动配置 PATH 环境变量
#   - 支持 bash 和 zsh
#===============================================================================

set -e

# 颜色定义
if [ -t 1 ] && command -v tput &>/dev/null && [ "$(tput colors)" -ge 8 ]; then
    RED='\033[0;31m'
    GREEN='\033[0;32m'
    YELLOW='\033[1;33m'
    BLUE='\033[0;34m'
    CYAN='\033[0;36m'
    MAGENTA='\033[0;35m'
    NC='\033[0m'
else
    RED=''; GREEN=''; YELLOW=''; BLUE=''; CYAN=''; MAGENTA=''; NC=''
fi

# 日志函数
log_info() { echo -e "${BLUE}➜${NC} $1"; }
log_success() { echo -e "${GREEN}✓${NC} $1"; }
log_warn() { echo -e "${YELLOW}⚠${NC} $1"; }
log_error() { echo -e "${RED}✗${NC} $1"; exit 1; }
log_step() { echo -e "\n${CYAN}▶${NC} $1"; }
log_title() {
    echo -e "\n${MAGENTA}════════════════════════════════════════════════════════════${NC}"
    echo -e "${MAGENTA} $1${NC}"
    echo -e "${MAGENTA}════════════════════════════════════════════════════════════${NC}"
}

# 配置
GO_VERSION="1.26.0"
GO_DOWNLOAD_URL="https://go.dev/dl"
INSTALL_DIR="/usr/local/go"
GOPATH_DIR="$HOME/code"
BACKUP_DIR="/usr/local/go.bak.$(date +%Y%m%d_%H%M%S)"

# 检测 Shell 类型
detect_shell() {
    local shell_name=$(basename "$SHELL")

    case "$shell_name" in
        bash)
            SHELL_CONFIG="$HOME/.bashrc"
            SHELL_TYPE="bash"
            ;;
        zsh)
            SHELL_CONFIG="$HOME/.zshrc"
            SHELL_TYPE="zsh"
            ;;
        *)
            # 默认使用 bash
            SHELL_CONFIG="$HOME/.bashrc"
            SHELL_TYPE="bash"
            log_warn "未识别的 Shell: $shell_name，默认使用 bash"
            ;;
    esac

    log_info "Shell 类型: $SHELL_TYPE"
    log_info "配置文件: $SHELL_CONFIG"
}

# 检测系统
detect_system() {
    log_step "检测系统环境..."

    if [[ "$OSTYPE" == "darwin"* ]]; then
        OS="macos"
        OS_NAME="macOS"
        ARCH=$(uname -m)

        # macOS 架构转换
        if [[ "$ARCH" == "x86_64" ]]; then
            GO_ARCH="amd64"
            GO_FILENAME="go${GO_VERSION}.darwin-amd64.tar.gz"
        elif [[ "$ARCH" == "arm64" ]]; then
            GO_ARCH="arm64"
            GO_FILENAME="go${GO_VERSION}.darwin-arm64.tar.gz"
        else
            log_error "不支持的架构: $ARCH"
        fi

        SYSTEM="${OS}-${GO_ARCH}"
        log_info "系统: macOS $(sw_vers -productVersion 2>/dev/null || echo '未知')"
        log_info "架构: $ARCH (Go: $GO_ARCH)"

    elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
        OS="linux"
        OS_NAME="Linux"
        ARCH=$(uname -m)

        # Linux 架构转换
        if [[ "$ARCH" == "x86_64" ]]; then
            GO_ARCH="amd64"
            GO_FILENAME="go${GO_VERSION}.linux-amd64.tar.gz"
        elif [[ "$ARCH" == "aarch64" ]] || [[ "$ARCH" == "arm64" ]]; then
            GO_ARCH="arm64"
            GO_FILENAME="go${GO_VERSION}.linux-arm64.tar.gz"
        elif [[ "$ARCH" == "armv7l" ]]; then
            GO_ARCH="armv6l"
            GO_FILENAME="go${GO_VERSION}.linux-armv6l.tar.gz"
        else
            log_error "不支持的架构: $ARCH"
        fi

        SYSTEM="${OS}-${GO_ARCH}"

        # 检测发行版
        if [ -f /etc/os-release ]; then
            . /etc/os-release
            log_info "系统: $PRETTY_NAME"
        elif [ -f /etc/redhat-release ]; then
            log_info "系统: $(cat /etc/redhat-release)"
        else
            log_info "系统: Linux"
        fi
        log_info "架构: $ARCH (Go: $GO_ARCH)"

    else
        log_error "不支持的操作系统: $OSTYPE"
    fi

    DOWNLOAD_URL="${GO_DOWNLOAD_URL}/${GO_FILENAME}"
    log_info "下载地址: $DOWNLOAD_URL"
}

# 检查权限
check_permissions() {
    log_step "检查权限..."

    if [ "$EUID" -eq 0 ]; then
        SUDO=""
        log_info "以 root 用户运行"
    else
        if command -v sudo &>/dev/null; then
            SUDO="sudo"
            log_info "使用 sudo 获取权限"
        else
            log_error "需要 root 权限安装到 /usr/local"
        fi
    fi

    # 检查目录权限
    if [ ! -w "/usr/local" ] && [ -z "$SUDO" ]; then
        log_error "没有权限写入 /usr/local，请使用 sudo 运行"
    fi
}

# 检查并备份现有 Go 安装
check_existing_go() {
    log_step "检查现有 Go 安装..."

    if [ -d "$INSTALL_DIR" ]; then
        log_warn "发现已存在的 Go 安装: $INSTALL_DIR"

        if command -v go &>/dev/null; then
            OLD_VERSION=$(go version 2>/dev/null | awk '{print $3}' || echo "未知")
            log_info "旧版本: $OLD_VERSION"
        fi

        read -p "是否备份并替换？ [y/N] " -n 1 -r
        echo
        if [[ $REPLY =~ ^[Yy]$ ]]; then
            log_info "备份到: $BACKUP_DIR"
            $SUDO mv "$INSTALL_DIR" "$BACKUP_DIR"
            log_success "备份完成"
        else
            log_info "取消安装"
            exit 0
        fi
    else
        log_info "未发现现有 Go 安装"
    fi
}

# 创建 GOPATH 目录
create_gopath_dir() {
    log_step "创建 GOPATH 目录..."

    if [ ! -d "$GOPATH_DIR" ]; then
        log_info "创建目录: $GOPATH_DIR"
        mkdir -p "$GOPATH_DIR"
        mkdir -p "$GOPATH_DIR/src"
        mkdir -p "$GOPATH_DIR/bin"
        mkdir -p "$GOPATH_DIR/pkg"
        log_success "GOPATH 目录创建完成"
    else
        log_info "GOPATH 目录已存在: $GOPATH_DIR"

        # 确保子目录存在
        mkdir -p "$GOPATH_DIR/src"
        mkdir -p "$GOPATH_DIR/bin"
        mkdir -p "$GOPATH_DIR/pkg"
    fi
}

# 下载并安装 Go
install_go() {
    log_step "下载并安装 Go ${GO_VERSION}..."

    # 创建临时目录
    TMP_DIR=$(mktemp -d 2>/dev/null || mktemp -d -t go_install)
    cd "$TMP_DIR"

    # 下载
    log_info "下载 ${GO_FILENAME}..."
    if command -v curl &>/dev/null; then
        curl -fsSL -o "${GO_FILENAME}" "$DOWNLOAD_URL" || {
            log_error "下载失败，请检查网络连接"
        }
    elif command -v wget &>/dev/null; then
        wget -q -O "${GO_FILENAME}" "$DOWNLOAD_URL" || {
            log_error "下载失败，请检查网络连接"
        }
    else
        log_error "未找到 curl 或 wget，请安装其中之一"
    fi

    # 验证下载
    if [ ! -f "${GO_FILENAME}" ]; then
        log_error "下载文件不存在"
    fi

    FILE_SIZE=$(du -h "${GO_FILENAME}" | cut -f1)
    log_info "下载完成，文件大小: $FILE_SIZE"

    # 解压
    log_info "解压到 $INSTALL_DIR..."
    $SUDO tar -C /usr/local -xzf "${GO_FILENAME}"

    # 清理临时文件
    cd /
    rm -rf "$TMP_DIR"

    log_success "Go 安装完成"
}

# 配置环境变量
configure_environment() {
    log_step "配置环境变量..."

    # 检查是否已经配置
    if grep -q "GOROOT=/usr/local/go" "$SHELL_CONFIG" 2>/dev/null; then
        log_warn "环境变量已在 $SHELL_CONFIG 中配置"
        read -p "是否重新配置？ [y/N] " -n 1 -r
        echo
        if [[ ! $REPLY =~ ^[Yy]$ ]]; then
            log_info "跳过环境变量配置"
            return
        fi
        # 删除旧的配置
        $SUDO sed -i.bak '/# Go environment variables/d' "$SHELL_CONFIG"
        $SUDO sed -i.bak '/export GOROOT=\/usr\/local\/go/d' "$SHELL_CONFIG"
        $SUDO sed -i.bak '/export GOPATH=.*code/d' "$SHELL_CONFIG"
        $SUDO sed -i.bak '/export PATH=.*GOROOT/d' "$SHELL_CONFIG"
        $SUDO sed -i.bak '/export PATH=.*GOPATH/d' "$SHELL_CONFIG"
    fi

    # 创建环境变量配置
    cat << EOF | $SUDO tee -a "$SHELL_CONFIG" > /dev/null

# Go environment variables
export GOROOT=/usr/local/go
export GOPATH=$GOPATH_DIR
export PATH=\$PATH:\$GOROOT/bin
export PATH=\$PATH:\$GOPATH/bin
export GOPROXY=https://goproxy.io

# Go proxy (optional, speeds up downloads)
export GO111MODULE=on
export CGO_ENABLED=1
export GOPROXY=https://goproxy.cn,direct

EOF

    log_success "环境变量已添加到 $SHELL_CONFIG"

    # 同时添加到 .profile (兼容性)
    if [ -f "$HOME/.profile" ] && [ "$SHELL_CONFIG" != "$HOME/.profile" ]; then
        if ! grep -q "GOROOT=/usr/local/go" "$HOME/.profile" 2>/dev/null; then
            cat << EOF | $SUDO tee -a "$HOME/.profile" > /dev/null

# Go environment variables (from install script)
export GOROOT=/usr/local/go
export GOPATH=$GOPATH_DIR
export PATH=\$PATH:\$GOROOT/bin
export PATH=\$PATH:\$GOPATH/bin
export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct

EOF
            log_info "也添加到了 ~/.profile"
        fi
    fi
}

# 验证安装
verify_installation() {
    log_step "验证安装..."

    # 临时设置环境变量
    export GOROOT="$INSTALL_DIR"
    export GOPATH="$GOPATH_DIR"
    export PATH="$PATH:$GOROOT/bin"
    export PATH="$PATH:$GOPATH/bin"

    if command -v go &>/dev/null; then
        GO_VERSION_INSTALLED=$(go version)
        log_success "Go 安装成功: $GO_VERSION_INSTALLED"

        log_info "GOROOT: $GOROOT"
        log_info "GOPATH: $GOPATH"
        log_info "GOBIN: $GOPATH/bin"

        # 显示 Go 环境信息
        echo
        log_info "Go 环境信息:"
        go env | grep -E "GOROOT|GOPATH|GO111MODULE|GOPROXY"

        return 0
    else
        log_error "Go 安装验证失败"
    fi
}

# 创建测试项目
create_test_project() {
    log_step "创建测试项目..."

    local test_dir="$GOPATH_DIR/src/hello"
    mkdir -p "$test_dir"

    cat << 'EOF' > "$test_dir/main.go"
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go 1.26!")
    fmt.Println("GOPATH:", "~/code")
    fmt.Println("GOROOT:", "/usr/local/go")
}
EOF

    log_info "测试项目已创建: $test_dir/main.go"

    # 编译测试
    cd "$test_dir"
    go mod init hello 2>/dev/null || true
    go build -o hello 2>/dev/null || log_warn "编译测试失败，请手动测试"

    if [ -f "./hello" ]; then
        log_success "编译成功！运行 ./hello 查看效果"
        ./hello || true
    fi
}

# 显示安装信息
show_installation_info() {
    log_title "Go 1.26 安装完成！"

    echo
    log_info "安装位置:"
    echo "  GOROOT: $INSTALL_DIR"
    echo "  GOPATH: $GOPATH_DIR"
    echo
    log_info "环境配置:"
    echo "  配置文件: $SHELL_CONFIG"
    echo "  已添加 Go 环境变量"
    echo
    log_info "下一步:"
    echo "  1. 重新加载配置: source $SHELL_CONFIG"
    echo "  2. 验证安装: go version"
    echo "  3. 查看环境: go env"
    echo "  4. 测试项目: cd $GOPATH_DIR/src/hello && go run main.go"
    echo
    log_info "Go 代理已设置为: https://goproxy.cn,direct"
    echo "  如需修改，请编辑 $SHELL_CONFIG 中的 GOPROXY"
    echo
    log_success "安装完成！请执行以下命令使配置生效:"
    echo -e "\n  ${GREEN}source $SHELL_CONFIG${NC}\n"
}

# 主函数
main() {
    log_title "Go ${GO_VERSION} 安装脚本"

    # 检测系统
    detect_system

    # 检测 Shell
    detect_shell

    # 检查权限
    check_permissions

    # 检查现有安装
    check_existing_go

    # 创建 GOPATH
    create_gopath_dir

    # 安装 Go
    install_go

    # 配置环境变量
    configure_environment

    # 验证安装
    verify_installation

    # 创建测试项目
    create_test_project

    # 显示信息
    show_installation_info
}

# 执行主函数
main "$@"