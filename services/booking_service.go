package services

import (
	"github.com/xytosis/learning-go-lib/db/dao"
	"github.com/xytosis/learning-go-lib/db/entity"
)

type BookingServiceImpl struct {
}

type BookingService interface {
	ListBooks(cxt *BookingServiceContext) []entity.Book
	DeleteBook(ctx *BookingServiceContext, id string) bool
}

type BookingServiceContext struct {
	BookDao dao.BookDao
}

func (b BookingServiceImpl) ListBooks(ctx *BookingServiceContext) []entity.Book {
	return ctx.BookDao.ListBooks()
}

func (b BookingServiceImpl) DeleteBook(ctx *BookingServiceContext, id string) bool {
	return ctx.BookDao.DeleteBooks(id)
}
