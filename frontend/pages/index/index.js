// pages/index/index.js
const api = require('../../utils/api')
const util = require('../../utils/util')

Page({
  data: {
    greeting: '你好',
    todayRecommendation: null,
    recentPosts: [],
    loading: false
  },

  onLoad() {
    this.setGreeting()
    this.loadData()
  },

  onShow() {
    // 检查登录状态
    const app = getApp()
    if (!app.globalData.isLogin) {
      wx.redirectTo({
        url: '/pages/login/login'
      })
      return
    }
    
    // 刷新数据
    this.loadData()
  },

  onPullDownRefresh() {
    this.loadData().finally(() => {
      wx.stopPullDownRefresh()
    })
  },

  // 设置问候语
  setGreeting() {
    const hour = new Date().getHours()
    let greeting = '你好'
    
    if (hour < 6) {
      greeting = '夜深了'
    } else if (hour < 12) {
      greeting = '早上好'
    } else if (hour < 18) {
      greeting = '下午好'
    } else {
      greeting = '晚上好'
    }
    
    this.setData({ greeting })
  },

  // 加载数据
  async loadData() {
    this.setData({ loading: true })
    
    try {
      // 并行加载推荐内容和最新动态
      const [recommendation, posts] = await Promise.all([
        this.loadTodayRecommendation(),
        this.loadRecentPosts()
      ])
      
      this.setData({
        todayRecommendation: recommendation,
        recentPosts: posts
      })
    } catch (error) {
      console.error('加载数据失败:', error)
      util.showToast('加载失败，请重试')
    } finally {
      this.setData({ loading: false })
    }
  },

  // 加载今日推荐
  async loadTodayRecommendation() {
    // 模拟推荐内容
    const recommendations = [
      {
        id: 1,
        title: '接受不完美的自己',
        category: '心理成长',
        content: '完美主义往往是焦虑的源头。学会接受自己的不完美，是心理健康的重要一步。'
      },
      {
        id: 2,
        title: '正念冥想的力量',
        category: '正念练习',
        content: '正念冥想可以帮助我们专注当下，减少对过去的悔恨和对未来的担忧。'
      },
      {
        id: 3,
        title: '建立健康的边界',
        category: '人际关系',
        content: '学会说"不"，建立健康的人际边界，保护自己的心理空间。'
      }
    ]
    
    // 随机选择一个推荐
    const randomIndex = Math.floor(Math.random() * recommendations.length)
    return recommendations[randomIndex]
  },

  // 加载最新动态
  async loadRecentPosts() {
    try {
      const response = await api.post.getList({ limit: 3 })
      
      // 格式化时间
      const posts = response.map(post => ({
        ...post,
        created_at: util.formatRelativeTime(post.created_at)
      }))
      
      return posts
    } catch (error) {
      console.error('加载最新动态失败:', error)
      return []
    }
  },

  // 跳转到心理测评
  goToAssessment() {
    util.showToast('心理测评功能开发中')
  },

  // 跳转到社区
  goToCommunity() {
    wx.switchTab({
      url: '/pages/posts/posts'
    })
  },

  // 跳转到心情日记
  goToJournal() {
    util.showToast('心情日记功能开发中')
  },

  // 跳转到正念练习
  goToExercise() {
    util.showToast('正念练习功能开发中')
  },

  // 阅读推荐内容
  readRecommendation() {
    const { todayRecommendation } = this.data
    if (todayRecommendation) {
      wx.showModal({
        title: todayRecommendation.title,
        content: todayRecommendation.content,
        showCancel: false,
        confirmText: '知道了'
      })
    }
  },

  // 跳转到帖子详情
  goToPostDetail(e) {
    const { id } = e.currentTarget.dataset
    wx.navigateTo({
      url: `/pages/post-detail/post-detail?id=${id}`
    })
  }
})