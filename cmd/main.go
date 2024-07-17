package main

import (
	"database/sql"
	"firstproject/cmd/api"
	"firstproject/cmd/config"
	"firstproject/cmd/db"

	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

func main(){
	db, err := db.MySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
		Timeout:              5 * time.Second,
	})

	if err != nil {
		log.Fatal("Failed to initialize database storage: ", err)
	}

	initStorage(db)

	server := api.NewApiServer(config.Envs.Port, db)
	if err := server.Run(); err != nil {
		log.Fatal("An error occurred while running the server: ", err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal("An error occurred while connecting to database: ", err)
	}

	log.Println("DB connected successfully")
}