package user

import (
	"context"
	"cryptoChallenges/entity"
	"cryptoChallenges/pkg/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	gormDB *gorm.DB
}

func NewUserRepository(gormDB *gorm.DB) *userRepository {
	return &userRepository{
		gormDB: gormDB,
	}
}

var _ entity.UserRepository = (*userRepository)(nil)

func (ur *userRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	const op errors.Op = "user/repository/createUser"

	err := ur.gormDB.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, errors.E(op, errors.Internal, err)
	}

	return user, nil
}

func (ur *userRepository) ReadUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	const op errors.Op = "user/repository/readUser"

	result := ur.gormDB.WithContext(ctx).Take(user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.E(op, errors.NotExist, "user does not exists")
	}
	if result.Error != nil {
		return nil, errors.E(op, errors.Internal, result.Error)
	}

	return user, nil
}

func (ur *userRepository) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {

	return &entity.User{}, nil
}

func (ur *userRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return nil
}
