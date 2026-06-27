# 创建数据目录
mkdir -p mysql_data mysql_logs

# 删除已运行容器
docker rm -f mysql 2>/dev/null

# 启动容器(设置root初始密码为123456)
docker run -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 --restart always \
        -e TZ=Asia/Shanghai \
        -v mysql-files:/var/lib/mysql-files \
		-v ./mysql_conf:/etc/mysql \
        -v ./mysql_data:/var/lib/mysql \
        -v ./mysql_logs:/var/log/mysql \
        --name mysql -d  mysql:8.0
