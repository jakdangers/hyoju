package dto

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}
