package repository

type User struct {
	UserID    int    `db:"user_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Username  string `db:"username"`
	Password  string `db:"password"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetById(int) (*User, error)
}
