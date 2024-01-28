package unit

import (
	"github.com/xytosis/learning-go-lib/db/entity"
	"testing"
)

var fakeBooks = []entity.Book{{
	ID:               "1",
	Title:            "7 Habits of Highly Effective People",
	Author:           "Stephen Covey",
	PublishedDate:    "15/08/1989",
	OriginalLanguage: "English",
}}

type fakeStorage struct {
}

func (b fakeStorage) ListBooks() []entity.Book {
	return fakeBooks
}

func (b fakeStorage) DeleteBooks(id string) bool {
	return true
}

func TestListBooks(t *testing.T) {
	books := fakeStorage{}.ListBooks()
	if books[0].Title != "7 Habits of Highly Effective People" {
		t.Errorf("expected ABC got %v", books[0])
	}
}
