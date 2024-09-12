package handlers

import (
	"ascii-art-web/ascii-art/intern"
	"html/template"
	"net/http"
)

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/ascii-art-web.html")
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		Errorhandler(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		Errorhandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	input := r.FormValue("input")
	if err := intern.IsASCII(input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		Errorhandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	font := r.FormValue("font")
	fonts, err := intern.MakeMap(font)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		Errorhandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	print := intern.PrintArt(input, fonts)
	if input == "" {
		w.WriteHeader(http.StatusBadRequest)
		Errorhandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	t.Execute(w, print)
}
