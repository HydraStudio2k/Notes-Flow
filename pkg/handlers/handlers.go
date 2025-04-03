package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"handlers/sqlite" // database methods
)

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		filepath.Join("..", "..", "ui", "html", "main_page", "main_page.html"),
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil { // Mandatory practice if you catch a error
		log.Printf("Error while parsing files: %e\n", err) // When logging, to get information, be sure to write where it happened
		http.Error(w, "Internal Server Error", 500)        // If the error occurred in the handler, you can simply report the error on the server side and return code 500
		return
	}
	sqlite.Test_connection()

	err = tmpl.ExecuteTemplate(w, "main_page.html", nil)
	if err != nil {
		log.Printf("Error while executing template: %e\n", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

}
