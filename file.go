package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		_, header, err := r.FormFile("file")
		if err != nil {
			log.Fatal()
		}
		s, _ := header.Open()
		p := filepath.Join("files", header.Filename)
		buf, _ := ioutil.ReadAll(s)
		fmt.Println(buf)
		http.Redirect(w, r, "/"+p, 301)
	} else {
		http.Redirect(w, r, "/", 301)
	}
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if ok, err := path.Match("/data/*.html", r.URL.Path); err != nil || !ok {
			http.NotFound(w, r)
			return
		}
		name := filepath.Join(cwd, "data", filepath.Base(r.URL.Path))

		f, err := os.Open(name)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		defer f.Close()
		io.Copy(w, f)
	})
	http.ListenAndServe(":8080", nil)
}
