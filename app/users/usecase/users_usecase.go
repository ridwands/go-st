package usecase

import (
	"github.com/gocraft/dbr/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"golang/domain"
	"golang/pkg/helpers/responder"
	"net/http"
)

type UsersUseCase struct {
	mysql     *dbr.Session
	UsersRepo domain.UsersRepository
}

func NewUsersUsecase(sess *dbr.Session, usersRepo domain.UsersRepository) domain.UsersUseCase {
	usecase := &UsersUseCase{
		mysql:     sess,
		UsersRepo: usersRepo,
	}
	return usecase
}

func (u UsersUseCase) Login(w http.ResponseWriter, payload domain.Users) (res interface{}, err error) {
	data, err := u.UsersRepo.GetUsersByNISN(u.mysql, payload.NISN)
	if err != nil {
		if dbr.ErrNotFound == err {
			responder.Error(w, 400, "Wrong Login")
			return nil, err
		}
		responder.Error(w, 500, err)
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(payload.Password))
	if err != nil {
		responder.Error(w, 400, "Wrong Login")
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nisn": data.NISN,
	})
	tokenString, err := token.SignedString([]byte(viper.GetString("JWT_SECRET_KET")))
	if err != nil {
		return nil, err
	}

	return tokenString, err
}
