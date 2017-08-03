package utility

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// InitDB initializes the database connection
func InitDB() {
	db, err := sql.Open("mysql", "PLACEHOLDER")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	} else {
		log.Printf("Connection established")
	}
}
