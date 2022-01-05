package database

import (
	"database/sql"
	"technical-test-aino-golang/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {

	// get configuration
	conf := config.LoadConfiguration()

	// create connection string for database connection
	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME
	// connect to mysql database
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic("connectionString Error...")
	}
	err = db.Ping()
	if err != nil {
		panic("DSN Invalid")
	}
}

func CreateCon() *sql.DB {
	// return db connection
	return db
}
