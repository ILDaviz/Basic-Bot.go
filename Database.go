package main

import (
	"database/sql"  //package database/sql
	"errors"		 	 //package errors
	
	"github.com/andersfylling/disgord" //lib disgord
	_ "github.com/mattn/go-sqlite3"	  //driver sqlite3
)

var db *sql.DB //variable where the connection to the database will be stored


/* Create the default server data only if it does not exist */
func Create(id disgord.Snowflake, prefix string) error{
	db := GetConnection() //Connection to the database

	query := `INSERT INTO settings (GuildId, Prefix)
					VALUES(?,?)` //Sql statement
 
	stmt, err := db.Prepare(query) //queries prepared
	if err != nil{
		 return err //If an error occurs we will return it
	}
	defer stmt.Close() //We make sure to close the query prepared at the end of the function execution

	rows, err := stmt.Exec(id, prefix) //We execute
	if err != nil{ //If an error occurs while inserting data, we return the error
		 return err
	}

	if i, err := rows.RowsAffected(); err != nil || i != 1{
		//If an error occurs or the affected rows were 0, we return this error
		 return errors.New("ERROR: An affected row was expected")
	}

	return nil //At this point, if it gets here everything went well so we returned nil
}

/* Get the prefix from the database */
func GetPrefix(gId disgord.Snowflake) string{
	db := GetConnection() //Connection to the database
	var Dataprefix string //var Dataprefix
	query := `SELECT prefix FROM settings WHERE GuildId = ?` //Sql statement

	err := db.QueryRow(query, gId).Scan(&Dataprefix) //We execute the query and pass the required parameter (server id)
	if err != nil{
		//If an error occurs it can usually be because I did not find a row linked to the server id,
		// so I inserted new data
		
		 _ = Create(gId, config.Prefix)
		 Dataprefix = config.Prefix
		 /* As we insert the default prefix (the one that is in the config.json) 
		 we will return that prefix and for the next query it will be obtaining 
		 the prefix from the database */
	}

	return Dataprefix //There will always be a returned prefix
}

/*--------- Prefix update ---------*/
func UpdatePrefix(id disgord.Snowflake, newPrefix string) error{
	db := GetConnection() //Connection to the database
	query := `UPDATE settings SET Prefix = ? WHERE GuildId = ?` //Sql statement
	stmt, err := db.Prepare(query) //Preparing the sentence
	if err != nil{
		//If an error occurs we will return it
		 return err
	}
	
	defer stmt.Close() //we ensure that you close the query at the end

	row, err := stmt.Exec(newPrefix, id)//We pass the necessary parameters for the shift consultation
	if err != nil{
		/* If an error occurs normally it is that I did not find the row linked to the server id, so we insert the data but with the new prefix */
		 Create(id, newPrefix)
		 return nil
	}

	if i, err := row.RowsAffected(); err != nil || i != 1{
		//If an error occurs or the affected rows were 0, we return this error
		 return errors.New("ERROR: An affected row was expected")
	}


	return nil //At this point, if it gets here everything went well so we returned nil
}

func GetConnection() *sql.DB { //Connection to the database
	if db != nil { //if there is already an open connection to the database we return that connection
		return db
	}

	var err error
	db, err = sql.Open("sqlite3", "database.db")
	//driver: sqlite3
	//databasename: database.db

	if err != nil {
		panic(err) //If an error occurs while opening the connection to the database
	}
	return db //We return the connection
}


// Create db tables if they don't exist
func CreateTableIfNotExist() error { 
	 db := GetConnection() //Connection to the database
	 
    query := `CREATE TABLE IF NOT EXISTS settings (
        		GuildId INTEGER PRIMARY KEY NOT NULL,
        		Prefix TEXT
				 );`//sql statement
				 
	_, err := db.Exec(query) //execution of sentence
   if err != nil{
		//if an error occurs while inserting the data, we will return it (normally it may be because there is no connection to the database)
      return err
    }

    return nil //At this point, if it gets here everything went well so we returned nil
}