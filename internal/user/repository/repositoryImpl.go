package repository

import (
	"context"
	"cryptoChallenges/internal/user/entity"
	"cryptoChallenges/pkg/errors"
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

func (ur *userRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	const op errors.Op = "user/createUser"
	err := ur.gormDB.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, errors.E(op, errors.Internal, err)
	}
	return user, nil
}

func (ur *userRepository) ReadUser(ctx context.Context, id uuid.UUID) (entity.User, error) {
	return entity.User{}, nil
}

func (ur *userRepository) UpdateUser(ctx context.Context, user *entity.User) (entity.User, error) {

	return entity.User{}, nil
}

func (ur *userRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return nil
}
