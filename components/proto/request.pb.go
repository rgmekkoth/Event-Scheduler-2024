// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: proto/request.proto

package proto

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

type Time struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start     int64   `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	Divisions []int64 `protobuf:"varint,2,rep,packed,name=divisions,proto3" json:"divisions,omitempty"`
}

func (x *Time) Reset() {
	*x = Time{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Time) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Time) ProtoMessage() {}

func (x *Time) ProtoReflect() protoreflect.Message {
	mi := &file_proto_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Time.ProtoReflect.Descriptor instead.
func (*Time) Descriptor() ([]byte, []int) {
	return file_proto_request_proto_rawDescGZIP(), []int{0}
}

func (x *Time) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Time) GetDivisions() []int64 {
	if x != nil {
		return x.Divisions
	}
	return nil
}

type Constraints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupSize int32 `protobuf:"varint,1,opt,name=groupSize,proto3" json:"groupSize,omitempty"`
}

func (x *Constraints) Reset() {
	*x = Constraints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Constraints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Constraints) ProtoMessage() {}

func (x *Constraints) ProtoReflect() protoreflect.Message {
	mi := &file_proto_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Constraints.ProtoReflect.Descriptor instead.
func (*Constraints) Descriptor() ([]byte, []int) {
	return file_proto_request_proto_rawDescGZIP(), []int{1}
}

func (x *Constraints) GetGroupSize() int32 {
	if x != nil {
		return x.GroupSize
	}
	return 0
}

type Registration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Students []*Student `protobuf:"bytes,1,rep,name=students,proto3" json:"students,omitempty"`
	Judges   []*Judge   `protobuf:"bytes,2,rep,name=judges,proto3" json:"judges,omitempty"`
	Rooms    []*Room    `protobuf:"bytes,3,rep,name=rooms,proto3" json:"rooms,omitempty"`
	Events   []*Event   `protobuf:"bytes,4,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *Registration) Reset() {
	*x = Registration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registration) ProtoMessage() {}

func (x *Registration) ProtoReflect() protoreflect.Message {
	mi := &file_proto_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registration.ProtoReflect.Descriptor instead.
func (*Registration) Descriptor() ([]byte, []int) {
	return file_proto_request_proto_rawDescGZIP(), []int{2}
}

func (x *Registration) GetStudents() []*Student {
	if x != nil {
		return x.Students
	}
	return nil
}

func (x *Registration) GetJudges() []*Judge {
	if x != nil {
		return x.Judges
	}
	return nil
}

func (x *Registration) GetRooms() []*Room {
	if x != nil {
		return x.Rooms
	}
	return nil
}

func (x *Registration) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type StudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event string   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Group []string `protobuf:"bytes,2,rep,name=group,proto3" json:"group,omitempty"`
}

func (x *StudentRequest) Reset() {
	*x = StudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentRequest) ProtoMessage() {}

func (x *StudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentRequest.ProtoReflect.Descriptor instead.
func (*StudentRequest) Descriptor() ([]byte, []int) {
	return file_proto_request_proto_rawDescGZIP(), []int{3}
}

func (x *StudentRequest) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *StudentRequest) GetGroup() []string {
	if x != nil {
		return x.Group
	}
	return nil
}

var File_proto_request_proto protoreflect.FileDescriptor

var file_proto_request_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x09, 0x64, 0x69, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x2b, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x69, 0x7a,
	0x65, 0x22, 0x91, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x08,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x06, 0x6a, 0x75, 0x64, 0x67,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4a, 0x75, 0x64, 0x67, 0x65,
	0x52, 0x06, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x05,
	0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x1e, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x3c, 0x0a, 0x0e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x42, 0x12, 0x5a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_request_proto_rawDescOnce sync.Once
	file_proto_request_proto_rawDescData = file_proto_request_proto_rawDesc
)

func file_proto_request_proto_rawDescGZIP() []byte {
	file_proto_request_proto_rawDescOnce.Do(func() {
		file_proto_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_request_proto_rawDescData)
	})
	return file_proto_request_proto_rawDescData
}

var file_proto_request_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_request_proto_goTypes = []interface{}{
	(*Time)(nil),           // 0: Time
	(*Constraints)(nil),    // 1: Constraints
	(*Registration)(nil),   // 2: Registration
	(*StudentRequest)(nil), // 3: StudentRequest
	(*Student)(nil),        // 4: Student
	(*Judge)(nil),          // 5: Judge
	(*Room)(nil),           // 6: Room
	(*Event)(nil),          // 7: Event
}
var file_proto_request_proto_depIdxs = []int32{
	4, // 0: Registration.students:type_name -> Student
	5, // 1: Registration.judges:type_name -> Judge
	6, // 2: Registration.rooms:type_name -> Room
	7, // 3: Registration.events:type_name -> Event
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_request_proto_init() }
func file_proto_request_proto_init() {
	if File_proto_request_proto != nil {
		return
	}
	file_proto_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Time); i {
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
		file_proto_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Constraints); i {
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
		file_proto_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registration); i {
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
		file_proto_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentRequest); i {
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
			RawDescriptor: file_proto_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_request_proto_goTypes,
		DependencyIndexes: file_proto_request_proto_depIdxs,
		MessageInfos:      file_proto_request_proto_msgTypes,
	}.Build()
	File_proto_request_proto = out.File
	file_proto_request_proto_rawDesc = nil
	file_proto_request_proto_goTypes = nil
	file_proto_request_proto_depIdxs = nil
}
