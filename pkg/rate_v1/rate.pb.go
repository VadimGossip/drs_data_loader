// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: rate.proto

package rate_v1

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

type FindRateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GwgrId  int64  `protobuf:"varint,1,opt,name=gwgrId,proto3" json:"gwgrId,omitempty"`
	DateAt  int64  `protobuf:"varint,2,opt,name=dateAt,proto3" json:"dateAt,omitempty"`
	Dir     uint32 `protobuf:"varint,3,opt,name=dir,proto3" json:"dir,omitempty"`
	ANumber string `protobuf:"bytes,4,opt,name=aNumber,proto3" json:"aNumber,omitempty"`
	BNumber string `protobuf:"bytes,5,opt,name=bNumber,proto3" json:"bNumber,omitempty"`
}

func (x *FindRateRequest) Reset() {
	*x = FindRateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindRateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindRateRequest) ProtoMessage() {}

func (x *FindRateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindRateRequest.ProtoReflect.Descriptor instead.
func (*FindRateRequest) Descriptor() ([]byte, []int) {
	return file_rate_proto_rawDescGZIP(), []int{0}
}

func (x *FindRateRequest) GetGwgrId() int64 {
	if x != nil {
		return x.GwgrId
	}
	return 0
}

func (x *FindRateRequest) GetDateAt() int64 {
	if x != nil {
		return x.DateAt
	}
	return 0
}

func (x *FindRateRequest) GetDir() uint32 {
	if x != nil {
		return x.Dir
	}
	return 0
}

func (x *FindRateRequest) GetANumber() string {
	if x != nil {
		return x.ANumber
	}
	return ""
}

func (x *FindRateRequest) GetBNumber() string {
	if x != nil {
		return x.BNumber
	}
	return ""
}

type FindRateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RmsrId    int64   `protobuf:"varint,1,opt,name=rmsrId,proto3" json:"rmsrId,omitempty"`
	PriceBase float64 `protobuf:"fixed64,2,opt,name=priceBase,proto3" json:"priceBase,omitempty"`
}

func (x *FindRateResponse) Reset() {
	*x = FindRateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindRateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindRateResponse) ProtoMessage() {}

func (x *FindRateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindRateResponse.ProtoReflect.Descriptor instead.
func (*FindRateResponse) Descriptor() ([]byte, []int) {
	return file_rate_proto_rawDescGZIP(), []int{1}
}

func (x *FindRateResponse) GetRmsrId() int64 {
	if x != nil {
		return x.RmsrId
	}
	return 0
}

func (x *FindRateResponse) GetPriceBase() float64 {
	if x != nil {
		return x.PriceBase
	}
	return 0
}

var File_rate_proto protoreflect.FileDescriptor

var file_rate_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x61,
	0x74, 0x65, 0x5f, 0x76, 0x31, 0x22, 0x87, 0x01, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x77, 0x67,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x67, 0x77, 0x67, 0x72, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x64, 0x69, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22,
	0x48, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6d, 0x73, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6d, 0x73, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x42, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x61, 0x73, 0x65, 0x32, 0x49, 0x0a, 0x06, 0x52, 0x61, 0x74,
	0x65, 0x56, 0x31, 0x12, 0x3f, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x61, 0x74, 0x65, 0x12,
	0x18, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x61, 0x74, 0x65,
	0x5f, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x56, 0x61, 0x64, 0x69, 0x6d, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2f, 0x64,
	0x72, 0x73, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x31, 0x3b, 0x72, 0x61, 0x74, 0x65, 0x5f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rate_proto_rawDescOnce sync.Once
	file_rate_proto_rawDescData = file_rate_proto_rawDesc
)

func file_rate_proto_rawDescGZIP() []byte {
	file_rate_proto_rawDescOnce.Do(func() {
		file_rate_proto_rawDescData = protoimpl.X.CompressGZIP(file_rate_proto_rawDescData)
	})
	return file_rate_proto_rawDescData
}

var file_rate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rate_proto_goTypes = []interface{}{
	(*FindRateRequest)(nil),  // 0: rate_v1.FindRateRequest
	(*FindRateResponse)(nil), // 1: rate_v1.FindRateResponse
}
var file_rate_proto_depIdxs = []int32{
	0, // 0: rate_v1.RateV1.FindRate:input_type -> rate_v1.FindRateRequest
	1, // 1: rate_v1.RateV1.FindRate:output_type -> rate_v1.FindRateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rate_proto_init() }
func file_rate_proto_init() {
	if File_rate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindRateRequest); i {
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
		file_rate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindRateResponse); i {
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
			RawDescriptor: file_rate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rate_proto_goTypes,
		DependencyIndexes: file_rate_proto_depIdxs,
		MessageInfos:      file_rate_proto_msgTypes,
	}.Build()
	File_rate_proto = out.File
	file_rate_proto_rawDesc = nil
	file_rate_proto_goTypes = nil
	file_rate_proto_depIdxs = nil
}
