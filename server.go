package main

import (
	"fmt"
	"net/http"
	"os"

	// "golang.org/x/crypto/bcrypt"
	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// the routers
	router := mux.NewRouter()

	// the handlers
	router.HandleFunc("/", homeHandler)
	router.Path("/robots.txt").Handler(http.FileServer(http.Dir("static")))
	router.Path("/favicon.ico").Handler(http.FileServer(http.Dir("static")))
	router.PathPrefix("/s/").Handler(http.FileServer(http.Dir("static")))

	// the server
	http.Handle("/", router)

	// listen on
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	http.ListenAndServe(":"+port, nil)
}
