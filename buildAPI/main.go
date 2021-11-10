package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// simulated DB
var courses []Course

// middleware, helper
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	r := mux.NewRouter()
	courses = append(courses, Course{CourseId: "2", CourseName: "Golang", CoursePrice: 200, Author: &Author{Fullname: "Subhra", Website: "youtube.com"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "C++", CoursePrice: 250, Author: &Author{Fullname: "HoneyBadger", Website: "twitch.com"}})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}

// controllers

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello world</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)
	var queriedId string = params["id"]

	for _, course := range courses {
		if course.CourseId == queriedId {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with the id = " + queriedId)
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var course Course
	if r.Body == nil {
		json.NewEncoder(w).Encode("No data found")
		return
	}

	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("empty data")
		return
	}

	// generate unique id
	// append course into courses

	rand.Seed(time.Now().Unix())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("No data found")
		return
	}

	// get id from req
	params := mux.Vars(r)

	// loop, id, remove, add with my ID

	for i, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:i], courses[i+1:]...)
			var updatedCourse Course
			json.NewDecoder(r.Body).Decode(&updatedCourse)
			if updatedCourse.IsEmpty() {
				json.NewEncoder(w).Encode("empty data")
				return
			}
			updatedCourse.CourseId = params["id"]
			courses = append(courses, updatedCourse)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("no course found with the id = " + params["id"])
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("No data found")
		return
	}

	// get id from req
	params := mux.Vars(r)

	// loop, id, remove, add with my ID

	for i, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:i], courses[i+1:]...)
			json.NewEncoder(w).Encode("course deleted")
			return
		}
	}

	json.NewEncoder(w).Encode("no course found with the id = " + params["id"])
}
