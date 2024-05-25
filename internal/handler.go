package internal

import (
	"log"
	"net/http"
)

func GetArticle(w http.ResponseWriter, r *http.Request) {
	log.Print("get")
	w.Write([]byte("get"))
}

func RegisterArticle(w http.ResponseWriter, r *http.Request) {
	log.Print("register")
	w.Write([]byte("register"))
}
