package main

import (
	"api/internal/adapter/web"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", web.Handler)
	log.Println(http.ListenAndServe(":8080", nil))
}
