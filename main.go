package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/whatacotton/entale-assignment/internal"
)

func main() {
	time.Sleep(3 * time.Second)
	db := internal.ConnectSQL()
	err := internal.DBInit(db)
	if err != nil {
		log.Fatal("DB connection error")
	}
	r := chi.NewRouter()
	h := internal.Handler{DB: db}
	r.Get("/register", h.RegisterArticle)
	r.Get("/", h.GetArticle)
	http.ListenAndServe(":8080", r)
}
