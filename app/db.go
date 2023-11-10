package app

import (
	"database/sql"

	"github.com/Alfian57/belajar_golang_restfull_api/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/belajar_golang_restfull_api")
	helper.PanicIfError(err)

	return db
}
