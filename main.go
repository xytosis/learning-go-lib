package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/xytosis/learning-go-lib/db/dao"
	"github.com/xytosis/learning-go-lib/resources"
	"github.com/xytosis/learning-go-lib/services"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// create daos
	bookDao := dao.BookDao{}

	// create services
	bookService := services.BookingService{}

	// create contexts
	bookServiceCtx := services.BookingServiceContext{&bookDao}
	bookingResourceCtx := resources.BookingResourceContext{r, &bookService, &bookServiceCtx}
	healthcheckResourceCtx := resources.HealthcheckResourceContext{r}

	// register resources
	h := resources.HealthcheckResource{}
	h.RegisterResource(healthcheckResourceCtx)
	b := resources.BookingResource{}
	b.RegisterResource(bookingResourceCtx)

	http.ListenAndServe(":3000", r)
}
