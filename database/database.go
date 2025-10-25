package database

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

func ConnectDB() *sqlx.DB{
	connStr := os.Getenv("DB_URL")
	if connStr == ""{
		log.Fatal("DB_URL isnt set")
	}

	db ,err := sqlx.Connect("psotgres", connStr)
	if err != nil{
		log.Fatal("failed to connect to database: ", err)
	}

	log.Println("Database connected succssfully")
	return db
}