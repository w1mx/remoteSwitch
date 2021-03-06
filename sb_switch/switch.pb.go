// Code generated by protoc-gen-go. DO NOT EDIT.
// source: switch.proto

/*
Package shackbus is a generated protocol buffer package.

It is generated from these files:
	switch.proto

It has these top-level messages:
	None
	Terminal
	PortName
	PortRequest
	Port
	Device
*/
package shackbus

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type None struct {
}

func (m *None) Reset()                    { *m = None{} }
func (m *None) String() string            { return proto.CompactTextString(m) }
func (*None) ProtoMessage()               {}
func (*None) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Terminal struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Index int32  `protobuf:"varint,2,opt,name=index" json:"index,omitempty"`
	State bool   `protobuf:"varint,3,opt,name=state" json:"state,omitempty"`
}

func (m *Terminal) Reset()                    { *m = Terminal{} }
func (m *Terminal) String() string            { return proto.CompactTextString(m) }
func (*Terminal) ProtoMessage()               {}
func (*Terminal) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Terminal) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Terminal) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Terminal) GetState() bool {
	if m != nil {
		return m.State
	}
	return false
}

type PortName struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *PortName) Reset()                    { *m = PortName{} }
func (m *PortName) String() string            { return proto.CompactTextString(m) }
func (*PortName) ProtoMessage()               {}
func (*PortName) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PortName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PortRequest struct {
	Name      string      `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Terminals []*Terminal `protobuf:"bytes,2,rep,name=terminals" json:"terminals,omitempty"`
}

func (m *PortRequest) Reset()                    { *m = PortRequest{} }
func (m *PortRequest) String() string            { return proto.CompactTextString(m) }
func (*PortRequest) ProtoMessage()               {}
func (*PortRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PortRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PortRequest) GetTerminals() []*Terminal {
	if m != nil {
		return m.Terminals
	}
	return nil
}

type Port struct {
	Name      string      `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Index     int32       `protobuf:"varint,2,opt,name=index" json:"index,omitempty"`
	Terminals []*Terminal `protobuf:"bytes,3,rep,name=terminals" json:"terminals,omitempty"`
	Exclusive bool        `protobuf:"varint,4,opt,name=exclusive" json:"exclusive,omitempty"`
}

func (m *Port) Reset()                    { *m = Port{} }
func (m *Port) String() string            { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()               {}
func (*Port) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Port) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Port) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Port) GetTerminals() []*Terminal {
	if m != nil {
		return m.Terminals
	}
	return nil
}

func (m *Port) GetExclusive() bool {
	if m != nil {
		return m.Exclusive
	}
	return false
}

type Device struct {
	Name      string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Index     int32   `protobuf:"varint,2,opt,name=index" json:"index,omitempty"`
	Ports     []*Port `protobuf:"bytes,3,rep,name=ports" json:"ports,omitempty"`
	Exclusive bool    `protobuf:"varint,4,opt,name=exclusive" json:"exclusive,omitempty"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Device) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Device) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Device) GetPorts() []*Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *Device) GetExclusive() bool {
	if m != nil {
		return m.Exclusive
	}
	return false
}

func init() {
	proto.RegisterType((*None)(nil), "shackbus.None")
	proto.RegisterType((*Terminal)(nil), "shackbus.Terminal")
	proto.RegisterType((*PortName)(nil), "shackbus.PortName")
	proto.RegisterType((*PortRequest)(nil), "shackbus.PortRequest")
	proto.RegisterType((*Port)(nil), "shackbus.Port")
	proto.RegisterType((*Device)(nil), "shackbus.Device")
}

func init() { proto.RegisterFile("switch.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xdb, 0x6a, 0xb3, 0x40,
	0x10, 0xc7, 0xd9, 0x78, 0x88, 0x4e, 0x3e, 0x3e, 0xca, 0xd2, 0x82, 0x84, 0x52, 0x64, 0xe9, 0x85,
	0x37, 0x91, 0x90, 0xbe, 0x42, 0x21, 0xd0, 0x8b, 0x50, 0xb4, 0x2f, 0xa0, 0x76, 0x20, 0xd2, 0x44,
	0x53, 0x77, 0xb4, 0xb9, 0xec, 0x73, 0xf4, 0x69, 0xcb, 0xba, 0x8a, 0x8d, 0x84, 0xd2, 0xdc, 0xed,
	0xec, 0x1c, 0x7e, 0xff, 0x39, 0xc0, 0x3f, 0xf9, 0x91, 0x53, 0xb6, 0x0d, 0x0f, 0x55, 0x49, 0x25,
	0x77, 0xe4, 0x36, 0xc9, 0xde, 0xd2, 0x5a, 0x0a, 0x1b, 0xcc, 0x4d, 0x59, 0xa0, 0x78, 0x02, 0xe7,
	0x05, 0xab, 0x7d, 0x5e, 0x24, 0x3b, 0xce, 0xc1, 0x2c, 0x92, 0x3d, 0x7a, 0xcc, 0x67, 0x81, 0x1b,
	0xb5, 0x6f, 0x7e, 0x0d, 0x56, 0x5e, 0xbc, 0xe2, 0xd1, 0x9b, 0xf8, 0x2c, 0xb0, 0x22, 0x6d, 0xa8,
	0x5f, 0x49, 0x09, 0xa1, 0x67, 0xf8, 0x2c, 0x70, 0x22, 0x6d, 0x88, 0x3b, 0x70, 0x9e, 0xcb, 0x8a,
	0x36, 0x2a, 0xef, 0x4c, 0x2d, 0x11, 0xc3, 0x4c, 0xf9, 0x23, 0x7c, 0xaf, 0x51, 0xd2, 0x59, 0xdc,
	0x12, 0x5c, 0xea, 0xe4, 0x48, 0x6f, 0xe2, 0x1b, 0xc1, 0x6c, 0xc5, 0xc3, 0x5e, 0x74, 0xd8, 0x2b,
	0x8d, 0x86, 0x20, 0xf1, 0xc9, 0xc0, 0x54, 0x55, 0x2f, 0x50, 0x7f, 0x02, 0x31, 0xfe, 0x00, 0xe1,
	0xb7, 0xe0, 0xe2, 0x31, 0xdb, 0xd5, 0x32, 0x6f, 0xd0, 0x33, 0xdb, 0x9e, 0x87, 0x0f, 0xd1, 0x80,
	0xfd, 0x88, 0x4d, 0x9e, 0xe1, 0x05, 0x1a, 0xee, 0xc1, 0x3a, 0x94, 0x15, 0xf5, 0xfc, 0xff, 0x03,
	0xbf, 0x1d, 0x91, 0x76, 0xfe, 0xce, 0x5d, 0x7d, 0x31, 0x70, 0xe2, 0x34, 0x6e, 0x17, 0xcc, 0x17,
	0x30, 0x5d, 0x23, 0xe9, 0x49, 0x9c, 0x16, 0x53, 0xfb, 0x98, 0x8f, 0x00, 0x7c, 0x09, 0xd3, 0xb8,
	0x0b, 0xbf, 0x19, 0xb1, 0xf5, 0x7a, 0x7e, 0x66, 0xa8, 0x4b, 0xe1, 0x0b, 0x70, 0xd7, 0x48, 0x5d,
	0xa3, 0x23, 0xe7, 0xfc, 0x6a, 0xb0, 0x75, 0x44, 0x6a, 0xb7, 0x17, 0xf7, 0xf0, 0x1d, 0x00, 0x00,
	0xff, 0xff, 0x5e, 0x38, 0xc8, 0xcd, 0x81, 0x02, 0x00, 0x00,
}
