package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

/*
router.GET("/courses", getCourses)
	router.GET("/courses/:id", getCourseByID)
	router.POST("/courses", postCourse)
	router.POST("/courses/student/:id/:name", addStudent)
	router.POST("/courses/teacher/:id/:name", addTeacher)
*/
func main() {
	getCourses()
	postCourse()
}

func getCourses() {
	resp, err := http.Get("http://localhost:8080/courses")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
	log.Println()
}

func postCourse() {
	client := &http.Client{}
	var jsonstr = []byte(`{"id": 4, "name": "The Modern Sound of Betty Carter", "teachers":["Betty-Carter","kurt"], "students":["Heino","kent"]}`)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/courses", bytes.NewBuffer(jsonstr))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}
