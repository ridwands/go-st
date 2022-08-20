package database

import (
	dbr "github.com/gocraft/dbr/v2"
)

func CreateConnection(driver, descriptor string) (connection *dbr.Connection) {
	db, err := dbr.Open(driver, descriptor, nil)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(50)

	return db
}
