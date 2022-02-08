package router

import (
	"github.com/Keshav-Agrawal/seperate/controler"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controler.StartPage).Methods("GET")
	router.HandleFunc("/employees", controler.DisplayAllEmployees).Methods("GET")
	router.HandleFunc("/employee/{id}", controler.DisplaySingleEmployee).Methods("GET")
	router.HandleFunc("/employee", controler.CreateEmployee).Methods("POST")
	router.HandleFunc("/employee/{id}", controler.DeleteEmployee).Methods("DELETE")
	router.HandleFunc("/employee/{id}", controler.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/employee", controler.DeleteAllEmployee).Methods("DELETE")
	return router
}
