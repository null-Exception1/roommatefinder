package structs

import "time"

type RoomCache struct {
	People []*Person
	Expiry time.Time
}
