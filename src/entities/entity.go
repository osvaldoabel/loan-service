package entities

import (
	"errors"
	"regexp"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

//GenerateNewUuID Generate new Uuid and parse to string
func GenerateNewUuID() string {
	return uuid.NewV4().String()
}

// IsValidEmail - Verify if email is valid or not.
func IsValidEmail(email string) bool {
	pattern := regexp.MustCompile(`^(\w|\.)+@(\w)+(.(\w)+){1,2}$`)
	return pattern.MatchString(email)
}

func IsUuID(id string) bool {
	_, err := uuid.FromString(id)
	if err != nil {
		return false
	}

	return true
}

// Encrypt - Encript password
func Encrypt(pass string) (string, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		return "", err
	}
	return string(encrypted), nil
}

//General Errors

//ErrNotFoundItem not found
var ErrNotFoundItem = errors.New("Item Not found")

// ErrNotValidEmail - The entered email is not valid.
var ErrNotValidEmail = errors.New("Not a Valid E-mail")

//Book Errors

//ErrBookIsBorrowed cannot be borrowed
var ErrBookIsBorrowed = errors.New("The Book is already borrowed")

//ErrInvalidItem cannot be borrowed
var ErrInvalidItem = errors.New("Invalid Item: The Item you are trying to create is not valid.")
