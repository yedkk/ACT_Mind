package controllers

import (
	"net/http"
	"time"

	"act-mind-backend/database"

	"github.com/gin-gonic/gin"
)

type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Database  string    `json:"database"`
	Version   string    `json:"version"`
}

// HealthCheck 健康检查
// @Summary 健康检查
// @Description 检查服务器和数据库状态
// @Tags 系统
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	response := HealthResponse{
		Status:    "ok",
		Timestamp: time.Now(),
		Version:   "1.0.0",
	}

	// 检查数据库连接
	db := database.GetDB()
	sqlDB, err := db.DB()
	if err != nil {
		response.Database = "error"
		response.Status = "error"
	} else {
		if err := sqlDB.Ping(); err != nil {
			response.Database = "disconnected"
			response.Status = "error"
		} else {
			response.Database = "connected"
		}
	}

	if response.Status == "error" {
		c.JSON(http.StatusServiceUnavailable, response)
	} else {
		c.JSON(http.StatusOK, response)
	}
}