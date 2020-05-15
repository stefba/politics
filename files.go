package main

import (
	"regexp"
	"fmt"
	"log"
	"net/http"
	"politics/go/entry"
	"politics/go/entry/helper"
	"politics/go/server"
	"strings"
	"path/filepath"
	"time"
)

func files(s *server.Server, w http.ResponseWriter, r *http.Request) {
	e, err := getEntry(s.Files, r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		log.Println(err)
		return
	}
	serveSingleBlob(w, r, e)
}

func serveSingleBlob(w http.ResponseWriter, r *http.Request, e entry.Entry) error {
	blob, ok := e.(entry.Blob)
	if !ok {
		return fmt.Errorf("File to serve (%v) is no blob.", e.File().Name())
	}
	serveStatic(w, r, blob.Location(""))
	return nil
}

func getEntry(files entry.Entries, path string) (entry.Entry, error) {
	hash, err := getHash(path)
	if err != nil {
		return nil, err
	}
	id, err := helper.ParseHash(hash)
	if err != nil {
		return nil, err
	}
	for _, e := range files {
		if e.Id() == id {
			return e, nil
		}
	}
	return nil, fmt.Errorf("getEntry: Id %v (%v) not found.", id, helper.ToTimestamp(id))
}

func getHash(path string) (string, error) {
	p, err := validPath(path)
	if err != nil {
		return "", err
	}
	rel := p[len("/files/"):]
	i := strings.Index(rel, ".")
	if i < 1 {
		return "", fmt.Errorf("invalid hash")
	}
	return rel[:i], nil
}

var valid = regexp.MustCompile(`^\/[0-9a-z+-_.\/]*$`)

func validPath(foreign string) (string, error) {
	if valid.MatchString(foreign) {
		return foreign, nil
	}
	return "", fmt.Errorf("invalid Path: %v", foreign)
}

func serveStatic(w http.ResponseWriter, r *http.Request, p string) {
	if filepath.Ext(p) == ".vtt" {
		w.Header().Set("Content-Type", "text/vtt")
	}
	w.Header().Set("Expires", time.Now().AddDate(0, 3, 0).Format(time.RFC1123))
	http.ServeFile(w, r, p)
}


