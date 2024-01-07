package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello, World!")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("main.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "Film 1", Director: "Director 1"},
				{Title: "Film 2", Director: "Director 2"},
				{Title: "Film 3", Director: "Director 3"},
			},
		}
		tmpl.Execute(w, films)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")

		tmpl := template.Must(template.ParseFiles("main.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/film", h2)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
