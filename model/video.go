package model

import "github.com/jinzhu/gorm"

// User 视频模型
type Video struct {
	gorm.Model
	Title  string
	Info   string
	UserID uint `gorm:"column:user_id"`
}
