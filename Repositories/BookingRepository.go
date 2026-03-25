package Repositories

import (
	"9263-solution/Models"
	"database/sql"
	"fmt"

	"log"
	"os"
	"strconv"
)

type CreateBookingDto struct {
	PlaceId string
	UserId  string
	From    string
	To      string
}

var db = getConnection()

func GetBookingsByUserId(userId uint64) ([]Models.Booking, error) {
	rows, err := db.Query("select * from Bookings where user_id = ?", userId)

	defer rows.Close()

	var bookings []Models.Booking
	for rows.Next() {
		var booking Models.Booking
		err := rows.Scan(&booking.Id, &booking.UserId, &booking.PlaceId, &booking.TimeFrom, &booking.TimeTo)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}

	return bookings, err
}

func GetBookingsByPlaceId(placeId uint64) ([]Models.Booking, error) {
	rows, err := db.Query("select * from Bookings where place_id = ?", placeId)

	defer rows.Close()

	var bookings []Models.Booking
	for rows.Next() {
		var booking Models.Booking
		err := rows.Scan(&booking.Id, &booking.UserId, &booking.PlaceId, &booking.TimeFrom, &booking.TimeTo)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}

	return bookings, err
}

func CreateBooking() {

}

func getConnection() *sql.DB {
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	credentials := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	db, err := sql.Open("postgres", credentials)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
