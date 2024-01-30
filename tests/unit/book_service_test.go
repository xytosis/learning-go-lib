package unit

import (
	"github.com/xytosis/learning-go-lib/db/entity"
	"github.com/xytosis/learning-go-lib/services"
	"testing"
)

type fakeDao struct {
}

func TestListBookImpl(t *testing.T) {
	ctx := services.BookingServiceContext{BookDao: fakeDao{}}
	bookService := services.BookingServiceImpl{}
	books := bookService.ListBooks(&ctx)
	if len(books) > 0 {
		t.Errorf("expected empty books, got %v", books)
	}
}

func (b fakeDao) ListBooks() []entity.Book {
	return []entity.Book{}
}

func (b fakeDao) DeleteBooks(id string) bool {
	return true
}
