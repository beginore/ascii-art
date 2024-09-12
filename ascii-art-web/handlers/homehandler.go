package handlers

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		Errorhandler(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		Errorhandler(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	files := "templates/ascii-art-web.html"

	ts, err := template.ParseFiles(files)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		Errorhandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return

	}

	err = ts.ExecuteTemplate(w, "ascii-art-web.html", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		Errorhandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
}
