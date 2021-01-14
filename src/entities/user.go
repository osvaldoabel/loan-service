package entities

import (
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//User  - related to user data.
type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	Status    string
	Birthday  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
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
		Birthday:  birthday,
		Status:    status,
		CreatedAt: time.Now(),
	}

	u.Password = pass

	err = u.Validate()
	if err != nil {
		return nil, ErrInvalidItem
	}

	return u, nil
}

func (u *User) Validate() error {
	if u.Name == "" || u.Email == "" || u.Password == "" {
		return ErrInvalidItem
	}

	if !IsValidEmail(u.Email) {
		return ErrNotValidEmail
	}

	return nil
}

// Encrypt - Encript password
func Encrypt(pass string) (string, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		return "", err
	}
	return string(encrypted), nil
}

// IsValidEmail - Verify if email is valid or not.
func IsValidEmail(email string) bool {
	pattern := regexp.MustCompile(`^(\w|\.)+\@(\w)+(.(\w)+){1,2}$`)
	return pattern.MatchString(email)
}
