package web

import (
	"api/internal/entity"
	"encoding/json"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	user := entity.User{Name: "hello"}
	my := json.NewEncoder(w)
	my.Encode(user)
}
