package internal

import "database/sql"

func RegisterArticleToRepoitory(db *sql.DB, article Article) error {
	tx, _ := db.Begin()
	_, err := tx.Exec(`
	INSERT INTO
		articles
		(id,title,body,published_at)
	VALUES
		(?,?,?,?)`, article.Id, article.Title, article.Body, article.PublishedAt)
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, m := range article.Medias {
		_, err = tx.Exec(`
		INSERT INTO
			medias
			(id,article_id,content_url,content_type)
		VALUES
			(?,?,?,?)
		`, m.Id, article.Id, m.ContentUrl, m.ContentType)
	}
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	defer db.Close()
	return nil
}
func GetArticleFromRepository(db *sql.DB) (articles []Article, err error) {
	mappedArticle := map[int]Article{}
	rows, err := db.Query(`SELECT 
	articles.*,medias.*
FROM 
	articles
JOIN
	medias
ON
	articles.id = medias.article_id`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var a Article
		var media Media
		var id int
		err = rows.Scan(&a.Id, &a.Title, &a.Body, &a.PublishedAt, &media.Id, &id, &media.ContentUrl, &media.ContentType)
		if err != nil {
			return nil, err
		}
		aa, exist := mappedArticle[a.Id]
		if exist {
			aa.Medias = append(aa.Medias, media)
			mappedArticle[a.Id] = aa
		} else {
			a.Medias = append(a.Medias, media)
			mappedArticle[a.Id] = a
		}
	}
	for i := 0; i < len(mappedArticle); i++ {
		articles = append(articles, mappedArticle[i+1])
	}

	return articles, nil
}
