package resources

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/xytosis/learning-go-lib/services"
	"net/http"
)

type BookingResource struct {
}

type BookingResourceContext struct {
	R              *chi.Mux
	BookService    services.BookingService
	BookServiceCtx *services.BookingServiceContext
}

func (b BookingResource) RegisterResource(ctx BookingResourceContext) {
	b.listBookings(ctx)
	b.deleteBooking(ctx)
	b.updateBook(ctx)
}

func (b BookingResource) listBookings(ctx BookingResourceContext) {
	ctx.R.Get("/listBooks", func(w http.ResponseWriter, req *http.Request) {
		err := json.NewEncoder(w).Encode(ctx.BookService.ListBooks(ctx.BookServiceCtx))
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
	})
}

func (b BookingResource) deleteBooking(ctx BookingResourceContext) {
	ctx.R.Get("/deleteBook", func(w http.ResponseWriter, req *http.Request) {
		err := json.NewEncoder(w).Encode(ctx.BookService.DeleteBook(ctx.BookServiceCtx, ""))
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
	})
}

func (b BookingResource) updateBook(ctx BookingResourceContext) {
	ctx.R.Get("/updateBook", func(w http.ResponseWriter, req *http.Request) {
		err := json.NewEncoder(w).Encode(ctx.BookService.DeleteBook(ctx.BookServiceCtx, ""))
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
	})
}
