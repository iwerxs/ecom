package main

import (
	"database/sql"
	"ecom/cmd/api"
	"ecom/db"
	"log"
)

// load the DB
func main(){
	db, err := db.NewMySQLStorage(mysql.Config{
		// handle this with environment variables
		User: config.Envs.DBUser, // change, hard-coded user "root"
		Passwd: config.Envs.DBPassword, // change, hard-coded for demo
		Addr: config.Envs.DBAddress, // a default home address with port 3306
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	// call initialization
	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB){
	err := db.Ping()
	if err != nil{
		log.Fatal(err)
	}
	log.Println("DB: Successfully connected")
}