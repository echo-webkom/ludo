package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Listening at port :8080")
	http.ListenAndServe(":8080", http.FileServer(http.Dir("web")))
}
