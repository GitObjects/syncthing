// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ext.proto

package ext

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
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

var E_XmlTags = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         74001,
	Name:          "ext.xml_tags",
	Tag:           "varint,74001,opt,name=xml_tags",
	Filename:      "ext.proto",
}

var E_Xml = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         75005,
	Name:          "ext.xml",
	Tag:           "bytes,75005,opt,name=xml",
	Filename:      "ext.proto",
}

var E_Json = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         75006,
	Name:          "ext.json",
	Tag:           "bytes,75006,opt,name=json",
	Filename:      "ext.proto",
}

var E_Default = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         75007,
	Name:          "ext.default",
	Tag:           "bytes,75007,opt,name=default",
	Filename:      "ext.proto",
}

var E_Restart = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         75008,
	Name:          "ext.restart",
	Tag:           "varint,75008,opt,name=restart",
	Filename:      "ext.proto",
}

var E_DeviceId = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         75009,
	Name:          "ext.device_id",
	Tag:           "varint,75009,opt,name=device_id",
	Filename:      "ext.proto",
}

var E_Goname = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         75010,
	Name:          "ext.goname",
	Tag:           "bytes,75010,opt,name=goname",
	Filename:      "ext.proto",
}

var E_Enumgoname = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         76010,
	Name:          "ext.enumgoname",
	Tag:           "bytes,76010,opt,name=enumgoname",
	Filename:      "ext.proto",
}

func init() {
	proto.RegisterExtension(E_XmlTags)
	proto.RegisterExtension(E_Xml)
	proto.RegisterExtension(E_Json)
	proto.RegisterExtension(E_Default)
	proto.RegisterExtension(E_Restart)
	proto.RegisterExtension(E_DeviceId)
	proto.RegisterExtension(E_Goname)
	proto.RegisterExtension(E_Enumgoname)
}

func init() { proto.RegisterFile("ext.proto", fileDescriptor_95fe6908ffcf64d3) }

var fileDescriptor_95fe6908ffcf64d3 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xd1, 0x3d, 0x4b, 0x03, 0x31,
	0x18, 0xc0, 0x71, 0x8e, 0x16, 0xdb, 0x66, 0xec, 0x24, 0x82, 0xb5, 0x6e, 0x9d, 0xee, 0x10, 0x41,
	0x31, 0x74, 0x52, 0x14, 0x1c, 0x44, 0x28, 0xe2, 0xe0, 0x52, 0xd2, 0xcb, 0xd3, 0x34, 0x92, 0x97,
	0x72, 0x49, 0x24, 0x6e, 0xea, 0x37, 0xf0, 0x2b, 0x39, 0x69, 0x27, 0xfd, 0x06, 0xd2, 0xd1, 0xef,
	0xe0, 0x0b, 0x77, 0x97, 0x93, 0x42, 0x87, 0xdb, 0x42, 0xf8, 0xff, 0x1e, 0x12, 0x1e, 0xd4, 0x01,
	0x6f, 0xe3, 0x79, 0xa6, 0xad, 0xee, 0x36, 0xc0, 0xdb, 0xad, 0x3e, 0xd3, 0x9a, 0x09, 0x48, 0x8a,
	0xab, 0x89, 0x9b, 0x26, 0x14, 0x4c, 0x9a, 0xf1, 0xb9, 0xd5, 0x59, 0x99, 0xe1, 0x21, 0x6a, 0x7b,
	0x29, 0xc6, 0x96, 0x30, 0xd3, 0xdd, 0x89, 0xcb, 0x3c, 0xae, 0xf2, 0xf8, 0x02, 0x8c, 0x21, 0x0c,
	0x2e, 0xe7, 0x96, 0x6b, 0x65, 0x36, 0x9f, 0x5f, 0x9a, 0xfd, 0x68, 0xd0, 0x1e, 0xb5, 0xbc, 0x14,
	0x57, 0x84, 0x19, 0xbc, 0x87, 0x1a, 0x5e, 0x8a, 0xee, 0xf6, 0x1a, 0x3c, 0xe3, 0x20, 0x68, 0xc5,
	0xbe, 0xdf, 0x72, 0xd6, 0x19, 0xe5, 0x2d, 0xde, 0x47, 0xcd, 0x5b, 0xa3, 0x55, 0x9d, 0xf9, 0x09,
	0xa6, 0x88, 0xf1, 0x11, 0x6a, 0x51, 0x98, 0x12, 0x27, 0x6c, 0x9d, 0xfb, 0x0d, 0xae, 0xea, 0x73,
	0x9a, 0x81, 0xb1, 0x24, 0xab, 0xa5, 0x0f, 0x8b, 0xf0, 0xbb, 0xd0, 0xe3, 0x21, 0xea, 0x50, 0xb8,
	0xe3, 0x29, 0x8c, 0x39, 0xad, 0xc3, 0x8f, 0x01, 0xb7, 0x4b, 0x71, 0x4e, 0xf1, 0x21, 0xda, 0x60,
	0x5a, 0x11, 0x09, 0x75, 0xf4, 0x69, 0x51, 0x3e, 0x39, 0xe4, 0xf8, 0x04, 0x21, 0x50, 0x4e, 0x06,
	0xbc, 0xbb, 0x86, 0x4f, 0x95, 0x93, 0xd7, 0x44, 0xb8, 0xff, 0xb5, 0x7c, 0x7d, 0x94, 0x03, 0x56,
	0xd8, 0xf1, 0xc1, 0xeb, 0xb2, 0x17, 0xbd, 0x2f, 0x7b, 0xd1, 0xe7, 0xb2, 0x17, 0xdd, 0x0c, 0x18,
	0xb7, 0x33, 0x37, 0x89, 0x53, 0x2d, 0x13, 0x73, 0xaf, 0x52, 0x3b, 0xe3, 0x8a, 0xad, 0x9c, 0x8a,
	0xd9, 0x09, 0x78, 0xfb, 0x17, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x40, 0x4c, 0x73, 0x42, 0x02, 0x00,
	0x00,
}
