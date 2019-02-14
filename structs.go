package newsapi

import "time"

type articlesResponse struct {
	Status			string		`json:"status"`
	TotalResults		int		`json:"totalResults"`
	Articles		[]Article	`json:"articles"`
}

type sourcesResponse struct {
	Status			string		`json:"status"`
	Sources			[]Source	`json:"sources"`
}

type Article struct {
	Source			Source		`json:"source"`
	Author			string		`json:"author"`
	Title			string		`json:"title"`
	Description		string		`json:"description"`
	URL			string		`json:"url"`
	ImageURL		string		`json:"urlToImage"`
	PublishedAt		time.Time	`json:"publishedAt"`
	Content			string		`json:"content"`
}

type Source struct {
	Id			string 		`json:"id"`
	Name			string		`json:"name"`
}
