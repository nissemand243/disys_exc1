package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type course struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	Workload      int      `json:"workload"`
	SatisfacScore int      `json:"courseScore"`
	Teacher       []string `json:"teachers"`
	Student       []string `json:"students"`
}

var courses = []course{
	{ID: 1, Name: "BDSA", Teacher: []string{"Paolo", "holger"}, Student: []string{"Jens", "peter"}},
	{ID: 2, Name: "Distributed Systems", Teacher: []string{"Marco", "Laura-langpat"}, Student: []string{"Henning", "GOKKE-JOKKE"}},
}

func main() {
	router := gin.Default()
	router.GET("/courses", getCourses)
	router.GET("/courses/:id", getCourseByID)
	router.POST("/courses", postCourse)
	router.POST("/courses/student/:id/:name", addStudent)
	router.POST("/courses/teacher/:id/:name", addTeacher)
	router.DELETE("/courses/student/:id/:name", deleteStudent)
	router.DELETE("/courses/teacher/:id/:name", deleteTeacher)
	router.PUT("/courses/score/:id/:score", updateScore)
	router.PUT("/courses/workload/:id/:workload", updateWorkload)

	router.Run("localhost:8080")
}

func getCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

// postCourse adds a course from JSON received in the request body.
func postCourse(c *gin.Context) {
	var newCourse course

	// Call BindJSON to bind the received JSON to
	// newCourse.
	if err := c.BindJSON(&newCourse); err != nil {
		return
	}

	//check if course allready exists in the server
	if check(newCourse) {

		c.IndentedJSON(http.StatusAlreadyReported, "Course allready exists")
	} else {
		// Add the new course to the slice.
		courses = append(courses, newCourse)
		c.IndentedJSON(http.StatusCreated, newCourse)
	}

}

func check(newCourse course) bool {
	var checkBool bool
	for _, a := range courses {
		if a.ID == newCourse.ID && a.Name == newCourse.Name {
			checkBool = true
			break
		} else {
			checkBool = false
		}
	}
	return checkBool
}

// getCourseByID locates the course whose ID value matches the id
// parameter sent by the client, then returns that course as a response.
func getCourseByID(c *gin.Context) {
	id := c.Param("id")
	index, _ := strconv.Atoi(id)
	// Loop over the list of courses, looking for
	// a course whose ID value matches the parameter.
	for _, a := range courses {
		if a.ID == index {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "course not found"})
}

func deleteStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	name := c.Params.ByName("name")
	index, _ := strconv.Atoi(id)
	var newArr []string
	var deleted string

	for i := 0; i < len(courses); i++ {
		if courses[i].ID == index {
			for j := 0; j < len(courses[i].Student); j++ {
				if courses[i].Student[j] == name {
					deleted = courses[i].Student[j]
				} else {
					newArr = append(newArr, courses[i].Student[j])
				}
			}
			courses[i].Student = newArr
		}
	}

	c.IndentedJSON(http.StatusOK, deleted)
}

func deleteTeacher(c *gin.Context) {
	id := c.Params.ByName("id")
	name := c.Params.ByName("name")
	index, _ := strconv.Atoi(id)
	var newArr []string
	var deleted string

	for i := 0; i < len(courses); i++ {
		if courses[i].ID == index {
			for j := 0; j < len(courses[i].Teacher); j++ {
				if courses[i].Teacher[j] == name {
					deleted = courses[i].Teacher[j]
				} else {
					newArr = append(newArr, courses[i].Teacher[j])
				}
			}
			courses[i].Teacher = newArr
		}
	}

	c.IndentedJSON(http.StatusOK, deleted)
}

func addStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	name := c.Params.ByName("name")
	index, _ := strconv.Atoi(id)

	for i := 0; i < len(courses); i++ {
		if courses[i].ID == index {
			courses[i].Student = append(courses[i].Student, name)
		}
	}

	c.IndentedJSON(http.StatusOK, name)
}

func addTeacher(c *gin.Context) {
	id := c.Params.ByName("id")
	name := c.Params.ByName("name")
	index, _ := strconv.Atoi(id)

	for i := 0; i < len(courses); i++ {
		if courses[i].ID == index {
			courses[i].Teacher = append(courses[i].Teacher, name)
		}
	}
	c.IndentedJSON(http.StatusOK, name)
}

func updateScore(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	score, _ := strconv.Atoi(c.Params.ByName("score"))
	for i := 0; i < len(courses); i++ {
		if courses[i].ID == id {
			courses[i].SatisfacScore = score
			c.IndentedJSON(http.StatusOK, courses[i])
		}
	}

}

func updateWorkload(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	workload, _ := strconv.Atoi(c.Params.ByName("workload"))

	for i := 0; i < len(courses); i++ {
		if courses[i].ID == id {
			courses[i].Workload = workload
			c.IndentedJSON(http.StatusOK, courses[i])
		}
	}
}

// GET ALL COURSES USING CMD
// curl http://localhost:8080/courses

// GET COURSE BY ID USING CMD
// curl http://localhost8080/courses/2

//adding into slice via CMD ...... ONLY CMD!!!!!!! - with arrays
// curl http://localhost:8080/courses --include --Header "Content-Type:application/json" --request "POST" --data "{\"id\": 4, \"name\": \"The Modern Sound of Betty Carter\", \"teachers\":[\"Betty-Carter\",\"kurt\"], \"students\":[\"Heino\",\"kent\"]}"

//DELETE STUDENT USIGN CMD
// curl http://localhost:8080/courses/student/2/Henning --Header "Content-Type:application/json" --request "DELETE"

//DELETE TEACHER USING CMD
// curl http://localhost:8080/courses/teacher/2/Marco --Header "Content-Type:application/json" --request "DELETE"

// ENROLL STUDENT TO COURSE USING CMD
// curl http://localhost:8080/courses/student/2/HOLGER --Header "Content-Type:application/json" --request "POST"

// ENROLL TEACHER TO COUSE USING CMD
// curl http://localhost:8080/courses/teacher/2/Jytte --Header "Content-Type:application/json" --request "POST"

// UPDATE COURSE SCORE USING CMD
// curl http://localhost:8080/courses/score/2/10 --Header "Content-Type:application/json" --request "PUT"

// UPDATE COURSE WORKLOAD USING CMD
// curl http://localhost:8080/courses/workload/2/100 --Header "Content-Type:application/json" --request "PUT"
