package main

import (
	"context"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var tpl *template.Template

func main() {
	r := mux.NewRouter()

	// Template
	templ, err := tpl.ParseFiles("index.html")
	if err != nil {
		log.Error(err)
	}

	templ.Execute(w, "index.html", nil)

	//
	http.HandleFunc("/", UploadHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	srv.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)

}

func UploadHandler(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	file, header, err := r.FormFile("file")
	if err != nil {
		log.Error(err)
	}

	log.Info(file)
	log.Info(header)

}
