package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct {
	Name               string   `json"Name"`
	UID                string   `json:"UID"`
	ECTS               float64  `json:"ECTS"`
	StatisfactoryScore float64  `json:"StatisfactoryScore"`
	TeacherID          []string `json"TeacherID"`
	StudentID          []string `json"StudentID"`
}

var Courses = []Course{
	{UID: "1", Name: "BDSA", TeacherID: []string{"Palo2", "Rasmus34"}, StudentID: []string{"Mads123213", "Alexander7374", "Anton8888", "Jack999"}},
	{UID: "2", Name: "Distributed Systems", TeacherID: []string{"Marco1"}, StudentID: []string{"Mads123213", "Alexander7374", "Anton8888", "Jack999", "Henning69"}},
}

func main() {

	handleRequest()

}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", TestRouter).Methods("GET")
	router.HandleFunc("/Course", CreateCourse).Methods("POST")
	router.HandleFunc("/Course", GetCourses).Methods("GET")
	router.HandleFunc("/Course/Teacher/{id}/{TeacherId}", AddTeacher).Methods("POST")
	router.HandleFunc("/Course/Teacher/{id}/{TeacherId}", DeleteTeacher).Methods("DELETE")
	router.HandleFunc("/Course/Student/{id}/{StudentId}", AddStudent).Methods("POST")
	router.HandleFunc("/Course/Student/{id}/{StudentId}", DeleteStudent).Methods("DELETE")
	router.HandleFunc("/Course/ECTS/{id}/{ECTS}", UpdateECTS).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", router))

}
func courseSevers() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", TestRouter).Methods("GET")
	router.HandleFunc("/Course", CreateCourse).Methods("POST")
	router.HandleFunc("/Course", GetCourses).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))

}
func teacherServers() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/Course/Teacher/{id}/{TeacherId}", AddTeacher).Methods("POST")
	router.HandleFunc("/Course/Teacher/{id}/{TeacherId}", DeleteTeacher).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8001", router))
}
func studentServers() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/Course/Student/{id}/{StudentId}", AddStudent).Methods("POST")
	router.HandleFunc("/Course/Student/{id}/{StudentId}", DeleteStudent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8002", router))
}
func otherServers() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/Course/ECTS/{id}/{ECTS}", UpdateECTS).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8003", router))
}

func TestRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Endpoint called: TestRouter()")
}
func CreateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Function called: CreateCourse()")

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	Courses = append(Courses, course)

	json.NewEncoder(w).Encode(course)
}
func GetCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Function called: GetCourses()")
	json.NewEncoder(w).Encode(Courses)

}
func AddTeacher(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Function called: AddTeacher()\n")
	fmt.Fprint(w, "Not implementet Hello world")
}

func DeleteTeacher(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Function called: DeleteTeacher()\n")
	fmt.Fprint(w, "Not implementet Hello world")
}
func AddStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Function called: AddStudent()\n")
	fmt.Fprint(w, "Not implementet Hello world")
}
func DeleteStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Function called: DeleteStudent()\n")
	fmt.Fprint(w, "Not implementet Hello world")
}
func UpdateECTS(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Function called: UpdateECTS()\n")
	fmt.Fprint(w, "Not implementet Hello world")
}
