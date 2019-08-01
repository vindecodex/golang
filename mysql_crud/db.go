package main

import "database/sql"

var db *sql.DB

func connectDB() {
	db, err = sql.Open("mysql", "root:root@/crud")
}
