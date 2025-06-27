package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"time"
)

// Product data structure
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	InStock     bool    `json:"in_stock"`
	CreatedAt   string  `json:"created_at"`
}

// Homepage handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		path.Join("templates", "base.html"),
		path.Join("templates", "home.html"),
	))

	data := struct {
		Title   string
		Message string
	}{
		Title:   "My Go Website",
		Message: "Welcome to my Go web server!",
	}

	tmpl.ExecuteTemplate(w, "base.html", data)
}

// API handler
func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products := []Product{
		{
			ID:          1,
			Name:        "Go Mug",
			Price:       12.99,
			Description: "Official Go programming language mug",
			InStock:     true,
			CreatedAt:   time.Now().Format(time.RFC3339),
		},
		{
			ID:          2,
			Name:        "Go T-Shirt",
			Price:       24.99,
			Description: "Comfortable cotton t-shirt with Go logo",
			InStock:     true,
			CreatedAt:   time.Now().Format(time.RFC3339),
		},
	}

	json.NewEncoder(w).Encode(products)
}

func main() {
	// Serve static files (CSS, JS, images)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/products", apiHandler)

	// Start server
	fmt.Println("Server running on http://localhost:8080")
	fmt.Println("Homepage: http://localhost:8080")
	fmt.Println("API: http://localhost:8080/api/products")
	http.ListenAndServe(":8080", nil)
}
