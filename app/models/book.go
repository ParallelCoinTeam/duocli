package models

type Book struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Img  string `json:"img"`
}

type BookResponse struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Img  string `json:"img"`
}
