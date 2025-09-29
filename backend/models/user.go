package models

import (
	"time"

	"gorm.io/gorm"
)

// User 用户基础信息
type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	OpenID    string         `json:"openid" gorm:"uniqueIndex;size:100;not null"`
	Nickname  string         `json:"nickname" gorm:"size:50"`
	AvatarURL string         `json:"avatar_url" gorm:"size:255"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// 关联关系
	Profile  UserProfile `json:"profile,omitempty" gorm:"foreignKey:UserID"`
	Posts    []Post      `json:"posts,omitempty" gorm:"foreignKey:UserID"`
	Comments []Comment   `json:"comments,omitempty" gorm:"foreignKey:UserID"`
}

// UserProfile 用户详细档案
type UserProfile struct {
	ID                uint   `json:"id" gorm:"primaryKey"`
	UserID            uint   `json:"user_id" gorm:"not null;index"`
	Bio               string `json:"bio" gorm:"type:text"`
	PsychologicalData string `json:"psychological_data" gorm:"type:json"` // 存储心理测评数据
	PrivacySettings   string `json:"privacy_settings" gorm:"type:json"`   // 隐私设置
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`

	// 关联关系
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

func (UserProfile) TableName() string {
	return "user_profiles"
}