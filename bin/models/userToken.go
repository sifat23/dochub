package models

import "time"

type UserToken struct {
	Email     string     `db:"email"`
	Token     string     `db:"token"`
	ValidTill *time.Time `db:"valid_till"`
}
