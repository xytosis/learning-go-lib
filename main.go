package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/xytosis/learning-go-lib/db/dao"
	"github.com/xytosis/learning-go-lib/resources"
	"github.com/xytosis/learning-go-lib/services"
	"net/http"
	"os"
)

func main() {
	r := setupServer()

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}

func setupServer() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// create daos
	bookDao := dao.BookDaoImpl{}

	// create services
	bookService := services.BookingServiceImpl{}

	// create contexts
	bookServiceCtx := services.BookingServiceContext{BookDao: &bookDao}
	bookingResourceCtx := resources.BookingResourceContext{R: r, BookService: &bookService, BookServiceCtx: &bookServiceCtx}
	healthcheckResourceCtx := resources.HealthcheckResourceContext{R: r}

	// register resources
	h := resources.HealthcheckResource{}
	h.RegisterResource(healthcheckResourceCtx)
	b := resources.BookingResource{}
	b.RegisterResource(bookingResourceCtx)

	return r
}
