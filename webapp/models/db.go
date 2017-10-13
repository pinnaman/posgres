package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

//dataSourceName="postgres://postgres:secret@192.168.99.100/realtime"
func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Panic(err)
	}
	//defer db.Close()

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	fmt.Println("Successfully connected!")
}
