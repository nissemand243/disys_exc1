// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative course/course.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.0
// source: course/course.proto

package course

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Course struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId          int32    `protobuf:"varint,1,opt,name=CourseId,proto3" json:"CourseId,omitempty"`
	Name              string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Workload          int32    `protobuf:"varint,3,opt,name=Workload,proto3" json:"Workload,omitempty"`
	SatisfactoryScore int32    `protobuf:"varint,4,opt,name=SatisfactoryScore,proto3" json:"SatisfactoryScore,omitempty"`
	Teachers          []string `protobuf:"bytes,5,rep,name=Teachers,proto3" json:"Teachers,omitempty"`
	Students          []string `protobuf:"bytes,6,rep,name=Students,proto3" json:"Students,omitempty"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_course_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_course_course_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_course_course_proto_rawDescGZIP(), []int{0}
}

func (x *Course) GetCourseId() int32 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *Course) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Course) GetWorkload() int32 {
	if x != nil {
		return x.Workload
	}
	return 0
}

func (x *Course) GetSatisfactoryScore() int32 {
	if x != nil {
		return x.SatisfactoryScore
	}
	return 0
}

func (x *Course) GetTeachers() []string {
	if x != nil {
		return x.Teachers
	}
	return nil
}

func (x *Course) GetStudents() []string {
	if x != nil {
		return x.Students
	}
	return nil
}

type GetCourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCourseRequest) Reset() {
	*x = GetCourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_course_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCourseRequest) ProtoMessage() {}

func (x *GetCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_course_course_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCourseRequest.ProtoReflect.Descriptor instead.
func (*GetCourseRequest) Descriptor() ([]byte, []int) {
	return file_course_course_proto_rawDescGZIP(), []int{1}
}

type PostCourseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *PostCourseReply) Reset() {
	*x = PostCourseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_course_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostCourseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostCourseReply) ProtoMessage() {}

func (x *PostCourseReply) ProtoReflect() protoreflect.Message {
	mi := &file_course_course_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostCourseReply.ProtoReflect.Descriptor instead.
func (*PostCourseReply) Descriptor() ([]byte, []int) {
	return file_course_course_proto_rawDescGZIP(), []int{2}
}

func (x *PostCourseReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetCourseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *GetCourseReply) Reset() {
	*x = GetCourseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_course_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCourseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCourseReply) ProtoMessage() {}

func (x *GetCourseReply) ProtoReflect() protoreflect.Message {
	mi := &file_course_course_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCourseReply.ProtoReflect.Descriptor instead.
func (*GetCourseReply) Descriptor() ([]byte, []int) {
	return file_course_course_proto_rawDescGZIP(), []int{3}
}

func (x *GetCourseReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_course_course_proto protoreflect.FileDescriptor

var file_course_course_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0xba, 0x01,
	0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x57, 0x6f, 0x72, 0x6b,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x57, 0x6f, 0x72, 0x6b,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x53, 0x61, 0x74, 0x69, 0x73, 0x66, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x11, 0x53, 0x61, 0x74, 0x69, 0x73, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2b,
	0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x87, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x37, 0x0a, 0x0a, 0x50, 0x6f, 0x73, 0x74, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x0e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x3f, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x18, 0x2e,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2e, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_course_course_proto_rawDescOnce sync.Once
	file_course_course_proto_rawDescData = file_course_course_proto_rawDesc
)

func file_course_course_proto_rawDescGZIP() []byte {
	file_course_course_proto_rawDescOnce.Do(func() {
		file_course_course_proto_rawDescData = protoimpl.X.CompressGZIP(file_course_course_proto_rawDescData)
	})
	return file_course_course_proto_rawDescData
}

var file_course_course_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_course_course_proto_goTypes = []interface{}{
	(*Course)(nil),           // 0: course.Course
	(*GetCourseRequest)(nil), // 1: course.GetCourseRequest
	(*PostCourseReply)(nil),  // 2: course.PostCourseReply
	(*GetCourseReply)(nil),   // 3: course.GetCourseReply
}
var file_course_course_proto_depIdxs = []int32{
	0, // 0: course.CourseProto.PostCourse:input_type -> course.Course
	1, // 1: course.CourseProto.GetCourse:input_type -> course.GetCourseRequest
	2, // 2: course.CourseProto.PostCourse:output_type -> course.PostCourseReply
	3, // 3: course.CourseProto.GetCourse:output_type -> course.GetCourseReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_course_course_proto_init() }
func file_course_course_proto_init() {
	if File_course_course_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_course_course_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Course); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_course_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCourseRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_course_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostCourseReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_course_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCourseReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_course_course_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_course_course_proto_goTypes,
		DependencyIndexes: file_course_course_proto_depIdxs,
		MessageInfos:      file_course_course_proto_msgTypes,
	}.Build()
	File_course_course_proto = out.File
	file_course_course_proto_rawDesc = nil
	file_course_course_proto_goTypes = nil
	file_course_course_proto_depIdxs = nil
}
