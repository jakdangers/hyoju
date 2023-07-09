package repository

type UserRepository interface {
	GetUsers() (string, error)
}
