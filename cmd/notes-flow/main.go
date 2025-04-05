package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"main/handlers"

	"github.com/gorilla/mux"
)

var (
	PORT string = ":8080"
)

func main() {
	if len(os.Args) == 2 {
		PORT = ":" + os.Args[1] // Setting the port at startup
	}
	log.Printf("PORT %s\n", PORT)

	router := mux.NewRouter()
	router.HandleFunc("/main", handlers.MainPageHandler).Methods("GET")  // main page handler
	router.HandleFunc("/main/", handlers.MainPageHandler).Methods("GET") // main page handler

	router.HandleFunc("/user/register/", handlers.RegisterUserHandler).Methods("GET") // register user form handler
	router.HandleFunc("/user/get_uuid/", handlers.GetUSerUUIDHandler).Methods("GET")  // get user uuid form handler

	router.HandleFunc("/user/register/postform", handlers.PostformRegisterUserHandler).Methods("POST") // register user postform handler
	router.HandleFunc("/user/get_uuid/postform", handlers.GetUserUUIDPostformHandler).Methods("POST")  // get user uuid postform handler

	http.Handle("/", router)                                                                                           // all requests are processed by the router
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir(filepath.Join("..", "..", "ui", "css"))))) // we indicate where the statistical files (css) are located

	log.Println("Starting the server...\n---\nURLs:\n\t1. /main/ & /main - Main Page\n\t2. /user/register/ - Register New User\n\t3. /user/get_uuid/ - Get User UUID")

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Printf("Error while listening and serving server: %e\n", err)
		panic(err)
	}
}
