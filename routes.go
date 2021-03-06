package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"politics/go/server"
	"path/filepath"
	"log"
)

func Router(s *server.Server) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	if s.Flags.Reload {
		r.Use(makeMiddleware(s, reloadMiddle))
	}

	r.HandleFunc("/", makeHandler(s, front))
	r.HandleFunc("/rl/", makeHandler(s, reload))
	r.PathPrefix("/files/").HandlerFunc(makeHandler(s, files))
	r.PathPrefix("/static/").HandlerFunc(makeHandler(s, static))

	return r
}


func makeHandler(s *server.Server, fn func(*server.Server, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(s, w, r)
	}
}

func makeMiddleware(s *server.Server, fn func(*server.Server, http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return fn(s, next)
	}
}

func reloadMiddle(s *server.Server, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if filepath.Ext(r.URL.Path) == "" {
			err := s.Load()
			if err != nil {
				log.Println(err)
			}
		}
		next.ServeHTTP(w, r)
	})
}

func reload(s *server.Server, w http.ResponseWriter, r *http.Request) {
	err := s.Load()
	if err != nil {
		fmt.Fprint(w, err)
	}
	fmt.Fprint(w, "ok")
}
