# RentHelp 安装指南

## 系统要求

- **Docker** 20.10 或更高版本
- **Docker Compose** 1.29 或更高版本
- **Git** (用于克隆项目)

## 快速安装

### Windows 用户

1. **双击运行启动脚本**
   ```
   start.bat
   ```

### Linux/Mac 用户

1. **运行启动脚本**
   ```bash
   chmod +x start.sh
   ./start.sh
   ```

### 手动安装

1. **启动服务**
   ```bash
   cd docker
   docker-compose up -d
   ```

2. **访问应用**
   - 前端: http://localhost:3000
   - 后端: http://localhost:8080

## 开发环境设置

### 前端开发

1. **安装依赖**
   ```bash
   cd frontend
   npm install
   ```

2. **启动开发服务器**
   ```bash
   npm run dev
   ```

### 后端开发

1. **安装依赖**
   ```bash
   cd backend
   go mod tidy
   ```

2. **启动开发服务器**
   ```bash
   go run cmd/server/main.go
   ```

## 常见问题

### 端口被占用
如果端口 3000 或 8080 被占用，请修改 `docker-compose.yml` 中的端口映射。

### MongoDB 连接失败
确保 MongoDB 容器正常运行：
```bash
docker-compose logs mongodb
```

### 前端无法连接后端
检查 `NEXT_PUBLIC_API_URL` 环境变量是否正确设置。

## 服务管理

### 停止服务
```bash
docker-compose down
```

### 重启服务
```bash
docker-compose restart
```

### 查看日志
```bash
docker-compose logs -f
```

### 清理数据
```bash
docker-compose down -v
```

## 生产部署

1. **修改环境变量**
   - 更改 JWT 密钥
   - 配置生产数据库
   - 设置域名和 SSL

2. **构建生产镜像**
   ```bash
   docker-compose -f docker-compose.prod.yml up -d
   ```

## 获取帮助

如遇到问题，请查看：
- [GitHub Issues](https://github.com/your-repo/rent-help/issues)
- [项目文档](README.md)
- [API 文档](http://localhost:8080/swagger)

## 下一步

安装完成后，您可以：
1. 注册新用户账号
2. 发布房源信息
3. 搜索和预订房源
4. 管理个人资料
