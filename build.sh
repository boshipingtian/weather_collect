#!/bin/bash

# 构建脚本
set -e

# 定义变量
IMAGE_NAME="weather-colly"
TAG="latest"

echo "开始构建Docker镜像..."

# 构建Docker镜像
docker build -t ${IMAGE_NAME}:${TAG} -f build/Dockerfile .

echo "Docker镜像构建完成: ${IMAGE_NAME}:${TAG}"

# 显示镜像信息
docker images | grep ${IMAGE_NAME}

# 清除无用镜像
echo "清除无用镜像..."
docker image prune -f

echo "清除悬空镜像..."
docker rmi $(docker images -f "dangling=true" -q) 2>/dev/null || echo "没有悬空镜像需要清除"

echo "构建完成并已清理无用镜像"