package database

import (
	"log"
	"os"
	"path/filepath"

	"act-mind-backend/config"
	"act-mind-backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	var err error

	// 确保数据目录存在
	dbPath := config.AppConfig.DBPath
	dbDir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		log.Fatal("创建数据库目录失败:", err)
	}

	// 配置GORM日志级别
	var logLevel logger.LogLevel
	switch config.AppConfig.LogLevel {
	case "debug":
		logLevel = logger.Info
	case "info":
		logLevel = logger.Warn
	default:
		logLevel = logger.Error
	}

	// 连接数据库
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	log.Printf("数据库连接成功: %s", dbPath)

	// 自动迁移
	if err := autoMigrate(); err != nil {
		log.Fatal("数据库迁移失败:", err)
	}

	log.Println("数据库初始化完成")
}

func autoMigrate() error {
	return DB.AutoMigrate(
		&models.User{},
		&models.UserProfile{},
		&models.Post{},
		&models.Comment{},
	)
}

func GetDB() *gorm.DB {
	return DB
}