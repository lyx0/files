package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", upload)
	// http.HandleFunc("/upload", upload)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))
	http.ListenAndServe(":8080", nil)
}

// func index(rw http.ResponseWriter, r *http.Request) {
// 	tpl.ExecuteTemplate(rw, "index.gohtml", nil)
// }

// Create random string for the filename
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

func RandStringBytesMask(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return string(b)
}

func upload(rw http.ResponseWriter, r *http.Request) {
	var s string

	if r.Method == http.MethodPost {
		// Open file/header from id "q"
		f, h, err := r.FormFile("q")
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// fmt.Println("\nfile: ", f, "\nheader:", h, "\nerr:", err)

		// Read file into memory
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		// fn = string(bs)

		// Create filename
		rnd := RandStringBytesMask(6)
		// get file extension
		ext := filepath.Ext(h.Filename)
		s = fmt.Sprintf("%s%s", rnd, ext)

		// Store file on disk
		dest, err := os.Create(filepath.Join("./uploads/", s))
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
