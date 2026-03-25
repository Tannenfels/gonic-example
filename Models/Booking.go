package Models

import "time"

type Booking struct {
	Id       uint64    `json:"id"`
	UserId   uint64    `json:"user_id"`
	PlaceId  uint64    `json:"place_id"`
	TimeFrom time.Time `json:"time_from"`
	TimeTo   time.Time `json:"time_to"`
}
