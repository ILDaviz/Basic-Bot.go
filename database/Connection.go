package database

import (
	"database/sql"                  //package database/sql
	_ "github.com/mattn/go-sqlite3" //driver sqlite3
	"log"
)

var db *sql.DB //variable where the connection to the database will be stored

func GetConnection() *sql.DB { //Connection to the database
	if db != nil { //if there is already an open connection to the database we return that connection
		return db
	}

	var err error
	db, err = sql.Open("sqlite3", "./data/database.db")
	//driver: sqlite3
	//databasename: database.db

	if err != nil {
		panic(err) //If an error occurs while opening the connection to the database
	}
	return db //We return the connection
}

func init() {
	db = GetConnection() //Connection to the database

	query := `CREATE TABLE IF NOT EXISTS settings (
        		GuildId INTEGER PRIMARY KEY NOT NULL,
        		Prefix TEXT
				 );` //sql statement

	_, err := db.Exec(query) //execution of sentence
	if err != nil {
		//if an error occurs while inserting the data, we will return it (normally it may be because there is no connection to the database)
		log.Fatal(err)
	}
}
