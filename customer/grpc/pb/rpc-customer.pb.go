// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: customer/grpc/proto/rpc-customer.proto

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

type CustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customer *Customer `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
}

func (x *CustomerResponse) Reset() {
	*x = CustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_grpc_proto_rpc_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerResponse) ProtoMessage() {}

func (x *CustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customer_grpc_proto_rpc_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerResponse.ProtoReflect.Descriptor instead.
func (*CustomerResponse) Descriptor() ([]byte, []int) {
	return file_customer_grpc_proto_rpc_customer_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerResponse) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

type DetailCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DetailCustomerRequest) Reset() {
	*x = DetailCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_grpc_proto_rpc_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailCustomerRequest) ProtoMessage() {}

func (x *DetailCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_grpc_proto_rpc_customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailCustomerRequest.ProtoReflect.Descriptor instead.
func (*DetailCustomerRequest) Descriptor() ([]byte, []int) {
	return file_customer_grpc_proto_rpc_customer_proto_rawDescGZIP(), []int{1}
}

func (x *DetailCustomerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_customer_grpc_proto_rpc_customer_proto protoreflect.FileDescriptor

var file_customer_grpc_proto_rpc_customer_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x22, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x10, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x5a, 0x0a,
	0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x47, 0x0a, 0x0e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x12, 0x5a, 0x10, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customer_grpc_proto_rpc_customer_proto_rawDescOnce sync.Once
	file_customer_grpc_proto_rpc_customer_proto_rawDescData = file_customer_grpc_proto_rpc_customer_proto_rawDesc
)

func file_customer_grpc_proto_rpc_customer_proto_rawDescGZIP() []byte {
	file_customer_grpc_proto_rpc_customer_proto_rawDescOnce.Do(func() {
		file_customer_grpc_proto_rpc_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_customer_grpc_proto_rpc_customer_proto_rawDescData)
	})
	return file_customer_grpc_proto_rpc_customer_proto_rawDescData
}

var file_customer_grpc_proto_rpc_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_customer_grpc_proto_rpc_customer_proto_goTypes = []interface{}{
	(*CustomerResponse)(nil),      // 0: proto.CustomerResponse
	(*DetailCustomerRequest)(nil), // 1: proto.DetailCustomerRequest
	(*Customer)(nil),              // 2: proto.Customer
}
var file_customer_grpc_proto_rpc_customer_proto_depIdxs = []int32{
	2, // 0: proto.CustomerResponse.customer:type_name -> proto.Customer
	1, // 1: proto.CustomerService.DetailCustomer:input_type -> proto.DetailCustomerRequest
	0, // 2: proto.CustomerService.DetailCustomer:output_type -> proto.CustomerResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_customer_grpc_proto_rpc_customer_proto_init() }
func file_customer_grpc_proto_rpc_customer_proto_init() {
	if File_customer_grpc_proto_rpc_customer_proto != nil {
		return
	}
	file_customer_grpc_proto_customer_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_customer_grpc_proto_rpc_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerResponse); i {
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
		file_customer_grpc_proto_rpc_customer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailCustomerRequest); i {
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
			RawDescriptor: file_customer_grpc_proto_rpc_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_customer_grpc_proto_rpc_customer_proto_goTypes,
		DependencyIndexes: file_customer_grpc_proto_rpc_customer_proto_depIdxs,
		MessageInfos:      file_customer_grpc_proto_rpc_customer_proto_msgTypes,
	}.Build()
	File_customer_grpc_proto_rpc_customer_proto = out.File
	file_customer_grpc_proto_rpc_customer_proto_rawDesc = nil
	file_customer_grpc_proto_rpc_customer_proto_goTypes = nil
	file_customer_grpc_proto_rpc_customer_proto_depIdxs = nil
}