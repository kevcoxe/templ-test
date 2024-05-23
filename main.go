package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"templ-test/handlers"

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

	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	// Define your routes here
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetHandler(w, r, sessionManager)
	})

	mux.Post("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.PostHandler(w, r, sessionManager)
	})

	mux.Put("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.PutHandler(w, r, sessionManager)
	})

	mux.Delete("/logout", func(w http.ResponseWriter, r *http.Request) {
		handlers.LogoutHandler(w, r, sessionManager)
	})

	// Start the server.
	fmt.Println("listening on http://localhost:3400")
	if err := http.ListenAndServe("0.0.0.0:3400", mux); err != nil {
		log.Error("error listening %v", err)
	}
}
