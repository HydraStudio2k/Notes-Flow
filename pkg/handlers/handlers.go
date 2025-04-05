package handlers

import (
	"handlers/sqlite" // database methods
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// Main Page & NotFound Handlers

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		filepath.Join("..", "..", "ui", "html", "main_page", "main_page.html"),
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Printf("Error while parsing files: %e\n", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Error while executing template: %e\n", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

// Register Handlers

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		filepath.Join("..", "..", "ui", "html", "register", "form.html"),
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Printf("Error while parsing files: %e\n", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Error while executing template: %e\n", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func PostformRegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	username := r.FormValue("username")
	password := r.FormValue("password")

	result_status := sqlite.CreateUser(name, username, password) // If the result is false, it means that a user with this username already exists and you should re-register using a different username.

	if result_status {
		files := []string{
			filepath.Join("..", "..", "ui", "html", "register", "successful_postform.html"),
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			log.Printf("Error while parsing files: %e\n", err)
			http.Error(w, "Internal Server Error", 500)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Printf("Error while executing template: %e\n", err)
			http.Error(w, "Internal Server Error", 500)
			return
		}
	} else {
		files := []string{
			filepath.Join("..", "..", "ui", "html", "register", "error_postform.html"),
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			log.Printf("Error while parsing files: %e\n", err)
			http.Error(w, "Internal Server Error", 500)
			return
		}

		tmpl.Execute(w, nil)
		if err != nil {
			log.Printf("Error while executing template: %e\n", err)
			http.Error(w, "Internal Server Error", 500)
			return
		}
	}
}
