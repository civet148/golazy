#!/bin/bash
set -euo pipefail

# ===================== 可修改配置 =====================
# 默认protoc版本，调用脚本时可传参覆盖，如 ./install_protoc.sh 33.0
DEFAULT_PROTOC_VERSION="33.0"
# ======================================================

# 1. 接收版本参数
PROTOC_VERSION="${1:-$DEFAULT_PROTOC_VERSION}"
echo "=== 准备安装 protoc v${PROTOC_VERSION} ==="

# 2. 自动识别系统与架构
OS=""
ARCH=""
case "$(uname -s)" in
    Darwin)
        OS="osx"
        ;;
    Linux)
        OS="linux"
        ;;
    *)
        echo "不支持的操作系统：$(uname -s)"
        exit 1
        ;;
esac

case "$(uname -m)" in
    x86_64)
        ARCH="x86_64"
        ;;
    arm64|aarch64)
        ARCH="aarch_64"
        ;;
    *)
        echo "不支持的CPU架构：$(uname -m)"
        exit 1
        ;;
esac

PACKAGE_NAME="protoc-${PROTOC_VERSION}-${OS}-${ARCH}"
DOWNLOAD_URL="https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/${PACKAGE_NAME}.zip"
echo "检测系统：${OS}-${ARCH}，下载包：${DOWNLOAD_URL}"

# 3. 创建临时目录
TMP_DIR=$(mktemp -d)
cd "${TMP_DIR}"
echo "临时工作目录：${TMP_DIR}"

# 4. 按需安装依赖 curl unzip，不存在才安装
NEED_INSTALL=""
if ! command -v curl &> /dev/null; then
    NEED_INSTALL+="curl "
fi
if ! command -v unzip &> /dev/null; then
    NEED_INSTALL+="unzip "
fi

# 无缺失依赖直接跳过
if [ -z "${NEED_INSTALL}" ]; then
    echo "下载和解压工具已存在，跳过依赖安装"
else
    echo "检测到缺失工具：${NEED_INSTALL}，开始安装"
    if command -v apt &> /dev/null; then
        sudo apt update && sudo apt install -y ${NEED_INSTALL}
    elif command -v dnf &> /dev/null; then
        sudo dnf install -y ${NEED_INSTALL}
    elif command -v brew &> /dev/null; then
        brew install ${NEED_INSTALL}
    else
        echo "不支持的包管理器，请手动安装：${NEED_INSTALL}"
        exit 1
    fi
fi

# 5. 下载二进制包
echo "开始下载：${DOWNLOAD_URL}"
curl -fsSL "${DOWNLOAD_URL}" -o "${PACKAGE_NAME}.zip"

# 6. 解压
unzip -q "${PACKAGE_NAME}.zip"

# 7. 安装到 /usr/local (全局生效)
sudo cp bin/protoc /usr/local/bin/
sudo cp -r include/google /usr/local/include/
echo "protoc安装到/usr/local/bin/protoc"

# 8. 清理临时文件
cd - > /dev/null
rm -rf "${TMP_DIR}"

# 9. 验证安装
echo -e "\n=== 安装完成，验证版本 ==="
/usr/local/bin/protoc --version

