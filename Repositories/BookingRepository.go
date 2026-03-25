package Repositories

import (
	"9263-solution/Models"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
	"time"
)

type CreateBookingDto struct {
	PlaceId string `json:"place_id"`
	UserId  string `json:"user_id"`
	From    string `json:"from"`
	To      string `json:"to"`
}

var db = getConnection()

func GetBookingsByUserId(userId uint64) ([]Models.Booking, error) {
	rows, err := db.Query("select * from bookings where user_id = $1", userId)

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
	rows, err := db.Query("select * from bookings where place_id = $1", placeId)

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

func CreateBooking(d *CreateBookingDto) (err error) {
	from, _ := time.Parse("RFC3339", d.From)
	to, _ := time.Parse("RFC3339", d.To)

	existingRows, _ := db.Query("select * from bookings where (time_to >= $1 and time_to <= $2) or (time_from <= $2 and time_from>= $1)", from.GoString(), to.GoString())

	if existingRows != nil {
		return errors.New("err_exists")
	}

	_, err = db.Query(
		"insert into bookings (user_id, place_id, time_from, time_to) values (?, ?, ?, ?)",
		d.UserId, d.PlaceId, d.From, d.To)
	if err != nil {
		return err
	}
	return nil
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
