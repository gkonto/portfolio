package main

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, http.StatusOK, "home.tmpl", nil)
}

func (app *application) projectView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	data := struct {
		ProjectID int
	}{
		ProjectID: id,
	}

	app.render(w, http.StatusOK, "view.tmpl", data)
}
