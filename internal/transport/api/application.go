package api

import (
	"net/http"

	"github.com/fed-605/weatherApi/internal/app"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type Application struct {
	Service *app.Service
}

func NewApplication(service *app.Service) *Application {
	return &Application{
		Service: service,
	}
}

func (app *Application) Routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/weather/{location}", app.GetWeatherByCity)
	r.Get("/healthz", app.HealthzChecker)
	return r
}
