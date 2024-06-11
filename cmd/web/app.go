package main

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/crgeary/wishlist/views"
)

type App struct {
	mux *http.ServeMux
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{
		mux: http.NewServeMux(),
	}

	app.mux.Handle("/", templ.Handler(views.Index()))

	return app, nil
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
