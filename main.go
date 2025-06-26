package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Parse templates
	tmpl, err := template.ParseFiles(
		path.Join("templates", "base.html"),
		path.Join("templates", "home.html"),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute template with data
	data := struct {
		Title   string
		Message string
	}{
		Title:   "My new website with Go",
		Message: "Hello World! I'm deploying my new website with Dash!",
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle routes
	http.HandleFunc("/", handler)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
