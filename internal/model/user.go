package model

import "time"

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	User_role string
	UpdatedAt time.Time
	CreatedAt time.Time
}
