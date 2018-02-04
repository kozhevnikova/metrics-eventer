package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func connectToPostgre(config Config) (*sql.DB, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
		config.Database.User, config.Database.Password, config.Database.Name, config.Database.Host)
	db, err := sql.Open("postgres", dbinfo)
	fmt.Println(dbinfo)
	if err != nil {
		return db, err
	}
	fmt.Println("POSTGRESQL IS WORKING")
	return db, nil
}
