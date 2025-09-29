// utils/api.js
const app = getApp()

// API基础配置
const API_BASE = 'http://localhost:8080/api/v1'

// 请求拦截器
function request(options) {
  return new Promise((resolve, reject) => {
    // 显示加载提示
    if (options.showLoading !== false) {
      wx.showLoading({
        title: '加载中...',
        mask: true
      })
    }

    // 获取token
    const token = wx.getStorageSync('token')
    
    // 设置请求头
    const header = {
      'Content-Type': 'application/json',
      ...options.header
    }
    
    if (token) {
      header.Authorization = `Bearer ${token}`
    }

    wx.request({
      url: `${API_BASE}${options.url}`,
      method: options.method || 'GET',
      data: options.data || {},
      header: header,
      success: (res) => {
        wx.hideLoading()
        
        if (res.statusCode === 200) {
          resolve(res.data)
        } else if (res.statusCode === 401) {
          // token过期，清除登录状态
          wx.removeStorageSync('token')
          wx.removeStorageSync('userInfo')
          app.globalData.isLogin = false
          
          wx.showToast({
            title: '登录已过期',
            icon: 'none'
          })
          
          // 跳转到登录页面
          setTimeout(() => {
            wx.redirectTo({
              url: '/pages/login/login'
            })
          }, 1500)
          
          reject(new Error('登录已过期'))
        } else {
          const errorMsg = res.data?.error || '请求失败'
          wx.showToast({
            title: errorMsg,
            icon: 'none'
          })
          reject(new Error(errorMsg))
        }
      },
      fail: (error) => {
        wx.hideLoading()
        wx.showToast({
          title: '网络错误',
          icon: 'none'
        })
        reject(error)
      }
    })
  })
}

// API方法
const api = {
  // 认证相关
  auth: {
    // 登录
    login(code) {
      return request({
        url: '/auth/login',
        method: 'POST',
        data: { code }
      })
    },
    
    // 注册/更新用户信息
    register(userInfo) {
      return request({
        url: '/auth/register',
        method: 'POST',
        data: userInfo
      })
    }
  },

  // 用户相关
  user: {
    // 获取用户档案
    getProfile() {
      return request({
        url: '/users/profile',
        method: 'GET'
      })
    },
    
    // 更新用户档案
    updateProfile(data) {
      return request({
        url: '/users/profile',
        method: 'PUT',
        data: data
      })
    }
  },

  // 帖子相关
  post: {
    // 获取帖子列表
    getList(params = {}) {
      return request({
        url: '/posts',
        method: 'GET',
        data: params
      })
    },
    
    // 获取单个帖子
    getDetail(id) {
      return request({
        url: `/posts/${id}`,
        method: 'GET'
      })
    },
    
    // 创建帖子
    create(data) {
      return request({
        url: '/posts',
        method: 'POST',
        data: data
      })
    },
    
    // 更新帖子
    update(id, data) {
      return request({
        url: `/posts/${id}`,
        method: 'PUT',
        data: data
      })
    },
    
    // 删除帖子
    delete(id) {
      return request({
        url: `/posts/${id}`,
        method: 'DELETE'
      })
    }
  },

  // 评论相关
  comment: {
    // 获取帖子评论
    getByPost(postId) {
      return request({
        url: `/comments/post/${postId}`,
        method: 'GET'
      })
    },
    
    // 创建评论
    create(data) {
      return request({
        url: '/comments',
        method: 'POST',
        data: data
      })
    },
    
    // 删除评论
    delete(id) {
      return request({
        url: `/comments/${id}`,
        method: 'DELETE'
      })
    }
  }
}

module.exports = api