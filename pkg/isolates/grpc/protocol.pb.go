// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: pkg/isolates/grpc/protocol.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_pkg_isolates_grpc_protocol_proto protoreflect.FileDescriptor

var file_pkg_isolates_grpc_protocol_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x73, 0x6f, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xba, 0x03, 0x0a, 0x0f, 0x49, 0x73, 0x6f, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x73, 0x6f, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x45, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x73, 0x6f, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x49, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x49, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x47, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x12, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x6f, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x49, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47,
	0x65, 0x74, 0x49, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x73, 0x6f, 0x6c, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x11,
	0x53, 0x65, 0x74, 0x49, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x72, 0x74, 0x65, 0x69, 0x6c, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b,
	0x74, 0x69, 0x76, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x73, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_pkg_isolates_grpc_protocol_proto_goTypes = []interface{}{
	(*UpdateIsolateRequest)(nil), // 0: grpc.UpdateIsolateRequest
	(*CreateIsolateRequest)(nil), // 1: grpc.CreateIsolateRequest
	(*ListIsolatesRequest)(nil),  // 2: grpc.ListIsolatesRequest
	(*GetIsolateRequest)(nil),    // 3: grpc.GetIsolateRequest
	(*SetTrafficRequest)(nil),    // 4: grpc.SetTrafficRequest
	(*emptypb.Empty)(nil),        // 5: google.protobuf.Empty
	(*ListIsolatesResponse)(nil), // 6: grpc.ListIsolatesResponse
	(*GetIsolateResponse)(nil),   // 7: grpc.GetIsolateResponse
}
var file_pkg_isolates_grpc_protocol_proto_depIdxs = []int32{
	0, // 0: grpc.IsolatesService.UpdateIsolate:input_type -> grpc.UpdateIsolateRequest
	1, // 1: grpc.IsolatesService.CreateIsolate:input_type -> grpc.CreateIsolateRequest
	2, // 2: grpc.IsolatesService.DeleteIsolates:input_type -> grpc.ListIsolatesRequest
	2, // 3: grpc.IsolatesService.ListIsolates:input_type -> grpc.ListIsolatesRequest
	3, // 4: grpc.IsolatesService.GetIsolate:input_type -> grpc.GetIsolateRequest
	4, // 5: grpc.IsolatesService.SetIsolateTraffic:input_type -> grpc.SetTrafficRequest
	5, // 6: grpc.IsolatesService.UpdateIsolate:output_type -> google.protobuf.Empty
	5, // 7: grpc.IsolatesService.CreateIsolate:output_type -> google.protobuf.Empty
	5, // 8: grpc.IsolatesService.DeleteIsolates:output_type -> google.protobuf.Empty
	6, // 9: grpc.IsolatesService.ListIsolates:output_type -> grpc.ListIsolatesResponse
	7, // 10: grpc.IsolatesService.GetIsolate:output_type -> grpc.GetIsolateResponse
	5, // 11: grpc.IsolatesService.SetIsolateTraffic:output_type -> google.protobuf.Empty
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_isolates_grpc_protocol_proto_init() }
func file_pkg_isolates_grpc_protocol_proto_init() {
	if File_pkg_isolates_grpc_protocol_proto != nil {
		return
	}
	file_pkg_isolates_grpc_isolates_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_isolates_grpc_protocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_isolates_grpc_protocol_proto_goTypes,
		DependencyIndexes: file_pkg_isolates_grpc_protocol_proto_depIdxs,
	}.Build()
	File_pkg_isolates_grpc_protocol_proto = out.File
	file_pkg_isolates_grpc_protocol_proto_rawDesc = nil
	file_pkg_isolates_grpc_protocol_proto_goTypes = nil
	file_pkg_isolates_grpc_protocol_proto_depIdxs = nil
}
