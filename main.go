package main

import (
	"login-api/db"
	"login-api/server"
)

func main() {
	db.Init()
	server.Init()

	db.Close()
}
