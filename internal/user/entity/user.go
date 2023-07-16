package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `db:"name"`
	Email    string `db:"email"`
	UserID   string `db:"user_id"`
	Password string `db:"password"`
}
