package services

import (
	"fmt"
	"osvaldoabel/smartmei-service/src/entities"
	"osvaldoabel/smartmei-service/src/repositories"
	"osvaldoabel/smartmei-service/utils"
	"time"
)

type UserService struct {
	UserRepository repositories.UserRepository
}

func NewUserService() UserService {
	return UserService{UserRepository: repositories.NewUserRepository()}
}

//CreateUser - create new user
func (u *UserService) CreateUser(payload *utils.UserPayload) (*entities.User, error) {

	const (
		layoutISO = "2006-01-02"
	)

	date := "1999-12-31"
	birthday, err := time.Parse(layoutISO, date)

	// birthday, err := time.Parse("yyyy-mm-dd", date)
	if err != nil {
		return nil, err
	}

	user, err := entities.CreateUser(payload.Name, payload.Email, payload.Password, payload.Status, birthday)
	if err != nil {
		return nil, err
	}

	user, err = u.UserRepository.Insert(user)
	fmt.Println("oi")
	utils.Dd(user, false)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) Find(id string) (*entities.User, error) {
	user, err := u.UserRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}
