// pages/login/login.js
const api = require('../../utils/api')
const util = require('../../utils/util')

Page({
  data: {
    loginLoading: false
  },

  onLoad() {
    // 检查是否已经登录
    const app = getApp()
    if (app.globalData.isLogin) {
      this.redirectToHome()
    }
  },

  // 获取用户信息并登录
  async onGetUserInfo(e) {
    if (e.detail.errMsg !== 'getUserInfo:ok') {
      util.showToast('需要授权才能使用完整功能')
      return
    }

    this.setData({ loginLoading: true })

    try {
      // 获取微信登录凭证
      const loginRes = await this.wxLogin()
      
      // 调用后端登录接口
      const authRes = await api.auth.login(loginRes.code)
      
      // 保存登录信息
      wx.setStorageSync('token', authRes.token)
      wx.setStorageSync('userInfo', authRes.user)
      
      // 更新全局状态
      const app = getApp()
      app.globalData.isLogin = true
      app.globalData.userInfo = authRes.user
      app.globalData.token = authRes.token
      
      // 如果是新用户，更新用户信息
      if (!authRes.user.nickname || authRes.user.nickname === '新用户') {
        await this.updateUserInfo(e.detail.userInfo, loginRes.code)
      }
      
      util.showToast('登录成功', 'success')
      
      // 延迟跳转，让用户看到成功提示
      setTimeout(() => {
        this.redirectToHome()
      }, 1500)
      
    } catch (error) {
      console.error('登录失败:', error)
      util.showToast('登录失败，请重试')
    } finally {
      this.setData({ loginLoading: false })
    }
  },

  // 微信登录
  wxLogin() {
    return new Promise((resolve, reject) => {
      wx.login({
        success: resolve,
        fail: reject
      })
    })
  },

  // 更新用户信息
  async updateUserInfo(userInfo, code) {
    try {
      const updateData = {
        openid: code, // 临时使用code作为openid
        nickname: userInfo.nickName,
        avatar_url: userInfo.avatarUrl
      }
      
      const result = await api.auth.register(updateData)
      
      // 更新本地存储的用户信息
      wx.setStorageSync('userInfo', result.user)
      
      const app = getApp()
      app.globalData.userInfo = result.user
      
    } catch (error) {
      console.error('更新用户信息失败:', error)
      // 不阻断登录流程
    }
  },

  // 跳转到首页
  redirectToHome() {
    wx.switchTab({
      url: '/pages/index/index'
    })
  },

  // 显示隐私政策
  showPrivacyPolicy() {
    wx.showModal({
      title: '隐私政策',
      content: '我们重视您的隐私保护。我们会严格按照相关法律法规保护您的个人信息安全，不会将您的信息用于其他商业用途。',
      showCancel: false,
      confirmText: '知道了'
    })
  },

  // 显示用户协议
  showUserAgreement() {
    wx.showModal({
      title: '用户协议',
      content: '欢迎使用ACT Mind。请您仔细阅读并遵守用户协议，合理使用本应用提供的服务。本应用仅供心理健康辅助，不能替代专业医疗建议。',
      showCancel: false,
      confirmText: '知道了'
    })
  }
})