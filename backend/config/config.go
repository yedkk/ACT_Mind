package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Environment string `mapstructure:"ENVIRONMENT"`
	Port        string `mapstructure:"PORT"`
	DBPath      string `mapstructure:"DB_PATH"`
	JWTSecret   string `mapstructure:"JWT_SECRET"`
	LogLevel    string `mapstructure:"LOG_LEVEL"`
}

var AppConfig *Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// 设置默认值
	viper.SetDefault("ENVIRONMENT", "development")
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("DB_PATH", "./data/act_mind.db")
	viper.SetDefault("JWT_SECRET", "your-secret-key-change-in-production")
	viper.SetDefault("LOG_LEVEL", "info")

	// 允许从环境变量读取
	viper.AutomaticEnv()

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("配置文件读取失败，使用默认配置: %v", err)
	}

	// 解析配置到结构体
	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatal("配置解析失败:", err)
	}

	log.Printf("配置加载完成: Environment=%s, Port=%s", AppConfig.Environment, AppConfig.Port)
}