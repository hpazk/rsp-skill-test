package models

import (
	"time"

	"gorm.io/gorm"
)

// Room struct is...
type Room struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	RoomName    string         `json:"room_name"`
	RoomCapcity string         `json:"room_capacity"`
	Photo       string         `json:"photo"`
	CreatedAt   time.Time      `json:"created_at"  gorm:"time"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"time"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"time"`
}
