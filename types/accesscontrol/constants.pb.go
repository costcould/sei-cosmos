// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/accesscontrol/constants.proto

package accesscontrol

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type AccessType int32

const (
	AccessType_UNKNOWN AccessType = 0
	AccessType_READ    AccessType = 1
	AccessType_WRITE   AccessType = 2
	AccessType_COMMIT  AccessType = 3
)

var AccessType_name = map[int32]string{
	0: "UNKNOWN",
	1: "READ",
	2: "WRITE",
	3: "COMMIT",
}

var AccessType_value = map[string]int32{
	"UNKNOWN": 0,
	"READ":    1,
	"WRITE":   2,
	"COMMIT":  3,
}

func (x AccessType) String() string {
	return proto.EnumName(AccessType_name, int32(x))
}

func (AccessType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36568f7561081112, []int{0}
}

type ResourceType int32

const (
	ResourceType_ANY                       ResourceType = 0
	ResourceType_KV                        ResourceType = 1
	ResourceType_Mem                       ResourceType = 2
	ResourceType_DexMem                    ResourceType = 3
	ResourceType_KV_BANK                   ResourceType = 4
	ResourceType_KV_STAKING                ResourceType = 5
	ResourceType_KV_WASM                   ResourceType = 6
	ResourceType_KV_ORACLE                 ResourceType = 7
	ResourceType_KV_DEX                    ResourceType = 8
	ResourceType_KV_EPOCH                  ResourceType = 9
	ResourceType_KV_TOKENFACTORY           ResourceType = 10
	ResourceType_KV_ORACLE_VOTE_TARGETS    ResourceType = 11
	ResourceType_KV_ORACLE_AGGREGATE_VOTES ResourceType = 12
	ResourceType_KV_ORACLE_FEEDERS         ResourceType = 13
)

var ResourceType_name = map[int32]string{
	0:  "ANY",
	1:  "KV",
	2:  "Mem",
	3:  "DexMem",
	4:  "KV_BANK",
	5:  "KV_STAKING",
	6:  "KV_WASM",
	7:  "KV_ORACLE",
	8:  "KV_DEX",
	9:  "KV_EPOCH",
	10: "KV_TOKENFACTORY",
	11: "KV_ORACLE_VOTE_TARGETS",
	12: "KV_ORACLE_AGGREGATE_VOTES",
	13: "KV_ORACLE_FEEDERS",
}

var ResourceType_value = map[string]int32{
	"ANY":                       0,
	"KV":                        1,
	"Mem":                       2,
	"DexMem":                    3,
	"KV_BANK":                   4,
	"KV_STAKING":                5,
	"KV_WASM":                   6,
	"KV_ORACLE":                 7,
	"KV_DEX":                    8,
	"KV_EPOCH":                  9,
	"KV_TOKENFACTORY":           10,
	"KV_ORACLE_VOTE_TARGETS":    11,
	"KV_ORACLE_AGGREGATE_VOTES": 12,
	"KV_ORACLE_FEEDERS":         13,
}

func (x ResourceType) String() string {
	return proto.EnumName(ResourceType_name, int32(x))
}

func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36568f7561081112, []int{1}
}

func init() {
	proto.RegisterEnum("cosmos.accesscontrol.v1beta1.AccessType", AccessType_name, AccessType_value)
	proto.RegisterEnum("cosmos.accesscontrol.v1beta1.ResourceType", ResourceType_name, ResourceType_value)
}

func init() {
	proto.RegisterFile("cosmos/accesscontrol/constants.proto", fileDescriptor_36568f7561081112)
}

var fileDescriptor_36568f7561081112 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x93, 0x75, 0x4b, 0xdb, 0xb7, 0x0e, 0x1e, 0x46, 0x20, 0x81, 0x20, 0x27, 0x4e, 0x93,
	0x48, 0x98, 0xb8, 0x71, 0xf3, 0x12, 0x2f, 0x14, 0x93, 0x04, 0x39, 0xc6, 0x65, 0x5c, 0xac, 0x36,
	0x58, 0x80, 0xa0, 0x75, 0x55, 0x67, 0x88, 0x7d, 0x0b, 0x3e, 0x16, 0xc7, 0x1d, 0x39, 0xa2, 0xf6,
	0xce, 0x67, 0x40, 0x69, 0x26, 0x2a, 0x76, 0xb2, 0xf5, 0x7e, 0xff, 0xff, 0x7b, 0x87, 0x1f, 0x3c,
	0xa9, 0xad, 0x9b, 0x5b, 0x17, 0x4f, 0xeb, 0xda, 0x38, 0x57, 0xdb, 0x45, 0xb3, 0xb2, 0x5f, 0xe3,
	0xda, 0x2e, 0x5c, 0x33, 0x5d, 0x34, 0x2e, 0x5a, 0xae, 0x6c, 0x63, 0xc9, 0xa3, 0x2e, 0x15, 0xfd,
	0x97, 0x8a, 0xbe, 0x9d, 0xcc, 0x4c, 0x33, 0x3d, 0x39, 0x7e, 0x01, 0x40, 0xb7, 0x40, 0x5e, 0x2e,
	0x0d, 0x39, 0x84, 0xfe, 0xdb, 0x82, 0x17, 0xe5, 0xa4, 0x40, 0x8f, 0x0c, 0x60, 0x5f, 0x30, 0x9a,
	0xa2, 0x4f, 0x86, 0x70, 0x30, 0x11, 0x63, 0xc9, 0x70, 0x8f, 0x00, 0x04, 0x49, 0x99, 0xe7, 0x63,
	0x89, 0xbd, 0xe3, 0x3f, 0x3e, 0x8c, 0x84, 0x71, 0xf6, 0x62, 0x55, 0x9b, 0x6d, 0xbd, 0x0f, 0x3d,
	0x5a, 0x9c, 0xa3, 0x47, 0x02, 0xd8, 0xe3, 0x0a, 0xfd, 0x76, 0x90, 0x9b, 0x79, 0x57, 0x4b, 0xcd,
	0xf7, 0xf6, 0xdf, 0x6b, 0x8f, 0x70, 0xa5, 0x4f, 0x69, 0xc1, 0x71, 0x9f, 0xdc, 0x02, 0xe0, 0x4a,
	0x57, 0x92, 0xf2, 0x71, 0x91, 0xe1, 0xc1, 0x35, 0x9c, 0xd0, 0x2a, 0xc7, 0x80, 0x1c, 0xc1, 0x90,
	0x2b, 0x5d, 0x0a, 0x9a, 0xbc, 0x66, 0xd8, 0x6f, 0x97, 0x70, 0xa5, 0x53, 0xf6, 0x0e, 0x07, 0x64,
	0x04, 0x03, 0xae, 0x34, 0x7b, 0x53, 0x26, 0x2f, 0x71, 0x48, 0xee, 0xc2, 0x6d, 0xae, 0xb4, 0x2c,
	0x39, 0x2b, 0xce, 0x68, 0x22, 0x4b, 0x71, 0x8e, 0x40, 0x1e, 0xc2, 0xfd, 0x7f, 0x6d, 0xad, 0x4a,
	0xc9, 0xb4, 0xa4, 0x22, 0x63, 0xb2, 0xc2, 0x43, 0xf2, 0x18, 0x1e, 0xec, 0x18, 0xcd, 0x32, 0xc1,
	0x32, 0x2a, 0xbb, 0x54, 0x85, 0x23, 0x72, 0x0f, 0xee, 0xec, 0xf0, 0x19, 0x63, 0x29, 0x13, 0x15,
	0x1e, 0x9d, 0xbe, 0xfa, 0xb9, 0x0e, 0xfd, 0xab, 0x75, 0xe8, 0xff, 0x5e, 0x87, 0xfe, 0x8f, 0x4d,
	0xe8, 0x5d, 0x6d, 0x42, 0xef, 0xd7, 0x26, 0xf4, 0xde, 0x3f, 0xfb, 0xf8, 0xb9, 0xf9, 0x74, 0x31,
	0x8b, 0x6a, 0x3b, 0x8f, 0xaf, 0xad, 0x74, 0xcf, 0x53, 0xf7, 0xe1, 0x4b, 0xdc, 0x5c, 0x2e, 0xcd,
	0x0d, 0x4d, 0xb3, 0x60, 0x6b, 0xe7, 0xf9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb7, 0xb0, 0x9a,
	0x44, 0xc5, 0x01, 0x00, 0x00,
}
