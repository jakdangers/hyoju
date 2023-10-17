package user

import (
	"context"
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

func (ur *userRepository) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	const op cerrors.Op = "user/repository/updateUser"

	result := ur.gormDB.WithContext(ctx).Save(user)
	if result.Error != nil {
		return nil, cerrors.E(op, cerrors.Internal, result.Error)
	}

	return user, nil
}

func (ur *userRepository) DeleteUser(ctx context.Context, id entity.BinaryUUID) error {
	const op cerrors.Op = "user/repository/deleteUser"

	result := ur.gormDB.WithContext(ctx).Delete(&entity.User{}, id)
	if result.Error != nil {
		return cerrors.E(op, cerrors.Internal, result.Error)
	}

	return nil
}

func (ur *userRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User

	const op cerrors.Op = "user/repository/findByEmail"

	result := ur.gormDB.WithContext(ctx).Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, cerrors.E(op, cerrors.Internal, result.Error)
	}

	return &user, nil
}

func (ur *userRepository) FindByID(ctx context.Context, id entity.BinaryUUID) (*entity.User, error) {
	const op cerrors.Op = "user/repository/readUser"

	var user entity.User
	result := ur.gormDB.WithContext(ctx).Where("id = ?", id).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, cerrors.E(op, cerrors.Internal, result.Error)
	}

	return &user, nil
}

func (ur *userRepository) FindByFriendCode(ctx context.Context, friendCode string) (*entity.User, error) {
	const op cerrors.Op = "user/repository/findByFriendCode"

	var user entity.User
	result := ur.gormDB.WithContext(ctx).Where("friend_code = ?", friendCode).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, cerrors.E(op, cerrors.Internal, result.Error)
	}

	return &user, nil
}
