package entity

import "gorm.io/gorm"

type User struct {
	UUID      string `db:"uuid"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
	UserID    int64  `db:"user_id"`
	Password  string `db:"password"`
	gorm.Model
}
