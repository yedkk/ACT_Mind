package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateComment 创建评论
func CreateComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "创建评论 - 待实现",
	})
}

// GetCommentsByPost 获取帖子的评论列表
func GetCommentsByPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "获取帖子评论 - 待实现",
	})
}

// DeleteComment 删除评论
func DeleteComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "删除评论 - 待实现",
	})
}