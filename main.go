package main

import (
	"log"
	"net/http"
)

func main() {
	router := registerRouters()

	log.Println("The azure archives is listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
