package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/whatacotton/entale-assignment/internal"
)

func main() {
	time.Sleep(2 * time.Second)

	if err := internal.DBInit(); err != nil {
		log.Print(err)
	}
	log.Print("server start")

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/", internal.GetArticle)
	r.Get("/register", internal.RegisterArticle)
	http.ListenAndServe(":8080", r)
}
