package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

const (
	dbName = "todo_app"
	port   = 5432
)

var (
	user     = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	host     = os.Getenv("POSTGRES_HOST")
)

func main() {
	db, err := newDbConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := buildRouter()
	r.Get("/", homeHandler)
	r.Get("/todos", func(w http.ResponseWriter, r *http.Request) {
		getTodosHandler(w, r, db)
	})
	r.Post("/todos", func(w http.ResponseWriter, r *http.Request) {
		createTodoHandler(w, r, db)
	})

	log.Println("Starting server on port 3000...")
	http.ListenAndServe(":3000", r)
}

func buildRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	return r
}
