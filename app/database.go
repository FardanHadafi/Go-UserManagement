package app

import (
	"Go-UserManagement/helper"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	DSN := "host=localhost port=5432 user=postgres dbname=go_user_management sslmode=disable"
	db, err := sql.Open("postgres", DSN)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return  db
}