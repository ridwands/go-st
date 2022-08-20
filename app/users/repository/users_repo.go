package repository

import (
	"github.com/gocraft/dbr/v2"
	"github.com/sirupsen/logrus"
	"golang/domain"
)

type UsersRepo struct {
}

const (
	TableName = "users"
)

func NewUsersRepo() domain.UsersRepository {
	return &UsersRepo{}
}

func (u UsersRepo) GetUsersByNISN(mysql *dbr.Session, nisn string) (domain.Users, error) {
	var data domain.Users
	stmt := mysql.Select("*").From(TableName).Where("nisn =?", nisn)
	err := stmt.LoadOne(&data)
	if err != nil {
		logrus.Error(err)
		return data, err
	}
	return data, nil
}
