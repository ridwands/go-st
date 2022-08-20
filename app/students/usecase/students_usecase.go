package usecase

import (
	"github.com/gocraft/dbr/v2"
	"github.com/sirupsen/logrus"
	"golang/domain"
	"net/http"
)

type StudentsUsecase struct {
	mysql        *dbr.Session
	StudentsRepo domain.StudentRepository
}

func NewStudentsUsecase(sess *dbr.Session, sr domain.StudentRepository) domain.StudentsUsecase {
	usecase := &StudentsUsecase{
		mysql:        sess,
		StudentsRepo: sr,
	}
	return usecase
}

func (s *StudentsUsecase) GetStudents() ([]domain.Students, error) {
	data := s.StudentsRepo.GetStudents(s.mysql)
	//fmt.Println(StudentsData)
	return data, nil
}

func (s *StudentsUsecase) CreateStudents(payload domain.StudentsPayload, w http.ResponseWriter) (interface{}, error) {
	err := s.StudentsRepo.CreateStudents(s.mysql, payload)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return nil, nil
}

func (s *StudentsUsecase) UpdateStudents(id int, payload domain.StudentsPayload) (interface{}, error) {
	_, err := s.StudentsRepo.GetStudentsById(s.mysql, id)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	err = s.StudentsRepo.UpdateStudents(s.mysql, id, payload)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	return nil, nil
}
