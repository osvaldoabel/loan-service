package entities_test

import (
	"testing"
	"time"

	"osvaldoabel/smartmei-service/src/entities"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	user, err := entities.CreateUser("Osvaldo Abel", "osvaldo.abel@example.com", "123", "active", time.Now())
	require.Nil(t, err)
	require.Equal(t, "Osvaldo Abel", user.Name)
}

func TestCreateUserWithInvalidEmail(t *testing.T) {
	_, err := entities.CreateUser("Osvaldo Abel", "kjfhskjdfskjdf", "123", "active", time.Now())
	require.NotNil(t, err)
}

func TestCreateUserWithDuplication(t *testing.T) {
	user, err := entities.CreateUser("Osvaldo Abel", "osvaldo.abel@example.com", "123", "active", time.Now())
	require.Nil(t, err)
	require.Equal(t, "Osvaldo Abel", user.Name)

	_, err = entities.CreateUser("Osvaldo Abel", "osvaldo.abelexample.com", "123", "active", time.Now())
	require.NotNil(t, err)
}
