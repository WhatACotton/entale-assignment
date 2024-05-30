package internal

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

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
		var art Article
		var media Media
		var id int
		err = rows.Scan(&art.Id, &art.Title, &art.Body, &art.PublishedAt, &media.Id, &id, &media.ContentUrl, &media.ContentType)
		if err != nil {
			return nil, err
		}

		a, exists := mappedArticle[id]
		if !exists {
			a = art
			mappedArticle[id] = a
		}
		a.Medias = append(a.Medias, media)
		mappedArticle[id] = a

	}
	for i := 0; i < len(mappedArticle); i++ {
		articles = append(articles, mappedArticle[i+1])
	}
	return articles, nil
}

const articleInit = `
CREATE TABLE IF NOT EXISTS articles (
	id int(11) NOT NULL,
	title text NOT NULL,
	body text NOT NULL,
	published_at text NOT NULL,
	PRIMARY KEY (id)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
`

const mediaInit = `
CREATE TABLE IF NOT EXISTS medias (
	id int(11) NOT NULL,
	article_id int(11) NOT NULL,
	content_url text NOT NULL,
	content_type varchar(100) NOT NULL,
	PRIMARY KEY (id),
	KEY medias_articles_FK (article_id),
	CONSTRAINT medias_articles_FK FOREIGN KEY (article_id) REFERENCES articles (id)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
`

func ConnectSQL() (db *sql.DB, err error) {
	// データベースのハンドルを取得する
	db, err = sql.Open("mysql", "root:password@tcp(entaleAssignmentdb:3306)/entaleAssignment?parseTime=true")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func DBInit(db *sql.DB) error {
	_, err := db.Exec(articleInit)
	if err != nil {
		return err
	}
	_, err = db.Exec(mediaInit)
	if err != nil {
		return err
	}
	return nil
}
