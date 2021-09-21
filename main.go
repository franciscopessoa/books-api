package main

import (
	"github.com/franciscopessoa/books-api/database"
	"github.com/franciscopessoa/books-api/server"
)

func main() {
	database.Connect()

	server := server.NewServer()
	server.Run()
}
