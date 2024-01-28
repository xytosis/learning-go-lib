package dao

import "github.com/xytosis/learning-go-lib/db/entity"

type BookDao struct {
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

func (b BookDao) ListBooks() []entity.Book {
	return books
}

func (b BookDao) DeleteBooks(id string) bool {
	return true
}
