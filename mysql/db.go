package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"time"
)

func DBConnection() *sqlx.DB {
	dbDriver := os.Getenv("DB_DRIVER")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbSchema := os.Getenv("DB_SCHEMA")

	db, err := sqlx.Connect(dbDriver, dbUsername+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbSchema)

	// Stop runtime immediately
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
