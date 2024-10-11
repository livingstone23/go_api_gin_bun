package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Connection to a database using GUN
func Connection() *sql.DB {

	errorVars := godotenv.Load()
	if errorVars != nil {
		panic("Error loading .env file")
	}

	// Construir la cadena de conexión
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SERVER"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	sqldb, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	//Create the table, only for the first time or when you want to reset the table
	//and update the model

	//"context"
	//"go_gin_bun/model"
	//"github.com/uptrace/bun"
	//"github.com/uptrace/bun/dialect/mysqldialect"
	/*
		db:= bun.NewDB(sqldb, mysqldialect.New())
		erre := db.ResetModel(context.TODO(), &model.PerfilModel{})
		if erre != nil {
			panic(erre)
		}
	*/

	return sqldb

}
