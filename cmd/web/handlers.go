package main

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := homePageData{
		Projects: portfolioProjects,
	}

	app.render(w, http.StatusOK, "home.tmpl.html", data)
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	app.render(w, http.StatusOK, "contact.tmpl.html", nil)
}

func (app *application) projectView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	project, ok := getProjectByID(id)
	if !ok {
		app.notFound(w)
		return
	}

	app.render(w, http.StatusOK, "view.tmpl.html", project)
}
