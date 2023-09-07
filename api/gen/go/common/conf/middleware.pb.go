// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: common/conf/middleware.proto

package conf

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

type Middleware struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnableLogging        bool             `protobuf:"varint,1,opt,name=enable_logging,json=enableLogging,proto3" json:"enable_logging,omitempty"`
	EnableRecovery       bool             `protobuf:"varint,2,opt,name=enable_recovery,json=enableRecovery,proto3" json:"enable_recovery,omitempty"`
	EnableTracing        bool             `protobuf:"varint,3,opt,name=enable_tracing,json=enableTracing,proto3" json:"enable_tracing,omitempty"`
	EnableValidate       bool             `protobuf:"varint,4,opt,name=enable_validate,json=enableValidate,proto3" json:"enable_validate,omitempty"`
	EnableMetric         bool             `protobuf:"varint,5,opt,name=enable_metric,json=enableMetric,proto3" json:"enable_metric,omitempty"`
	EnableCircuitBreaker bool             `protobuf:"varint,6,opt,name=enable_circuit_breaker,json=enableCircuitBreaker,proto3" json:"enable_circuit_breaker,omitempty"`
	EnableLimiter        bool             `protobuf:"varint,7,opt,name=enable_limiter,json=enableLimiter,proto3" json:"enable_limiter,omitempty"`
	Auth                 *Middleware_Auth `protobuf:"bytes,10,opt,name=auth,proto3" json:"auth,omitempty"`
}

func (x *Middleware) Reset() {
	*x = Middleware{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_middleware_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Middleware) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Middleware) ProtoMessage() {}

func (x *Middleware) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_middleware_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Middleware.ProtoReflect.Descriptor instead.
func (*Middleware) Descriptor() ([]byte, []int) {
	return file_common_conf_middleware_proto_rawDescGZIP(), []int{0}
}

func (x *Middleware) GetEnableLogging() bool {
	if x != nil {
		return x.EnableLogging
	}
	return false
}

func (x *Middleware) GetEnableRecovery() bool {
	if x != nil {
		return x.EnableRecovery
	}
	return false
}

func (x *Middleware) GetEnableTracing() bool {
	if x != nil {
		return x.EnableTracing
	}
	return false
}

func (x *Middleware) GetEnableValidate() bool {
	if x != nil {
		return x.EnableValidate
	}
	return false
}

func (x *Middleware) GetEnableMetric() bool {
	if x != nil {
		return x.EnableMetric
	}
	return false
}

func (x *Middleware) GetEnableCircuitBreaker() bool {
	if x != nil {
		return x.EnableCircuitBreaker
	}
	return false
}

func (x *Middleware) GetEnableLimiter() bool {
	if x != nil {
		return x.EnableLimiter
	}
	return false
}

func (x *Middleware) GetAuth() *Middleware_Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

// JWT
type Middleware_Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"` // JWT signature algorithm, supported algorithm: HS256
	Key    string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`       // JWT secret
}

func (x *Middleware_Auth) Reset() {
	*x = Middleware_Auth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_middleware_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Middleware_Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Middleware_Auth) ProtoMessage() {}

func (x *Middleware_Auth) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_middleware_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Middleware_Auth.ProtoReflect.Descriptor instead.
func (*Middleware_Auth) Descriptor() ([]byte, []int) {
	return file_common_conf_middleware_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Middleware_Auth) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Middleware_Auth) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_common_conf_middleware_proto protoreflect.FileDescriptor

var file_common_conf_middleware_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x22, 0x92, 0x03, 0x0a, 0x0a,
	0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e,
	0x67, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e,
	0x67, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0c, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12,
	0x34, 0x0a, 0x16, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69,
	0x74, 0x5f, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x14, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x42, 0x72,
	0x65, 0x61, 0x6b, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x04,
	0x61, 0x75, 0x74, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x30,
	0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x65, 0x76, 0x65, 0x78, 0x70, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74,
	0x72, 0x61, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_conf_middleware_proto_rawDescOnce sync.Once
	file_common_conf_middleware_proto_rawDescData = file_common_conf_middleware_proto_rawDesc
)

func file_common_conf_middleware_proto_rawDescGZIP() []byte {
	file_common_conf_middleware_proto_rawDescOnce.Do(func() {
		file_common_conf_middleware_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_conf_middleware_proto_rawDescData)
	})
	return file_common_conf_middleware_proto_rawDescData
}

var file_common_conf_middleware_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_conf_middleware_proto_goTypes = []interface{}{
	(*Middleware)(nil),      // 0: common.conf.Middleware
	(*Middleware_Auth)(nil), // 1: common.conf.Middleware.Auth
}
var file_common_conf_middleware_proto_depIdxs = []int32{
	1, // 0: common.conf.Middleware.auth:type_name -> common.conf.Middleware.Auth
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_conf_middleware_proto_init() }
func file_common_conf_middleware_proto_init() {
	if File_common_conf_middleware_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_conf_middleware_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Middleware); i {
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
		file_common_conf_middleware_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Middleware_Auth); i {
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
			RawDescriptor: file_common_conf_middleware_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_conf_middleware_proto_goTypes,
		DependencyIndexes: file_common_conf_middleware_proto_depIdxs,
		MessageInfos:      file_common_conf_middleware_proto_msgTypes,
	}.Build()
	File_common_conf_middleware_proto = out.File
	file_common_conf_middleware_proto_rawDesc = nil
	file_common_conf_middleware_proto_goTypes = nil
	file_common_conf_middleware_proto_depIdxs = nil
}
