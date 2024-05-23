package handlers

import (
	"fmt"
	"net/http"

	"github.com/kevcoxe/templ-test/components"

	"github.com/alexedwards/scs/v2"
)

type GlobalState struct {
	Count int
}

var global GlobalState

func GetHandler(w http.ResponseWriter, r *http.Request, sm *scs.SessionManager) {
	userCount := sm.GetInt(r.Context(), "count")
	component := components.HomePage(global.Count, userCount, r.Context(), sm)
	component.Render(r.Context(), w)
}

func PostHandler(w http.ResponseWriter, r *http.Request, sm *scs.SessionManager) {
	// Update state.
	r.ParseForm()

	// Check to see if the global button was pressed.
	if r.Form.Has("global") {
		global.Count++
	}
	if r.Form.Has("user") {
		currentCount := sm.GetInt(r.Context(), "count")
		sm.Put(r.Context(), "count", currentCount+1)
	}

	// Display the form.
	GetHandler(w, r, sm)
}

func PutHandler(w http.ResponseWriter, r *http.Request, sm *scs.SessionManager) {
	// Update state.
	r.ParseForm()

	username := r.Form.Get("username")

	if username != "" {
		sm.Put(r.Context(), "username", username)
		sm.Remove(r.Context(), "usernameError")
		fmt.Printf("logged in with username: '%v'\n", username)
	} else {
		sm.Put(r.Context(), "usernameError", "Username cannot be empty")
	}

	GetHandler(w, r, sm)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request, sm *scs.SessionManager) {
	username := sm.GetString(r.Context(), "username")

	sm.Remove(r.Context(), "username")
	sm.Remove(r.Context(), "count")

	if username != "" {
		fmt.Printf("user '%v' has logged out\n", username)
	}
	GetHandler(w, r, sm)
}
