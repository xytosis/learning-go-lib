package dao

import "github.com/xytosis/learning-go-lib/db/entity"

type BookDaoImpl struct {
}

type BookDao interface {
	ListBooks() []entity.Book
	DeleteBooks(id string) bool
}

var books = []entity.Book{
	{
		ID:               "1",
		Title:            "7 habits of Highly Effective People",
		Author:           "Stephen Covey",
		PublishedDate:    "15/08/1989",
		OriginalLanguage: "English",
	},
}

func (b BookDaoImpl) ListBooks() []entity.Book {
	return books
}

func (b BookDaoImpl) DeleteBooks(id string) bool {
	return true
}
