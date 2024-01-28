package services

import (
	"github.com/xytosis/learning-go-lib/db/dao"
	"github.com/xytosis/learning-go-lib/db/entity"
)

type BookingService struct {
}

type BookingServiceContext struct {
	BookDao *dao.BookDao
}

func (b BookingService) ListBooks(ctx *BookingServiceContext) []entity.Book {
	return ctx.BookDao.ListBooks()
}

func (b BookingService) DeleteBook(ctx *BookingServiceContext, id string) bool {
	return ctx.BookDao.DeleteBooks(id)
}
