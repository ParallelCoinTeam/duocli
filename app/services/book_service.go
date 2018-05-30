package services

import (
	"github.com/revel/revel"
)

type BookService struct {
	CommonService
}

func (c BookService) GetBooks(session revel.Session) map[string]interface{} {
	return c.CallGetBackend("GET", "wp-json/coins/all", session)
}

func (c BookService) GetABook(book_id string, session revel.Session) map[string]interface{} {
	return c.CallGetBackend("GET", "book/"+slug, session)
}
