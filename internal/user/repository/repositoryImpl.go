package repository

import (
	"cryptoChallenges/internal/user/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type userRepository struct {
	gormDB *gorm.DB
	sqlxDB *sqlx.DB
}

func New(gormDB *gorm.DB, sqlxDB *sqlx.DB) *userRepository {
	return &userRepository{
		gormDB: gormDB,
		sqlxDB: sqlxDB,
	}
}

func (ur *userRepository) CreateUser(user *entity.User) (*entity.User, error) {
	err := ur.gormDB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) ReadUser(id uuid.UUID) (entity.User, error) {
	return entity.User{}, nil
}

func (ur *userRepository) UpdateUser(user *entity.User) (entity.User, error) {

	return entity.User{}, nil
}

func (ur *userRepository) DeleteUser(id uuid.UUID) error {
	return nil
}
