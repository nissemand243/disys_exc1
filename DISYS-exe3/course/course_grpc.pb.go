// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package go_course_grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CourseServiceClient is the client API for CourseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CourseServiceClient interface {
	PostCourse(ctx context.Context, in *NewCourse, opts ...grpc.CallOption) (*Course, error)
}

type courseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCourseServiceClient(cc grpc.ClientConnInterface) CourseServiceClient {
	return &courseServiceClient{cc}
}

func (c *courseServiceClient) PostCourse(ctx context.Context, in *NewCourse, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := c.cc.Invoke(ctx, "/couse.CourseService/PostCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourseServiceServer is the server API for CourseService service.
// All implementations must embed UnimplementedCourseServiceServer
// for forward compatibility
type CourseServiceServer interface {
	PostCourse(context.Context, *NewCourse) (*Course, error)
	mustEmbedUnimplementedCourseServiceServer()
}

// UnimplementedCourseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCourseServiceServer struct {
}

func (UnimplementedCourseServiceServer) PostCourse(context.Context, *NewCourse) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostCourse not implemented")
}
func (UnimplementedCourseServiceServer) mustEmbedUnimplementedCourseServiceServer() {}

// UnsafeCourseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourseServiceServer will
// result in compilation errors.
type UnsafeCourseServiceServer interface {
	mustEmbedUnimplementedCourseServiceServer()
}

func RegisterCourseServiceServer(s grpc.ServiceRegistrar, srv CourseServiceServer) {
	s.RegisterService(&CourseService_ServiceDesc, srv)
}

func _CourseService_PostCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewCourse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).PostCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/couse.CourseService/PostCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).PostCourse(ctx, req.(*NewCourse))
	}
	return interceptor(ctx, in, info, handler)
}

// CourseService_ServiceDesc is the grpc.ServiceDesc for CourseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CourseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "couse.CourseService",
	HandlerType: (*CourseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostCourse",
			Handler:    _CourseService_PostCourse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "course/course.proto",
}