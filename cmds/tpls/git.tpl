#!/bin/bash
#===============================================================================
# Git 最新版本安装脚本
# 支持: macOS (Homebrew) 和 Linux (基于发行版包管理器或源码编译)
# 版本: 1.0
#===============================================================================

set -e  # 遇到错误立即退出

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 打印函数
print_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

print_step() {
    echo -e "\n${GREEN}========================================${NC}"
    echo -e "${GREEN} $1${NC}"
    echo -e "${GREEN}========================================${NC}\n"
}

# 检测操作系统
detect_os() {
    if [[ "$OSTYPE" == "darwin"* ]]; then
        OS="macos"
        print_info "检测到操作系统: macOS"
    elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
        OS="linux"
        print_info "检测到操作系统: Linux"
    else
        print_error "不支持的操作系统: $OSTYPE"
        exit 1
    fi
}

# 检测包管理器
detect_package_manager() {
    if [[ "$OS" == "macos" ]]; then
        if command -v brew &> /dev/null; then
            PKG_MANAGER="homebrew"
            print_info "检测到包管理器: Homebrew"
        else
            print_error "未检测到 Homebrew，请先安装 Homebrew"
            print_info "安装命令: /bin/bash -c \"\$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)\""
            exit 1
        fi
    elif [[ "$OS" == "linux" ]]; then
        if command -v apt-get &> /dev/null; then
            PKG_MANAGER="apt"
            print_info "检测到包管理器: APT (Debian/Ubuntu)"
        elif command -v yum &> /dev/null; then
            PKG_MANAGER="yum"
            print_info "检测到包管理器: YUM (RHEL/CentOS)"
        elif command -v dnf &> /dev/null; then
            PKG_MANAGER="dnf"
            print_info "检测到包管理器: DNF (Fedora)"
        elif command -v zypper &> /dev/null; then
            PKG_MANAGER="zypper"
            print_info "检测到包管理器: Zypper (openSUSE)"
        elif command -v pacman &> /dev/null; then
            PKG_MANAGER="pacman"
            print_info "检测到包管理器: Pacman (Arch Linux)"
        elif command -v apk &> /dev/null; then
            PKG_MANAGER="apk"
            print_info "检测到包管理器: APK (Alpine Linux)"
        else
            print_error "未检测到支持的包管理器，将尝试从源码编译安装"
            PKG_MANAGER="source"
        fi
    fi
}

# 获取当前 Git 版本
get_current_git_version() {
    if command -v git &> /dev/null; then
        CURRENT_VERSION=$(git --version | awk '{print $3}')
        print_info "当前 Git 版本: $CURRENT_VERSION"
    else
        CURRENT_VERSION="未安装"
        print_info "当前 Git 版本: 未安装"
    fi
}

# 获取最新 Git 版本（从GitHub API）
get_latest_git_version() {
    print_info "正在获取最新 Git 版本信息..."

    # 使用 GitHub API 获取最新版本
    LATEST_VERSION=$(curl -s https://api.github.com/repos/git/git/tags | grep -o '"name": "v[0-9]*\.[0-9]*\.[0-9]*"' | head -1 | sed 's/"name": "v//' | sed 's/"//')

    if [[ -z "$LATEST_VERSION" ]]; then
        print_warning "无法从GitHub获取最新版本，尝试从Git官网获取..."
        LATEST_VERSION=$(curl -s https://www.kernel.org/pub/software/scm/git/ | grep -o 'git-[0-9]\+\.[0-9]\+\.[0-9]\+\.tar\.gz' | head -1 | sed 's/git-//' | sed 's/\.tar\.gz//')
    fi

    if [[ -z "$LATEST_VERSION" ]]; then
        print_error "无法获取最新版本信息，请检查网络连接"
        exit 1
    fi

    print_info "最新 Git 版本: $LATEST_VERSION"
}

# 检查是否需要更新
check_if_update_needed() {
    if [[ "$CURRENT_VERSION" == "未安装" ]]; then
        NEED_UPDATE="yes"
        print_info "Git 未安装，将进行安装"
        return
    fi

    # 比较版本号
    if [[ "$CURRENT_VERSION" == "$LATEST_VERSION" ]]; then
        print_success "Git 已是最新版本 ($CURRENT_VERSION)，无需更新"
        read -p "是否仍要重新安装？(y/N): " -n 1 -r
        echo
        if [[ ! $REPLY =~ ^[Yy]$ ]]; then
            exit 0
        fi
        NEED_UPDATE="yes"
    else
        print_info "发现新版本: $LATEST_VERSION (当前: $CURRENT_VERSION)"
        NEED_UPDATE="yes"
    fi
}

# 通过包管理器安装 Git
install_via_package_manager() {
    print_step "通过 $PKG_MANAGER 安装 Git $LATEST_VERSION"

    # 注意：包管理器可能没有最新版本，这里安装最新可用版本
    case $PKG_MANAGER in
        homebrew)
            brew update
            brew install git
            ;;
        apt)
            sudo apt-get update
            sudo apt-get install -y git
            ;;
        yum)
            sudo yum install -y git
            ;;
        dnf)
            sudo dnf install -y git
            ;;
        zypper)
            sudo zypper install -y git
            ;;
        pacman)
            sudo pacman -S --noconfirm git
            ;;
        apk)
            sudo apk add git
            ;;
        *)
            print_error "不支持的包管理器: $PKG_MANAGER"
            return 1
            ;;
    esac

    print_success "Git 安装完成"
}

# 从源码编译安装 Git
install_from_source() {
    print_step "从源码编译安装 Git $LATEST_VERSION"

    # 安装编译依赖
    print_info "安装编译依赖..."
    case $PKG_MANAGER in
        apt)
            sudo apt-get update
            sudo apt-get install -y make libssl-dev libghc-zlib-dev libcurl4-gnutls-dev libexpat1-dev gettext unzip
            ;;
        yum|dnf)
            sudo yum install -y curl-devel expat-devel gettext-devel openssl-devel zlib-devel gcc perl-ExtUtils-MakeMaker
            ;;
        zypper)
            sudo zypper install -y gcc make curl-devel expat-devel gettext-tools openssl-devel zlib-devel
            ;;
        pacman)
            sudo pacman -S --noconfirm base-devel curl expat gettext openssl zlib
            ;;
        apk)
            sudo apk add alpine-sdk curl-dev expat-dev openssl-dev zlib-dev gettext
            ;;
        *)
            print_info "尝试安装通用编译依赖..."
            sudo apt-get update 2>/dev/null || true
            sudo apt-get install -y make gcc libssl-dev libcurl4-gnutls-dev libexpat1-dev gettext unzip 2>/dev/null || \
            sudo yum install -y curl-devel expat-devel gettext-devel openssl-devel zlib-devel gcc perl-ExtUtils-MakeMaker 2>/dev/null || \
            print_warning "请手动安装编译依赖"
            ;;
    esac

    # 创建临时目录
    TMP_DIR=$(mktemp -d)
    cd "$TMP_DIR"

    # 下载源码
    print_info "下载 Git $LATEST_VERSION 源码..."
    SOURCE_URL="https://www.kernel.org/pub/software/scm/git/git-$LATEST_VERSION.tar.gz"
    curl -L -o "git-$LATEST_VERSION.tar.gz" "$SOURCE_URL" || {
        # 备用下载源
        SOURCE_URL="https://github.com/git/git/archive/v$LATEST_VERSION.tar.gz"
        curl -L -o "git-$LATEST_VERSION.tar.gz" "$SOURCE_URL"
    }

    # 解压源码
    print_info "解压源码..."
    tar -xzf "git-$LATEST_VERSION.tar.gz"
    cd "git-$LATEST_VERSION" || cd "git-$LATEST_VERSION" || {
        print_error "解压失败"
        exit 1
    }

    # 编译安装
    print_info "配置编译选项..."
    make configure
    ./configure --prefix=/usr/local

    print_info "开始编译 (这可能需要几分钟)..."
    make -j$(nproc 2>/dev/null || sysctl -n hw.ncpu 2>/dev/null || echo 2)

    print_info "安装 Git..."
    sudo make install

    # 清理临时目录
    cd /
    rm -rf "$TMP_DIR"

    print_success "Git $LATEST_VERSION 源码编译安装完成"
}

# 验证安装
verify_installation() {
    print_step "验证安装"

    if command -v git &> /dev/null; then
        INSTALLED_VERSION=$(git --version | awk '{print $3}')
        print_success "Git 安装成功！版本: $INSTALLED_VERSION"

        # 显示 Git 路径
        GIT_PATH=$(which git)
        print_info "Git 安装路径: $GIT_PATH"

        # 显示 Git 配置信息
        print_info "Git 配置信息:"
        git --help | head -5
    else
        print_error "Git 安装失败"
        exit 1
    fi
}

# 配置 Git（可选）
configure_git() {
    print_step "Git 基本配置"

    read -p "是否配置 Git 用户信息？(y/N): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        read -p "请输入 Git 用户名: " GIT_USER_NAME
        read -p "请输入 Git 用户邮箱: " GIT_USER_EMAIL

        git config --global user.name "$GIT_USER_NAME"
        git config --global user.email "$GIT_USER_EMAIL"

        print_success "Git 用户信息配置完成"

        # 显示当前配置
        print_info "当前 Git 配置:"
        git config --global --list | grep user
    fi

    # 设置默认编辑器
    read -p "是否设置默认编辑器？(y/N): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        read -p "请输入编辑器命令 (例如: vim, nano, code): " GIT_EDITOR
        git config --global core.editor "$GIT_EDITOR"
        print_success "默认编辑器设置为: $GIT_EDITOR"
    fi
}

# 主函数
main() {
    print_step "Git 最新版本安装工具"

    # 检测系统
    detect_os
    detect_package_manager

    # 获取版本信息
    get_current_git_version
    get_latest_git_version

    # 检查是否需要更新
    check_if_update_needed

    if [[ "$NEED_UPDATE" == "yes" ]]; then
        # 确认安装
        echo
        read -p "是否安装 Git $LATEST_VERSION？(Y/n): " -n 1 -r
        echo
        if [[ ! $REPLY =~ ^[Yy]$ ]] && [[ ! -z "$REPLY" ]]; then
            print_info "安装已取消"
            exit 0
        fi

        # 选择安装方式
        if [[ "$OS" == "macos" ]]; then
            # macOS 使用 Homebrew
            install_via_package_manager
        elif [[ "$PKG_MANAGER" != "source" ]]; then
            # Linux 有包管理器
            print_info "包管理器可能不包含最新版本，建议从源码编译安装以获得最新版本"
            read -p "选择安装方式: [1] 源码编译 (推荐) [2] 包管理器: " -n 1 -r
            echo
            case $REPLY in
                2)
                    install_via_package_manager
                    ;;
                *)
                    install_from_source
                    ;;
            esac
        else
            # 没有包管理器，源码编译
            install_from_source
        fi

        # 验证安装
        verify_installation

        # 配置 Git
        configure_git

        print_step "安装完成！"
        print_success "Git $LATEST_VERSION 已成功安装"
        print_info "运行 'git --version' 查看版本信息"
        print_info "运行 'git config --list' 查看配置信息"
    fi
}

# 执行主函数
main "$@"