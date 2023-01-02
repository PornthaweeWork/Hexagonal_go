package service

type UserResponse struct {
	UserID    int    `json:"user_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

type UserService interface {
	GetUsers() ([]UserResponse, error)
	GetUser(int) (*UserResponse, error)
}
