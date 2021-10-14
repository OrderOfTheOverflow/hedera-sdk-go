// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/system_delete.proto

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

//*
// Delete a file or smart contract - can only be done with a Hedera administrative multisignature.
// When it is deleted, it immediately disappears from the system as seen by the user, but is still
// stored internally until the expiration time, at which time it is truly and permanently deleted.
// Until that time, it can be undeleted by the Hedera administrative multisignature. When a smart
// contract is deleted, the cryptocurrency account within it continues to exist, and is not affected
// by the expiration time here.
type SystemDeleteTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Id:
	//	*SystemDeleteTransactionBody_FileID
	//	*SystemDeleteTransactionBody_ContractID
	Id isSystemDeleteTransactionBody_Id `protobuf_oneof:"id"`
	//*
	// The timestamp in seconds at which the "deleted" file should truly be permanently deleted
	ExpirationTime *TimestampSeconds `protobuf:"bytes,3,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`
}

func (x *SystemDeleteTransactionBody) Reset() {
	*x = SystemDeleteTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_system_delete_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemDeleteTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemDeleteTransactionBody) ProtoMessage() {}

func (x *SystemDeleteTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_proto_system_delete_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemDeleteTransactionBody.ProtoReflect.Descriptor instead.
func (*SystemDeleteTransactionBody) Descriptor() ([]byte, []int) {
	return file_proto_system_delete_proto_rawDescGZIP(), []int{0}
}

func (m *SystemDeleteTransactionBody) GetId() isSystemDeleteTransactionBody_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (x *SystemDeleteTransactionBody) GetFileID() *FileID {
	if x, ok := x.GetId().(*SystemDeleteTransactionBody_FileID); ok {
		return x.FileID
	}
	return nil
}

func (x *SystemDeleteTransactionBody) GetContractID() *ContractID {
	if x, ok := x.GetId().(*SystemDeleteTransactionBody_ContractID); ok {
		return x.ContractID
	}
	return nil
}

func (x *SystemDeleteTransactionBody) GetExpirationTime() *TimestampSeconds {
	if x != nil {
		return x.ExpirationTime
	}
	return nil
}

type isSystemDeleteTransactionBody_Id interface {
	isSystemDeleteTransactionBody_Id()
}

type SystemDeleteTransactionBody_FileID struct {
	//*
	// The file ID of the file to delete, in the format used in transactions
	FileID *FileID `protobuf:"bytes,1,opt,name=fileID,proto3,oneof"`
}

type SystemDeleteTransactionBody_ContractID struct {
	//*
	// The contract ID instance to delete, in the format used in transactions
	ContractID *ContractID `protobuf:"bytes,2,opt,name=contractID,proto3,oneof"`
}

func (*SystemDeleteTransactionBody_FileID) isSystemDeleteTransactionBody_Id() {}

func (*SystemDeleteTransactionBody_ContractID) isSystemDeleteTransactionBody_Id() {}

var File_proto_system_delete_proto protoreflect.FileDescriptor

var file_proto_system_delete_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x1b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f,
	0x64, 0x79, 0x12, 0x27, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49,
	0x44, 0x48, 0x00, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x33, 0x0a, 0x0a, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x49, 0x44, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44,
	0x12, 0x3f, 0x0a, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x73, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x42, 0x04, 0x0a, 0x02, 0x69, 0x64, 0x42, 0x4b, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x68,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x68, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_system_delete_proto_rawDescOnce sync.Once
	file_proto_system_delete_proto_rawDescData = file_proto_system_delete_proto_rawDesc
)

func file_proto_system_delete_proto_rawDescGZIP() []byte {
	file_proto_system_delete_proto_rawDescOnce.Do(func() {
		file_proto_system_delete_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_system_delete_proto_rawDescData)
	})
	return file_proto_system_delete_proto_rawDescData
}

var file_proto_system_delete_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_system_delete_proto_goTypes = []interface{}{
	(*SystemDeleteTransactionBody)(nil), // 0: proto.SystemDeleteTransactionBody
	(*FileID)(nil),                      // 1: proto.FileID
	(*ContractID)(nil),                  // 2: proto.ContractID
	(*TimestampSeconds)(nil),            // 3: proto.TimestampSeconds
}
var file_proto_system_delete_proto_depIdxs = []int32{
	1, // 0: proto.SystemDeleteTransactionBody.fileID:type_name -> proto.FileID
	2, // 1: proto.SystemDeleteTransactionBody.contractID:type_name -> proto.ContractID
	3, // 2: proto.SystemDeleteTransactionBody.expirationTime:type_name -> proto.TimestampSeconds
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_system_delete_proto_init() }
func file_proto_system_delete_proto_init() {
	if File_proto_system_delete_proto != nil {
		return
	}
	file_proto_basic_types_proto_init()
	file_proto_timestamp_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_system_delete_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemDeleteTransactionBody); i {
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
	file_proto_system_delete_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SystemDeleteTransactionBody_FileID)(nil),
		(*SystemDeleteTransactionBody_ContractID)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_system_delete_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_system_delete_proto_goTypes,
		DependencyIndexes: file_proto_system_delete_proto_depIdxs,
		MessageInfos:      file_proto_system_delete_proto_msgTypes,
	}.Build()
	File_proto_system_delete_proto = out.File
	file_proto_system_delete_proto_rawDesc = nil
	file_proto_system_delete_proto_goTypes = nil
	file_proto_system_delete_proto_depIdxs = nil
}
