// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: todo.proto

package todopb

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

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type TodoMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Task *Task  `protobuf:"bytes,2,opt,name=Task,proto3" json:"Task,omitempty"`
}

func (x *TodoMessageRequest) Reset() {
	*x = TodoMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoMessageRequest) ProtoMessage() {}

func (x *TodoMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoMessageRequest.ProtoReflect.Descriptor instead.
func (*TodoMessageRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{1}
}

func (x *TodoMessageRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TodoMessageRequest) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type TodoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List map[string]*Task `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TodoList) Reset() {
	*x = TodoList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoList) ProtoMessage() {}

func (x *TodoList) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoList.ProtoReflect.Descriptor instead.
func (*TodoList) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{2}
}

func (x *TodoList) GetList() map[string]*Task {
	if x != nil {
		return x.List
	}
	return nil
}

type TodoListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List *TodoList `protobuf:"bytes,1,opt,name=List,proto3" json:"List,omitempty"`
}

func (x *TodoListResponse) Reset() {
	*x = TodoListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoListResponse) ProtoMessage() {}

func (x *TodoListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoListResponse.ProtoReflect.Descriptor instead.
func (*TodoListResponse) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{3}
}

func (x *TodoListResponse) GetList() *TodoList {
	if x != nil {
		return x.List
	}
	return nil
}

var File_todo_proto protoreflect.FileDescriptor

var file_todo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x45, 0x0a, 0x12, 0x54, 0x6f, 0x64, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x22, 0x7f, 0x0a, 0x08, 0x54, 0x6f, 0x64, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x1a, 0x44, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x37, 0x0a, 0x10, 0x54, 0x6f, 0x64,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x16, 0x5a, 0x14, 0x2e, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_todo_proto_rawDescOnce sync.Once
	file_todo_proto_rawDescData = file_todo_proto_rawDesc
)

func file_todo_proto_rawDescGZIP() []byte {
	file_todo_proto_rawDescOnce.Do(func() {
		file_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_todo_proto_rawDescData)
	})
	return file_todo_proto_rawDescData
}

var file_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_todo_proto_goTypes = []any{
	(*Task)(nil),               // 0: proto.Task
	(*TodoMessageRequest)(nil), // 1: proto.TodoMessageRequest
	(*TodoList)(nil),           // 2: proto.TodoList
	(*TodoListResponse)(nil),   // 3: proto.TodoListResponse
	nil,                        // 4: proto.TodoList.ListEntry
}
var file_todo_proto_depIdxs = []int32{
	0, // 0: proto.TodoMessageRequest.Task:type_name -> proto.Task
	4, // 1: proto.TodoList.List:type_name -> proto.TodoList.ListEntry
	2, // 2: proto.TodoListResponse.List:type_name -> proto.TodoList
	0, // 3: proto.TodoList.ListEntry.value:type_name -> proto.Task
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_todo_proto_init() }
func file_todo_proto_init() {
	if File_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todo_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Task); i {
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
		file_todo_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TodoMessageRequest); i {
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
		file_todo_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TodoList); i {
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
		file_todo_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TodoListResponse); i {
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
			RawDescriptor: file_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_todo_proto_goTypes,
		DependencyIndexes: file_todo_proto_depIdxs,
		MessageInfos:      file_todo_proto_msgTypes,
	}.Build()
	File_todo_proto = out.File
	file_todo_proto_rawDesc = nil
	file_todo_proto_goTypes = nil
	file_todo_proto_depIdxs = nil
}
