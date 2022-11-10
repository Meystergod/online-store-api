package model

import "time"

type User struct {
	Id        uint32     `json:"id"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	FirstName *string    `json:"first_name"`
	LastName  *string    `json:"last_name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	IsAdmin   bool       `json:"is_admin"`
}
