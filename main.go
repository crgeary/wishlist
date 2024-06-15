package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/a-h/templ"
	"github.com/crgeary/wishlist/views"
)

func main() {
	app, err := NewApp(context.Background())

	if err != nil {
		log.Fatalf("failed to init new application: %v", err)
		return
	}

	s := &http.Server{
		Addr:         ":4000",
		Handler:      app,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 4 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGTERM)

	<-sc

	ctx, cancelCtx := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelCtx()

	s.Shutdown(ctx)
}

type App struct {
	mux *http.ServeMux
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{
		mux: http.NewServeMux(),
	}

	app.mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	app.mux.Handle("/", templ.Handler(views.Index()))
	app.mux.Handle("/register", templ.Handler(views.Register()))
	app.mux.Handle("/signin", templ.Handler(views.SignIn()))

	return app, nil
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
