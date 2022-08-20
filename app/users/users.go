package users

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gocraft/dbr/v2"
	"github.com/gorilla/mux"
	UsersHttp "golang/app/users/delevery/http"
	"golang/app/users/repository"
	"golang/app/users/usecase"
)

func InitUsers(router *mux.Router, sess *dbr.Session, validator *validator.Validate) {
	fmt.Println(sess)
	r := repository.NewUsersRepo()
	u := usecase.NewUsersUsecase(sess, r)
	UsersHttp.NewUsersHandler(router, validator, u)
	return
}
