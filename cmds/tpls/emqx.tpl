#!/bin/bash

# ============================================================
# EMQX 5.8 一键部署脚本
# 适用环境: Mac M 系列 (arm64) / Linux amd64
# ============================================================

set -e

# -------- 可配置变量 --------
EMQX_VERSION="5.8.8"  # EMQX 5.8 版本号
EMQX_PASSWD="public"
CONTAINER_NAME="emqx" # 容器名称
DATA_DIR="$PWD/data"  # 数据持久化目录（宿主机）
LOG_DIR="$PWD/log"    # 日志持久化目录（宿主机）
# ---------------------------

# 颜色输出
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

log_info()  { echo -e "${GREEN}[INFO]${NC} $1"; }
log_warn()  { echo -e "${YELLOW}[WARN]${NC} $1"; }
log_error() { echo -e "${RED}[ERROR]${NC} $1"; }

# 1. 检查 Docker 是否可用
if ! command -v docker &> /dev/null; then
    log_error "未检测到 Docker，请先安装 Docker Desktop 或 Docker Engine。"
    exit 1
fi

if ! docker info &> /dev/null; then
    log_error "Docker 守护进程未运行，请启动 Docker 后重试。"
    exit 1
fi

# 2. 检测当前系统架构
ARCH=$(uname -m)
log_info "检测到系统架构: $ARCH"

case "$ARCH" in
    x86_64)
        log_info "Linux amd64 架构，将拉取 linux/amd64 镜像"
        ;;
    arm64|aarch64)
        log_info "ARM64 架构 (Mac M 系列 / Linux arm64)，将拉取 linux/arm64/v8 镜像"
        ;;
    *)
        log_warn "未知架构: $ARCH，将尝试默认拉取"
        ;;
esac

# 3. 如果容器已存在，先移除
if docker ps -a --format '{{.Names}}' | grep -q "^${CONTAINER_NAME}$"; then
    log_warn "容器 ${CONTAINER_NAME} 已存在，正在移除..."
    docker rm -f "${CONTAINER_NAME}" > /dev/null 2>&1
fi

# 4. 创建持久化目录
log_info "创建持久化目录..."
mkdir -p "${DATA_DIR}" "${LOG_DIR}"

# 5. 拉取镜像（Docker 会自动匹配当前架构）
log_info "拉取 EMQX ${EMQX_VERSION} 镜像..."
docker pull "emqx/emqx:${EMQX_VERSION}"

# 6. 启动容器
log_info "启动 EMQX 容器..."
docker run -d \
    --name "${CONTAINER_NAME}" \
    --restart unless-stopped \
    -p 1883:1883 \
    -p 8083:8083 \
    -p 8084:8084 \
    -p 8883:8883 \
    -p 18083:18083 \
	-e EMQX_DASHBOARD__DEFAULT_PASSWORD=${EMQX_PASSWD} \
    -v "${DATA_DIR}:/opt/emqx/data" \
    -v "${LOG_DIR}:/opt/emqx/log" \
    "emqx/emqx:${EMQX_VERSION}"

# 7. 等待容器就绪
log_info "等待 EMQX 启动..."
sleep 5

if docker ps --format '{{.Names}}' | grep -q "^${CONTAINER_NAME}$"; then
    log_info "EMQX 启动成功！"
    echo ""
    echo "========================================="
    echo "  EMQX Dashboard: http://localhost:18083"
    echo "  默认账号: admin"
    echo "  默认密码: ${EMQX_PASSWD}"
    echo "========================================="
    echo ""
    echo "  MQTT TCP 端口:    1883"
    echo "  WebSocket 端口:   8083"
    echo "  WebSocket SSL:    8084"
    echo "  MQTT SSL 端口:    8883"
    echo ""
    echo "  数据目录: ${DATA_DIR}"
    echo "  日志目录: ${LOG_DIR}"
    echo "========================================="
else
    log_error "容器启动失败，请执行 'docker logs ${CONTAINER_NAME}' 查看错误日志。"
    exit 1
fi
