# ACT Mind 微信小程序

基于微信小程序原生开发的心理健康应用前端。

## 项目结构

```
frontend/
├── app.js                  # 小程序逻辑
├── app.json               # 小程序配置
├── app.wxss               # 小程序公共样式
├── sitemap.json           # 索引配置
├── project.config.json    # 项目配置
├── pages/                 # 页面目录
│   ├── index/            # 首页
│   │   ├── index.wxml    # 页面结构
│   │   ├── index.wxss    # 页面样式
│   │   ├── index.js      # 页面逻辑
│   │   └── index.json    # 页面配置
│   ├── login/            # 登录页
│   ├── profile/          # 个人中心
│   ├── posts/            # 社区帖子列表
│   ├── post-detail/      # 帖子详情
│   └── create-post/      # 发布帖子
├── utils/                # 工具函数
│   ├── api.js           # API调用封装
│   └── util.js          # 通用工具函数
└── images/              # 图片资源
    ├── logo.png
    ├── home.png
    ├── community.png
    └── profile.png
```

## 功能特性

### 已实现功能
- ✅ 用户登录/注册
- ✅ 首页展示
- ✅ API调用封装
- ✅ 全局状态管理
- ✅ 响应式设计

### 开发中功能
- 🚧 社区帖子列表
- 🚧 帖子详情页面
- 🚧 个人中心
- 🚧 发布帖子
- 🚧 心理测评

### 计划功能
- 📋 心情日记
- 📋 正念练习
- 📋 数据统计
- 📋 消息通知

## 开发指南

### 1. 环境准备

1. 下载并安装[微信开发者工具](https://developers.weixin.qq.com/miniprogram/dev/devtools/download.html)
2. 注册微信小程序账号并获取AppID
3. 在`project.config.json`中配置你的AppID

### 2. 项目配置

1. 修改`utils/api.js`中的API基础地址：
```javascript
const API_BASE = 'https://your-domain.com/api/v1'
```

2. 配置服务器域名：
   - 在微信公众平台配置服务器域名
   - 确保域名支持HTTPS

### 3. 开发调试

1. 用微信开发者工具打开项目
2. 启动后端服务
3. 在开发者工具中预览和调试

### 4. 页面说明

#### 首页 (pages/index)
- 显示欢迎信息和快速功能入口
- 展示今日推荐内容
- 显示最新社区动态

#### 登录页 (pages/login)
- 微信授权登录
- 用户信息获取和更新
- 隐私政策和用户协议

#### 个人中心 (pages/profile)
- 用户信息展示和编辑
- 个人设置
- 历史记录

## API接口

### 认证相关
- `POST /auth/login` - 用户登录
- `POST /auth/register` - 用户注册

### 用户相关
- `GET /users/profile` - 获取用户档案
- `PUT /users/profile` - 更新用户档案

### 帖子相关
- `GET /posts` - 获取帖子列表
- `POST /posts` - 创建帖子
- `GET /posts/:id` - 获取帖子详情

## 样式规范

### 颜色规范
- 主色调：`#4CAF50` (绿色)
- 辅助色：`#45a049` (深绿色)
- 文字色：`#333` (深灰)
- 辅助文字：`#666` (中灰)
- 提示文字：`#999` (浅灰)

### 尺寸规范
- 页面边距：`30rpx`
- 卡片圆角：`16rpx`
- 按钮圆角：`12rpx`
- 头像圆角：`50%`

### 字体规范
- 标题：`32rpx` / `font-weight: 600`
- 正文：`28rpx` / `font-weight: 400`
- 辅助文字：`24rpx` / `font-weight: 400`
- 小字：`22rpx` / `font-weight: 400`

## 组件规范

### 按钮组件
```html
<button class="btn btn-primary">主要按钮</button>
<button class="btn btn-secondary">次要按钮</button>
<button class="btn btn-outline">边框按钮</button>
```

### 卡片组件
```html
<view class="card">
  <view class="card-header">
    <view class="card-title">标题</view>
  </view>
  <view class="card-content">内容</view>
</view>
```

### 列表组件
```html
<view class="list">
  <view class="list-item">
    <view class="list-item-content">
      <view class="list-item-title">标题</view>
      <view class="list-item-desc">描述</view>
    </view>
  </view>
</view>
```

## 部署说明

### 1. 代码审核
- 确保代码符合微信小程序规范
- 检查敏感词和违规内容
- 测试所有功能正常

### 2. 提交审核
- 在微信公众平台提交代码审核
- 填写版本说明和功能介绍
- 等待审核结果

### 3. 发布上线
- 审核通过后发布上线
- 监控用户反馈和错误日志
- 及时修复问题

## 注意事项

1. **网络请求**：小程序要求使用HTTPS协议
2. **域名配置**：需要在微信公众平台配置合法域名
3. **用户授权**：获取用户信息需要用户主动授权
4. **性能优化**：注意包大小限制和页面加载速度
5. **兼容性**：测试不同机型和微信版本的兼容性

## 常见问题

### Q: 网络请求失败
A: 检查域名是否已在微信公众平台配置，确保使用HTTPS协议

### Q: 用户授权失败
A: 确保使用正确的授权方式，处理用户拒绝授权的情况

### Q: 页面跳转异常
A: 检查页面路径是否正确，确保页面已在app.json中注册

### Q: 样式显示异常
A: 检查rpx单位使用是否正确，确保样式文件路径正确

## 更新日志

### v1.0.0 (2024-01-XX)
- 初始版本发布
- 实现用户登录功能
- 完成首页基础布局
- 集成API调用功能