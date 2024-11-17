package main

import (
	"log"

	srv "github.com/MotyaSS/DB_CW/pkg/server"
)

const port string = ":8080"

func main() {
	server := srv.NewServer()
	log.Fatal(server.Run(port))
}
