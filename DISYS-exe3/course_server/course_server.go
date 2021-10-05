package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "example.com/go-course-grpc/course"
	"google.golang.org/grpc"
)

const (
	port = ":8008"
)

type CourseServiceServer struct {
	pb.UnimplementedCourseServiceServer
}

func (s *CourseServiceServer) PostCourse(ctx context.Context, in *pb.NewCourse) (*pb.Course, error) {
	log.Printf("Recieved: %v", in.GetName())
	var course_id int32 = int32(rand.Intn(1000))
	return &pb.Course{Name: in.GetName(), Students: in.GetStudents(), Teachers: in.GetTeachers(), Workload: in.GetWorkload(), Score: 0, CourseId: course_id}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCourseServiceServer(s, &CourseServiceServer{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
