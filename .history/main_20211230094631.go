package main

import (
	"strconv"

	"gihub.com/lmotyl/aqualog/config"
	"gihub.com/lmotyl/aqualog/database"
	_ "github.com/go-sql-driver/mysql"
)

func initConfig() {
	config.LoadConfig(".env")
}

func initDbConn() {
	dbConn, _ := config.GetEnv("DB_CONNECTION")
	dbHost, _ := config.GetEnv("DB_HOST")
	dbPort, _ := config.GetEnv("DB_PORT")
	dbUser, _ := config.GetEnv("DB_USERNAME")
	dbName, _ := config.GetEnv("DB_DATABASE")
	dbPassword, _ := config.GetEnv("DB_PASSWORD")
	dbMaxOpenConns, isMaxOpenCons := config.GetEnv("DB_MAX_OPEN_CONNS")
	dbMaxIdleConns, isMaxIdleConns := config.GetEnv("DB_MAX_IDLE_CONNS")
	dbConnMaxLifeTime, isConnMaxLifeTime := config.GetEnv("DB_CONN_MAX_LIFE_TIME")

	if isMaxOpenCons == false {
		dbMaxOpenConns = "4"
	}
	if isMaxIdleConns == false {
		dbMaxIdleConns = "4"
	}
	if isConnMaxLifeTime == false {
		dbConnMaxLifeTime = "60"
	}

	dbConfig := database.CreateDbConfig(
		dbConn,
		dbHost,
		dbPort,
		dbName,
		dbUser,
		dbPassword,
		strconv.Atoi(dbMaxOpenConns),
		strconv.Atoi(dbMaxIdleConns),
		strconv.Atoi(dbConnMaxLifeTime))

	database.SetupDatabase(*dbConfig)
}

func main() {
	initConfig()
	initDbConn()

}
