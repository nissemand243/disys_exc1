package main

import (
	"context"
	"log"
	"time"

	pb "example.com/go-course-grpc/course"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8008"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCourseServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var new_courses = make(map[string]int32)
	new_courses["bdsd"] = 15
	new_courses["daba"] = 7
	for name, workload := range new_courses {
		r, err := c.PostCourse(ctx, &pb.NewCourse{Name: name, Workload: workload})
		if err != nil {
			log.Fatalf("Could not create course: %v", err)
		}
		log.Printf("Course Details:\n"+
			"NAME: %s\n"+
			"WORKLOAD: %d\n"+
			"ID: %d", r.GetName(), r.GetWorkload(), r.GetCourseId())
	}
}
