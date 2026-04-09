package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) renderPage(page string, data any) ([]byte, error) {
	ts, ok := app.templateCache[page]
	if !ok {
		return nil, fmt.Errorf("the template %s does not exist", page)
	}

	buf := new(bytes.Buffer)

	if err := ts.ExecuteTemplate(buf, "base", data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (app *application) render(w http.ResponseWriter, status int, page string, data any) {
	content, err := app.renderPage(page, data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(status)
	w.Write(content)
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
