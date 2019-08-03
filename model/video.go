package model

import "github.com/jinzhu/gorm"

// User 用户模型
type Video struct {
	gorm.Model
	Title string
	Info  string
}
