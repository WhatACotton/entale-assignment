package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/whatacotton/entale-assignment/internal"
)

func main() {
	db, err := internal.ConnectSQL()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = internal.DBInit(db)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	r := chi.NewRouter()
	h := internal.Handler{DB: db}
	r.Get("/register", h.RegisterArticle)
	r.Get("/", h.GetArticle)
	log.Print("server started")
	http.ListenAndServe(":8080", r)
}
