package main

import (
	"fmt"

	"github.com/Keshav-Agrawal/seperate/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()

	fmt.Println("Welcome to seperate ones application")
	log.Fatal(http.ListenAndServe(":4000", r))
}
