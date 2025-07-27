@echo off
echo 🚀 启动 RentHelp 租房服务平台...

REM 检查是否安装了 Docker
docker --version >nul 2>&1
if %errorlevel% neq 0 (
    echo ❌ Docker 未安装，请先安装 Docker
    pause
    exit /b 1
)

REM 检查是否安装了 Docker Compose
docker-compose --version >nul 2>&1
if %errorlevel% neq 0 (
    echo ❌ Docker Compose 未安装，请先安装 Docker Compose
    pause
    exit /b 1
)

REM 进入 docker 目录
cd docker

REM 启动服务
echo 📦 正在启动容器...
docker-compose up -d

REM 等待服务启动
echo ⏳ 等待服务启动...
timeout /t 10 /nobreak >nul

REM 检查服务状态
echo 📊 检查服务状态...
docker-compose ps

echo.
echo ✅ 服务启动完成！
echo.
echo 🌐 访问地址：
echo    前端: http://localhost:3000
echo    后端 API: http://localhost:8080
echo    API 健康检查: http://localhost:8080/health
echo.
echo 🛑 停止服务: docker-compose down
echo 🔄 重启服务: docker-compose restart
echo 📜 查看日志: docker-compose logs -f
echo.
pause
