package presenters

import (
	"osvaldoabel/smartmei-service/src/entities"
)

//ToCollection - Receives a slice of User formats it as necessary and returns a Slice of UsesrPayloads
func ToCollection(items []*entities.User) []*UserCollection {
	var userCollection []*UserCollection

	for _, item := range items {
		userCollection = append(userCollection, ToArray(item))
	}

	return userCollection
}

func ToArray(item *entities.User) *UserCollection {

	return &UserCollection{
		ID:        item.ID,
		Name:      item.Email,
		Email:     item.Email,
		Status:    item.Status,
		Birthday:  item.Birthday.Format("dd-mm-yyyy"),
		CreatedAt: item.CreatedAt.Format("dd-mm-yyyy"),
		UpdatedAt: item.UpdatedAt.Format("dd-mm-yyyy"),
	}
}
