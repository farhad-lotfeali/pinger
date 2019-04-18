package database

import (
	"database/sql"
	"os"
)

//NewDB create new db
func NewDB() *sql.DB {
	// IP := os.Getenv("db_ip")
	DataBase := os.Getenv("db_database")
	Username := os.Getenv("db_username")
	Password := os.Getenv("db_password")

	db, err := sql.Open("mysql", Username+":"+Password+"@/"+DataBase)
	if err != nil {
		panic(err)
	}
	return db
}
