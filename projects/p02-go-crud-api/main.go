package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movie

func main() {
	fmt.Println("Golang Basic CRUD API")

	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "12345", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "54321", Title: "Movie Two", Director: &Director{FirstName: "Doe", LastName: "John"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	log.Println("Starting server at port: 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Add("Content-Type", "application/json")

	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100_000_000))
	movies = append(movies, movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Add("Content-Type", "application/json")

	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)

	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			item.Director = movie.Director
			item.Isbn = movie.Isbn
			item.Title = movie.Title
			movies = append(movies[:index], item)
			movies = append(movies, movies[index+1:]...)
			break
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}
