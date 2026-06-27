#!/bin/bash

# Redis 单机版启动脚本
echo "=========================================="
echo "启动 Redis 容器（单机版）"
echo "=========================================="

# 设置密码
PASSWD=123456

# 创建数据目录
mkdir -p redis_data

# 停止并删除已存在的容器
docker rm -f redis-server 2>/dev/null

# 启动 Redis 容器
docker run -d \
  --name redis-server \
  --restart unless-stopped \
  -p 6379:6379 \
  -v ./redis_data:/data \
  redis:7-alpine \
  redis-server \
  --appendonly yes \
  --requirepass "${PASSWD}" \
  --maxmemory 512mb \
  --maxmemory-policy allkeys-lru

echo "✅ Redis 已启动！"
echo "🔗 连接地址: redis://localhost:6379"
echo "🔑 密码: ${PASSWD}"
echo "💾 数据持久化: 已启用 (AOF)"
echo "🗑️  内存策略: 最大512MB，LRU淘汰"
echo ""
echo "📋 常用命令:"
echo "  连接Redis: docker exec -it redis-server redis-cli -a ${PASSWD}"
echo "  查看日志: docker logs redis-server"
echo "  停止服务: docker stop redis-server"
echo "  删除数据: docker volume rm redis_data"
echo "=========================================="
