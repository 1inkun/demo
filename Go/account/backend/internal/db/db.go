package db

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectToDB () (*sql.DB,error){
	db, err := sql.Open("sqlite3","./data/backend.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	if err = db.Ping(); err != nil {
		return nil,err
	}
	return db,err
}
