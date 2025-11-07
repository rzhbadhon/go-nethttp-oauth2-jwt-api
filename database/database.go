package database

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)
// this func connects db which is postgres
func ConnectDB() *sqlx.DB{
	connStr := os.Getenv("DB_URL")
	if connStr == ""{
		log.Fatal("DB_URL isnt set")
	}

	db ,err := sqlx.Connect("postgres", connStr)
	if err != nil{
		log.Fatal("failed to connect to database: ", err)
	}

	log.Println("Database connected succssfully")
	return db
}