package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

var DbConn *sql.DB

type DbConfig struct {
	sqlConnection 	string
	connString      string
	maxOpenConns    int
	maxIdleConns    int
	connMaxLifeTime int
}

func CreateDbConfig(
	dbConnection string, 
	dbHost string, 
	dbPort string, 
	dbName string, 
	dbUserName string, 
	dbPassword string,
	dbMaxOpenConns int,
	dbMaxIdleConns int,
	dbConnMaxLifeTime int
	) DbConfig {
		
	dbConfig := DbConfig{
		sqlConnection: dbConnection,
		connString: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUserName, dbPassword, dbHost, dbPort, dbName),
		maxOpenConns: dbMaxOpenConns,
		maxIdleConns: dbMaxIdleConns,
		connMaxLifeTime: dbConnMaxLifeTime,
	}	
		

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
