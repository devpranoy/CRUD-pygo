package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func read(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
func update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
func delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/create", create)
	router.HandleFunc("/read", read)
	router.HandleFunc("/update", update)
	router.HandleFunc("/delete", delete)
	log.Fatal(http.ListenAndServe(":8080", router))
}
