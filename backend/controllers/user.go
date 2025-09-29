package controllers

import (
	"net/http"

	"act-mind-backend/database"
	"act-mind-backend/middleware"
	"act-mind-backend/models"

	"github.com/gin-gonic/gin"
)

type UpdateProfileRequest struct {
	Nickname          string `json:"nickname"`
	Bio               string `json:"bio"`
	PsychologicalData string `json:"psychological_data"`
	PrivacySettings   string `json:"privacy_settings"`
}

// GetUserProfile 获取用户档案
// @Summary 获取用户档案
// @Description 获取当前用户的详细档案信息
// @Tags 用户
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.User
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /users/profile [get]
func GetUserProfile(c *gin.Context) {
	userID, exists := middleware.GetCurrentUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "用户未认证",
		})
		return
	}

	db := database.GetDB()
	var user models.User

	// 预加载用户档案
	result := db.Preload("Profile").First(&user, userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "用户不存在",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUserProfile 更新用户档案
// @Summary 更新用户档案
// @Description 更新当前用户的档案信息
// @Tags 用户
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body UpdateProfileRequest true "更新请求"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /users/profile [put]
func UpdateUserProfile(c *gin.Context) {
	userID, exists := middleware.GetCurrentUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "用户未认证",
		})
		return
	}

	var req UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请求参数错误",
			"details": err.Error(),
		})
		return
	}

	db := database.GetDB()

	// 更新用户基本信息
	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "用户不存在",
		})
		return
	}

	// 更新昵称
	if req.Nickname != "" {
		user.Nickname = req.Nickname
		db.Save(&user)
	}

	// 更新或创建用户档案
	var profile models.UserProfile
	result := db.Where("user_id = ?", userID).First(&profile)
	
	if result.Error != nil {
		// 档案不存在，创建新档案
		profile = models.UserProfile{
			UserID:            userID,
			Bio:               req.Bio,
			PsychologicalData: req.PsychologicalData,
			PrivacySettings:   req.PrivacySettings,
		}
		db.Create(&profile)
	} else {
		// 档案存在，更新信息
		if req.Bio != "" {
			profile.Bio = req.Bio
		}
		if req.PsychologicalData != "" {
			profile.PsychologicalData = req.PsychologicalData
		}
		if req.PrivacySettings != "" {
			profile.PrivacySettings = req.PrivacySettings
		}
		db.Save(&profile)
	}

	// 重新加载用户信息
	db.Preload("Profile").First(&user, userID)

	c.JSON(http.StatusOK, user)
}