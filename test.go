package main

import (
	"database/sql"
	"time"
    "fmt"
    "log"
	"context"
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

	ctx := context.Background()
	sql := "select * from speda_fleet"

    rows, err := db.QueryContext(ctx, sql)

    if err != nil {
        log.Fatal(err)
    }

    for rows.Next(){
		var id int
		var is_available bool
		var fleet_name string
		var created_date string

		err := rows.Scan(&id, &is_available, &fleet_name, &created_date)
		if err != nil {
			panic(err)
		}

		fmt.Println(id, is_available, fleet_name, created_date)
	}
}