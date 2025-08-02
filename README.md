# RentHelp - 现代化租房服务平台

[![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)](https://github.com/phuhao00/rent_help)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/go-1.21-blue.svg)](https://golang.org/)
[![Node Version](https://img.shields.io/badge/node-18+-green.svg)](https://nodejs.org/)

一个使用最前沿技术栈构建的全栈租房服务平台，连接房东和租客，提供安全、便捷、透明的租房服务。

## ✨ 功能特性

### 🏠 核心功能
- **用户认证**: 注册、登录、JWT 令牌管理、角色权限控制
- **房源管理**: 发布、编辑、删除房源信息，图片上传
- **智能搜索**: 按地区、类型、价格、特性等条件筛选房源
- **预订管理**: 租客预订房源，房东管理预订状态
- **用户资料**: 完整的个人信息管理系统
- **评价系统**: 房源和用户双向评价机制

### 👥 用户角色
- **🏠 租客**: 搜索和预订房源，管理预订记录
- **🏡 房东**: 发布和管理房源，处理预订请求
- **⚙️ 管理员**: 平台管理，用户和内容审核

## 🛠 技术栈

### 前端技术
- **React 18.2.0** + **TypeScript 4.9.5** - 现代化前端框架
- **Next.js** - 服务端渲染和路由
- **Material-UI v5** - 高质量 UI 组件库
- **Tailwind CSS** - 原子化 CSS 框架
- **Zustand** - 轻量级状态管理
- **React Hook Form** - 表单处理
- **Axios** - HTTP 客户端

### 后端技术
- **Go 1.21** + **Gin** - 高性能 Web 框架
- **MongoDB 7.0** - NoSQL 文档数据库
- **Redis 7** - 缓存和会话管理
- **JWT** - 身份认证和授权
- **Air** - 开发热重载

### 开发运维
- **Docker** + **Docker Compose** - 容器化部署
- **GitHub Actions** - CI/CD 流水线
- **Nginx** - 反向代理和负载均衡

## 📁 项目架构

```
rent_help/
├── 📂 backend/                    # Go 后端服务
│   ├── 📂 cmd/server/            # 应用程序入口点
│   ├── 📂 internal/              # 内部业务逻辑
│   │   ├── 📂 config/           # 配置管理
│   │   ├── 📂 handlers/         # HTTP 请求处理器
│   │   ├── 📂 middleware/       # 中间件
│   │   ├── 📂 models/           # 数据模型定义
│   │   └── 📂 services/         # 业务逻辑层
│   ├── 📂 pkg/                  # 共享包
│   │   ├── 📂 database/         # 数据库连接
│   │   ├── 📂 errors/           # 错误处理
│   │   ├── 📂 utils/            # 工具函数
│   │   └── 📂 validation/       # 输入验证
│   ├── 📂 tests/                # 测试文件
│   ├── 📂 uploads/              # 文件上传存储
│   ├── 🐳 Dockerfile            # 后端容器配置
│   └── 📄 go.mod                # Go 模块依赖
├── 📂 frontend/                  # React 前端应用
│   ├── 📂 app/                  # Next.js App Router 页面
│   │   ├── 📂 (auth)/           # 认证相关页面
│   │   ├── 📂 (dashboard)/      # 仪表板页面
│   │   ├── 📄 layout.tsx        # 根布局
│   │   └── 📄 page.tsx          # 首页
│   ├── 📂 components/           # React 组件
│   │   ├── 📂 ui/               # UI 基础组件
│   │   └── 📄 Navigation.tsx    # 导航组件
│   ├── 📂 hooks/                # 自定义 React Hooks
│   ├── 📂 lib/                  # 工具函数和配置
│   ├── 📂 store/                # 状态管理
│   ├── 📂 types/                # TypeScript 类型定义
│   ├── 🐳 Dockerfile            # 前端容器配置
│   └── 📄 package.json          # Node.js 依赖
├── 📂 docker/                   # Docker 配置文件
│   ├── 📂 mongodb/              # MongoDB 初始化脚本
│   └── 🐳 docker-compose.yml    # 容器编排配置
├── 📂 .github/                  # GitHub 工作流和说明
│   ├── 📂 workflows/            # CI/CD 流水线
│   └── 📂 instructions/         # 项目说明文档
├── 🐳 docker-compose.dev.yml    # 开发环境配置
├── 🐳 docker-compose.prod.yml   # 生产环境配置
├── 📄 API.md                    # API 接口文档
├── 📄 DEPLOYMENT.md             # 部署配置文档
├── 📄 INSTALL.md                # 安装指南
└── 📄 README.md                 # 项目说明文档
```

## 🚀 快速开始

### 方式一：一键启动（推荐）

#### Windows 用户
```cmd
# 双击运行启动脚本
start.bat
```

#### Linux/Mac 用户
```bash
# 赋予执行权限并运行
chmod +x start.sh
./start.sh
```

### 方式二：Docker Compose

1. **克隆项目**
```bash
git clone https://github.com/phuhao00/rent_help.git
cd rent_help
```

2. **启动开发环境**
```bash
docker-compose -f docker-compose.dev.yml up -d
```

3. **访问应用**
- 🌐 前端应用: http://localhost:3001
- 🔗 后端 API: http://localhost:8081
- 📊 MongoDB: localhost:27018
- 🗄️ Redis: localhost:6380

### 方式三：本地开发

#### 前端开发
```bash
cd frontend
npm install
npm run dev
# 访问: http://localhost:3001
```

#### 后端开发
```bash
cd backend
go mod tidy
go run cmd/server/main.go
# API 运行在: http://localhost:8081
```

### 🔧 系统要求

- **Docker Desktop** 27.5.1+ 
- **Docker Compose** 2.0+
- **Git** 2.30+
- **Go** 1.21+ (本地开发)
- **Node.js** 18+ (本地开发)

## 🎯 示例账户

系统预置了以下测试账户供快速体验：

| 角色 | 邮箱 | 密码 | 说明 |
|------|------|------|------|
| 🔧 管理员 | admin@renthelp.com | admin123 | 系统管理权限 |
| 🏡 房东 | john.landlord@email.com | landlord123 | 房源管理 |
| 🏡 房东 | mike.landlord@email.com | landlord123 | 房源管理 |
| 🏠 租客 | jane.tenant@email.com | tenant123 | 房源搜索预订 |

## 📖 API 文档

### 🔐 认证接口

| 方法 | 端点 | 描述 |
|------|------|------|
| POST | `/api/v1/auth/register` | 用户注册 |
| POST | `/api/v1/auth/login` | 用户登录 |
| POST | `/api/v1/auth/refresh` | 刷新令牌 |

### 👤 用户接口

| 方法 | 端点 | 描述 | 权限 |
|------|------|------|------|
| GET | `/api/v1/users/profile` | 获取用户资料 | 🔒 需要认证 |
| PUT | `/api/v1/users/profile` | 更新用户资料 | 🔒 需要认证 |

### 🏠 房源接口

| 方法 | 端点 | 描述 | 权限 |
|------|------|------|------|
| GET | `/api/v1/properties` | 获取房源列表 | 🔒 需要认证 |
| GET | `/api/v1/properties/:id` | 获取房源详情 | 🔒 需要认证 |
| POST | `/api/v1/properties` | 创建房源 | 🏡 仅房东 |
| PUT | `/api/v1/properties/:id` | 更新房源 | 🏡 仅所有者 |
| DELETE | `/api/v1/properties/:id` | 删除房源 | 🏡 仅所有者 |

### 📋 预订接口

| 方法 | 端点 | 描述 | 权限 |
|------|------|------|------|
| GET | `/api/v1/bookings` | 获取预订列表 | 🔒 需要认证 |
| GET | `/api/v1/bookings/:id` | 获取预订详情 | 🔒 需要认证 |
| POST | `/api/v1/bookings` | 创建预订 | 🏠 仅租客 |
| PUT | `/api/v1/bookings/:id` | 更新预订 | 🔒 相关用户 |
| DELETE | `/api/v1/bookings/:id` | 取消预订 | 🔒 相关用户 |

### 📊 健康检查

| 方法 | 端点 | 描述 |
|------|------|------|
| GET | `/api/v1/health` | 服务状态检查 |

### 🔧 请求示例

#### 用户注册
```bash
curl -X POST http://localhost:8081/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123",
    "full_name": "用户姓名",
    "phone": "13800138000",
    "role": "tenant"
  }'
```

#### 用户登录
```bash
curl -X POST http://localhost:8081/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

#### 获取房源列表
```bash
curl -X GET "http://localhost:8081/api/v1/properties?limit=10&city=北京" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

#### 创建房源（房东）
```bash
curl -X POST http://localhost:8081/api/v1/properties \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "市中心精装公寓",
    "description": "位于市中心，交通便利",
    "type": "apartment",
    "price": 3500,
    "bedrooms": 2,
    "bathrooms": 1,
    "area": 850,
    "address": {
      "street": "建国门外大街1号",
      "city": "北京",
      "state": "北京市",
      "zip_code": "100020",
      "country": "中国"
    },
    "amenities": ["WiFi", "空调", "洗衣机", "停车位"],
    "images": ["https://example.com/image1.jpg"]
  }'
```

### 📋 状态码说明

| 状态码 | 含义 |
|--------|------|
| 200 | ✅ 请求成功 |
| 201 | ✅ 创建成功 |
| 400 | ❌ 请求参数错误 |
| 401 | ❌ 未认证或认证失败 |
| 403 | ❌ 权限不足 |
| 404 | ❌ 资源不存在 |
| 409 | ❌ 资源冲突 |
| 500 | ❌ 服务器内部错误 |

## ⚙️ 环境配置

### 后端环境变量 (.env)

```bash
# 🌐 服务器配置
PORT=8080
GIN_MODE=debug

# 🗄️ 数据库配置
MONGODB_URI=mongodb://admin:password123@mongodb:27017/renthelp?authSource=admin
DB_NAME=rent_help

# 🔐 JWT 配置
JWT_SECRET=your-super-secret-jwt-key-here
JWT_EXPIRES_IN=24h

# 📁 文件上传配置
UPLOAD_PATH=/app/uploads
MAX_FILE_SIZE=10MB

# 📧 邮件配置
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your-email@gmail.com
SMTP_PASSWORD=your-app-password

# 🗄️ Redis 配置
REDIS_URL=redis://redis:6379

# 💳 支付配置
STRIPE_SECRET_KEY=sk_test_...
STRIPE_WEBHOOK_SECRET=whsec_...

# ☁️ 云存储配置
CLOUDINARY_URL=cloudinary://...

# 📊 监控配置
SENTRY_DSN=https://...
```

### 前端环境变量 (.env.local)

```bash
# 🔗 API 配置
NEXT_PUBLIC_API_URL=http://localhost:8081
NEXT_PUBLIC_WEBSOCKET_URL=ws://localhost:8081

# 🔐 认证配置
NEXTAUTH_URL=http://localhost:3001
NEXTAUTH_SECRET=your-nextauth-secret

# 🗺️ 地图服务
NEXT_PUBLIC_GOOGLE_MAPS_API_KEY=AIza...

# 💳 支付服务
NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY=pk_test_...

# 📊 分析服务
NEXT_PUBLIC_GA_ID=G-...
NEXT_PUBLIC_HOTJAR_ID=...

# 🚀 功能开关
NEXT_PUBLIC_ENABLE_CHAT=true
NEXT_PUBLIC_ENABLE_PAYMENTS=true
```

## 🗄️ 数据模型

### 👤 用户模型 (User)
```go
type User struct {
    ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Email     string            `json:"email" bson:"email"`
    Password  string            `json:"-" bson:"password"`
    FullName  string            `json:"fullName" bson:"full_name"`
    Phone     string            `json:"phone" bson:"phone"`
    Avatar    string            `json:"avatar" bson:"avatar"`
    Role      string            `json:"role" bson:"role"` // "tenant", "landlord", "admin"
    IsVerified bool             `json:"isVerified" bson:"is_verified"`
    IsActive   bool             `json:"isActive" bson:"is_active"`
    CreatedAt  time.Time        `json:"createdAt" bson:"created_at"`
    UpdatedAt  time.Time        `json:"updatedAt" bson:"updated_at"`
}
```

### 🏠 房源模型 (Property)
```go
type Property struct {
    ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Title       string            `json:"title" bson:"title"`
    Description string            `json:"description" bson:"description"`
    Type        string            `json:"type" bson:"type"` // "apartment", "house", "condo", "townhouse", "studio"
    Price       float64           `json:"price" bson:"price"`
    Bedrooms    int               `json:"bedrooms" bson:"bedrooms"`
    Bathrooms   int               `json:"bathrooms" bson:"bathrooms"`
    Area        float64           `json:"area" bson:"area"`
    Address     Address           `json:"address" bson:"address"`
    Location    Location          `json:"location" bson:"location"`
    Amenities   []string          `json:"amenities" bson:"amenities"`
    Images      []string          `json:"images" bson:"images"`
    OwnerID     primitive.ObjectID `json:"ownerId" bson:"owner_id"`
    Available   bool              `json:"available" bson:"available"`
    AvailableFrom time.Time       `json:"availableFrom" bson:"available_from"`
    LeaseTerms    []string        `json:"leaseTerms" bson:"lease_terms"`
    PetsAllowed   bool            `json:"petsAllowed" bson:"pets_allowed"`
    SmokingAllowed bool           `json:"smokingAllowed" bson:"smoking_allowed"`
    UtilitiesIncluded []string    `json:"utilitiesIncluded" bson:"utilities_included"`
    CreatedAt   time.Time         `json:"createdAt" bson:"created_at"`
    UpdatedAt   time.Time         `json:"updatedAt" bson:"updated_at"`
}

type Address struct {
    Street  string `json:"street" bson:"street"`
    City    string `json:"city" bson:"city"`
    State   string `json:"state" bson:"state"`
    ZipCode string `json:"zipCode" bson:"zip_code"`
    Country string `json:"country" bson:"country"`
}

type Location struct {
    Type        string    `json:"type" bson:"type"` // "Point"
    Coordinates []float64 `json:"coordinates" bson:"coordinates"` // [longitude, latitude]
}
```

### 📋 预订模型 (Booking)
```go
type Booking struct {
    ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    PropertyID      primitive.ObjectID `json:"propertyId" bson:"property_id"`
    TenantID        primitive.ObjectID `json:"tenantId" bson:"tenant_id"`
    LandlordID      primitive.ObjectID `json:"landlordId" bson:"landlord_id"`
    StartDate       time.Time         `json:"startDate" bson:"start_date"`
    EndDate         time.Time         `json:"endDate" bson:"end_date"`
    MonthlyRent     float64           `json:"monthlyRent" bson:"monthly_rent"`
    SecurityDeposit float64           `json:"securityDeposit" bson:"security_deposit"`
    Status          string            `json:"status" bson:"status"` // "pending", "confirmed", "cancelled", "completed"
    LeaseTerms      string            `json:"leaseTerms" bson:"lease_terms"`
    SpecialRequests string            `json:"specialRequests" bson:"special_requests"`
    CreatedAt       time.Time         `json:"createdAt" bson:"created_at"`
    UpdatedAt       time.Time         `json:"updatedAt" bson:"updated_at"`
}
```

### ⭐ 评价模型 (Review)
```go
type Review struct {
    ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    PropertyID primitive.ObjectID `json:"propertyId" bson:"property_id"`
    ReviewerID primitive.ObjectID `json:"reviewerId" bson:"reviewer_id"`
    RevieweeID primitive.ObjectID `json:"revieweeId" bson:"reviewee_id"`
    Rating     int               `json:"rating" bson:"rating"` // 1-5
    Title      string            `json:"title" bson:"title"`
    Comment    string            `json:"comment" bson:"comment"`
    ReviewType string            `json:"reviewType" bson:"review_type"` // "property", "tenant", "landlord"
    CreatedAt  time.Time         `json:"createdAt" bson:"created_at"`
    UpdatedAt  time.Time         `json:"updatedAt" bson:"updated_at"`
}
```

## 🚀 部署指南

### 🐳 Docker 部署（推荐）

#### 开发环境
```bash
# 启动开发环境
docker-compose -f docker-compose.dev.yml up -d

# 查看日志
docker-compose -f docker-compose.dev.yml logs -f

# 停止服务
docker-compose -f docker-compose.dev.yml down
```

#### 生产环境
```bash
# 构建生产镜像
docker-compose -f docker-compose.prod.yml build

# 启动生产环境
docker-compose -f docker-compose.prod.yml up -d

# 水平扩展后端服务
docker-compose -f docker-compose.prod.yml up -d --scale backend=3
```

### ☸️ Kubernetes 部署

#### 前置条件
- Kubernetes 集群 (1.18+)
- kubectl 已配置
- Ingress 控制器 (nginx-ingress)

#### 部署步骤
```bash
# 创建命名空间
kubectl create namespace renthelp

# 应用配置
kubectl apply -f k8s/

# 检查部署状态
kubectl get pods -n renthelp

# 获取服务 URL
kubectl get ingress -n renthelp
```

### 🔧 性能优化

#### 数据库优化
```javascript
// MongoDB 索引优化
db.properties.createIndex({ location: "2dsphere" })
db.properties.createIndex({ type: 1, price: 1 })
db.properties.createIndex({ owner_id: 1, created_at: -1 })
db.bookings.createIndex({ property_id: 1, start_date: 1, end_date: 1 })
db.users.createIndex({ email: 1 }, { unique: true })
```

#### 缓存策略
```go
// Redis 缓存实现
func (s *PropertyService) GetProperty(id string) (*models.Property, error) {
    // 先查缓存
    if cached := s.cache.Get("property:" + id); cached != nil {
        return cached.(*models.Property), nil
    }
    
    // 数据库查询
    property, err := s.db.GetProperty(id)
    if err != nil {
        return nil, err
    }
    
    // 缓存1小时
    s.cache.Set("property:"+id, property, time.Hour)
    return property, nil
}
```

### 📊 监控和日志

#### 健康检查配置
```yaml
# docker-compose.yml 健康检查
services:
  backend:
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8080/api/v1/health"]
      interval: 30s
      timeout: 10s
      retries: 3
      
  frontend:
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:3000/api/health"]
      interval: 30s
      timeout: 10s
      retries: 3
```

#### Prometheus 监控
```yaml
# prometheus.yml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'renthelp-backend'
    static_configs:
      - targets: ['backend:8080']
    metrics_path: /metrics
    
  - job_name: 'renthelp-frontend'
    static_configs:
      - targets: ['frontend:3000']
    metrics_path: /api/metrics
```

### 🔐 安全配置

#### SSL/TLS 设置
```nginx
# nginx.conf
server {
    listen 443 ssl http2;
    server_name renthelp.com;

    ssl_certificate /etc/ssl/certs/renthelp.crt;
    ssl_certificate_key /etc/ssl/private/renthelp.key;
    
    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_ciphers ECDHE-RSA-AES256-GCM-SHA512;
    
    location / {
        proxy_pass http://frontend:3000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
    
    location /api {
        proxy_pass http://backend:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

#### 防火墙规则
```bash
# UFW 配置
sudo ufw allow ssh
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
sudo ufw deny 3000/tcp  # 阻止直接访问前端
sudo ufw deny 8080/tcp  # 阻止直接访问后端
sudo ufw enable
```

### 💾 备份策略

#### 数据库备份
```bash
#!/bin/bash
# backup.sh

DATE=$(date +%Y%m%d_%H%M%S)
BACKUP_DIR="/backups"
MONGO_CONTAINER="renthelp_mongodb_1"

# 创建备份
docker exec $MONGO_CONTAINER mongodump --db renthelp --out /tmp/backup_$DATE

# 从容器复制备份文件
docker cp $MONGO_CONTAINER:/tmp/backup_$DATE $BACKUP_DIR/

# 上传到云存储 (AWS S3)
aws s3 cp $BACKUP_DIR/backup_$DATE s3://renthelp-backups/ --recursive

# 清理旧备份 (保留30天)
find $BACKUP_DIR -name "backup_*" -mtime +30 -delete
```

## 🔧 开发指南

### 📝 代码规范

#### Go 后端规范
```go
// 命名约定
var userName string          // 变量使用 camelCase
func GetUserProfile()       // 导出函数使用 PascalCase
type UserService struct{}   // 类型使用 PascalCase

// 错误处理
func (s *UserService) CreateUser(req CreateUserRequest) (*User, error) {
    if err := s.validateUser(req); err != nil {
        return nil, fmt.Errorf("validation failed: %w", err)
    }
    
    user, err := s.repository.Create(req)
    if err != nil {
        return nil, fmt.Errorf("failed to create user: %w", err)
    }
    
    return user, nil
}
```

#### React 前端规范
```tsx
// 组件定义
interface UserProfileProps {
  userId: string;
  onUpdate?: (user: User) => void;
}

const UserProfile: React.FC<UserProfileProps> = ({ userId, onUpdate }) => {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchUser = async () => {
      try {
        setLoading(true);
        const userData = await userService.getUser(userId);
        setUser(userData);
      } catch (err) {
        setError('Failed to load user profile');
        console.error('User fetch error:', err);
      } finally {
        setLoading(false);
      }
    };

    fetchUser();
  }, [userId]);

  if (loading) return <CircularProgress />;
  if (error) return <Alert severity="error">{error}</Alert>;
  if (!user) return <Typography>User not found</Typography>;

  return (
    <Card>
      <CardContent>
        <Typography variant="h5">{user.fullName}</Typography>
        <Typography variant="body2" color="textSecondary">
          {user.email}
        </Typography>
      </CardContent>
    </Card>
  );
};

export default UserProfile;
```

### 🧪 测试指南

#### 后端测试
```go
// user_service_test.go
func TestUserService_CreateUser(t *testing.T) {
    tests := []struct {
        name    string
        request CreateUserRequest
        want    *User
        wantErr bool
    }{
        {
            name: "valid user creation",
            request: CreateUserRequest{
                Email:    "test@example.com",
                Password: "password123",
                FullName: "Test User",
                Role:     "tenant",
            },
            want: &User{
                Email:    "test@example.com",
                FullName: "Test User",
                Role:     "tenant",
            },
            wantErr: false,
        },
        {
            name: "invalid email",
            request: CreateUserRequest{
                Email:    "invalid-email",
                Password: "password123",
                FullName: "Test User",
                Role:     "tenant",
            },
            want:    nil,
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            service := NewUserService(mockRepo, mockValidator)
            got, err := service.CreateUser(tt.request)
            
            if (err != nil) != tt.wantErr {
                t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            
            if !tt.wantErr && got.Email != tt.want.Email {
                t.Errorf("CreateUser() email = %v, want %v", got.Email, tt.want.Email)
            }
        })
    }
}
```

#### 前端测试
```tsx
// UserProfile.test.tsx
import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import UserProfile from './UserProfile';
import * as userService from '../services/userService';

jest.mock('../services/userService');
const mockUserService = userService as jest.Mocked<typeof userService>;

describe('UserProfile', () => {
  const mockUser = {
    id: '1',
    email: 'test@example.com',
    fullName: 'Test User',
    role: 'tenant',
  };

  beforeEach(() => {
    jest.clearAllMocks();
  });

  it('renders user profile correctly', async () => {
    mockUserService.getUser.mockResolvedValue(mockUser);

    render(<UserProfile userId="1" />);

    expect(screen.getByRole('progressbar')).toBeInTheDocument();

    await waitFor(() => {
      expect(screen.getByText('Test User')).toBeInTheDocument();
      expect(screen.getByText('test@example.com')).toBeInTheDocument();
    });
  });

  it('displays error when user fetch fails', async () => {
    mockUserService.getUser.mockRejectedValue(new Error('Network error'));

    render(<UserProfile userId="1" />);

    await waitFor(() => {
      expect(screen.getByRole('alert')).toBeInTheDocument();
      expect(screen.getByText(/Failed to load user profile/i)).toBeInTheDocument();
    });
  });
});
```

### 🚀 CI/CD 流水线

```yaml
# .github/workflows/ci.yml
name: CI/CD Pipeline

on:
  push:
    branches: [main, develop]
  pull_request:
    branches: [main]

jobs:
  test-backend:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Setup Go
        uses: actions/setup-go@v3
        with:
          go-version: 1.21
          
      - name: Cache Go modules
        uses: actions/cache@v3
        with:
          path: ~/go/pkg/mod
          key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
          
      - name: Run tests
        run: |
          cd backend
          go test -v -race -coverprofile=coverage.out ./...
          
      - name: Upload coverage
        uses: codecov/codecov-action@v3
        with:
          file: ./backend/coverage.out

  test-frontend:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Setup Node.js
        uses: actions/setup-node@v3
        with:
          node-version: 18
          cache: 'npm'
          cache-dependency-path: frontend/package-lock.json
          
      - name: Install dependencies
        run: |
          cd frontend
          npm ci
          
      - name: Run tests
        run: |
          cd frontend
          npm run test:coverage
          
      - name: Upload coverage
        uses: codecov/codecov-action@v3
        with:
          file: ./frontend/coverage/lcov.info

  deploy:
    needs: [test-backend, test-frontend]
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'
    
    steps:
      - uses: actions/checkout@v3
      
      - name: Deploy to production
        run: |
          echo "🚀 Deploying to production..."
          # 添加部署命令
```

## ❓ 常见问题

### 🐳 Docker 相关

**Q: 端口被占用怎么办？**
```bash
# 检查端口占用
netstat -ano | findstr :3001
netstat -ano | findstr :8081

# 修改 docker-compose.dev.yml 中的端口映射
ports:
  - "3002:3000"  # 修改前端端口
  - "8082:8080"  # 修改后端端口
```

**Q: MongoDB 连接失败？**
```bash
# 检查 MongoDB 容器状态
docker-compose -f docker-compose.dev.yml logs mongodb

# 重启 MongoDB 容器
docker-compose -f docker-compose.dev.yml restart mongodb
```

**Q: 容器内存不足？**
```yaml
# 在 docker-compose.dev.yml 中增加内存限制
services:
  backend:
    mem_limit: 512m
    mem_reservation: 256m
```

### 🌐 网络相关

**Q: 前端无法连接后端？**
```bash
# 检查环境变量
echo $NEXT_PUBLIC_API_URL

# 确保后端服务正常运行
curl http://localhost:8081/api/v1/health
```

**Q: CORS 错误？**
```go
// 在 backend/cmd/server/main.go 中配置 CORS
r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:3001"},
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"*"},
    AllowCredentials: true,
}))
```

### 🗄️ 数据库相关

**Q: 如何重置数据库？**
```bash
# 停止服务并删除数据卷
docker-compose -f docker-compose.dev.yml down -v

# 重新启动服务
docker-compose -f docker-compose.dev.yml up -d
```

**Q: 如何查看数据库内容？**
```bash
# 连接到 MongoDB 容器
docker exec -it rent_help_mongodb_1 mongosh -u admin -p password123

# 切换到应用数据库
use renthelp

# 查看集合
show collections

# 查看用户数据
db.users.find().pretty()
```

### 🔐 认证相关

**Q: JWT 令牌过期？**
```typescript
// 前端自动刷新令牌
const refreshToken = async () => {
  try {
    const response = await api.post('/auth/refresh');
    const { access_token } = response.data;
    localStorage.setItem('token', access_token);
    return access_token;
  } catch (error) {
    // 重定向到登录页
    window.location.href = '/login';
  }
};
```

**Q: 忘记管理员密码？**
```bash
# 连接到 MongoDB 并重置管理员密码
docker exec -it rent_help_mongodb_1 mongosh -u admin -p password123

use renthelp
db.users.updateOne(
  { email: "admin@renthelp.com" },
  { $set: { password: "$2b$12$LQv3c1yqBwqQ4QxQ.3K.oeqJ5UQzQJ5P5QhP5QhP5QhP5QhP5QhP5Q" } }
)
```

## 🛠️ 故障排除

### 📋 服务状态检查

```bash
# 检查所有容器状态
docker-compose -f docker-compose.dev.yml ps

# 查看具体服务日志
docker-compose -f docker-compose.dev.yml logs backend
docker-compose -f docker-compose.dev.yml logs frontend
docker-compose -f docker-compose.dev.yml logs mongodb

# 检查容器资源使用情况
docker stats
```

### 🔍 调试模式

```bash
# 启用后端调试模式
export GIN_MODE=debug

# 启用前端开发模式
export NODE_ENV=development

# 查看详细错误信息
docker-compose -f docker-compose.dev.yml logs -f --tail=100
```

### 🧹 清理和重置

```bash
# 清理未使用的 Docker 资源
docker system prune -f

# 重建所有容器
docker-compose -f docker-compose.dev.yml up -d --build --force-recreate

# 完全重置环境
docker-compose -f docker-compose.dev.yml down -v --remove-orphans
docker system prune -af
```

## 🚦 开发路线图

### ✅ 已完成功能
- [x] 用户认证系统（注册、登录、JWT）
- [x] 房源管理（CRUD 操作）
- [x] 基础搜索和筛选
- [x] 预订管理系统
- [x] 用户角色权限控制
- [x] Docker 容器化部署
- [x] MongoDB 数据持久化
- [x] 响应式 UI 设计

### 🔄 开发中功能
- [ ] 高级搜索和地图集成
- [ ] 实时聊天系统
- [ ] 文件上传和图片管理
- [ ] 邮件通知系统
- [ ] 用户评价和评分

### 📋 计划功能
- [ ] 支付集成 (Stripe/PayPal)
- [ ] 移动端应用 (React Native)
- [ ] 多语言支持 (i18n)
- [ ] 推荐算法
- [ ] 数据分析仪表板
- [ ] API 限流和监控
- [ ] 自动化测试套件
- [ ] 性能优化和缓存

### 🎯 长期目标
- [ ] 微服务架构重构
- [ ] 人工智能房源推荐
- [ ] 区块链智能合约
- [ ] IoT 设备集成
- [ ] 虚拟现实看房
- [ ] 社交功能扩展

## 🤝 贡献指南

我们欢迎所有形式的贡献！无论是代码、文档、设计还是想法。

### 🌟 如何贡献

1. **Fork 项目**
   ```bash
   git clone https://github.com/phuhao00/rent_help.git
   cd rent_help
   ```

2. **创建功能分支**
   ```bash
   git checkout -b feature/amazing-new-feature
   ```

3. **提交更改**
   ```bash
   git add .
   git commit -m "✨ Add amazing new feature"
   ```

4. **推送分支**
   ```bash
   git push origin feature/amazing-new-feature
   ```

5. **创建 Pull Request**
   - 详细描述你的更改
   - 包含相关的 issue 编号
   - 确保所有测试通过

### 📝 提交规范

使用 [Conventional Commits](https://www.conventionalcommits.org/) 规范：

```bash
# 功能添加
git commit -m "✨ feat: add user profile page"

# 错误修复  
git commit -m "🐛 fix: resolve login validation issue"

# 文档更新
git commit -m "📚 docs: update API documentation"

# 性能优化
git commit -m "⚡ perf: optimize database queries"

# 重构代码
git commit -m "♻️ refactor: simplify user service logic"

# 测试添加
git commit -m "✅ test: add user registration tests"
```

### 🔍 代码审查

所有 Pull Request 都需要经过代码审查：

- ✅ 代码符合项目规范
- ✅ 包含适当的测试
- ✅ 文档已更新
- ✅ 通过所有 CI 检查
- ✅ 无安全漏洞

### 🏷️ Issue 标签

| 标签 | 描述 |
|------|------|
| `bug` | 🐛 错误报告 |
| `enhancement` | ✨ 功能增强 |
| `documentation` | 📚 文档改进 |
| `good first issue` | 👶 适合新手 |
| `help wanted` | 🙋 需要帮助 |
| `priority: high` | 🔥 高优先级 |
| `priority: low` | ❄️ 低优先级 |

## 📄 许可证

本项目采用 [MIT 许可证](LICENSE) - 查看 LICENSE 文件了解详情。

```
MIT License

Copyright (c) 2025 RentHelp Team

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

## 📞 联系我们

如有任何问题、建议或合作意向，欢迎联系我们：

### 🔗 项目链接
- **GitHub**: https://github.com/phuhao00/rent_help
- **Issues**: https://github.com/phuhao00/rent_help/issues
- **Discussions**: https://github.com/phuhao00/rent_help/discussions

### 📧 联系方式
- **项目维护者**: [@phuhao00](https://github.com/phuhao00)
- **邮箱**: support@renthelp.com
- **官网**: https://renthelp.com

### 💬 社区支持
- **Discord**: [加入我们的 Discord 服务器](https://discord.gg/renthelp)
- **微信群**: 添加微信 `renthelp-support` 备注"加群"
- **QQ群**: 123456789

### 🌐 社交媒体
- **Twitter**: [@RentHelpApp](https://twitter.com/RentHelpApp)
- **LinkedIn**: [RentHelp](https://linkedin.com/company/renthelp)
- **知乎**: [RentHelp团队](https://zhihu.com/org/renthelp-team)

---

<div align="center">

**⭐ 如果这个项目对你有帮助，请给我们一个 Star！**

**Made with ❤️ by RentHelp Team**

![Visitor Count](https://visitor-badge.laobi.icu/badge?page_id=phuhao00.rent_help)
![GitHub last commit](https://img.shields.io/github/last-commit/phuhao00/rent_help)
![GitHub issues](https://img.shields.io/github/issues/phuhao00/rent_help)
![GitHub pull requests](https://img.shields.io/github/issues-pr/phuhao00/rent_help)

</div>