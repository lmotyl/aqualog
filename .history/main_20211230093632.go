package main

import (
	"gihub.com/lmotyl/aqualog/config"
	"gihub.com/lmotyl/aqualog/database"
	_ "github.com/go-sql-driver/mysql"
)

func initConfig() {
	config.LoadConfig(".env")
}

func initDbConn() {
	dbConn := config.GetEnv("DB_CONNECTION")
	dbHost := config.GetEnv("DB_HOST")
	dPort  := config.GetEnv("DB_PORT")
	dbUser := config.GetEnv("DB_USERNAME")
	dbDatabase := config.GetEnv("DB_DATABASE")
	dbPassword := config.GetEnv("DB_PASSWORD")
	dbMaxOpenConns := config.GetEnv("DB_MAX_OPEN_CONNS")
	dbMaxIdleConns := config.GetEnv("DB_MAX_IDLE_CONNS")
	dbConnMaxLifeTime := config.GetEnv("DB_CONN_MAX_LIFE_TIME")

	

	// maxOpenConns    int
	// maxIdleConns    int
	// connMaxLifeTime int


	dbConfig := database.CreateDbConfig(
		dbConn,
		dbHost,
		dbPort,
		dbName,
		dbUser,
		dbPassword
	)

}

func main() {
	initConfig()
	initDbConn()

}
