package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func CreateConn() *sql.DB {

	// Load env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// Set postgres url for connecting
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", "localhost", 5432, os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	// Connect db
	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		panic(err)
	}

	// Check db connection status
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected succesfully!")
	return db
}
