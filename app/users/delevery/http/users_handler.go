package UsersHttp

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"golang/domain"
	customVal "golang/pkg/helpers/validator"
)

type UsersHandler struct {
	validator    customVal.CustomValidator
	UsersUseCase domain.UsersUseCase
}

func NewUsersHandler(router *mux.Router, v *validator.Validate, usecase domain.UsersUseCase) {
	handler := &UsersHandler{
		validator:    customVal.NewCustomValidator(v),
		UsersUseCase: usecase,
	}

	v1 := router.PathPrefix("/v1/users").Subrouter()
	v1.HandleFunc("/login", handler.Login)
}
