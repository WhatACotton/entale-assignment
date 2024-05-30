package internal

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetArticleFromRemote(url string) (article *[]Article, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &article); err != nil {
		return nil, err
	}
	return article, nil
}
