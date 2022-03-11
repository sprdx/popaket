package users

import (
	"github.com/go-playground/validator/v10"
)

func (data *Domain) ValidateUserData() string {
	validate := validator.New()
	err := validate.Var(data.Username, "required,min=3,max=35,startsnotwith= ,endsnotwith= ")
	if err != nil {
		return "Invalid name"
	}
	err = validate.Var(data.Password, "required,min=8")
	if err != nil {
		return "Invalid password"
	}
	return "VALID"
}
