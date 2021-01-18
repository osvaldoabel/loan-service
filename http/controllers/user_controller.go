package controllers

import (
	"encoding/json"
	"net/http"
	"osvaldoabel/smartmei-service/http/presenters"
	"osvaldoabel/smartmei-service/src/services"
	"osvaldoabel/smartmei-service/utils"
)

type UserController struct {
	UserService services.UserService
}

// NewUserController - Create new UserController Struct
func NewUserController() *UserController {
	return &UserController{UserService: services.NewUserService()}
}

func (u *UserController) Create(w http.ResponseWriter, r *http.Request) {

	payload, err := getPayloads(r, utils.UserPayload{})
	if err != nil {
		utils.JsonResponse(w, nil, 400)
	}

	m := payload.(map[string]interface{})
	p := utils.UserPayload{
		Name:     m["Name"].(string),
		Email:    m["Email"].(string),
		Password: m["Password"].(string),
		Birthday: m["Birthday"].(string),
		Status:   m["Status"].(string),
	}

	user, err := u.UserService.CreateUser(&p)

	if err != nil {
		utils.JsonResponse(w, nil, 500)
		return
	}

	result, err := json.Marshal(presenters.ToArray(user))
	if err != nil {
		utils.JsonResponse(w, nil, 500)
		return
	}

	utils.JsonResponse(w, result, 200)
}
