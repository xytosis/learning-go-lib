package resources

import "github.com/go-chi/chi/v5"

type Resource interface {
	RegisterResource(r *chi.Mux)
}
