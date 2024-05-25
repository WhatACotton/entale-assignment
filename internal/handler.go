package internal

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Handler struct {
	DB *sql.DB
}

func (h *Handler) RegisterArticle(w http.ResponseWriter, r *http.Request) {
	article, err := RegisterArticle("https://gist.githubusercontent.com/gotokatsuya/cc78c04d3af15ebe43afe5ad970bc334/raw/dc39bacb834105c81497ba08940be5432ed69848/articles.json")
	if err != nil {
		w.Write([]byte("fetch error"))
	}
	if article != nil {
		for _, a := range *article {
			err = RegisterArticleToRepoitory(h.DB, a)
			if err != nil {
				log.Print(err.Error())
				if strings.HasPrefix(err.Error(), "Error 1062") {
					w.Write([]byte("the articles have already been registered"))
					return
				}
				w.Write([]byte("DB error"))
				return
			}
		}
		w.Write([]byte("article was successfully registered"))
	}
}
func (h *Handler) GetArticle(w http.ResponseWriter, r *http.Request) {
	article, err := GetArticleFromRepository(h.DB)
	if err != nil {
		w.Write([]byte("fetch error"))
	}

	b, err := json.Marshal(article)
	if err != nil {
		w.Write([]byte("fetch error"))
	}
	w.Write(b)
}
