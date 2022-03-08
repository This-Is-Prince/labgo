package main

import (
	"encoding/json"
	"fmt"
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
	Price  string  `json:"price"`
	Author *Author `json:"author"`
}

//
type Author struct {
	FullName string `json:"fullname"`
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

}

// controllers - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by LearnCodeOnline</h1>"))
}

// get all courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(courses)
	if err != nil {
		panic(err)
	}
}

// get one course
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
	json.NewEncoder(w).Encode("No Course found with given id.")
	return
}

// create one course
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		err := json.NewEncoder(w).Encode("Please send some data.")
		checkNilError(err)
		return
	}
	// what about - {}

	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	checkNilError(err)
	if course.IsEmpty() {
		err := json.NewEncoder(w).Encode("No data inside JSON")
		checkNilError(err)
		return
	}

	// generate unique id, string
	// append course into courses
	rand.Seed(time.Now().UnixNano())
	course.Id = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	err = json.NewEncoder(w).Encode(course)
	checkNilError(err)
	return
}

// update one course
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loops, id, remove, add with my ID
	for index, course := range courses {
		if course.Id == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var newCourse Course
			err := json.NewDecoder(r.Body).Decode(&newCourse)
			checkNilError(err)

			newCourse.Id = params["id"]
			courses = append(courses, newCourse)

			err = json.NewEncoder(w).Encode(newCourse)
			checkNilError(err)
			return
		}
	}
	json.NewEncoder(w).Encode(fmt.Sprintf("No course with id `%v`", params["id"]))
	return
}

// Check nil error
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
