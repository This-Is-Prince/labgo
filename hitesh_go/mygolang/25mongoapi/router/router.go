package router

import (
	"github.com/This-Is-Prince/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	r.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	r.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	r.HandleFunc("/api/deleteallmovie", controller.DeleteAllMovies).Methods("DELETE")
	return r
}
