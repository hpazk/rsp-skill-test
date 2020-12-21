package operation

import (
	"github.com/hpazk/rsp-skill-test/app/database"
	"github.com/hpazk/rsp-skill-test/app/models"
)

// BookInsert is...
func BookInsert(book *models.Book) bool {
	res := database.DB.Create(&book)

	if res.Error == nil {
		return true
	}

	return false
}

// // BookList is...
// func BookList() []Book {
// 	books := []Book{}
// 	res := database.DBConn.Find(&books)
// 	if res.Error == nil {
// 		return books
// 	}
// 	return nil
// }

// BookFetchAll is...
func BookFetchAll() []models.Book {
	var books []models.Book
	// Get all records
	res := database.DBConn().Find(&books)
	if res.Error == nil {
		return books
	}
	return nil
}

/*
===================
ROOM QUERYING
===================
*/

// SelectAllRooms is...
func SelectAllRooms() []models.Room {
	var rooms []models.Room
	// Get all records
	result := database.DB.Find(&rooms)
	if result.Error == nil {
		return rooms
	}
	return nil
}

// SelectRoom is...
func SelectRoom(id int) *models.Room {
	var room = new(models.Room)
	result := database.DB.First(&room, id)
	if result.Error == nil {
		return room
	}
	return nil
}

// InsertRoom is...
func InsertRoom(room *models.Room) bool {
	result := database.DB.Create(&room)
	if result.Error == nil {
		return true
	}

	return false
}

/*
===================
BOOKING QUERYING
===================
*/

// SelectAllBookings is...
func SelectAllBookings() []models.Booking {
	var bookings []models.Booking
	// Get all records
	result := database.DB.Find(&bookings)
	if result.Error == nil {
		return bookings
	}
	return nil
}

// InsertBooking is...
func InsertBooking(booking *models.Booking) bool {
	result := database.DB.Create(&booking)
	if result.Error == nil {
		return true
	}

	return false
}
