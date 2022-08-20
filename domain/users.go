package domain

import (
	"github.com/gocraft/dbr/v2"
	"net/http"
)

type Users struct {
	NISN     string `json:"nisn" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UsersUseCase interface {
	Login(w http.ResponseWriter, payload Users) (res interface{}, err error)
}

type UsersRepository interface {
	GetUsersByNISN(mysql *dbr.Session, nisn string) (data Users, err error)
}
