package database

import (
	"database/sql"
	"log"
	"time"
)

var DbConn *sql.DB

type DbConfig struct {
	connString      string
	maxOpenConns    int
	maxIdleConns    int
	connMaxLifeTime int
}

func SetupDatabase(dbConfig *DbConfig) {
	var err error

	DbConn, err = sql.Open("mysql", dbConfig.connString)
	if err != nil {
		log.Fatal(err)
	}

	DbConn.SetMaxOpenConns(dbConfig.maxOpenConns)
	DbConn.SetMaxIdleConns(dbConfig.maxIdleConns)
	DbConn.SetConnMaxLifetime(time.Duration(dbConfig.connMaxLifeTime) * time.Second)
}
