package controllers

import (
	"net/http"

	"act-mind-backend/database"
	"act-mind-backend/models"
	"act-mind-backend/utils"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Code string `json:"code" binding:"required"` // 微信小程序登录凭证
}

type RegisterRequest struct {
	OpenID    string `json:"openid" binding:"required"`
	Nickname  string `json:"nickname"`
	AvatarURL string `json:"avatar_url"`
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

// Login 用户登录
// @Summary 用户登录
// @Description 微信小程序用户登录
// @Tags 认证
// @Accept json
// @Produce json
// @Param request body LoginRequest true "登录请求"
// @Success 200 {object} AuthResponse
// @Failure 400 {object} map[string]interface{}
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请求参数错误",
			"details": err.Error(),
		})
		return
	}

	// TODO: 这里应该调用微信API验证code并获取openid
	// 现在为了演示，直接使用code作为openid
	openID := req.Code

	// 查找或创建用户
	var user models.User
	db := database.GetDB()
	
	result := db.Where("open_id = ?", openID).First(&user)
	if result.Error != nil {
		// 用户不存在，创建新用户
		user = models.User{
			OpenID:   openID,
			Nickname: "新用户",
		}
		
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "用户创建失败",
			})
			return
		}

		// 创建用户档案
		profile := models.UserProfile{
			UserID: user.ID,
		}
		db.Create(&profile)
	}

	// 生成JWT token
	token, err := utils.GenerateToken(user.ID, user.OpenID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "token生成失败",
		})
		return
	}

	c.JSON(http.StatusOK, AuthResponse{
		Token: token,
		User:  user,
	})
}

// Register 用户注册
// @Summary 用户注册
// @Description 用户注册（更新用户信息）
// @Tags 认证
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "注册请求"
// @Success 200 {object} AuthResponse
// @Failure 400 {object} map[string]interface{}
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请求参数错误",
			"details": err.Error(),
		})
		return
	}

	db := database.GetDB()
	var user models.User

	// 查找用户
	result := db.Where("open_id = ?", req.OpenID).First(&user)
	if result.Error != nil {
		// 用户不存在，创建新用户
		user = models.User{
			OpenID:    req.OpenID,
			Nickname:  req.Nickname,
			AvatarURL: req.AvatarURL,
		}
		
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "用户创建失败",
			})
			return
		}

		// 创建用户档案
		profile := models.UserProfile{
			UserID: user.ID,
		}
		db.Create(&profile)
	} else {
		// 用户存在，更新信息
		user.Nickname = req.Nickname
		user.AvatarURL = req.AvatarURL
		db.Save(&user)
	}

	// 生成JWT token
	token, err := utils.GenerateToken(user.ID, user.OpenID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "token生成失败",
		})
		return
	}

	c.JSON(http.StatusOK, AuthResponse{
		Token: token,
		User:  user,
	})
}