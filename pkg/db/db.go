package db

import (
	"database/sql"
	_ "github.com/lib/pq" // Import the pq package
	"github.com/sirupsen/logrus"
	"log"
)

func GetDBConnection() *sql.DB {
	logrus.Info("connecting")
	connStr := "user=username dbname=mydb sslmode=disable"

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return db
}
