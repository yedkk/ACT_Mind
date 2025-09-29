package routes

import (
	"act-mind-backend/controllers"
	"act-mind-backend/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine) {
	// Swagger文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 健康检查
	r.GET("/health", controllers.HealthCheck)

	// API v1 路由组
	v1 := r.Group("/api/v1")
	{
		// 认证相关路由（无需JWT验证）
		auth := v1.Group("/auth")
		{
			auth.POST("/login", controllers.Login)
			auth.POST("/register", controllers.Register)
		}

		// 需要JWT验证的路由
		protected := v1.Group("/")
		protected.Use(middleware.JWTAuth())
		{
			// 用户相关
			users := protected.Group("/users")
			{
				users.GET("/profile", controllers.GetUserProfile)
				users.PUT("/profile", controllers.UpdateUserProfile)
			}

			// 帖子相关
			posts := protected.Group("/posts")
			{
				posts.GET("/", controllers.GetPosts)
				posts.POST("/", controllers.CreatePost)
				posts.GET("/:id", controllers.GetPost)
				posts.PUT("/:id", controllers.UpdatePost)
				posts.DELETE("/:id", controllers.DeletePost)
			}

			// 评论相关
			comments := protected.Group("/comments")
			{
				comments.POST("/", controllers.CreateComment)
				comments.GET("/post/:post_id", controllers.GetCommentsByPost)
				comments.DELETE("/:id", controllers.DeleteComment)
			}
		}
	}
}