package server

import (
	"database/sql"
	_ "mysql-master"
)

func Koneksi() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost)/db_pizza")

	if err != nil {
		return nil, err
	}

	return db, nil
}
