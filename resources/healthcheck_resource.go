package resources

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type HealthcheckResource struct {
}

type HealthcheckResourceContext struct {
	R *chi.Mux
}

func (h HealthcheckResource) RegisterResource(ctx HealthcheckResourceContext) {
	h.healthCheck(ctx)
}

func (h HealthcheckResource) healthCheck(ctx HealthcheckResourceContext) {
	ctx.R.Get("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("OK"))
	})
}
