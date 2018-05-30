package controllers

import (
	"github.com/ParallelCoinTeam/duocli/app/services"
	"github.com/revel/revel"
)

type BookController struct {
	*revel.Controller
}

func (c BookController) PublishedBooks() revel.Result {

	resultMap := make(map[string]interface{})
	resultMap = services.BookService{}.GetBooks(c.Session)
	data := resultMap["coins"]
	server := c.Session["server"]
	return c.Render(server, data)
}

// func (c BookController) EditBook(book models.Book) revel.Result {

// 	bookJson, err := json.Marshal(book)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	bookId := strconv.Itoa(book.Id)
// 	_ = services.BookService{}.EditBook(bookId, bookJson, c.Session)

// 	return c.Redirect(routes.BookController.MyBooks())
// }
