package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func RegisterArticle(url string) (article *[]Article, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &article); err != nil {
		fmt.Println(err)
		return
	}
	return article, nil
}
