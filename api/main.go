package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Name string
}

func handler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "hello"}
	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println(http.ListenAndServe(":8080", nil))
}
