package main

import "github.com/franciscopessoa/books-api/server"

func main() {
	server := server.NewServer()
	server.Run()
}
