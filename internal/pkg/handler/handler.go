package handler

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/hi20160616/fetcher_web/internal/pkg/render"
)

var validPath = regexp.MustCompile("^/(home|sites|search)/(.*?)$")

type Handler struct {
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, *render.Page)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, &render.Page{})
	}
}

func GetHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		homeHandler(w, req)
		// fmt.Fprintf(w, "Welcome to the home page!")
	})
	mux.Handle("/s/", http.StripPrefix("/s/", http.FileServer(http.Dir("templates/default"))))
	// mux.HandleFunc("/search/", makeHandler(searchHandler))
	return mux
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// 1. get activities videos by channelIds
	fmt.Fprintf(w, "Hello world!")
}
