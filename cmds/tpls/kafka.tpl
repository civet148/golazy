#!/bin/bash

# Kafka 启动脚本（单节点）
echo "=========================================="
echo "启动 Kafka 容器（单节点版）"
echo "=========================================="

#创建数据目录
mkdir -p kafka_data

# 清理旧容器
docker rm -f kafka zookeeper 2>/dev/null

# 1. 先启动 Zookeeper
echo "启动 Zookeeper..."
docker run -d \
  --name zookeeper \
  -p 2181:2181 \
  -e ZOOKEEPER_CLIENT_PORT=2181 \
  -e ZOOKEEPER_TICK_TIME=2000 \
  -v zookeeper_data:/data \
  -v zookeeper_log:/datalog \
  --restart unless-stopped \
  confluentinc/cp-zookeeper:7.5.0

sleep 5

# 2. 启动 Kafka
echo "启动 Kafka..."
docker run -d \
  --name kafka \
  -p 9092:9092 \
  -p 29092:29092 \
  -e KAFKA_BROKER_ID=1 \
  -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 \
  -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092,PLAINTEXT_INTERNAL://kafka:29092 \
  -e KAFKA_LISTENER_SECURITY_PROTOCOL_MAP=PLAINTEXT:PLAINTEXT,PLAINTEXT_INTERNAL:PLAINTEXT \
  -e KAFKA_INTER_BROKER_LISTENER_NAME=PLAINTEXT_INTERNAL \
  -e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 \
  -e KAFKA_TRANSACTION_STATE_LOG_MIN_ISR=1 \
  -e KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR=1 \
  -e KAFKA_NUM_PARTITIONS=3 \
  -e KAFKA_DELETE_TOPIC_ENABLE=true \
  -e KAFKA_AUTO_CREATE_TOPICS_ENABLE=true \
  -e KAFKA_LOG_RETENTION_HOURS=168 \
  -e KAFKA_LOG_RETENTION_BYTES=1073741824 \
  -v ./kafka_data:/var/lib/kafka/data \
  --link zookeeper \
  --restart unless-stopped \
  confluentinc/cp-kafka:7.5.0

echo "✅ Kafka 已启动！"
echo "🔗 Kafka 地址: localhost:9092"
echo "🔗 Zookeeper 地址: localhost:2181"
echo ""
echo "📋 常用命令:"
echo "  查看主题: docker exec kafka kafka-topics --list --bootstrap-server localhost:9092"
echo "  创建主题: docker exec kafka kafka-topics --create --topic test --partitions 3 --replication-factor 1 --bootstrap-server localhost:9092"
echo "  生产消息: docker exec -it kafka kafka-console-producer --topic test --bootstrap-server localhost:9092"
echo "  消费消息: docker exec -it kafka kafka-console-consumer --topic test --from-beginning --bootstrap-server localhost:9092"
echo "  查看日志: docker logs kafka"
echo "  停止服务: docker stop kafka zookeeper"
echo "=========================================="
