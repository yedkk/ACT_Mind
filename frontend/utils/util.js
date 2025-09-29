// utils/util.js

// 格式化时间
const formatTime = date => {
  const year = date.getFullYear()
  const month = date.getMonth() + 1
  const day = date.getDate()
  const hour = date.getHours()
  const minute = date.getMinutes()
  const second = date.getSeconds()

  return `${[year, month, day].map(formatNumber).join('/')} ${[hour, minute, second].map(formatNumber).join(':')}`
}

// 格式化日期
const formatDate = date => {
  const year = date.getFullYear()
  const month = date.getMonth() + 1
  const day = date.getDate()

  return `${[year, month, day].map(formatNumber).join('-')}`
}

// 相对时间格式化
const formatRelativeTime = dateString => {
  const date = new Date(dateString)
  const now = new Date()
  const diff = now.getTime() - date.getTime()
  
  const minute = 60 * 1000
  const hour = 60 * minute
  const day = 24 * hour
  const week = 7 * day
  const month = 30 * day
  
  if (diff < minute) {
    return '刚刚'
  } else if (diff < hour) {
    return `${Math.floor(diff / minute)}分钟前`
  } else if (diff < day) {
    return `${Math.floor(diff / hour)}小时前`
  } else if (diff < week) {
    return `${Math.floor(diff / day)}天前`
  } else if (diff < month) {
    return `${Math.floor(diff / week)}周前`
  } else {
    return formatDate(date)
  }
}

const formatNumber = n => {
  n = n.toString()
  return n[1] ? n : `0${n}`
}

// 防抖函数
const debounce = (func, wait) => {
  let timeout
  return function executedFunction(...args) {
    const later = () => {
      clearTimeout(timeout)
      func(...args)
    }
    clearTimeout(timeout)
    timeout = setTimeout(later, wait)
  }
}

// 节流函数
const throttle = (func, limit) => {
  let inThrottle
  return function() {
    const args = arguments
    const context = this
    if (!inThrottle) {
      func.apply(context, args)
      inThrottle = true
      setTimeout(() => inThrottle = false, limit)
    }
  }
}

// 深拷贝
const deepClone = obj => {
  if (obj === null || typeof obj !== 'object') return obj
  if (obj instanceof Date) return new Date(obj.getTime())
  if (obj instanceof Array) return obj.map(item => deepClone(item))
  if (typeof obj === 'object') {
    const clonedObj = {}
    for (const key in obj) {
      if (obj.hasOwnProperty(key)) {
        clonedObj[key] = deepClone(obj[key])
      }
    }
    return clonedObj
  }
}

// 生成随机ID
const generateId = () => {
  return Math.random().toString(36).substr(2, 9)
}

// 验证手机号
const validatePhone = phone => {
  const reg = /^1[3-9]\d{9}$/
  return reg.test(phone)
}

// 验证邮箱
const validateEmail = email => {
  const reg = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  return reg.test(email)
}

// 截取字符串
const truncateString = (str, length = 50) => {
  if (str.length <= length) return str
  return str.substring(0, length) + '...'
}

// 获取文件扩展名
const getFileExtension = filename => {
  return filename.slice((filename.lastIndexOf('.') - 1 >>> 0) + 2)
}

// 格式化文件大小
const formatFileSize = bytes => {
  if (bytes === 0) return '0 Bytes'
  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

// 检查是否为空对象
const isEmptyObject = obj => {
  return Object.keys(obj).length === 0 && obj.constructor === Object
}

// 数组去重
const uniqueArray = arr => {
  return [...new Set(arr)]
}

// 随机打乱数组
const shuffleArray = arr => {
  const newArr = [...arr]
  for (let i = newArr.length - 1; i > 0; i--) {
    const j = Math.floor(Math.random() * (i + 1))
    ;[newArr[i], newArr[j]] = [newArr[j], newArr[i]]
  }
  return newArr
}

// 获取URL参数
const getUrlParams = url => {
  const params = {}
  const urlSearchParams = new URLSearchParams(url.split('?')[1])
  for (const [key, value] of urlSearchParams) {
    params[key] = value
  }
  return params
}

// 存储管理
const storage = {
  set(key, value) {
    try {
      wx.setStorageSync(key, JSON.stringify(value))
    } catch (e) {
      console.error('存储失败:', e)
    }
  },
  
  get(key) {
    try {
      const value = wx.getStorageSync(key)
      return value ? JSON.parse(value) : null
    } catch (e) {
      console.error('读取存储失败:', e)
      return null
    }
  },
  
  remove(key) {
    try {
      wx.removeStorageSync(key)
    } catch (e) {
      console.error('删除存储失败:', e)
    }
  },
  
  clear() {
    try {
      wx.clearStorageSync()
    } catch (e) {
      console.error('清空存储失败:', e)
    }
  }
}

// 显示提示信息
const showToast = (title, icon = 'none', duration = 2000) => {
  wx.showToast({
    title,
    icon,
    duration
  })
}

// 显示确认对话框
const showConfirm = (content, title = '提示') => {
  return new Promise((resolve) => {
    wx.showModal({
      title,
      content,
      success: (res) => {
        resolve(res.confirm)
      }
    })
  })
}

module.exports = {
  formatTime,
  formatDate,
  formatRelativeTime,
  debounce,
  throttle,
  deepClone,
  generateId,
  validatePhone,
  validateEmail,
  truncateString,
  getFileExtension,
  formatFileSize,
  isEmptyObject,
  uniqueArray,
  shuffleArray,
  getUrlParams,
  storage,
  showToast,
  showConfirm
}