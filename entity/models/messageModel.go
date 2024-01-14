package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Content string `json:"content" gorm:"type:varchar(500);not null"`
	Seen    bool   `json:"seen" gorm:"type:bit;not null;default:false"`
	// Relations
	UserID string `json:"userId" gorm:"type:varchar(50);not null"`
	User   *User  `json:"userDtos" gorm:"foreignKey:UserID;references:ID"`
}
