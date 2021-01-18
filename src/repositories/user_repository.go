package repositories

import (
	"fmt"

	"osvaldoabel/smartmei-service/database"
	"osvaldoabel/smartmei-service/src/entities"
	"osvaldoabel/smartmei-service/utils"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

type UserRepository interface {
	Insert(user *entities.User) (*entities.User, error)
	Find(id string) (*entities.User, error)
}

type UserRepositoryDb struct {
	Db *gorm.DB
}

func init() {
	godotenv.Load("./../../.env")
}

// Description: "Instanceate" a new and clean UserRepositoryDb. It receives the Entity User (*entities.User) and returns
func NewUserRepository() *UserRepositoryDb {
	db := database.NewDbConnection()

	return &UserRepositoryDb{Db: db}
}

// Description: Create a new User registry. It receives the Entity User (*entities.User) and returns
func (r *UserRepositoryDb) Insert(user *entities.User) (*entities.User, error) {
	message := ""
	err := r.Db.Create(user).Error

	if err != nil {
		message = fmt.Sprintf("Coulden't  Insert User. [Name: %s].", user.Name)
		return nil, err
	}

	message = fmt.Sprintf("User Inserted. [ID: %s].", user.ID)
	utils.App_log(message)
	return user, nil
}

func (r *UserRepositoryDb) Find(id string) (*entities.User, error) {
	var user entities.User

	r.Db.Find(&user, "id=?", id)

	if user.ID == "" {
		return nil, entities.ErrNotFoundItem
	}

	return &user, nil
}
