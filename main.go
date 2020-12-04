package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// fmt.Println("Drivers: ", sql.Drivers())
	db, err := sql.Open("mysql", "root:password1@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal("Unable to connect to Database...")
	}
	defer db.Close()
	results, err := db.Query("SELECT * FROM products")
	if err != nil {
		log.Fatal("Error while fetching table rows")
	}
	for results.Next() {
		var (
			id    int
			name  string
			price int
		)
		if err = results.Scan(&id, &name, &price); err != nil {
			log.Fatal("unable to parse row")
		}
		fmt.Printf("ID: %d, Name: '%s', Price: %d\n", id, name, price)
	}

}
