// app.js
App({
  onLaunch() {
    console.log('ACT Mind 小程序启动')
    
    // 检查登录状态
    this.checkLoginStatus()
    
    // 获取系统信息
    this.getSystemInfo()
  },

  onShow() {
    console.log('ACT Mind 小程序显示')
  },

  onHide() {
    console.log('ACT Mind 小程序隐藏')
  },

  // 检查登录状态
  checkLoginStatus() {
    const token = wx.getStorageSync('token')
    const userInfo = wx.getStorageSync('userInfo')
    
    if (token && userInfo) {
      this.globalData.isLogin = true
      this.globalData.userInfo = userInfo
      this.globalData.token = token
    } else {
      this.globalData.isLogin = false
    }
  },

  // 获取系统信息
  getSystemInfo() {
    wx.getSystemInfo({
      success: (res) => {
        this.globalData.systemInfo = res
        console.log('系统信息:', res)
      }
    })
  },

  // 登录方法
  login() {
    return new Promise((resolve, reject) => {
      wx.login({
        success: (res) => {
          if (res.code) {
            // 调用后端登录接口
            wx.request({
              url: `${this.globalData.apiBase}/auth/login`,
              method: 'POST',
              data: {
                code: res.code
              },
              success: (response) => {
                if (response.statusCode === 200) {
                  const { token, user } = response.data
                  
                  // 保存登录信息
                  wx.setStorageSync('token', token)
                  wx.setStorageSync('userInfo', user)
                  
                  this.globalData.isLogin = true
                  this.globalData.userInfo = user
                  this.globalData.token = token
                  
                  resolve(response.data)
                } else {
                  reject(new Error('登录失败'))
                }
              },
              fail: (error) => {
                reject(error)
              }
            })
          } else {
            reject(new Error('获取登录凭证失败'))
          }
        },
        fail: (error) => {
          reject(error)
        }
      })
    })
  },

  // 退出登录
  logout() {
    wx.removeStorageSync('token')
    wx.removeStorageSync('userInfo')
    
    this.globalData.isLogin = false
    this.globalData.userInfo = null
    this.globalData.token = null
    
    // 跳转到登录页面
    wx.redirectTo({
      url: '/pages/login/login'
    })
  },

  // 全局数据
  globalData: {
    apiBase: 'http://localhost:8080/api/v1', // 后端API地址
    isLogin: false,
    userInfo: null,
    token: null,
    systemInfo: null
  }
})