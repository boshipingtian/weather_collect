#!/bin/bash

# 部署脚本
set -e

# 定义变量
IMAGE_NAME="weather-colly"
TAG="latest"
CONTAINER_NAME="weather-colly-app"

echo "开始部署应用..."

# 调用构建脚本
echo "执行构建..."
./build.sh

# 停止并删除现有容器
echo "停止现有容器..."
docker stop ${CONTAINER_NAME} 2>/dev/null || echo "容器 ${CONTAINER_NAME} 未运行"
docker rm ${CONTAINER_NAME} 2>/dev/null || echo "容器 ${CONTAINER_NAME} 不存在"

# 启动应用容器
echo "启动应用容器..."
docker run -d \
  --name ${CONTAINER_NAME} \
  -e TZ=Asia/Shanghai \
  -v $(pwd)/settings.yaml:/build/settings.yaml \
  --restart unless-stopped \
  ${IMAGE_NAME}:${TAG}

# 显示运行状态
echo "服务状态:"
docker ps | grep -E "(${CONTAINER_NAME}|${MYSQL_CONTAINER_NAME})"

echo "部署完成!"
echo "查看应用日志: docker logs -f ${CONTAINER_NAME}"
echo "停止服务: docker stop ${CONTAINER_NAME}"