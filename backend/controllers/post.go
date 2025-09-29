package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPosts 获取帖子列表
func GetPosts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "获取帖子列表 - 待实现",
	})
}

// CreatePost 创建帖子
func CreatePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "创建帖子 - 待实现",
	})
}

// GetPost 获取单个帖子
func GetPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "获取单个帖子 - 待实现",
	})
}

// UpdatePost 更新帖子
func UpdatePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "更新帖子 - 待实现",
	})
}

// DeletePost 删除帖子
func DeletePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "删除帖子 - 待实现",
	})
}