# ACT Mind 后端服务

基于 Go + Gin + GORM + SQLite 的心理健康应用后端服务。

## 项目结构

```
backend/
├── main.go                 # 应用入口
├── go.mod                  # Go模块依赖
├── config/                 # 配置管理
│   ├── config.go          # 配置加载逻辑
│   └── config.yaml        # 配置文件
├── database/              # 数据库相关
│   └── database.go        # 数据库连接和初始化
├── models/                # 数据模型
│   ├── user.go           # 用户模型
│   └── post.go           # 帖子和评论模型
├── controllers/           # 控制器
│   ├── auth.go           # 认证控制器
│   ├── user.go           # 用户控制器
│   ├── post.go           # 帖子控制器
│   ├── comment.go        # 评论控制器
│   └── health.go         # 健康检查
├── middleware/            # 中间件
│   └── jwt.go            # JWT认证中间件
├── routes/               # 路由配置
│   └── routes.go         # 路由设置
├── utils/                # 工具函数
│   └── jwt.go            # JWT工具
└── data/                 # 数据存储目录
    └── act_mind.db       # SQLite数据库文件
```

## 快速开始

### 1. 安装依赖

```bash
go mod tidy
```

### 2. 配置环境

复制并修改配置文件：
```bash
cp config/config.yaml.example config/config.yaml
```

### 3. 运行服务

```bash
go run main.go
```

服务将在 `http://localhost:8080` 启动。

### 4. 查看API文档

访问 `http://localhost:8080/swagger/index.html` 查看Swagger API文档。

## API 端点

### 认证相关
- `POST /api/v1/auth/login` - 用户登录
- `POST /api/v1/auth/register` - 用户注册

### 用户相关（需要JWT认证）
- `GET /api/v1/users/profile` - 获取用户档案
- `PUT /api/v1/users/profile` - 更新用户档案

### 帖子相关（需要JWT认证）
- `GET /api/v1/posts` - 获取帖子列表
- `POST /api/v1/posts` - 创建帖子
- `GET /api/v1/posts/:id` - 获取单个帖子
- `PUT /api/v1/posts/:id` - 更新帖子
- `DELETE /api/v1/posts/:id` - 删除帖子

### 评论相关（需要JWT认证）
- `POST /api/v1/comments` - 创建评论
- `GET /api/v1/comments/post/:post_id` - 获取帖子评论
- `DELETE /api/v1/comments/:id` - 删除评论

### 系统相关
- `GET /health` - 健康检查

## 数据库

项目使用SQLite作为数据库，数据文件存储在 `data/act_mind.db`。

### 数据表结构

- `users` - 用户基础信息
- `user_profiles` - 用户详细档案
- `posts` - 帖子信息
- `comments` - 评论信息

## 开发说明

### JWT认证

所有需要认证的API都需要在请求头中包含JWT token：

```
Authorization: Bearer <your-jwt-token>
```

### 错误处理

API返回标准的HTTP状态码和JSON格式的错误信息：

```json
{
  "error": "错误描述",
  "details": "详细错误信息（可选）"
}
```

### 日志

应用使用logrus进行日志记录，日志级别可在配置文件中设置。

## 部署

### 构建

```bash
go build -o act-mind-backend main.go
```

### 运行

```bash
./act-mind-backend
```

## 环境变量

可以通过环境变量覆盖配置文件中的设置：

- `ENVIRONMENT` - 运行环境（development/production）
- `PORT` - 服务端口
- `DB_PATH` - 数据库文件路径
- `JWT_SECRET` - JWT密钥
- `LOG_LEVEL` - 日志级别

## 待完成功能

- [ ] 完善帖子CRUD操作
- [ ] 完善评论CRUD操作
- [ ] 添加心理测评模块
- [ ] 添加数据分析功能
- [ ] 集成微信小程序登录
- [ ] 添加文件上传功能
- [ ] 添加缓存机制
- [ ] 添加单元测试