package students

import (
	"github.com/go-playground/validator/v10"
	"github.com/gocraft/dbr/v2"
	"github.com/gorilla/mux"
	studentsHandler "golang/app/students/delevery/http"
	studentRepo "golang/app/students/repository"
	studentUsecase "golang/app/students/usecase"
)

func InitStudents(router *mux.Router, sess *dbr.Session, v *validator.Validate) {
	sr := studentRepo.NewStudentRepo()
	su := studentUsecase.NewStudentsUsecase(sess, sr)
	studentsHandler.NewStudentsHandler(router, su, v)
	return
}
