package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"pixelix/entity"
	"pixelix/pkg/cerrors"
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
	const op cerrors.Op = "user/repository/createUser"

	err := ur.gormDB.WithContext(ctx).Create(user).Error
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err)
	}

	return user, nil
}

func (ur *userRepository) ReadUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	const op cerrors.Op = "user/repository/readUser"

	result := ur.gormDB.WithContext(ctx).Take(user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, cerrors.E(op, cerrors.Internal, result.Error)
	}

	return user, nil
}

func (ur *userRepository) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	const op cerrors.Op = "user/repository/updateUser"

	ur.gormDB.Save(user)

	return user, nil
}

func (ur *userRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return nil
}
