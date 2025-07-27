# API 文档

## 基础信息

- **Base URL**: `http://localhost:8080/api/v1`
- **Content-Type**: `application/json`
- **认证方式**: Bearer Token (JWT)

## 错误响应格式

```json
{
  "error": "错误描述信息"
}
```

## 认证接口

### 用户注册
- **URL**: `POST /auth/register`
- **请求体**:
```json
{
  "email": "user@example.com",
  "password": "password123",
  "name": "用户姓名",
  "phone": "13800138000",
  "role": "tenant" // 或 "landlord"
}
```
- **响应**:
```json
{
  "message": "User created successfully",
  "user": {
    "id": "user_id",
    "email": "user@example.com",
    "name": "用户姓名",
    "role": "tenant"
  }
}
```

### 用户登录
- **URL**: `POST /auth/login`
- **请求体**:
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```
- **响应**:
```json
{
  "access_token": "jwt_token",
  "expires_in": 86400
}
```

### 刷新令牌
- **URL**: `POST /auth/refresh`
- **Header**: `Authorization: Bearer <token>`
- **响应**: 与登录响应相同

## 用户接口

### 获取用户资料
- **URL**: `GET /users/profile`
- **Header**: `Authorization: Bearer <token>`
- **响应**:
```json
{
  "id": "user_id",
  "email": "user@example.com",
  "name": "用户姓名",
  "phone": "13800138000",
  "role": "tenant",
  "verified": false,
  "created_at": "2025-01-01T00:00:00Z",
  "updated_at": "2025-01-01T00:00:00Z"
}
```

### 更新用户资料
- **URL**: `PUT /users/profile`
- **Header**: `Authorization: Bearer <token>`
- **请求体**:
```json
{
  "name": "新姓名",
  "phone": "新手机号"
}
```
- **响应**:
```json
{
  "message": "Profile updated successfully"
}
```

## 房源接口

### 获取房源列表
- **URL**: `GET /properties`
- **Header**: `Authorization: Bearer <token>`
- **查询参数**:
  - `limit`: 返回数量限制 (默认: 10)
  - `skip`: 跳过数量 (默认: 0)
  - `city`: 城市筛选
  - `type`: 房源类型筛选 (`apartment`, `house`, `room`)
- **响应**:
```json
[
  {
    "id": "property_id",
    "title": "房源标题",
    "description": "房源描述",
    "type": "apartment",
    "price": 3000,
    "currency": "CNY",
    "location": {
      "address": "详细地址",
      "city": "城市",
      "state": "省份",
      "country": "国家",
      "latitude": 39.9042,
      "longitude": 116.4074
    },
    "features": ["WiFi", "空调", "停车位"],
    "images": ["image_url1", "image_url2"],
    "available": true,
    "owner_id": "owner_id",
    "created_at": "2025-01-01T00:00:00Z",
    "updated_at": "2025-01-01T00:00:00Z"
  }
]
```

### 获取单个房源详情
- **URL**: `GET /properties/{id}`
- **Header**: `Authorization: Bearer <token>`
- **响应**: 单个房源对象 (格式同上)

### 创建房源
- **URL**: `POST /properties`
- **Header**: `Authorization: Bearer <token>`
- **请求体**:
```json
{
  "title": "房源标题",
  "description": "房源描述",
  "type": "apartment",
  "price": 3000,
  "currency": "CNY",
  "location": {
    "address": "详细地址",
    "city": "城市",
    "state": "省份",
    "country": "国家",
    "latitude": 39.9042,
    "longitude": 116.4074
  },
  "features": ["WiFi", "空调", "停车位"],
  "images": ["image_url1", "image_url2"]
}
```
- **响应**: 创建的房源对象

### 更新房源
- **URL**: `PUT /properties/{id}`
- **Header**: `Authorization: Bearer <token>`
- **权限**: 仅房源所有者
- **请求体**: 与创建房源相同 (可部分更新)
- **响应**:
```json
{
  "message": "Property updated successfully"
}
```

### 删除房源
- **URL**: `DELETE /properties/{id}`
- **Header**: `Authorization: Bearer <token>`
- **权限**: 仅房源所有者
- **响应**:
```json
{
  "message": "Property deleted successfully"
}
```

## 预订接口

### 获取预订列表
- **URL**: `GET /bookings`
- **Header**: `Authorization: Bearer <token>`
- **查询参数**:
  - `limit`: 返回数量限制 (默认: 10)
  - `skip`: 跳过数量 (默认: 0)
  - `status`: 状态筛选 (`pending`, `confirmed`, `cancelled`, `completed`)
- **响应**:
```json
[
  {
    "id": "booking_id",
    "property_id": "property_id",
    "tenant_id": "tenant_id",
    "start_date": "2025-02-01T00:00:00Z",
    "end_date": "2025-03-01T00:00:00Z",
    "status": "pending",
    "total_price": 3000,
    "message": "预订备注",
    "created_at": "2025-01-01T00:00:00Z",
    "updated_at": "2025-01-01T00:00:00Z"
  }
]
```

### 获取单个预订详情
- **URL**: `GET /bookings/{id}`
- **Header**: `Authorization: Bearer <token>`
- **权限**: 仅预订所有者
- **响应**: 单个预订对象 (格式同上)

### 创建预订
- **URL**: `POST /bookings`
- **Header**: `Authorization: Bearer <token>`
- **请求体**:
```json
{
  "property_id": "property_id",
  "start_date": "2025-02-01T00:00:00Z",
  "end_date": "2025-03-01T00:00:00Z",
  "total_price": 3000,
  "message": "预订备注"
}
```
- **响应**: 创建的预订对象

### 更新预订
- **URL**: `PUT /bookings/{id}`
- **Header**: `Authorization: Bearer <token>`
- **权限**: 仅预订所有者
- **请求体**:
```json
{
  "status": "confirmed",
  "message": "更新备注"
}
```
- **响应**:
```json
{
  "message": "Booking updated successfully"
}
```

### 删除预订
- **URL**: `DELETE /bookings/{id}`
- **Header**: `Authorization: Bearer <token>`
- **权限**: 仅预订所有者
- **响应**:
```json
{
  "message": "Booking deleted successfully"
}
```

## 健康检查

### 服务状态检查
- **URL**: `GET /health`
- **响应**:
```json
{
  "status": "ok"
}
```

## 状态码说明

- `200`: 成功
- `201`: 创建成功
- `400`: 请求参数错误
- `401`: 未认证或认证失败
- `403`: 权限不足
- `404`: 资源不存在
- `409`: 资源冲突 (如邮箱已存在)
- `500`: 服务器内部错误

## 使用示例

### 完整的用户注册登录流程

1. **注册用户**:
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "tenant@example.com",
    "password": "password123",
    "name": "张三",
    "phone": "13800138000",
    "role": "tenant"
  }'
```

2. **登录获取令牌**:
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "tenant@example.com",
    "password": "password123"
  }'
```

3. **使用令牌访问受保护接口**:
```bash
curl -X GET http://localhost:8080/api/v1/users/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 房源操作示例

1. **获取房源列表**:
```bash
curl -X GET "http://localhost:8080/api/v1/properties?city=北京&limit=5" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

2. **创建房源** (房东):
```bash
curl -X POST http://localhost:8080/api/v1/properties \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "市中心精装公寓",
    "description": "位于市中心，交通便利，精装修",
    "type": "apartment",
    "price": 3500,
    "currency": "CNY",
    "location": {
      "address": "朝阳区建国门外大街1号",
      "city": "北京",
      "state": "北京市",
      "country": "中国",
      "latitude": 39.9042,
      "longitude": 116.4074
    },
    "features": ["WiFi", "空调", "洗衣机", "停车位"],
    "images": ["https://example.com/image1.jpg"]
  }'
```

3. **创建预订** (租客):
```bash
curl -X POST http://localhost:8080/api/v1/bookings \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "property_id": "PROPERTY_ID",
    "start_date": "2025-02-01T00:00:00Z",
    "end_date": "2025-03-01T00:00:00Z",
    "total_price": 3500,
    "message": "希望能尽快入住"
  }'
```
