// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: pb/exercise.proto

package pb

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

type Exercise struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Exercise *NewExercise `protobuf:"bytes,2,opt,name=exercise,proto3" json:"exercise,omitempty"`
}

func (x *Exercise) Reset() {
	*x = Exercise{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_exercise_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Exercise) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Exercise) ProtoMessage() {}

func (x *Exercise) ProtoReflect() protoreflect.Message {
	mi := &file_pb_exercise_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Exercise.ProtoReflect.Descriptor instead.
func (*Exercise) Descriptor() ([]byte, []int) {
	return file_pb_exercise_proto_rawDescGZIP(), []int{0}
}

func (x *Exercise) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Exercise) GetExercise() *NewExercise {
	if x != nil {
		return x.Exercise
	}
	return nil
}

type NewExercise struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	MuscleGroup string `protobuf:"bytes,3,opt,name=muscle_group,json=muscleGroup,proto3" json:"muscle_group,omitempty"`
	UserId      int32  `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *NewExercise) Reset() {
	*x = NewExercise{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_exercise_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewExercise) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewExercise) ProtoMessage() {}

func (x *NewExercise) ProtoReflect() protoreflect.Message {
	mi := &file_pb_exercise_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewExercise.ProtoReflect.Descriptor instead.
func (*NewExercise) Descriptor() ([]byte, []int) {
	return file_pb_exercise_proto_rawDescGZIP(), []int{1}
}

func (x *NewExercise) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewExercise) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NewExercise) GetMuscleGroup() string {
	if x != nil {
		return x.MuscleGroup
	}
	return ""
}

func (x *NewExercise) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_pb_exercise_proto protoreflect.FileDescriptor

var file_pb_exercise_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x62, 0x2f, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x47, 0x0a, 0x08, 0x45, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x08, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x65, 0x77, 0x45, 0x78,
	0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x08, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65,
	0x22, 0x7f, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x75, 0x73, 0x63, 0x6c, 0x65, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x75, 0x73,
	0x63, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x32, 0x42, 0x0a, 0x0f, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78,
	0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x65, 0x77, 0x45,
	0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x78, 0x65,
	0x72, 0x63, 0x69, 0x73, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x63, 0x68, 0x61, 0x65, 0x6c, 0x2d, 0x6d, 0x2d, 0x74, 0x72,
	0x75, 0x6f, 0x6e, 0x67, 0x2f, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2d, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_exercise_proto_rawDescOnce sync.Once
	file_pb_exercise_proto_rawDescData = file_pb_exercise_proto_rawDesc
)

func file_pb_exercise_proto_rawDescGZIP() []byte {
	file_pb_exercise_proto_rawDescOnce.Do(func() {
		file_pb_exercise_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_exercise_proto_rawDescData)
	})
	return file_pb_exercise_proto_rawDescData
}

var file_pb_exercise_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_exercise_proto_goTypes = []interface{}{
	(*Exercise)(nil),    // 0: pb.Exercise
	(*NewExercise)(nil), // 1: pb.NewExercise
}
var file_pb_exercise_proto_depIdxs = []int32{
	1, // 0: pb.Exercise.exercise:type_name -> pb.NewExercise
	1, // 1: pb.ExerciseService.CreateExercise:input_type -> pb.NewExercise
	0, // 2: pb.ExerciseService.CreateExercise:output_type -> pb.Exercise
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_exercise_proto_init() }
func file_pb_exercise_proto_init() {
	if File_pb_exercise_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_exercise_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Exercise); i {
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
		file_pb_exercise_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewExercise); i {
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
			RawDescriptor: file_pb_exercise_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_exercise_proto_goTypes,
		DependencyIndexes: file_pb_exercise_proto_depIdxs,
		MessageInfos:      file_pb_exercise_proto_msgTypes,
	}.Build()
	File_pb_exercise_proto = out.File
	file_pb_exercise_proto_rawDesc = nil
	file_pb_exercise_proto_goTypes = nil
	file_pb_exercise_proto_depIdxs = nil
}