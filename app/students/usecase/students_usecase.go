package usecase

import (
	"encoding/json"
	"github.com/gocraft/dbr/v2"
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang/domain"
	"net/http"
)

type StudentsUsecase struct {
	mysql        *dbr.Session
	StudentsRepo domain.StudentRepository
	NSQProducer  *nsq.Producer
}

func NewStudentsUsecase(sess *dbr.Session, sr domain.StudentRepository, p *nsq.Producer) domain.StudentsUsecase {
	usecase := &StudentsUsecase{
		mysql:        sess,
		StudentsRepo: sr,
		NSQProducer:  p,
	}
	return usecase
}

func (s *StudentsUsecase) GetStudents() ([]domain.Students, error) {
	data := s.StudentsRepo.GetStudents(s.mysql)
	//fmt.Println(StudentsData)
	return data, nil
}

func (s *StudentsUsecase) CreateStudents(payload domain.Students, w http.ResponseWriter) (interface{}, error) {
	err := s.StudentsRepo.CreateStudents(s.mysql, payload)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	//Publish Message To NSQ
	data := map[string]interface{}{
		"nisn":  payload.NISN,
		"name":  payload.Name,
		"email": payload.Email,
	}
	NSQPayload, err := json.Marshal(data)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	s.NSQProducer.Publish(viper.GetString("NSQ_TOPIC"), NSQPayload)
	logrus.Info("Success To Publish")

	return nil, nil
}

func (s *StudentsUsecase) UpdateStudents(id int, payload domain.Students) (interface{}, error) {
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
