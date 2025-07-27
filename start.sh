#!/bin/bash

echo "🚀 启动 RentHelp 租房服务平台..."

# 检查是否安装了 Docker
if ! command -v docker &> /dev/null; then
    echo "❌ Docker 未安装，请先安装 Docker"
    exit 1
fi

# 检查是否安装了 Docker Compose
if ! command -v docker-compose &> /dev/null; then
    echo "❌ Docker Compose 未安装，请先安装 Docker Compose"
    exit 1
fi

# 进入 docker 目录
cd docker

# 启动服务
echo "📦 正在启动容器..."
docker-compose up -d

# 等待服务启动
echo "⏳ 等待服务启动..."
sleep 10

# 检查服务状态
echo "📊 检查服务状态..."
docker-compose ps

echo ""
echo "✅ 服务启动完成！"
echo ""
echo "🌐 访问地址："
echo "   前端: http://localhost:3000"
echo "   后端 API: http://localhost:8080"
echo "   API 健康检查: http://localhost:8080/health"
echo ""
echo "🛑 停止服务: docker-compose down"
echo "🔄 重启服务: docker-compose restart"
echo "📜 查看日志: docker-compose logs -f"
