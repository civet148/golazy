# 创建数据目录
mkdir -p postgres_data

# 删除已启动容器
docker rm -f postgres  2>/dev/null

# 启动容器
docker run -d \
 --name postgres \
 -e POSTGRES_PASSWORD=123456 \
 -v ./postgres_data:/var/lib/postgresql \
 -p 5432:5432 \
 swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/postgres:18.0-alpine
 