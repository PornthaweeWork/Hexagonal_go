package service

import (
	"bank/repository"
	"database/sql"
	"errors"
	"log"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) userService {
	return userService{userRepo: userRepo}
}

func (u userService) GetUsers() ([]UserResponse, error) {
	users, err := u.userRepo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	userResponses := []UserResponse{}
	for _, user := range users {
		userResponse := UserResponse{
			UserID:    user.UserID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}
		userResponses = append(userResponses, userResponse)
	}

	return userResponses, nil
}

func (u userService) GetUser(id int) (*UserResponse, error) {
	user, err := u.userRepo.GetById(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}

		log.Println(err)
		return nil, err
	}

	userResponse := UserResponse{
		UserID:    user.UserID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	return &userResponse, nil
}
