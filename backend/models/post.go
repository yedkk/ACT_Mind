package models

import (
	"time"

	"gorm.io/gorm"
)

// Post 帖子模型
type Post struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"user_id" gorm:"not null;index"`
	Title     string         `json:"title" gorm:"size:200;not null"`
	Content   string         `json:"content" gorm:"type:text;not null"`
	Category  string         `json:"category" gorm:"size:50;index"`
	ViewCount int            `json:"view_count" gorm:"default:0"`
	LikeCount int            `json:"like_count" gorm:"default:0"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// 关联关系
	User     User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Comments []Comment `json:"comments,omitempty" gorm:"foreignKey:PostID"`
}

// Comment 评论模型
type Comment struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	PostID    uint           `json:"post_id" gorm:"not null;index"`
	UserID    uint           `json:"user_id" gorm:"not null;index"`
	Content   string         `json:"content" gorm:"type:text;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// 关联关系
	Post Post `json:"post,omitempty" gorm:"foreignKey:PostID"`
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// TableName 指定表名
func (Post) TableName() string {
	return "posts"
}

func (Comment) TableName() string {
	return "comments"
}