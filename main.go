package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lyx0/files/handlers"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	nu := handlers.NewUploader(l)
	ni := handlers.NewIndex(l)

	sm := http.NewServeMux()
	sm.Handle("/upload", nu)
	sm.Handle("/", ni)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Received terminate", sig)

	tc, err := context.WithTimeout(context.Background(), 30*time.Second)
	if err != nil {
		l.Fatal(err)
	}
	s.Shutdown(tc)

	// Fix dumb warning
	_ = tpl
}
