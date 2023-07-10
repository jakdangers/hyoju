package entity

import "cryptoChallenges/pkg/common"

type User struct {
	common.Model
	Name     string `db:"name"`
	Email    string `db:"email"`
	UserID   int64  `db:"user_id"`
	Password string `db:"password"`
}
