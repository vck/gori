package main

import (
	"database/sql"
	"time"
    "fmt"
    "log"
	_ "github.com/go-sql-driver/mysql"
)



func main(){

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/speda")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	var version string

    err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

    if err2 != nil {
        log.Fatal(err2)
    }

    fmt.Println(version)
}