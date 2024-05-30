package internal

import (
	"database/sql"
	"log"
	"testing"
)

func TestRepository(t *testing.T) {
	db, err := connectSQLforTest()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = DBInit(db)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	cases := []struct {
		name     string
		article  []Article
		expected error
	}{
		{
			name: "正常",
			article: []Article{
				{
					Id:          1,
					Title:       "title",
					Body:        "body",
					Medias:      []Media{{Id: 1, ContentUrl: "contentUrl", ContentType: "contentType"}},
					PublishedAt: "2021-01-01",
				},
				{
					Id:          2,
					Title:       "title",
					Body:        "body",
					Medias:      []Media{{Id: 2, ContentUrl: "contentUrl", ContentType: "contentType"}},
					PublishedAt: "2021-01-01",
				},
			},
			expected: nil,
		},
		{
			name: "記事が重複している",
			article: []Article{
				{
					Id:          1,
					Title:       "title",
					Body:        "body",
					Medias:      []Media{{Id: 1, ContentUrl: "contentUrl", ContentType: "contentType"}},
					PublishedAt: "2021-01-01",
				},
				{
					Id:          1,
					Title:       "title",
					Body:        "body",
					Medias:      []Media{{Id: 2, ContentUrl: "contentUrl", ContentType: "contentType"}},
					PublishedAt: "2021-01-01",
				},
			},
			expected: &Error{Message: "article has already been registered"},
		},
		{
			name: "メディアが重複している",
			article: []Article{
				{
					Id:          1,
					Title:       "title",
					Body:        "body",
					Medias:      []Media{{Id: 1, ContentUrl: "contentUrl", ContentType: "contentType"}},
					PublishedAt: "2021-01-01",
				},
				{
					Id:          2,
					Title:       "title",
					Body:        "body",
					Medias:      []Media{{Id: 1, ContentUrl: "contentUrl", ContentType: "contentType"}},
					PublishedAt: "2021-01-01",
				},
			},
			expected: &Error{Message: "failed to register article"},
		},
	}
	for _, tt := range cases {
		err = resetDB(db)
		if err != nil {
			t.Fatal(err)
		}
		t.Run(tt.name, func(t *testing.T) {
			err := RegisterArticleToRepoitory(db, tt.article)
			if err != nil {
				if err.Error() != tt.expected.Error() {
					t.Errorf("expected %v, but got %v", tt.expected, err)
				}
			}

		})
	}

}

func connectSQLforTest() (db *sql.DB, err error) {
	// データベースのハンドルを取得する
	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/entaleAssignment?parseTime=true")
	if err != nil {
		log.Print("サーバーに接続できませんでした。サーバーが起動しているか確認して下さい。")
		return nil, err
	}
	return db, nil
}
func resetDB(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM medias")
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM articles")
	if err != nil {
		return err
	}
	return nil
}
