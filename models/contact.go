package models

import "time"
type Contact struct {
	ID int
	Nama string
	Email string
	Phone string
	Created_At time.Time
	Updated_at time.Time
}