package main

import (
	"net/http"
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

func (app *application) inclens(w http.ResponseWriter, r *http.Request) {
	app.render(w, http.StatusOK, "inclens.tmpl.html", inclensPage)
}

func (app *application) monkeyInCPP(w http.ResponseWriter, r *http.Request) {
	app.render(w, http.StatusOK, "monkey-in-cpp.tmpl.html", monkeyInCPPPage)
}
