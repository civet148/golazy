# 创建数据和配置目录
mkdir -p minio01/{data,config}

# 删除已启动容器
docker rm -f minio01 2>/dev/null

# 运行容器
#-p  19000 为API访问端口
#-p  19090 为MinIO控制台的地址
#-e "MINIO_ROOT_USER=minio" 设置登录名
#-e "MINIO_ROOT_PASSWORD=12345678"   设置登录密码 (一定要大于8位，否则容器将启动失败)
#-v ./minio01/data:/data  挂载数据目录到宿主机
#-v ./minio01/config:/root/.minio  挂载配置文件
docker run -p 19000:9000 -p 19090:9090 \
 --name minio01 \
 -d --restart=always \
 -e "MINIO_ROOT_USER=minio" \
 -e "MINIO_ROOT_PASSWORD=12345678" \
 -v ./minio01/data:/data \
 -v ./minio01/config:/root/.minio  \
 minio/minio server \
 /data --console-address ":9090"
