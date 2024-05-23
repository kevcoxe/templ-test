package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"kevcoxe/templ-test/handlers"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
)

var sessionManager *scs.SessionManager

func main() {
	// Initialize the session.
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour

	mux := chi.NewRouter()
	mux.Use(sessionManager.LoadAndSave)

	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	// Define your routes here
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetHandler(w, r, sessionManager)
	})

	mux.Post("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.PostHandler(w, r, sessionManager)
	})

	// Start the server.
	fmt.Println("listening on http://localhost:3400")
	if err := http.ListenAndServe("0.0.0.0:3400", mux); err != nil {
		log.Printf("error listening: %v", err)
	}
}
