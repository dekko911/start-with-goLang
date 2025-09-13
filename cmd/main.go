package main

import (
	"database/sql"
	"log"

	"github.com/dekko911/start-with-goLang/cmd/api"
	"github.com/dekko911/start-with-goLang/config"
	"github.com/dekko911/start-with-goLang/db"
	"github.com/go-sql-driver/mysql"
)

// For getting database connection.
func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database Successfully Connected!")
}

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Env.DBUser,
		Passwd:               config.Env.DBPassword,
		Addr:                 config.Env.DBAddress,
		DBName:               config.Env.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	// checking error, Always.
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
