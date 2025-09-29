# ACT Mind 心理学应用技术栈实施方案

## 项目概述

ACT Mind 是一个基于接受与承诺疗法（Acceptance and Commitment Therapy）的心理健康应用，旨在为用户提供心理测评、社区讨论和个人成长工具。

### 确定技术栈

经过综合考虑，项目采用以下技术栈：

- **前端**：微信小程序原生开发
- **后端**：Go + Gin 框架 + GORM
- **数据库**：SQLite（初期）→ PostgreSQL（扩展期）
- **部署**：云服务器 + Nginx

## 技术栈详细分析

### 前端：微信小程序原生开发

#### 选择理由
- **用户基数大**：微信12亿+用户，获客成本低
- **社交传播**：适合心理测试分享和社区讨论
- **学习成本低**：基于Web技术，语法简单
- **开发工具完善**：微信开发者工具功能强大

#### 核心能力
- 丰富的内置组件（30+）
- 完善的动画支持（CSS3 + API）
- 社交分享能力
- 用户授权和支付功能

### 后端：Go + Gin + GORM

#### 选择理由
- **高性能**：编译型语言，并发处理能力强
- **类型安全**：编译时错误检查，减少运行时问题
- **简洁部署**：单文件部署，运维成本低
- **学习价值**：提升Go语言技能，符合职业发展目标

#### 技术组合
- **Web框架**：Gin（轻量级、高性能）
- **ORM**：GORM（功能完整、代码生成）
- **配置管理**：Viper
- **日志系统**：Logrus
- **认证**：JWT + 中间件
- **API文档**：Swagger

### 数据库：SQLite → PostgreSQL

#### 阶段规划
- **初期（0-1万用户）**：SQLite
  - 零配置，文件型数据库
  - 适合快速原型和MVP验证
  - 部署简单，维护成本低

- **扩展期（1万+用户）**：PostgreSQL
  - 更好的并发支持
  - 丰富的数据类型和扩展
  - 成熟的生态系统

## 项目架构设计

### 整体架构图

```
┌─────────────────┐    ┌─────────────────┐
│   微信小程序     │    │   管理后台      │
│   (前端界面)     │    │   (Web界面)     │
└─────────┬───────┘    └─────────┬───────┘
          │                      │
          │ HTTPS API 调用       │
          │                      │
    ┌─────▼──────────────────────▼─────┐
    │         Nginx 反向代理           │
    │      (SSL终止 + 负载均衡)        │
    └─────────────┬───────────────────┘
                  │
    ┌─────────────▼───────────────────┐
    │        Go Gin 应用服务器        │
    │  ┌─────────┬─────────┬────────┐ │
    │  │ 用户管理 │ 内容管理 │ 社区管理│ │
    │  └─────────┴─────────┴────────┘ │
    └─────────────┬───────────────────┘
                  │
    ┌─────────────▼───────────────────┐
    │         SQLite 数据库           │
    │    (用户数据 + 内容数据)        │
    └─────────────────────────────────┘
```

### 数据库设计

#### 核心数据模型

```sql
-- 用户表
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    openid VARCHAR(100) UNIQUE NOT NULL,
    nickname VARCHAR(50),
    avatar_url VARCHAR(255),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- 用户档案表
CREATE TABLE user_profiles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    bio TEXT,
    psychological_data JSON,
    privacy_settings JSON,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- 帖子表
CREATE TABLE posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    title VARCHAR(200) NOT NULL,
    content TEXT NOT NULL,
    category VARCHAR(50),
    view_count INTEGER DEFAULT 0,
    like_count INTEGER DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- 评论表
CREATE TABLE comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    post_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (post_id) REFERENCES posts(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
```

## 开发计划

### 第一阶段：项目基础搭建（1-2周）

#### 后端开发
- [ ] Go项目初始化和依赖管理
- [ ] 数据库模型定义和迁移
- [ ] 基础API路由和中间件
- [ ] JWT认证系统
- [ ] 用户注册登录API

#### 前端开发
- [ ] 小程序项目初始化
- [ ] 基础页面结构搭建
- [ ] 用户授权和登录流程
- [ ] API调用封装

### 第二阶段：核心功能开发（2-3周）

#### 用户系统
- [ ] 用户档案管理
- [ ] 个人信息编辑
- [ ] 隐私设置

#### 内容系统
- [ ] 帖子发布和编辑
- [ ] 内容浏览和搜索
- [ ] 分类管理

#### 社区功能
- [ ] 评论系统
- [ ] 点赞功能
- [ ] 用户互动

### 第三阶段：高级功能（3-4周）

#### 心理测评
- [ ] 测评问卷系统
- [ ] 结果分析和展示
- [ ] 历史记录管理

#### 数据分析
- [ ] 用户行为统计
- [ ] 内容热度分析
- [ ] 个人成长报告

### 第四阶段：优化和部署（1-2周）

#### 性能优化
- [ ] 数据库查询优化
- [ ] 缓存策略实施
- [ ] 前端性能优化

#### 部署上线
- [ ] 服务器环境配置
- [ ] CI/CD流程搭建
- [ ] 监控和日志系统
- [ ] 小程序审核和发布

## 开发环境配置

### 后端环境
```bash
# Go版本要求
go version >= 1.19

# 主要依赖
github.com/gin-gonic/gin
gorm.io/gorm
gorm.io/driver/sqlite
github.com/golang-jwt/jwt/v4
github.com/spf13/viper
github.com/sirupsen/logrus
```

### 前端环境
```bash
# 微信开发者工具
版本 >= 1.06.0

# 基础库版本
微信小程序基础库 >= 2.10.0
```

## 部署方案

### 服务器配置
- **推荐配置**：2核4G，5M带宽
- **操作系统**：Ubuntu 20.04 LTS
- **年度成本**：约500-800元

### 域名和证书
- **域名**：.com域名，约60元/年
- **SSL证书**：Let's Encrypt免费证书
- **备案**：ICP备案（免费，10-20天）

### 部署架构
```
用户请求 → CDN → Nginx → Go应用 → SQLite
```

## 风险评估和缓解

### 技术风险
1. **Go学习曲线**
   - 缓解：充分利用现有编程经验，重点学习Go特性
   - 时间：预计1-2周掌握基础开发

2. **小程序政策变化**
   - 缓解：同步开发H5版本作为备选方案
   - 关注：定期查看微信官方政策更新

3. **数据库扩展**
   - 缓解：设计时考虑迁移方案，使用GORM便于切换
   - 时机：用户量达到5000时考虑迁移

### 业务风险
1. **用户增长超预期**
   - 缓解：提前准备扩容方案
   - 监控：设置用户量和性能指标告警

2. **内容合规要求**
   - 缓解：建立内容审核机制
   - 预防：了解心理健康类应用相关法规

## 成功指标

### 技术指标
- API响应时间 < 200ms
- 小程序启动时间 < 3s
- 数据库查询效率 > 95%
- 系统可用性 > 99%

### 业务指标
- 用户注册转化率 > 60%
- 日活跃用户留存 > 30%
- 内容发布活跃度 > 20%
- 用户满意度评分 > 4.0

## 下一步行动

1. **立即开始**：搭建开发环境和项目结构
2. **第一周目标**：完成用户认证系统
3. **第二周目标**：实现基础内容发布功能
4. **第三周目标**：完成社区讨论功能
5. **第四周目标**：集成心理测评模块

---

*本文档将随着项目进展持续更新，确保技术方案与实际开发保持同步。*