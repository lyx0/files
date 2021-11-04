package handlers

import (
	"log"
	"net/http"
)

type Upload struct {
	l *log.Logger
}

func NewUploader(l *log.Logger) *Upload {
	return &Upload{l}
}

func (u *Upload) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("upload")

}
