package model

import "time"

type User struct {
	ID           int
	Name         string
	Email        string
	Password     string
	User_role    string
	Update_date  time.Time
	Created_date time.Time
}
