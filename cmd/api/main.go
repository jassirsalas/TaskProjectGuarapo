package main

import (
	"fmt"
	"log"
	"net/http"

	"taskproject/internal/handlers"
)

func main() {

	server := handlers.NewServer()

	log.Println("Server listening on port", server.Addr)

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}
}
