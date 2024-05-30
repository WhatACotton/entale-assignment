package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestHandler(t *testing.T) {
	db, err := connectSQLforTest()
	if err != nil {
		t.Error("failed to connect sql")
	}
	err = resetDB(db)
	if err != nil {
		t.Error("failed to reset db")
	}

	article, err := GetArticleFromRemote("https://gist.githubusercontent.com/gotokatsuya/cc78c04d3af15ebe43afe5ad970bc334/raw/dc39bacb834105c81497ba08940be5432ed69848/articles.json")
	if err != nil {
		t.Error("failed to fetch article")
	}
	if article != nil {
		err = RegisterArticleToRepoitory(db, *article)
		if err != nil {
			t.Error("failed to register article")
		}
	}
	articles, err := GetArticleFromRepository(db)
	if err != nil {
		t.Error("failed to get article")
	}
	resp, err := http.Get("https://gist.githubusercontent.com/gotokatsuya/cc78c04d3af15ebe43afe5ad970bc334/raw/dc39bacb834105c81497ba08940be5432ed69848/articles.json")
	if err != nil {
		t.Error("failed to fetch article")
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if body == nil {
		t.Error("failed to fetch article")
	}
	if err = json.Unmarshal(body, &article); err != nil {
		t.Error("failed to unmarshal")
	}

	if !reflect.DeepEqual(articles, *article) {
		t.Error("failed to get article")

	}
}
