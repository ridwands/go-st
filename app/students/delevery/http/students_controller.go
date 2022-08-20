package StudentsHttp

import (
	"encoding/json"
	"github.com/gocraft/dbr/v2"
	"github.com/sirupsen/logrus"
	"golang/domain"
	"golang/pkg/helpers/responder"
	"net/http"
)

func (h StudentsHandler) GetStudents(res http.ResponseWriter, req *http.Request) {
	s, err := h.StudentsUsecase.GetStudents()
	if err != nil {
		logrus.Error(err)
	}
	responder.Success(res, s, "Success")
}

func (h StudentsHandler) CreateStudents(res http.ResponseWriter, req *http.Request) {
	var payload domain.StudentsPayload
	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		responder.Error(res, 400, err)
		return
	}

	err = h.validator.Validate(payload)
	if err != nil {
		m := h.validator.Message(err)
		logrus.Error(m)
		responder.Error(res, 422, m)
		return
	}

	_, err = h.StudentsUsecase.CreateStudents(payload, res)
	if err != nil {
		responder.Error(res, 500, err.Error())
		return
	}
	responder.Success(res, payload, "success")
}

func (h StudentsHandler) UpdateStudents(res http.ResponseWriter, req *http.Request) {
	var payload domain.StudentsPayload
	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		responder.Error(res, 400, err)
		return
	}
	_, err = h.StudentsUsecase.UpdateStudents(payload.ID, payload)
	if err != nil {
		if err == dbr.ErrNotFound {
			responder.Error(res, 400, "Data Not Found")
			return
		}
		responder.Error(res, 500, err.Error())
		return
	}
	responder.Success(res, payload, "Success")

}
