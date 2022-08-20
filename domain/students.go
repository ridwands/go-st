package domain

import (
	"github.com/gocraft/dbr/v2"
	"net/http"
)

type Students struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	NISN string `json:"nisn"`
}

type StudentsPayload struct {
	Name string `json:"name" validate:"required"`
	NISN string `json:"nisn" validate:"required,gt=3"`
	ID   int    `json:"id"`
}

type StudentsUsecase interface {
	GetStudents() ([]Students, error)
	CreateStudents(payload StudentsPayload, w http.ResponseWriter) (interface{}, error)
	UpdateStudents(id int, payload StudentsPayload) (interface{}, error)
}

type StudentRepository interface {
	GetStudents(mysql *dbr.Session) []Students
	CreateStudents(mysql *dbr.Session, payload StudentsPayload) error
	UpdateStudents(mysql *dbr.Session, id int, payload StudentsPayload) error
	GetStudentsById(mysql *dbr.Session, id int) (res interface{}, err error)
}
