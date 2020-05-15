package main

import (
	"politics/go/server"
	"text/template"
	"log"
	"net/http"
)

func front(s *server.Server, w http.ResponseWriter, r *http.Request) {
	nt := template.New("")//.Funcs(funcs)

	t, err := nt.ParseGlob(s.Paths.Root + "/html/*.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = t.ExecuteTemplate(w, "front", s.Recent)
	if err != nil {
		log.Println(err)
	}
}

