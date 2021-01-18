package repositories_test

import (
	"osvaldoabel/smartmei-service/database"
	"osvaldoabel/smartmei-service/src/entities"
	"osvaldoabel/smartmei-service/src/repositories"

	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestUserRepositoryDbInsert(t *testing.T) {

	db := database.NewDbTest()
	defer db.Close()

	user, err := entities.CreateUser("Osvaldo Abel", "osvaldo.abel@example.com", "123", "active", time.Now())
	require.Nil(t, err)

	repo := repositories.NewUserRepository()
	user, err = repo.Insert(user)

	require.Nil(t, err)
}

func TestUserRepositoryDbFind(t *testing.T) {

	db := database.NewDbTest()
	defer db.Close()

	user, err := entities.CreateUser("Osvaldo Abel", "osvaldo.abel@example.com", "123", "active", time.Now())
	require.Nil(t, err)

	repo := repositories.NewUserRepository()
	user, err = repo.Insert(user)
	require.Nil(t, err)

	_, err = repo.Find(user.ID)
	require.Nil(t, err)
}
