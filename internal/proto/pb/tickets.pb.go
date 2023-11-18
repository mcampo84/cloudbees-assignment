// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: internal/proto/tickets.proto

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

type PurchaseTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	From   string  `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To     string  `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Price  float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *PurchaseTicketRequest) Reset() {
	*x = PurchaseTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_tickets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseTicketRequest) ProtoMessage() {}

func (x *PurchaseTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_tickets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseTicketRequest.ProtoReflect.Descriptor instead.
func (*PurchaseTicketRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_tickets_proto_rawDescGZIP(), []int{0}
}

func (x *PurchaseTicketRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PurchaseTicketRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *PurchaseTicketRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *PurchaseTicketRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type PurchaseTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From  string  `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To    string  `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	User  *User   `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Price float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *PurchaseTicketResponse) Reset() {
	*x = PurchaseTicketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_tickets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseTicketResponse) ProtoMessage() {}

func (x *PurchaseTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_tickets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseTicketResponse.ProtoReflect.Descriptor instead.
func (*PurchaseTicketResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_tickets_proto_rawDescGZIP(), []int{1}
}

func (x *PurchaseTicketResponse) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *PurchaseTicketResponse) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *PurchaseTicketResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *PurchaseTicketResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastNae   string `protobuf:"bytes,2,opt,name=last_nae,json=lastNae,proto3" json:"last_nae,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_tickets_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_tickets_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_internal_proto_tickets_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastNae() string {
	if x != nil {
		return x.LastNae
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_internal_proto_tickets_proto protoreflect.FileDescriptor

var file_internal_proto_tickets_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x6a,
	0x0a, 0x15, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x7c, 0x0a, 0x16, 0x50, 0x75,
	0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x28, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x56, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x32, 0x75, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x0e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x25, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x26, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_tickets_proto_rawDescOnce sync.Once
	file_internal_proto_tickets_proto_rawDescData = file_internal_proto_tickets_proto_rawDesc
)

func file_internal_proto_tickets_proto_rawDescGZIP() []byte {
	file_internal_proto_tickets_proto_rawDescOnce.Do(func() {
		file_internal_proto_tickets_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_tickets_proto_rawDescData)
	})
	return file_internal_proto_tickets_proto_rawDescData
}

var file_internal_proto_tickets_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_proto_tickets_proto_goTypes = []interface{}{
	(*PurchaseTicketRequest)(nil),  // 0: trainTicketing.PurchaseTicketRequest
	(*PurchaseTicketResponse)(nil), // 1: trainTicketing.PurchaseTicketResponse
	(*User)(nil),                   // 2: trainTicketing.User
}
var file_internal_proto_tickets_proto_depIdxs = []int32{
	2, // 0: trainTicketing.PurchaseTicketResponse.user:type_name -> trainTicketing.User
	0, // 1: trainTicketing.TrainTicketService.PurchaseTicket:input_type -> trainTicketing.PurchaseTicketRequest
	1, // 2: trainTicketing.TrainTicketService.PurchaseTicket:output_type -> trainTicketing.PurchaseTicketResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_proto_tickets_proto_init() }
func file_internal_proto_tickets_proto_init() {
	if File_internal_proto_tickets_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_tickets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchaseTicketRequest); i {
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
		file_internal_proto_tickets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchaseTicketResponse); i {
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
		file_internal_proto_tickets_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_internal_proto_tickets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_tickets_proto_goTypes,
		DependencyIndexes: file_internal_proto_tickets_proto_depIdxs,
		MessageInfos:      file_internal_proto_tickets_proto_msgTypes,
	}.Build()
	File_internal_proto_tickets_proto = out.File
	file_internal_proto_tickets_proto_rawDesc = nil
	file_internal_proto_tickets_proto_goTypes = nil
	file_internal_proto_tickets_proto_depIdxs = nil
}
