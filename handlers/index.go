package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type Index struct {
	l *log.Logger
}

func NewIndex(l *log.Logger) *Index {
	return &Index{l}
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func (u *Index) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(rw, "index.html", nil)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

}
