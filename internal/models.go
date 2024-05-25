package internal

type Article struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Body   string  `json:"body"`
	Medias []Media `json:"medias"`
}

type Media struct {
	Id          int    `json:"id"`
	ContentUrl  string `json:"contentUrl"`
	ContentType string `json:"contentType"`
}
