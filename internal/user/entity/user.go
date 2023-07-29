package entity

import "cryptoChallenges/pkg/models"

type User struct {
	models.Base
	Name     string `db:"name"`
	Email    string `db:"email"`
	UserID   string `db:"user_id"`
	Password string `db:"password"`
}
