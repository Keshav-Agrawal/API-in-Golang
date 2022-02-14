package router

import (
	"github.com/Keshav-Agrawal/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/tasks", controller.GetMyAllTask).Methods("GET")
	router.HandleFunc("/api/task", controller.CreateTask).Methods("POST")
	router.HandleFunc("/api/task/{id}", controller.MarkAsDone).Methods("PUT")
	router.HandleFunc("/api/task/{id}", controller.DeleteATask).Methods("DELETE")
	router.HandleFunc("/api/task", controller.DeleteAllTask).Methods("DELETE")

	return router
}
