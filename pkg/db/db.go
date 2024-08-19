package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Import the pq package
	"github.com/sirupsen/logrus"
	"log"
)

func GetDBConnection() *sql.DB {
	logrus.Info("connecting")
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		"localhost", 5432, "example", "example", "postgres", "disable")

	// Open a connection to the database
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return db
}
