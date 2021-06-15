package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
	"time"

	"github.com/hi20160616/fetchnews/configs"
	"github.com/yuin/goldmark"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Page struct {
	Title string
	Data  interface{}
}

var templates = template.New("")

func init() {
	templates.Funcs(template.FuncMap{
		"summary":   summary,
		"smartTime": smartTime,
		"smartDate": smartDate,
		"markdown":  markdown,
	})
	// tmplPath := filepath.Join("../../../templates", "default") // for test
	tmplPath := filepath.Join(configs.Data.WebServer.Tmpl, "default")
	pattern := filepath.Join(tmplPath, "*.html")
	templates = template.Must(templates.ParseGlob(pattern))
}

func Derive(w http.ResponseWriter, tmpl string, p *Page) {
	if err := templates.ExecuteTemplate(w, tmpl+".html", p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("err template: %s.html\n\terror: %#v", tmpl, err)
	}
}

func summary(des string) string {
	dRune := []rune(des)
	if len(dRune) <= 300 {
		return des
	}
	return string(dRune[:300])
}

func parseWithZone(t time.Time) time.Time {
	loc := time.FixedZone("UTC", 8*60*60)
	return t.In(loc)

}

func smartTime(t *timestamppb.Timestamp, site string) string {
	if site == "cna" || site == "dw" || site == "kabar" || site == "ucpnz" {
		return t.AsTime().Format("15:04")
	}
	return parseWithZone(t.AsTime()).Format("15:04")
}

func smartDate(t *timestamppb.Timestamp, site string) string {
	if site == "cna" || site == "dw" || site == "kabar" || site == "ucpnz" {
		return t.AsTime().Format("01.02")
	}
	return parseWithZone(t.AsTime()).Format("01.02")
}

func markdown(in string) (string, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(in), &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}
