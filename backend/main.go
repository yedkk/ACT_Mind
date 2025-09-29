package main

import (
	"log"

	"act-mind-backend/config"
	"act-mind-backend/database"
	"act-mind-backend/routes"

	"github.com/gin-gonic/gin"
)

// @title ACT Mind API
// @version 1.0
// @description 心理健康应用后端API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @schemes http https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// 加载配置
	config.LoadConfig()

	// 初始化数据库
	database.InitDB()

	// 设置Gin模式
	if config.AppConfig.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建路由
	r := gin.Default()

	// 设置中间件
	setupMiddleware(r)

	// 设置路由
	routes.SetupRoutes(r)

	// 启动服务器
	port := config.AppConfig.Port
	if port == "" {
		port = "8080"
	}

	log.Printf("服务器启动在端口 %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}

func setupMiddleware(r *gin.Engine) {
	// CORS中间件
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// 日志中间件
	r.Use(gin.Logger())

	// 恢复中间件
	r.Use(gin.Recovery())
}