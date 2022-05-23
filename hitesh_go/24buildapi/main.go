package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	// return c.Id == "" && c.Name == ""
	return c.Name == ""
}

func main() {
	fmt.Println("Building API")

	r := mux.NewRouter()

	// seeding
	courses = append(courses,
		Course{Id: "2", Name: "ReactJS", Price: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}},
		Course{Id: "4", Name: "MERN Stack", Price: 199, Author: &Author{Fullname: "Hitesh Choudhary", Website: "go.dev"}},
	)

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<h1>Welcome to API by LearnCodeOnline</h1>`))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.Id == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// What about - {}

	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// generate unique id, string
	rand.Seed(time.Now().UnixNano())
	course.Id = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req
	params := mux.Vars(r)

	id, isPresent := params["id"]
	if !isPresent {
		json.NewEncoder(w).Encode("Please send course id")
		return
	}

	// loops, id
	for index, course := range courses {
		if course.Id == id {
			courses = append(courses[:index], courses[index+1:]...)
			var newCourse Course
			err := json.NewDecoder(r.Body).Decode(&newCourse)
			if err != nil {
				json.NewEncoder(w).Encode(err)
				return
			}
			newCourse.Id = id
			courses = append(courses, newCourse)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// TODO: send a response when id is not found
	json.NewEncoder(w).Encode("Course not found with id " + id)
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Context-Type", "application/json")

	params := mux.Vars(r)

	id, isPresent := params["id"]
	if !isPresent {
		json.NewEncoder(w).Encode("Please send the id")
		return
	}
	// loop, id, remove (index, index+1)
	for index, course := range courses {
		if course.Id == id {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found with id " + id)
}
