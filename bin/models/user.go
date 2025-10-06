package models

import "time"

type User struct {
	Id        int64
	Name      string
	Email     string
	Password  string
	Status    int8
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
