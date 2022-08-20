package repository

import (
	"fmt"
	"github.com/gocraft/dbr/v2"
	"github.com/sirupsen/logrus"
	"golang/domain"
)

type StudentsRepo struct {
}

var (
	TableName = "students"
)

func NewStudentRepo() domain.StudentRepository {
	return &StudentsRepo{}
}

func (sr StudentsRepo) GetStudents(mysql *dbr.Session) []domain.Students {
	var data []domain.Students
	_, err := mysql.Select("*").From(TableName).Load(&data)
	if err != nil {
		logrus.Error(err)
	}

	return data
}

func (sr StudentsRepo) CreateStudents(mysql *dbr.Session, payload domain.StudentsPayload) error {
	_, err := mysql.InsertInto(TableName).Columns("name", "nisn").Record(payload).Exec()
	if err != nil {
		return err
	}

	return nil
}

func (sr StudentsRepo) UpdateStudents(mysql *dbr.Session, id int, payload domain.StudentsPayload) error {
	_, err := mysql.Update(TableName).Set("name", payload.Name).Set("nisn", payload.NISN).Where("id=?", id).Exec()
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (sr StudentsRepo) GetStudentsById(mysql *dbr.Session, id int) (interface{}, error) {
	var data domain.Students
	err := mysql.Select("*").Where("id=?", id).From(TableName).LoadOne(&data)
	fmt.Println(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}
