package main

import (
	"net/http"
	"time"

	"github.com/giorgiodots/todo-go-api/routes"
	"github.com/giorgiodots/todo-go-api/store/memory"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	store := memory.NewInMemoryStore()
	tdr := routes.NewTodosResource(store)
	r.Mount("/todos", tdr.Routes())

	http.ListenAndServe(":3333", r)
}
