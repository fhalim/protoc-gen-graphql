// Code generated by protoc-gen-go. DO NOT EDIT.
// source: options.proto

package graphqlpb // import "github.com/martinxsliu/protoc-gen-graphql/graphqlpb"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

var E_Operation = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         82731,
	Name:          "graphql.operation",
	Tag:           "bytes,82731,opt,name=operation",
	Filename:      "options.proto",
}

func init() {
	proto.RegisterExtension(E_Operation)
}

func init() { proto.RegisterFile("options.proto", fileDescriptor_options_06e7ca03491b66cf) }

var fileDescriptor_options_06e7ca03491b66cf = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x2f, 0x28, 0xc9,
	0xcc, 0xcf, 0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x2f, 0x4a, 0x2c, 0xc8,
	0x28, 0xcc, 0x91, 0x52, 0x48, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x0b, 0x27, 0x95, 0xa6,
	0xe9, 0xa7, 0xa4, 0x16, 0x27, 0x17, 0x65, 0x16, 0x94, 0xe4, 0x17, 0x41, 0x94, 0x5a, 0xd9, 0x71,
	0x71, 0xe6, 0x17, 0xa4, 0x16, 0x25, 0x82, 0xb4, 0x0b, 0xc9, 0xe9, 0x41, 0xd4, 0xeb, 0xc1, 0xd4,
	0xeb, 0xf9, 0xa6, 0x96, 0x64, 0xe4, 0xa7, 0xf8, 0x43, 0x4c, 0x97, 0x58, 0xdd, 0xc6, 0xaa, 0xc0,
	0xa8, 0xc1, 0x19, 0x84, 0xd0, 0xe2, 0x64, 0x1a, 0x65, 0x9c, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4,
	0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x9b, 0x58, 0x54, 0x92, 0x99, 0x57, 0x51, 0x9c, 0x93, 0x59, 0x0a,
	0xb1, 0x33, 0x59, 0x37, 0x3d, 0x35, 0x4f, 0x17, 0xea, 0x1c, 0x7d, 0x28, 0x5d, 0x90, 0x94, 0xc4,
	0x06, 0x96, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x57, 0xb3, 0x24, 0xb0, 0xb9, 0x00, 0x00,
	0x00,
}