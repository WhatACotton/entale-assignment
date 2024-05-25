package internal

import (
	"log"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

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

func ConnectSQL() (db *sql.DB) {
	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", "root:password@tcp(entaleAssignmentdb:3306)/entaleAssignment?parseTime=true")

	if err != nil {
		// ここではエラーを返さない
		log.Fatal(err)
	}
	return db
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
