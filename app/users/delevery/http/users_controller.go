package UsersHttp

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"golang/domain"
	"golang/pkg/helpers/responder"
	"net/http"
)

func (c UsersHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload domain.Users
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		logrus.Error(err)
		responder.Error(w, 500, err)
		return
	}

	//Validator
	err = c.validator.Validate(payload)
	if err != nil {
		m := c.validator.Message(err)
		logrus.Error(m)
		responder.Error(w, 422, m)
		return
	}

	login, err := c.UsersUseCase.Login(w, payload)
	if err != nil {
		return
	}
	data := map[string]interface{}{
		"token": login,
	}
	responder.Success(w, data, "Success")
}
