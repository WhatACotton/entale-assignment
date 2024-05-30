package internal

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Handler struct {
	DB *sql.DB
}

func (h *Handler) RegisterArticle(w http.ResponseWriter, r *http.Request) {
	article, err := GetArticleFromRemote("https://gist.githubusercontent.com/gotokatsuya/cc78c04d3af15ebe43afe5ad970bc334/raw/dc39bacb834105c81497ba08940be5432ed69848/articles.json")
	if err != nil {
		w.Write([]byte("fetch error"))
	}
	if article != nil {
		err = RegisterArticleToRepoitory(h.DB, *article)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		w.Write([]byte("article was successfully registered"))
	}
}
func (h *Handler) GetArticle(w http.ResponseWriter, r *http.Request) {
	article, err := GetArticleFromRepository(h.DB)
	if err != nil {
		w.Write([]byte("fetch error"))
		return
	}
	if article == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("article not found"))
		http.NotFound(w, r)
		return
	}
	b, err := json.Marshal(article)
	if err != nil {
		w.Write([]byte("fetch error"))
		return
	}
	w.Write(b)
}
