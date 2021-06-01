package handler

import (
	"context"
	"log"
	"net/http"
	"regexp"

	pb "github.com/hi20160616/fetchnews-api/proto/v1"
	"github.com/hi20160616/fetchnews/internal/data"
	"github.com/hi20160616/fetchnews/internal/pkg/render"
)

var validPath = regexp.MustCompile("^/(list|search)/(.*?)$")

type Handler struct {
}

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
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
	// for static resource request
	// mux.Handle("/s/", http.StripPrefix("/s/", http.FileServer(http.Dir("templates/default"))))
	mux.HandleFunc("/list/", makeHandler(listHandler))
	return mux
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	render.Derive(w, "home", &render.Page{Title: "Home", Data: data.NewsSites()})
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	ds, err := data.List(context.Background(), &pb.ListArticlesRequest{})
	if err != nil {
		log.Println(err)
	}

	render.Derive(w, "list", &render.Page{Title: "List", Data: ds})
}
