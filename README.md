# RentHelp - 现代化租房服务平台

一个使用最前沿技术栈构建的租房服务平台，连接房东和租客，提供安全、便捷、透明的租房服务。

## 技术栈

### 前端
- **Next.js 14** - React 框架，使用 App Router
- **TypeScript** - 类型安全
- **Tailwind CSS** - 现代化样式框架
- **Shadcn/ui** - 高质量的 UI 组件库
- **Zustand** - 轻量级状态管理
- **React Hook Form** - 表单处理
- **Axios** - HTTP 客户端

### 后端
- **Golang 1.21** - 高性能后端语言
- **Gin** - 轻量级 Web 框架
- **MongoDB** - NoSQL 数据库
- **JWT** - 身份认证
- **CORS** - 跨域支持

### 开发工具
- **Docker** - 容器化部署
- **Docker Compose** - 多容器编排

## 项目结构

```
rent_help/
├── backend/                    # Go 后端
│   ├── cmd/server/            # 应用入口
│   ├── internal/              # 内部包
│   │   ├── config/           # 配置管理
│   │   ├── handlers/         # HTTP 处理器
│   │   ├── middleware/       # 中间件
│   │   ├── models/           # 数据模型
│   │   └── services/         # 业务逻辑
│   ├── pkg/database/         # 数据库连接
│   ├── Dockerfile
│   ├── go.mod
│   └── .env
├── frontend/                  # Next.js 前端
│   ├── app/                  # App Router 页面
│   ├── components/           # React 组件
│   ├── lib/                  # 工具函数
│   ├── store/                # 状态管理
│   ├── types/                # TypeScript 类型
│   ├── hooks/                # 自定义 Hooks
│   ├── Dockerfile
│   ├── package.json
│   └── tailwind.config.js
├── docker/                   # Docker 配置
│   └── docker-compose.yml
└── README.md
```

## 功能特性

### 核心功能
- **用户认证**: 注册、登录、JWT 令牌管理
- **房源管理**: 发布、编辑、删除房源信息
- **房源搜索**: 按地区、类型、价格等条件搜索
- **预订管理**: 租客预订房源，房东管理预订
- **用户资料**: 个人信息管理

### 用户角色
- **租客**: 搜索和预订房源
- **房东**: 发布和管理房源
- **管理员**: 平台管理

## 快速开始

### 使用 Docker（推荐）

1. **克隆项目**
```bash
git clone <repository-url>
cd rent_help
```

2. **启动服务**
```bash
cd docker
docker-compose up -d
```

3. **访问应用**
- 前端: http://localhost:3000
- 后端 API: http://localhost:8080
- MongoDB: localhost:27017

### 本地开发

#### 前端开发
```bash
cd frontend
npm install
npm run dev
```

#### 后端开发
```bash
cd backend
go mod tidy
go run cmd/server/main.go
```

#### MongoDB
确保 MongoDB 运行在 localhost:27017

## API 文档

### 认证相关
```
POST /api/v1/auth/register    # 用户注册
POST /api/v1/auth/login       # 用户登录
POST /api/v1/auth/refresh     # 刷新令牌
```

### 用户相关
```
GET  /api/v1/users/profile    # 获取用户资料
PUT  /api/v1/users/profile    # 更新用户资料
```

### 房源相关
```
GET    /api/v1/properties     # 获取房源列表
GET    /api/v1/properties/:id # 获取房源详情
POST   /api/v1/properties     # 创建房源
PUT    /api/v1/properties/:id # 更新房源
DELETE /api/v1/properties/:id # 删除房源
```

### 预订相关
```
GET    /api/v1/bookings       # 获取预订列表
GET    /api/v1/bookings/:id   # 获取预订详情
POST   /api/v1/bookings       # 创建预订
PUT    /api/v1/bookings/:id   # 更新预订
DELETE /api/v1/bookings/:id   # 取消预订
```

## 环境变量

### 后端 (.env)
```
PORT=8080
MONGO_URI=mongodb://localhost:27017
JWT_SECRET=your-secret-key
DB_NAME=rent_help
```

### 前端
```
NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
```

## 数据模型

### 用户 (User)
```go
type User struct {
    ID        primitive.ObjectID
    Email     string
    Password  string
    Name      string
    Phone     string
    Avatar    string
    Role      string  // "tenant", "landlord", "admin"
    Verified  bool
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

### 房源 (Property)
```go
type Property struct {
    ID          primitive.ObjectID
    Title       string
    Description string
    Type        string  // "apartment", "house", "room"
    Price       float64
    Currency    string
    Location    Location
    Features    []string
    Images      []string
    Available   bool
    OwnerID     primitive.ObjectID
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

### 预订 (Booking)
```go
type Booking struct {
    ID         primitive.ObjectID
    PropertyID primitive.ObjectID
    TenantID   primitive.ObjectID
    StartDate  time.Time
    EndDate    time.Time
    Status     string  // "pending", "confirmed", "cancelled", "completed"
    TotalPrice float64
    Message    string
    CreatedAt  time.Time
    UpdatedAt  time.Time
}
```

## 开发计划

- [ ] 用户认证系统
- [ ] 房源管理功能
- [ ] 搜索和筛选
- [ ] 预订管理
- [ ] 用户评价系统
- [ ] 支付集成
- [ ] 实时聊天
- [ ] 地图集成
- [ ] 移动端适配
- [ ] 多语言支持

## 贡献指南

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 联系方式

如有问题或建议，请提交 Issue 或联系开发团队。