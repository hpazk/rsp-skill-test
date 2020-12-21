package models

import (
	"time"

	"gorm.io/gorm"
)

// Booking struct is...
type Booking struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id"`
	// User         User           `gorm:"foreignKey:UserID"`
	RoomID uint `json:"room_id"`
	// Room         Room           `gorm:"foreignKey:RoomID"`
	TotalPerson  int            `json:"total_person"`
	BookingTime  time.Time      `json:"booking_time"`
	Noted        string         `json:"noted"`
	CheckInTime  time.Time      `json:"check_in_time"`
	CheckOutTime time.Time      `json:"check_out_time"`
	CreatedAt    time.Time      `json:"created_at"  gorm:"time"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"time"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"time"`
}
