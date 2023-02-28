package main

import (
	"github.com/it-02/dormitory/db"
	"github.com/it-02/dormitory/server"
)

func main() {
	db.InitDB()
	server.RunHttpServer()
}
