# 删除已启动容器
docker rm -f influxdb  2>/dev/null

#启动容器 开启8083、8086端口
docker run -p 8083:8083 -p 8086:8086 --name influxdb -td tutum/influxdb:latest
