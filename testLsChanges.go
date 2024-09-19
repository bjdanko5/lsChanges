package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"rsc.io/quote"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println(quote.Go())
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("testLsChanges.html"))
		films := map[string][]Film{ //ключ string - название таблицы, []Film - набор данных
			"Films": {
				{Title: "The Matrix", Director: "The Wachowskis"},
				{Title: "The Matrix Reloaded", Director: "The Wachowskis"},
				{Title: "The Matrix Revolutions", Director: "The Wachowskis"},
			},
		}
		tmpl.Execute(w, films)
		/* 	io.WriteString(w, "Hello, World!")
		io.WriteString(w, r.Method) */
	}
	h2 := func(w http.ResponseWriter, r *http.Request) {
		log.Print("HTMX request recieved")
		log.Print(r.Header.Get("HX-Request"))
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)

	}
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
