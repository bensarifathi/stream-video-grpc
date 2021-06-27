// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.6.1
// source: streamVod/streamVodpb/streamVod.proto

package streamVodpb

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

type Status int32

const (
	Status_GoodStream Status = 0
	Status_BadStream  Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "GoodStream",
		1: "BadStream",
	}
	Status_value = map[string]int32{
		"GoodStream": 0,
		"BadStream":  1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_streamVod_streamVodpb_streamVod_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_streamVod_streamVodpb_streamVod_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_streamVod_streamVodpb_streamVod_proto_rawDescGZIP(), []int{0}
}

type ImageFrameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rows  int32  `protobuf:"varint,1,opt,name=rows,proto3" json:"rows,omitempty"`
	Cols  int32  `protobuf:"varint,2,opt,name=cols,proto3" json:"cols,omitempty"`
	Type  int32  `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Frame []byte `protobuf:"bytes,4,opt,name=frame,proto3" json:"frame,omitempty"`
}

func (x *ImageFrameRequest) Reset() {
	*x = ImageFrameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streamVod_streamVodpb_streamVod_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageFrameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageFrameRequest) ProtoMessage() {}

func (x *ImageFrameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_streamVod_streamVodpb_streamVod_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageFrameRequest.ProtoReflect.Descriptor instead.
func (*ImageFrameRequest) Descriptor() ([]byte, []int) {
	return file_streamVod_streamVodpb_streamVod_proto_rawDescGZIP(), []int{0}
}

func (x *ImageFrameRequest) GetRows() int32 {
	if x != nil {
		return x.Rows
	}
	return 0
}

func (x *ImageFrameRequest) GetCols() int32 {
	if x != nil {
		return x.Cols
	}
	return 0
}

func (x *ImageFrameRequest) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *ImageFrameRequest) GetFrame() []byte {
	if x != nil {
		return x.Frame
	}
	return nil
}

type ImageFrameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode Status `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3,enum=streamVod.Status" json:"status_code,omitempty"`
}

func (x *ImageFrameResponse) Reset() {
	*x = ImageFrameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streamVod_streamVodpb_streamVod_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageFrameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageFrameResponse) ProtoMessage() {}

func (x *ImageFrameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_streamVod_streamVodpb_streamVod_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageFrameResponse.ProtoReflect.Descriptor instead.
func (*ImageFrameResponse) Descriptor() ([]byte, []int) {
	return file_streamVod_streamVodpb_streamVod_proto_rawDescGZIP(), []int{1}
}

func (x *ImageFrameResponse) GetStatusCode() Status {
	if x != nil {
		return x.StatusCode
	}
	return Status_GoodStream
}

var File_streamVod_streamVodpb_streamVod_proto protoreflect.FileDescriptor

var file_streamVod_streamVodpb_streamVod_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x56, 0x6f, 0x64, 0x2f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x56, 0x6f, 0x64, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x56, 0x6f,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x56,
	0x6f, 0x64, 0x22, 0x65, 0x0a, 0x11, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x6c, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x12, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x56, 0x6f, 0x64,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x2a, 0x27, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a,
	0x0a, 0x47, 0x6f, 0x6f, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x10, 0x00, 0x12, 0x0d, 0x0a,
	0x09, 0x42, 0x61, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x10, 0x01, 0x32, 0x61, 0x0a, 0x12,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x56, 0x6f, 0x64, 0x2e, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x56, 0x6f, 0x64, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x42,
	0x19, 0x5a, 0x17, 0x2e, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x56, 0x6f, 0x64, 0x2f, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x56, 0x6f, 0x64, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_streamVod_streamVodpb_streamVod_proto_rawDescOnce sync.Once
	file_streamVod_streamVodpb_streamVod_proto_rawDescData = file_streamVod_streamVodpb_streamVod_proto_rawDesc
)

func file_streamVod_streamVodpb_streamVod_proto_rawDescGZIP() []byte {
	file_streamVod_streamVodpb_streamVod_proto_rawDescOnce.Do(func() {
		file_streamVod_streamVodpb_streamVod_proto_rawDescData = protoimpl.X.CompressGZIP(file_streamVod_streamVodpb_streamVod_proto_rawDescData)
	})
	return file_streamVod_streamVodpb_streamVod_proto_rawDescData
}

var file_streamVod_streamVodpb_streamVod_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_streamVod_streamVodpb_streamVod_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_streamVod_streamVodpb_streamVod_proto_goTypes = []interface{}{
	(Status)(0),                // 0: streamVod.Status
	(*ImageFrameRequest)(nil),  // 1: streamVod.ImageFrameRequest
	(*ImageFrameResponse)(nil), // 2: streamVod.ImageFrameResponse
}
var file_streamVod_streamVodpb_streamVod_proto_depIdxs = []int32{
	0, // 0: streamVod.ImageFrameResponse.status_code:type_name -> streamVod.Status
	1, // 1: streamVod.videoStreamService.ImageFrame:input_type -> streamVod.ImageFrameRequest
	2, // 2: streamVod.videoStreamService.ImageFrame:output_type -> streamVod.ImageFrameResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_streamVod_streamVodpb_streamVod_proto_init() }
func file_streamVod_streamVodpb_streamVod_proto_init() {
	if File_streamVod_streamVodpb_streamVod_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_streamVod_streamVodpb_streamVod_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageFrameRequest); i {
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
		file_streamVod_streamVodpb_streamVod_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageFrameResponse); i {
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
			RawDescriptor: file_streamVod_streamVodpb_streamVod_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_streamVod_streamVodpb_streamVod_proto_goTypes,
		DependencyIndexes: file_streamVod_streamVodpb_streamVod_proto_depIdxs,
		EnumInfos:         file_streamVod_streamVodpb_streamVod_proto_enumTypes,
		MessageInfos:      file_streamVod_streamVodpb_streamVod_proto_msgTypes,
	}.Build()
	File_streamVod_streamVodpb_streamVod_proto = out.File
	file_streamVod_streamVodpb_streamVod_proto_rawDesc = nil
	file_streamVod_streamVodpb_streamVod_proto_goTypes = nil
	file_streamVod_streamVodpb_streamVod_proto_depIdxs = nil
}