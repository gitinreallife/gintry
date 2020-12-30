package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

// Replace with your own connection parameters
var server = "172.18.133.110"
var port = 1433
var user = "sa"
var password = "starbuck"
var database = "LAS"
var db *sql.DB

func ConnectLASDB() {
	var err error

	// Create connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating LAS connection pool: " + err.Error())
	}
	log.Printf("Connected!\n")
}
