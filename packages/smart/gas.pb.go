// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gas.proto

package smart

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

type PaymentType int32

const (
	PaymentType_INVALID          PaymentType = 0
	PaymentType_ContractCaller   PaymentType = 1
	PaymentType_ContractBinder   PaymentType = 2
	PaymentType_EcosystemAddress PaymentType = 3
)

var PaymentType_name = map[int32]string{
	0: "INVALID",
	1: "ContractCaller",
	2: "ContractBinder",
	3: "EcosystemAddress",
}

var PaymentType_value = map[string]int32{
	"INVALID":          0,
	"ContractCaller":   1,
	"ContractBinder":   2,
	"EcosystemAddress": 3,
}

func (x PaymentType) String() string {
	return proto.EnumName(PaymentType_name, int32(x))
}

func (PaymentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_df176b4a803aa869, []int{0}
}

type GasScenesType int32

const (
	GasScenesType_Unknown    GasScenesType = 0
	GasScenesType_Reward     GasScenesType = 1
	GasScenesType_Taxes      GasScenesType = 2
	GasScenesType_Direct     GasScenesType = 15
	GasScenesType_Combustion GasScenesType = 16
)

var GasScenesType_name = map[int32]string{
	0:  "Unknown",
	1:  "Reward",
	2:  "Taxes",
	15: "Direct",
	16: "Combustion",
}

var GasScenesType_value = map[string]int32{
	"Unknown":    0,
	"Reward":     1,
	"Taxes":      2,
	"Direct":     15,
	"Combustion": 16,
}

func (x GasScenesType) String() string {
	return proto.EnumName(GasScenesType_name, int32(x))
}

func (GasScenesType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_df176b4a803aa869, []int{1}
}

type GasPayAbleType int32

const (
	GasPayAbleType_Invalid GasPayAbleType = 0
	GasPayAbleType_Unable  GasPayAbleType = 1
	GasPayAbleType_Capable GasPayAbleType = 2
)

var GasPayAbleType_name = map[int32]string{
	0: "Invalid",
	1: "Unable",
	2: "Capable",
}

var GasPayAbleType_value = map[string]int32{
	"Invalid": 0,
	"Unable":  1,
	"Capable": 2,
}

func (x GasPayAbleType) String() string {
	return proto.EnumName(GasPayAbleType_name, int32(x))
}

func (GasPayAbleType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_df176b4a803aa869, []int{2}
}

type FuelType int32

const (
	FuelType_UNKNOWN      FuelType = 0
	FuelType_vmCost_fee   FuelType = 1
	FuelType_storage_fee  FuelType = 2
	FuelType_element_fee  FuelType = 3
	FuelType_expedite_fee FuelType = 4
)

var FuelType_name = map[int32]string{
	0: "UNKNOWN",
	1: "vmCost_fee",
	2: "storage_fee",
	3: "element_fee",
	4: "expedite_fee",
}

var FuelType_value = map[string]int32{
	"UNKNOWN":      0,
	"vmCost_fee":   1,
	"storage_fee":  2,
	"element_fee":  3,
	"expedite_fee": 4,
}

func (x FuelType) String() string {
	return proto.EnumName(FuelType_name, int32(x))
}

func (FuelType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_df176b4a803aa869, []int{3}
}

type Arithmetic int32

const (
	Arithmetic_NATIVE Arithmetic = 0
	Arithmetic_MUL    Arithmetic = 3
	Arithmetic_DIV    Arithmetic = 4
)

var Arithmetic_name = map[int32]string{
	0: "NATIVE",
	3: "MUL",
	4: "DIV",
}

var Arithmetic_value = map[string]int32{
	"NATIVE": 0,
	"MUL":    3,
	"DIV":    4,
}

func (x Arithmetic) String() string {
	return proto.EnumName(Arithmetic_name, int32(x))
}

func (Arithmetic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_df176b4a803aa869, []int{4}
}

func init() {
	proto.RegisterEnum("smart.PaymentType", PaymentType_name, PaymentType_value)
	proto.RegisterEnum("smart.GasScenesType", GasScenesType_name, GasScenesType_value)
	proto.RegisterEnum("smart.GasPayAbleType", GasPayAbleType_name, GasPayAbleType_value)
	proto.RegisterEnum("smart.FuelType", FuelType_name, FuelType_value)
	proto.RegisterEnum("smart.Arithmetic", Arithmetic_name, Arithmetic_value)
}

func init() { proto.RegisterFile("gas.proto", fileDescriptor_df176b4a803aa869) }

var fileDescriptor_df176b4a803aa869 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xed, 0xa4, 0x7f, 0xe8, 0x04, 0xd2, 0xd1, 0x8a, 0xb3, 0xef, 0x58, 0x6a, 0x73, 0x40,
	0xe2, 0xee, 0x38, 0xa5, 0xb2, 0x28, 0x6e, 0x05, 0x49, 0xa8, 0xe0, 0x80, 0xc6, 0xf6, 0xe0, 0xae,
	0x6a, 0xef, 0x5a, 0xbb, 0x9b, 0x36, 0x79, 0x0b, 0x1e, 0x8b, 0x63, 0x8f, 0x1c, 0x51, 0xf2, 0x22,
	0x68, 0x5d, 0x11, 0x71, 0xdb, 0xfd, 0xcd, 0xcc, 0x37, 0x9f, 0xe6, 0x83, 0x93, 0x9a, 0xec, 0x79,
	0x67, 0xb4, 0xd3, 0xe2, 0xd0, 0xb6, 0x64, 0x5c, 0x7c, 0x0b, 0xa3, 0x1b, 0xda, 0xb4, 0xac, 0xdc,
	0x7c, 0xd3, 0xb1, 0x18, 0xc1, 0x71, 0x96, 0x2f, 0x93, 0xab, 0x6c, 0x86, 0x81, 0x10, 0x30, 0x4e,
	0xb5, 0x72, 0x86, 0x4a, 0x97, 0x52, 0xd3, 0xb0, 0xc1, 0xf0, 0x7f, 0x36, 0x95, 0xaa, 0x62, 0x83,
	0x03, 0xf1, 0x1a, 0xf0, 0xa2, 0xd4, 0x76, 0x63, 0x1d, 0xb7, 0x49, 0x55, 0x19, 0xb6, 0x16, 0x87,
	0xf1, 0x35, 0xbc, 0xba, 0x24, 0xfb, 0xb9, 0x64, 0xc5, 0xf6, 0x9f, 0xf6, 0x42, 0xdd, 0x2b, 0xfd,
	0xa8, 0x30, 0x10, 0x00, 0x47, 0x9f, 0xf8, 0x91, 0x4c, 0x85, 0xa1, 0x38, 0x81, 0xc3, 0x39, 0xad,
	0xd9, 0xe2, 0xc0, 0xe3, 0x99, 0x34, 0x5c, 0x3a, 0x3c, 0x15, 0x63, 0x80, 0x54, 0xb7, 0xc5, 0xca,
	0x3a, 0xa9, 0x15, 0x62, 0xfc, 0x0e, 0xc6, 0x97, 0x64, 0x6f, 0x68, 0x93, 0x14, 0x0d, 0xef, 0xdd,
	0xaa, 0x07, 0x6a, 0x64, 0xf5, 0xac, 0xb8, 0x50, 0x54, 0x34, 0x8c, 0xa1, 0x2f, 0xa4, 0xd4, 0xf5,
	0x9f, 0x41, 0xfc, 0x0d, 0x5e, 0xbc, 0x5f, 0x71, 0xb3, 0xf7, 0x90, 0x7f, 0xc8, 0xaf, 0xbf, 0xe4,
	0x18, 0xf8, 0x05, 0x0f, 0x6d, 0xaa, 0xad, 0xfb, 0xfe, 0x83, 0xfd, 0xd4, 0x29, 0x8c, 0xac, 0xd3,
	0x86, 0x6a, 0xee, 0xc1, 0xc0, 0x03, 0x6e, 0xd8, 0x1f, 0xa7, 0x07, 0x43, 0x81, 0xf0, 0x92, 0xd7,
	0x1d, 0x57, 0xd2, 0x3d, 0xb7, 0x1c, 0xc4, 0x31, 0x40, 0x62, 0xa4, 0xbb, 0x6b, 0xd9, 0xc9, 0xd2,
	0x7b, 0xc8, 0x93, 0x79, 0xb6, 0xbc, 0xc0, 0x40, 0x1c, 0xc3, 0xf0, 0xe3, 0xe2, 0x0a, 0x87, 0xfe,
	0x31, 0xcb, 0x96, 0x78, 0x30, 0x4d, 0x7f, 0x6d, 0xa3, 0xf0, 0x69, 0x1b, 0x85, 0x7f, 0xb6, 0x51,
	0xf8, 0x73, 0x17, 0x05, 0x4f, 0xbb, 0x28, 0xf8, 0xbd, 0x8b, 0x82, 0xaf, 0x6f, 0x6a, 0xe9, 0xee,
	0x56, 0xc5, 0x79, 0xa9, 0xdb, 0x49, 0x36, 0x4d, 0x6e, 0xcf, 0xa4, 0x9e, 0xd4, 0xfa, 0x4c, 0x16,
	0xb4, 0x9e, 0x74, 0x54, 0xde, 0x53, 0xcd, 0x76, 0xd2, 0x07, 0x56, 0x1c, 0xf5, 0xf1, 0xbd, 0xfd,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0x16, 0x16, 0x89, 0x01, 0xcb, 0x01, 0x00, 0x00,
}