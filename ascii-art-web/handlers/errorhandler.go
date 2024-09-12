package handlers

import (
	"ascii-art-web/models"
	"html/template"
	"net/http"
)

func Errorhandler(w http.ResponseWriter, r *http.Request, errorcode int, errormsg string) {
	data := models.Error{
		Errormsg:  errormsg,
		Errorcode: errorcode,
	}
	t, err := template.ParseFiles("templates/error.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
