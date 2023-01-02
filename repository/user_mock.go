package repository

import "errors"

type userRepositoryMock struct {
	users []User
}

func NewUserRepositoryMock() userRepositoryMock {
	users := []User{
		{UserID: 2, FirstName: "test", LastName: "test", Username: "test", Password: "test"},
		{UserID: 3, FirstName: "test", LastName: "test", Username: "test", Password: "test"},
	}

	return userRepositoryMock{users: users}
}

func (r userRepositoryMock) GetAll() ([]User, error) {
	return r.users, nil
}

func (r userRepositoryMock) GetById(id int) (*User, error) {
	for _, user := range r.users {
		if user.UserID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}
