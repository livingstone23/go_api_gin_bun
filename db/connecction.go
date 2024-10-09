package db

import (
	"database/sql"
	"os"
	"github.com/joho/godotenv"
	
	_ "github.com/go-sql-driver/mysql"
)

// Connection to a database using GUN
func Connection() *sql.DB {

	errorVars := godotenv.Load()
	if errorVars != nil {
		panic("Error loading .env file")
	}

	sqldb, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		panic(err.Error())
	}


	//Create the table, only for the first time or when you want to reset the table
	//and update the model
	/*
	//"go_gin_bun/model"
	//"context"
	//"github.com/uptrace/bun"
	//"github.com/uptrace/bun/dialect/mysqldialect"
	db:= bun.NewDB(sqldb, mysqldialect.New()) 
	erre := db.ResetModel(context.TODO(), &model.TematicModel{})
	if erre != nil {
		panic(erre)
	}
	*/
		
	

	return sqldb

}
