package dto

import (
	"github.com/google/uuid"
	"pixelix/pkg/cerrors"
)

type UserDTO struct {
	ID     uuid.UUID `json:"ID"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	UserID string    `json:"userID"`
}

type CreateUserRequest struct {
	Email       string `json:"email"`
	FirebaseUID string `json:"firebaseUID"`
	Provider    string `json:"provider"`
	NickName    string `json:"nickName"`
}

type CreateUserResponse struct {
	ID          string `json:"ID"`
	Email       string `json:"email"`
	FirebaseUID string `json:"firebaseUID"`
	Provider    string `json:"provider"`
	NickName    string `json:"nickName"`
}

type ReadUserRequest struct {
	ID string `form:"ID"`
}

type ReadUserResponse struct {
	ID       string `form:"ID"`
	Email    string `json:"email"`
	NickName string `json:"nickName"`
}

type UpdateUserRequest struct {
	ID       string `json:"ID"`
	NickName string `json:"nickName"`
}

func (ur UpdateUserRequest) Valid() error {
	const op cerrors.Op = "/user/controller/valid"

	if ur.NickName == "" {
		return cerrors.E(op, cerrors.Invalid, "invalid input")
	}

	return nil
}

type UpdateUserResponse struct {
	ID       string `json:"ID"`
	Email    string `json:"email"`
	NickName string `json:"nickName"`
}

type DeleteUserRequest struct {
	ID string `json:"ID" uri:"ID"`
}

type OAuthLoginUserRequest struct {
	Email       string `json:"email"`
	FirebaseUID string `json:"firebaseUID"`
	Provider    string `json:"provider"`
}

type OAuthLoginUserResponse struct {
	AccessToken string `json:"accessToken"`
	// TODO refreshToken 추가
}
