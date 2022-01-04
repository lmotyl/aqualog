package user

type User struct {
	UserID  int
	name    string
	created int
}

func CreateUser(id int, name string, created int) User {
	// dbConfig := &DbConfig{
	// 	sqlConnection:   dbConnection,
	// 	connString:      fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUserName, dbPassword, dbHost, dbPort, dbName),
	// 	maxOpenConns:    dbMaxOpenConns,
	// 	maxIdleConns:    dbMaxIdleConns,
	// 	connMaxLifeTime: dbConnMaxLifeTime,
	// }

	// return dbConfig
	user := User{
		UserID:  id,
		name:    name,
		created: created,
	}

	return user
}
