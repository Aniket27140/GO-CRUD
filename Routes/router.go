package routes

import (
	controller "CRUD/Controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/apii/movies", controller.Getallmymovies).Methods("GET")
	// router.HandleFunc("/movies/{id}", controller.Createmovie).Methods("GET")
	router.HandleFunc("/api/movies", controller.Createmovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controller.MarkWatch).Methods("PUT")
	// router.HandleFunc("/movies/{id}", deletemovie).Methods("DELETE")

	return router
}
