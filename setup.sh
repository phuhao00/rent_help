#!/bin/bash

echo "📦 初始化 RentHelp 项目..."

# 检查是否在项目根目录
if [ ! -f "README.md" ]; then
    echo "❌ 请在项目根目录运行此脚本"
    exit 1
fi

echo "🔧 安装后端依赖..."
cd backend
if command -v go &> /dev/null; then
    go mod tidy
    echo "✅ Go 依赖安装完成"
else
    echo "⚠️  Go 未安装，跳过后端依赖安装"
fi
cd ..

echo "🎨 安装前端依赖..."
cd frontend
if command -v npm &> /dev/null; then
    npm install
    echo "✅ npm 依赖安装完成"
elif command -v yarn &> /dev/null; then
    yarn install
    echo "✅ yarn 依赖安装完成"
else
    echo "⚠️  npm/yarn 未安装，跳过前端依赖安装"
fi
cd ..

echo "🐳 检查 Docker 环境..."
if command -v docker &> /dev/null; then
    echo "✅ Docker 已安装"
    if command -v docker-compose &> /dev/null; then
        echo "✅ Docker Compose 已安装"
    else
        echo "⚠️  Docker Compose 未安装"
    fi
else
    echo "⚠️  Docker 未安装"
fi

echo ""
echo "🎉 项目初始化完成！"
echo ""
echo "📋 接下来你可以："
echo "   1. 运行 './start.sh' 启动完整服务"
echo "   2. 或者分别启动："
echo "      - 后端: cd backend && go run cmd/server/main.go"
echo "      - 前端: cd frontend && npm run dev"
echo "      - 数据库: docker run -d -p 27017:27017 mongo"
echo ""
echo "📚 更多信息请查看 README.md"
