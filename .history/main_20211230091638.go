package main

import (
	"gihub.com/lmotyl/aqualog/config"
	"gihub.com/lmotyl/aqualog/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config.LoadConfig(".env")
	dbConfig := database.CreateDbConfig(
		config.GetEnv()
	)

}
