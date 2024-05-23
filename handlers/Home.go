package handlers

import (
	"fmt"
	"net/http"
	"time"

	"templ-test/components"

	"github.com/alexedwards/scs/v2"
)

var global components.GlobalState = components.GlobalState{
	Count: 0,
	Messages: []components.Message{{
		User:    "FirstMate",
		Message: "Well here we are this is the first message.",
		Time:    time.Now().Format("Mon 15:04:05"),
	}},
}

func GetHandler(w http.ResponseWriter, r *http.Request, sm *scs.SessionManager) {
	userCount := sm.GetInt(r.Context(), "count")
	component := components.HomePage(global, userCount, r.Context(), sm)
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

func MessageHandler(w http.ResponseWriter, r *http.Request, sm *scs.SessionManager) {
	// Update state.
	r.ParseForm()

	message := r.Form.Get("message")
	if message == "" {
		GetHandler(w, r, sm)
		return
	}

	username := sm.GetString(r.Context(), "username")

	if username != "" {

		global.Messages = append(global.Messages, components.Message{
			User:    username,
			Message: message,
			Time:    time.Now().Format("Mon 15:04:05"),
		})
		fmt.Printf("logged in with username: '%v'\n", username)
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
