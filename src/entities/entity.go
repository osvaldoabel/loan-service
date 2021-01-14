package entities

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

//GenerateNewUuID Generate new Uuid and parse to string
func GenerateNewUuID() string {
	return uuid.NewV4().String()
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
