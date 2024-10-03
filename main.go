package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func main() {
	// Parse HTML templates
	tmplPath := filepath.Join("templates", "form.html")
	greetPath := filepath.Join("templates", "greet.html")
	templates := template.Must(template.ParseFiles(tmplPath, greetPath))

	// Serve the HTML form at the root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "form.html", nil)
	})

	// Handle the form submission and show the greeting
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			name := r.FormValue("name")
			templates.ExecuteTemplate(w, "greet.html", name)
		}
	})

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
