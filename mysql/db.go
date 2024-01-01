package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

func DBConnection() *sqlx.DB {
	db, err := sqlx.Connect("mysql", "root:root@tcp(127.0.0.1:3306)/latihan_api_go")

	// Stop runtime immediately
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
