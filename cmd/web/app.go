package main

import (
	"context"
	"net/http"
)

type App struct {
	mux *http.ServeMux
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{
		mux: http.NewServeMux(),
	}

	return app, nil
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
