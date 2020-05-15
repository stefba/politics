package main

import (
	"politics/go/server"
	"log"
	"net/http"
)

func main() {
	s := server.NewServer()

	err := s.Load()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", Router(s))
	http.ListenAndServe(":8553", nil)
}
