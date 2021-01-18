package entities

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

//User  - related to user data.
type User struct {
	ID        string    `valid:"uuid" gorm:"type:uuid;primary_key"`
	Name      string    `valid:"notnull"`
	Email     string    `valid:"notnull"`
	Password  string    `valid:"notnull"`
	Status    string    `valid:"notnull"`
	Birthday  time.Time `valid:"-"`
	CreatedAt time.Time `valid:"-"`
	UpdatedAt time.Time `valid:"-"`
}

func CreateUser(name string, email string, password string, status string, birthday time.Time) (*User, error) {
	pass, err := Encrypt(password)

	if err != nil {
		return nil, err
	}

	u := &User{
		ID:        GenerateNewUuID(),
		Name:      name,
		Email:     email,
		Password:  password,
		Birthday:  birthday,
		Status:    status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	u.Password = pass

	err = u.Validate()
	if err != nil {
		fmt.Println(err)
		return nil, ErrInvalidItem
	}

	return u, nil
}

func (u *User) Validate() error {

	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return err
	}

	if !IsValidEmail(u.Email) {
		return ErrNotValidEmail
	}

	return nil
}
