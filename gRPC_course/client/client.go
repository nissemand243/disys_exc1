package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	serverAddr = "localhost:8008"
)

func main() {
	//var opts []grpc.DialOption
	// Set up a connection to the server.
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := NewCourseProtoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	course, err := client.PostCourse(ctx, &Course{
		CourseId:          1,
		Name:              "DØSYS",
		Workload:          15,
		SatisfactoryScore: 10,
		Teachers:          []string{"Tom"},
		Students:          []string{"Bente"},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Course: %s", course.GetMessage())
}
