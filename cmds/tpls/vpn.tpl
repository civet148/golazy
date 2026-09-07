#!/bin/bash
# 支持 Linux 和 macOS 的 Tailscale 安装脚本

set -e

# 颜色输出
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
NC='\033[0m' # No Color

# 检测操作系统
detect_os() {
    case "$(uname -s)" in
        Linux*)     echo "Linux";;
        Darwin*)    echo "macOS";;
        *)          echo "Unknown";;
    esac
}

# 检测 Linux 发行版
detect_distro() {
    if [ -f /etc/os-release ]; then
        . /etc/os-release
        echo "$ID"
    else
        echo "unknown"
    fi
}

OS=$(detect_os)
echo -e "${GREEN}检测到操作系统: $OS${NC}"

# 检查是否已安装
if command -v tailscale &> /dev/null; then
    echo -e "${YELLOW}Tailscale 已安装，当前版本:${NC}"
    tailscale version 2>/dev/null || echo "无法获取版本信息"
    read -p "是否重新安装？(y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        exit 0
    fi
fi

# 安装 Tailscale
if [ "$OS" = "Linux" ]; then
    echo -e "${GREEN}使用官方脚本安装 Tailscale (Linux)...${NC}"
    curl -fsSL https://tailscale.com/install.sh | sh
    echo -e "${GREEN}安装完成！启动 tailscaled 服务...${NC}"
    sudo systemctl enable --now tailscaled
    echo -e "${GREEN}服务状态:${NC}"
    sudo systemctl status tailscaled --no-pager || true

elif [ "$OS" = "macOS" ]; then
    # 检查 Homebrew
    if ! command -v brew &> /dev/null; then
        echo -e "${RED}错误: 未检测到 Homebrew。请先安装 Homebrew:${NC}"
        echo "  /bin/bash -c \"\$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)\""
        exit 1
    fi
    echo -e "${GREEN}使用 Homebrew 安装 Tailscale (macOS)...${NC}"
    brew install --cask tailscale
    echo -e "${GREEN}安装完成！请从应用程序中启动 Tailscale 并登录。${NC}"
    echo "  或使用命令行: sudo tailscale up"
else
    echo -e "${RED}不支持的操作系统: $OS${NC}"
    exit 1
fi

echo -e "${GREEN}Tailscale 安装完成！${NC}"
echo "连接设备:"
echo "  - 交互式登录: sudo tailscale up"
echo "  - 使用 Auth Key: sudo tailscale up --authkey=tskey-auth-xxxxx"
