package services_test

import (
	"osvaldoabel/smartmei-service/src/entities"
	"osvaldoabel/smartmei-service/src/services"
	"osvaldoabel/smartmei-service/utils"
	"osvaldoabel/smartmei-service/utils/database"
	"time"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserServiceInsert(t *testing.T) {

	user, err := entities.CreateUser("Osvaldo Abel", "osvaldo.abel@example.com", "123", "active", time.Now())
	require.Nil(t, err)
	require.NotNil(t, user)

	db := database.NewDbTest()
	defer db.Close()

	userService := services.NewUserService()

	payload := &utils.UserPayload{
		Name:     "Osvaldo Abel",
		Email:    "teste@example.com",
		Status:   "active",
		Address:  "My  Street , 15-30",
		Birthday: "1992-01-01",
		Password: "123456",
	}

	user, err = userService.CreateUser(payload)
	require.Nil(t, err)
}
