package main

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type projectPageData struct {
	ProjectID    int
	ProjectTitle string
}

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

	projects := map[int]string{
		1: "Project One",
		2: "Project Two",
	}

	title, ok := projects[id]
	if !ok {
		app.notFound(w)
		return
	}

	data := projectPageData{
		ProjectID:    id,
		ProjectTitle: title,
	}

	app.render(w, http.StatusOK, "view.tmpl", data)
}
