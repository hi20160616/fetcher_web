package handler

import (
	"context"
	"log"
	"net/http"
	"strings"

	pb "github.com/hi20160616/fetchnews-api/proto/v1"
	"github.com/hi20160616/fetchnews/config"
	"github.com/hi20160616/fetchnews/internal/data"
	"github.com/hi20160616/fetchnews/internal/pkg/render"
)

var part1 = "list|search"
var part2 = make([]string, 0) // microservice title arranged in config

func init() {
	for _, v := range config.Data.MS {
		part2 = append(part2, v.Title)
	}
}

func validReq(r *http.Request) []string {
	invalidPart1 := func(s string) bool {
		for _, v := range strings.Split(part1, "|") {
			if v == s {
				return false
			}
		}
		return true
	}
	invalidPart2 := func(t string) bool {
		for _, tt := range part2 {
			if t == tt {
				return false
			}
		}
		return true
	}
	parts := strings.Split(r.URL.Path, "/")
	if parts == nil || invalidPart1(parts[1]) || invalidPart2(parts[2]) {
		return nil
	}

	return parts
}

// makeHandler invoke fn after path valided, and arrange args from url to object: `&render.Page{}`
func makeHandler(fn func(http.ResponseWriter, *http.Request, *render.Page)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := &render.Page{}
		parts := validReq(r)
		if len(parts) == 0 {
			http.NotFound(w, r)
			return
		}
		switch {
		case len(parts) == 3:
			p.Title = parts[2]
		case len(parts) == 4:
			p.Title, p.Data = parts[2], parts[3]
		}
		fn(w, r, p)
	}
}

// GetHandler is a handler merger and a router for mutipl handler
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
	mux.HandleFunc("/list/", makeHandler(listArticlesHandler))
	return mux
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	render.Derive(w, "home", &render.Page{Title: "Home", Data: config.Data.MS})
}

func listArticlesHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	ds, err := data.ListArticles(context.Background(), &pb.ListArticlesRequest{}, p.Title)
	if err != nil {
		log.Println(err)
	}
	p.Data = ds.Articles
	render.Derive(w, "list", p)
}

func getArticleHandler(w http.ResponseWriter, r *http.Request, p *render.Page) {
	// `/list/bbc-1_c/s123adfasdf`
	// p.Data arranged by makeHandler: validReq
	id, ok := p.Data.(string)
	if !ok {
		http.Error(w, "data type assertion error", http.StatusInternalServerError)
	}
	a, err := data.GetArticle(context.Background(), &pb.GetArticleRequest{Id: id}, p.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	p.Data = a
	render.Derive(w, "content", p)
}
