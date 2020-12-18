package models

import "github.com/jinzhu/gorm"

// User struct is...
type User struct {
	gorm.Model
	Email     string `json:"email" form:"email" query:"email"`
	Password  string `json:"password" form:"password" query:"password"`
	Photo     string `json:"photo" form:"photo" query:"photo"`
	CreatedAt string `json:"created_at" form:"created_at" query:"created_at"`
	UpdatedAt string `json:"updated_at" form:"updated_at" query:"updated_at"`
	DeletedAt string `json:"deleted_at" form:"deleted_at" query:"deleted_at"`
}

// // Room struct is...
// type Room struct {
// 	RoomName     string `json:"room_name" form:"room_name" query:"room_name"`
// 	RoomCapacity string `json:"room_capacity" form:"room_capacity" query:"room_capacity"`
// 	Photo        string `json:"photo" form:"photo" query:"photo"`
// 	CreatedAt    string `json:"created_at" form:"created_at" query:"created_at"`
// 	UpdatedAt    string `json:"updated_at" form:"updated_at" query:"updated_at"`
// 	DeletedAt    string `json:"deleted_at" form:"deleted_at" query:"deleted_at"`
// }

// // Booking struct is...
// type Booking struct {
// 	TotalPerson  string `json:"total_person" form:"total_person" query:"total_person"`
// 	BookingTime  string `json:"booking_time" form:"booking_time" query:"booking_time"`
// 	Noted        string `json:"noted" form:"noted" query:"noted"`
// 	ChekInTime   string `json:"check_in_time" form:"check_in_time" query:"check_in_time"`
// 	CheckOutTime string `json:"check_out_time" form:"check_out_time" query:"check_out_time"`
// 	CreatedAt    string `json:"created_at" form:"created_at" query:"created_at"`
// 	UpdatedAt    string `json:"updated_at" form:"updated_at" query:"updated_at"`
// 	DeletedAt    string `json:"deleted_at" form:"deleted_at" query:"deleted_at"`
// }
