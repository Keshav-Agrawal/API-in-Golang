package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func firstpage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("1st page")
	json.NewEncoder(w).Encode("This is my first page")
}
func secondpage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("2nd Page")
	json.NewEncoder(w).Encode("This is my second page")
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", firstpage).Methods("GET")
	r.HandleFunc("/second", secondpage).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}
