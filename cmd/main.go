package main

import (
	"database/sql"
	"firstproject/cmd/api"

	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "soulwalletadmin"
	password = "soulwalletadmin%812"
	hostname = "soulwallet-uat-database-df.csbrrsmepyza.us-east-2.rds.amazonaws.com"
	dbname   = "soulwallet_uat"
)

func main(){
	db, err := sql.Open("mysql", dsn("soulwallet_uat"))
    if err != nil {
	    log.Printf("Error %s when opening DB\n", err)
	    return
    }
	defer db.Close()
	server := api.NewApiServer(":8080");
	if err := server.Run(); err != nil {
		log.Fatal("Some error accourd", err);
	}
}

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}