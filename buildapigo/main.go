package main // Declares the main package, the entry point of the Go program

import (
	"encoding/json" // For encoding and decoding JSON data
	"fmt"           // For printing debug messages
	"log"
	"math/rand"
	"net/http" // For creating HTTP server and handling routes
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for course-file
// Course struct defines the schema for a course with JSON tags for API interactions
type Course struct {
	CourseId    string  `json:"courseid"`   // Unique ID for the course
	CourseName  string  `json:"coursename"` // Name of the course
	CoursePrice int     `json:"-"`          // Price of the course
	Author      *Author `json:"author"`     // Pointer to the Author struct (nested object)
}

// Author struct defines the details of the course author
type Author struct {
	Fullname string `json:"fullname"` // Full name of the author
	Website  string `json:"website"`  // Author's website
}

// fake db
// Simulates a database by using an in-memory slice of Course structs
var courses []Course

// middleware, helper
// IsEmpty checks if the course has empty ID and Name (used for validation)
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

// main function - entry point of the program
func main() {
	// For now, it's empty.
	// You can initialize fake data or start your web server here.
	fmt.Println("Course API service is starting...")
	r := mux.NewRouter()
	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN STACK", CoursePrice: 199, Author: &Author{Fullname: "Hitesh Choudhary", Website: "go.dev"}})
	//listen to a port
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers-file

// serveHome handles the root (/) route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello this is me</h1>"))
}

// getAllCourses returns all courses as JSON
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// new function
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request params
	params := mux.Vars(r)

	// loop through courses, find matching id and return response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("no course found")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// Check if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// Decode the JSON body into a Course struct
	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		json.NewEncoder(w).Encode("Error decoding data")
		return
	}

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}
    //check only if title matches then duplicated created or tell exists
	// Generate a unique ID and append to the course list
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)

	// Return the created course as JSON
	json.NewEncoder(w).Encode(course)
}
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// Grab ID from URL
	params := mux.Vars(r)

	// Loop through courses to find the matching ID
	for index, course := range courses {
		if course.CourseId == params["id"] {
			// Remove the old course
			courses = append(courses[:index], courses[index+1:]...)

			// Decode new course data from request body
			var updatedCourse Course
			_ = json.NewDecoder(r.Body).Decode(&updatedCourse)

			// Assign the same ID
			updatedCourse.CourseId = params["id"]

			// Append updated course
			courses = append(courses, updatedCourse)

			// Respond with the updated course
			json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}

	// Respond if no course with given ID is found
	http.Error(w, "Course not found", http.StatusNotFound)
}
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			// Remove course by slicing it out
			courses = append(courses[:index], courses[index+1:]...)
			//send a confirm msg that it is deleted
			break
		}
	}

	// Optionally return confirmation
	json.NewEncoder(w).Encode(map[string]string{"message": "Course deleted successfully"})
}
