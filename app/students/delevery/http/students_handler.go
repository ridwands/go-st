package StudentsHttp

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"golang/domain"
	"golang/pkg/middleware"
	customVal "golang/pkg/validator"
)

type StudentsHandler struct {
	StudentsUsecase domain.StudentsUsecase
	validator       customVal.CustomValidator
}

func NewStudentsHandler(router *mux.Router, su domain.StudentsUsecase, v *validator.Validate) {
	handler := &StudentsHandler{
		StudentsUsecase: su,
		validator:       customVal.NewCustomValidator(v),
	}

	v1 := router.PathPrefix("/v1/students").Subrouter()
	v1.Use(middleware.AuthMiddleware)
	v1.HandleFunc("/list", handler.GetStudents).Methods("GET")
	v1.HandleFunc("/create", handler.CreateStudents).Methods("POST")
	v1.HandleFunc("/update", handler.UpdateStudents).Methods("PUT")
}
