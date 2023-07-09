package repository

import (
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type userRepository struct {
	gorm *gorm.DB
	sqlx *sqlx.DB
}

func New(gorm *gorm.DB, sqlx *sqlx.DB) *userRepository {
	return &userRepository{
		gorm: gorm,
		sqlx: sqlx,
	}
}

func (u *userRepository) GetUsers() (string, error) {
	//TODO implement me
	panic("implement me")
}
