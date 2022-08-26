package domain

import (
	"github.com/gocraft/dbr/v2"
	"net/http"
)

type Students struct {
	Name  string `json:"name" validate:"required"`
	NISN  string `json:"nisn" validate:"required,gt=3"`
	ID    int    `json:"id"`
	Email string `json:"email" validate:"required,email"`
}

type StudentsUsecase interface {
	GetStudents() ([]Students, error)
	CreateStudents(payload Students, w http.ResponseWriter) (interface{}, error)
	UpdateStudents(id int, payload Students) (interface{}, error)
}

type StudentRepository interface {
	GetStudents(mysql *dbr.Session) []Students
	CreateStudents(mysql *dbr.Session, payload Students) error
	UpdateStudents(mysql *dbr.Session, id int, payload Students) error
	GetStudentsById(mysql *dbr.Session, id int) (res interface{}, err error)
}
