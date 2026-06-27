#!/bin/bash

# RabbitMQ 启动脚本
echo "=========================================="
echo "启动 RabbitMQ 容器"
echo "=========================================="

USER=admin
PASSWORD='12345678'

# 创建数据目录
mkdir -p rabbitmq_data

# 停止并删除已存在的容器
docker rm -f rabbitmq  2>/dev/null

# 启动 RabbitMQ 容器
docker run -d \
  --name rabbitmq \
  --hostname rabbitmq-host \
  -p 5672:5672 \
  -p 15672:15672 \
  -p 15692:15692 \
  -p 25672:25672 \
  -e RABBITMQ_DEFAULT_USER=${USER}\
  -e RABBITMQ_DEFAULT_PASS=${PASSWORD}\
  -e RABBITMQ_DEFAULT_VHOST=/ \
  -v ./rabbitmq_data:/var/lib/rabbitmq \
  -v ./rabbitmq.conf:/etc/rabbitmq/rabbitmq.conf:ro \
  --restart unless-stopped \
  rabbitmq:3.13-management-alpine

echo "✅ RabbitMQ 已启动！"
echo "🔗 AMQP 连接地址: amqp://${USER}:${PASSWORD}@localhost:5672/"
echo "🌐 管理界面: http://localhost:15672"
echo "📊 监控界面: http://localhost:15692"
echo ""
echo "📋 常用命令:"
echo "  查看状态: docker logs rabbitmq"
echo "  进入容器: docker exec -it rabbitmq bash"
echo "  停止服务: docker stop rabbitmq"
echo "  删除容器: docker rm rabbitmq"
echo "=========================================="
