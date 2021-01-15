package entities_test

import (
	"testing"
	"time"

	"osvaldoabel/smartmei-service/src/entities"

	"github.com/stretchr/testify/require"
)

func TestCreateBook(t *testing.T) {
	user, err := entities.CreateUser("Osvaldo Abel", "osvaldo.abel@example.com", "123", "active", time.Now())
	require.Nil(t, err)

	book, err := entities.CreateBook(user.ID, "Book 1", "available", 200)
	require.Nil(t, err)
	require.Equal(t, book.Title, "Book 1")
}

func TestCreateBookWithInvalidUserID(t *testing.T) {
	_, err := entities.CreateBook("213123", "Book 1", "available", 200)
	require.NotNil(t, err)
}
