package main

import (
	hnd "github.com/MotyaSS/DB_CW/pkg/handler"
	srvr "github.com/MotyaSS/DB_CW/pkg/server"
	srvc "github.com/MotyaSS/DB_CW/pkg/service"
	strg "github.com/MotyaSS/DB_CW/pkg/storage"
)

const port string = ":8080"

func main() {
	storage := strg.New()
	service := srvc.New(storage)
	handler := hnd.New(service)
	server := srvr.New(port, handler.InitRouter())
	server.Run()
}
