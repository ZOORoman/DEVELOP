package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Go!")
	t, err := template.ParseFiles("templates/index.html")
	// t, err := template.ParseFiles( "templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil) // Чтобы рвботали шаблоны
}

func handleFunc() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3030", nil)
}

func main() {
	handleFunc()
}
