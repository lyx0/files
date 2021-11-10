package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(rw http.ResponseWriter, r *http.Request) {
	var s string

	if r.Method == http.MethodPost {
		// Open file/header from id "q"
		f, h, err := r.FormFile("q")
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		fmt.Println("\nfile: ", "\nheader:", h, "\nerr:", err)

		// Read file into memory
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)

		// Store file on disk
		dest, err := os.Create(filepath.Join("./uploads/", h.Filename))
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dest.Close()

		_, err = dest.Write(bs)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl.ExecuteTemplate(rw, "index.gohtml", s)

}
