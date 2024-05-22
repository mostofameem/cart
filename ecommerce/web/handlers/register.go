package handlers

import (
	"ecommerce/db"
	"ecommerce/web/utils"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type NewUser struct {
	Name     string `json:"name" validate:"required,min=3,max=20,alpha"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

func Register(w http.ResponseWriter, r *http.Request) {

	var user NewUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	validate := validator.New()

	err = validate.Struct(user)

	if err != nil {
		utils.SendError(w, 404, err, "Validation error")
		return
	}

	err = db.Register(user.Name, user.Email, user.Password)
	if err != nil {
		utils.SendError(w, 404, err, "User Already Exists")
		return
	}
	utils.SendData(w, user)
}
