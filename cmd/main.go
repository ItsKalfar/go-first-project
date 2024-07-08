package main

import (
	"firstproject/cmd/api"
	"log"
)

func main(){
	server := api.NewApiServer(":8080");
	if err := server.Run(); err != nil {
		log.Fatal("Some error accourd", err);
	}
}