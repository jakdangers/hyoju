package service

type UserService interface {
	GetUsers() (string, error)
}
