@echo off
echo ====================================
echo     租房助手平台 - 状态检查
echo ====================================
echo.

echo 检查 Docker 容器状态...
docker-compose -f docker-compose.dev.yml ps
echo.

echo 检查端口占用情况...
netstat -an | findstr ":3001 :8081 :27018 :6380"
echo.

echo 检查容器日志（最后10行）...
echo --- MongoDB 日志 ---
docker logs --tail 10 renthelp_mongodb_dev 2>nul
echo.
echo --- Backend 日志 ---
docker logs --tail 10 renthelp_backend_dev 2>nul
echo.
echo --- Frontend 日志 ---
docker logs --tail 10 renthelp_frontend_dev 2>nul
echo.

echo ====================================
echo 如果所有服务正常运行，访问：
echo 前端应用: http://localhost:3001
echo 后端 API: http://localhost:8081
echo ====================================
