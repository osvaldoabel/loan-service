package entities

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

//Book  - related to book data.
type Book struct {
	ID        string    `valid:"uuid" gorm:"type:uuid;primary_key"`
	UserID    string    `valid:"uuid" gorm:"type:uuid"`
	Title     string    `valid:"notnull"`
	Pages     int       `valid:"notnull"`
	Status    string    `valid:"notnull"`
	CreatedAt time.Time `valid:"-"`
	UpdatedAt time.Time `valid:"-"`
}

func CreateBook(userID string, title string, status string, pages int) (*Book, error) {
	b := &Book{
		ID:        GenerateNewUuID(),
		UserID:    userID,
		Title:     title,
		Pages:     pages,
		Status:    status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := b.Validate()
	if err != nil {
		fmt.Println(err)
		return nil, ErrInvalidItem
	}

	return b, nil
}

func (b *Book) Validate() error {

	_, err := govalidator.ValidateStruct(b)

	if err != nil {
		return err
	}

	return nil
}
