package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type Upload struct {
	l *log.Logger
}

func NewUploader(l *log.Logger) *Upload {
	return &Upload{l}
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func (u *Upload) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(rw, "upload.html", nil)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

}
