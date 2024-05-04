// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: program.proto

package program_v1

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

type Set struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Quantity int64   `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Weight   float64 `protobuf:"fixed64,3,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *Set) Reset() {
	*x = Set{}
	if protoimpl.UnsafeEnabled {
		mi := &file_program_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Set) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Set) ProtoMessage() {}

func (x *Set) ProtoReflect() protoreflect.Message {
	mi := &file_program_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Set.ProtoReflect.Descriptor instead.
func (*Set) Descriptor() ([]byte, []int) {
	return file_program_proto_rawDescGZIP(), []int{0}
}

func (x *Set) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Set) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Set) GetWeight() float64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type Exercise struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ExerciseName string   `protobuf:"bytes,2,opt,name=exercise_name,json=exerciseName,proto3" json:"exercise_name,omitempty"`
	Pictures     []string `protobuf:"bytes,3,rep,name=pictures,proto3" json:"pictures,omitempty"`
	Description  string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Sets         []*Set   `protobuf:"bytes,5,rep,name=sets,proto3" json:"sets,omitempty"`
}

func (x *Exercise) Reset() {
	*x = Exercise{}
	if protoimpl.UnsafeEnabled {
		mi := &file_program_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Exercise) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Exercise) ProtoMessage() {}

func (x *Exercise) ProtoReflect() protoreflect.Message {
	mi := &file_program_proto_msgTypes[1]
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
	return file_program_proto_rawDescGZIP(), []int{1}
}

func (x *Exercise) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Exercise) GetExerciseName() string {
	if x != nil {
		return x.ExerciseName
	}
	return ""
}

func (x *Exercise) GetPictures() []string {
	if x != nil {
		return x.Pictures
	}
	return nil
}

func (x *Exercise) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Exercise) GetSets() []*Set {
	if x != nil {
		return x.Sets
	}
	return nil
}

type TrainDay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DayName     string      `protobuf:"bytes,2,opt,name=day_name,json=dayName,proto3" json:"day_name,omitempty"`
	Description string      `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Exercises   []*Exercise `protobuf:"bytes,4,rep,name=exercises,proto3" json:"exercises,omitempty"`
}

func (x *TrainDay) Reset() {
	*x = TrainDay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_program_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrainDay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrainDay) ProtoMessage() {}

func (x *TrainDay) ProtoReflect() protoreflect.Message {
	mi := &file_program_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrainDay.ProtoReflect.Descriptor instead.
func (*TrainDay) Descriptor() ([]byte, []int) {
	return file_program_proto_rawDescGZIP(), []int{2}
}

func (x *TrainDay) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TrainDay) GetDayName() string {
	if x != nil {
		return x.DayName
	}
	return ""
}

func (x *TrainDay) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TrainDay) GetExercises() []*Exercise {
	if x != nil {
		return x.Exercises
	}
	return nil
}

type TrainProgram struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProgramName string      `protobuf:"bytes,2,opt,name=program_name,json=programName,proto3" json:"program_name,omitempty"`
	Description string      `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Status      string      `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	TrainDays   []*TrainDay `protobuf:"bytes,5,rep,name=train_days,json=trainDays,proto3" json:"train_days,omitempty"`
}

func (x *TrainProgram) Reset() {
	*x = TrainProgram{}
	if protoimpl.UnsafeEnabled {
		mi := &file_program_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrainProgram) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrainProgram) ProtoMessage() {}

func (x *TrainProgram) ProtoReflect() protoreflect.Message {
	mi := &file_program_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrainProgram.ProtoReflect.Descriptor instead.
func (*TrainProgram) Descriptor() ([]byte, []int) {
	return file_program_proto_rawDescGZIP(), []int{3}
}

func (x *TrainProgram) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TrainProgram) GetProgramName() string {
	if x != nil {
		return x.ProgramName
	}
	return ""
}

func (x *TrainProgram) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TrainProgram) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TrainProgram) GetTrainDays() []*TrainDay {
	if x != nil {
		return x.TrainDays
	}
	return nil
}

type TrainPrograms struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrainPrograms []*TrainProgram `protobuf:"bytes,1,rep,name=train_programs,json=trainPrograms,proto3" json:"train_programs,omitempty"`
}

func (x *TrainPrograms) Reset() {
	*x = TrainPrograms{}
	if protoimpl.UnsafeEnabled {
		mi := &file_program_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrainPrograms) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrainPrograms) ProtoMessage() {}

func (x *TrainPrograms) ProtoReflect() protoreflect.Message {
	mi := &file_program_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrainPrograms.ProtoReflect.Descriptor instead.
func (*TrainPrograms) Descriptor() ([]byte, []int) {
	return file_program_proto_rawDescGZIP(), []int{4}
}

func (x *TrainPrograms) GetTrainPrograms() []*TrainProgram {
	if x != nil {
		return x.TrainPrograms
	}
	return nil
}

var File_program_proto protoreflect.FileDescriptor

var file_program_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x76, 0x31, 0x22, 0x49, 0x0a, 0x03, 0x53,
	0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0xa2, 0x01, 0x0a, 0x08, 0x45, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x78, 0x65, 0x72,
	0x63, 0x69, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x69, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x69, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x04, 0x73, 0x65, 0x74, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x04, 0x73, 0x65, 0x74, 0x73, 0x22, 0x8b, 0x01, 0x0a, 0x08,
	0x54, 0x72, 0x61, 0x69, 0x6e, 0x44, 0x61, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x61, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x61, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x09, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x61, 0x6d, 0x5f, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x09,
	0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x22, 0xb0, 0x01, 0x0a, 0x0c, 0x54, 0x72,
	0x61, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x44, 0x61,
	0x79, 0x52, 0x09, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x44, 0x61, 0x79, 0x73, 0x22, 0x50, 0x0a, 0x0d,
	0x54, 0x72, 0x61, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x3f, 0x0a,
	0x0e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x5f,
	0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52,
	0x0d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x3b,
	0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x69, 0x72,
	0x69, 0x6c, 0x6c, 0x6d, 0x63, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x65,
	0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x76, 0x31,
	0x3b, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_program_proto_rawDescOnce sync.Once
	file_program_proto_rawDescData = file_program_proto_rawDesc
)

func file_program_proto_rawDescGZIP() []byte {
	file_program_proto_rawDescOnce.Do(func() {
		file_program_proto_rawDescData = protoimpl.X.CompressGZIP(file_program_proto_rawDescData)
	})
	return file_program_proto_rawDescData
}

var file_program_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_program_proto_goTypes = []interface{}{
	(*Set)(nil),           // 0: program_v1.Set
	(*Exercise)(nil),      // 1: program_v1.Exercise
	(*TrainDay)(nil),      // 2: program_v1.TrainDay
	(*TrainProgram)(nil),  // 3: program_v1.TrainProgram
	(*TrainPrograms)(nil), // 4: program_v1.TrainPrograms
}
var file_program_proto_depIdxs = []int32{
	0, // 0: program_v1.Exercise.sets:type_name -> program_v1.Set
	1, // 1: program_v1.TrainDay.exercises:type_name -> program_v1.Exercise
	2, // 2: program_v1.TrainProgram.train_days:type_name -> program_v1.TrainDay
	3, // 3: program_v1.TrainPrograms.train_programs:type_name -> program_v1.TrainProgram
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_program_proto_init() }
func file_program_proto_init() {
	if File_program_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_program_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Set); i {
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
		file_program_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_program_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrainDay); i {
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
		file_program_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrainProgram); i {
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
		file_program_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrainPrograms); i {
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
			RawDescriptor: file_program_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_program_proto_goTypes,
		DependencyIndexes: file_program_proto_depIdxs,
		MessageInfos:      file_program_proto_msgTypes,
	}.Build()
	File_program_proto = out.File
	file_program_proto_rawDesc = nil
	file_program_proto_goTypes = nil
	file_program_proto_depIdxs = nil
}