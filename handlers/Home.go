package handlers

import (
	"kevcoxe/templ-test/components"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

type GlobalState struct {
	Count int
}

var global GlobalState

func GetHandler(w http.ResponseWriter, r *http.Request, sm *scs.SessionManager) {
	userCount := sm.GetInt(r.Context(), "count")
	component := components.HomePage(global.Count, userCount)
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
