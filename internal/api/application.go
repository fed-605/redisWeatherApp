package api

import (
	"net/http"

	"github.com/fed-605/weatherApi/internal/cache"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type Application struct {
	cache cache.Cache
}

func NewApplication(cache cache.Cache) *Application {
	return &Application{
		cache: cache,
	}
}

func (app *Application) Routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/weather/{location}", app.GetWeatherByCity)
	r.Get("/healthz", app.HealthzChecker)
	return r
}
