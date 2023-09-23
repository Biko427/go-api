package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"github.com/Biko427/go-api/models"
	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB


func main() {
	// database configuration section
	cfg := mysql.Config {
		User: "root",
		Passwd: "Bikobiko59",
		Net: "tcp",
		Addr:   "127.0.0.1:3306",
        DBName: "savings",
	}
	// get a database handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Mysql connected successfully")
	// this is our logic to create our table
	err = models.CreateUserTable(db)
	err = models.CreateAccountTable(db)
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	// r.HandleFunc("/", handlers.SaveAccount(rw, r))
	fmt.Println("Started server successfully at http://localhost:8080")
	
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}