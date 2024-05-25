package internal

import (
	"log"

	"github.com/jmoiron/sqlx"
)

const articleInit = `
CREATE TABLE IF NOT EXISTS articles (
	id int(11) NOT NULL,
	title text NOT NULL,
	body text NOT NULL,
	PRIMARY KEY (id)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
`

const mediaInit = `
CREATE TABLE IF NOT EXISTS medias (
	id int(11) NOT NULL,
	article_id int(11) NOT NULL,
	contentUrl text NOT NULL,
	content_type varchar(100) NOT NULL,
	PRIMARY KEY (id),
	KEY medias_articles_FK (article_id),
	CONSTRAINT medias_articles_FK FOREIGN KEY (article_id) REFERENCES articles (id)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
`

func ConnectSQL() (db *sqlx.DB) {
	// データベースのハンドルを取得する
	db, err := sqlx.Connect("mysql", "root:password@tcp(entaleAssignmentdb:3306)/entaleAssignment")

	if err != nil {
		// ここではエラーを返さない
		log.Fatal(err)
	}
	return db
}

func DBInit() error {
	db := ConnectSQL()
	_, err := db.Exec(articleInit)
	if err != nil {
		return err
	}
	log.Print("article")
	_, err = db.Exec(mediaInit)
	if err != nil {
		return err
	}
	log.Print("media")
	return nil
}
