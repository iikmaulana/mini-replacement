// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vehicle_history.proto

package packets

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	anypb "google.golang.org/protobuf/types/known/anypb"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type VehHistoryRequest struct {
	Data                 *anypb.Any `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *VehHistoryRequest) Reset()         { *m = VehHistoryRequest{} }
func (m *VehHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*VehHistoryRequest) ProtoMessage()    {}
func (*VehHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a111f33d71f9ac4e, []int{0}
}

func (m *VehHistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VehHistoryRequest.Unmarshal(m, b)
}
func (m *VehHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VehHistoryRequest.Marshal(b, m, deterministic)
}
func (m *VehHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VehHistoryRequest.Merge(m, src)
}
func (m *VehHistoryRequest) XXX_Size() int {
	return xxx_messageInfo_VehHistoryRequest.Size(m)
}
func (m *VehHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VehHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VehHistoryRequest proto.InternalMessageInfo

func (m *VehHistoryRequest) GetData() *anypb.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

type VehHistoryRequestByChassis struct {
	ChassisNumber        string   `protobuf:"bytes,1,opt,name=ChassisNumber,proto3" json:"ChassisNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VehHistoryRequestByChassis) Reset()         { *m = VehHistoryRequestByChassis{} }
func (m *VehHistoryRequestByChassis) String() string { return proto.CompactTextString(m) }
func (*VehHistoryRequestByChassis) ProtoMessage()    {}
func (*VehHistoryRequestByChassis) Descriptor() ([]byte, []int) {
	return fileDescriptor_a111f33d71f9ac4e, []int{1}
}

func (m *VehHistoryRequestByChassis) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VehHistoryRequestByChassis.Unmarshal(m, b)
}
func (m *VehHistoryRequestByChassis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VehHistoryRequestByChassis.Marshal(b, m, deterministic)
}
func (m *VehHistoryRequestByChassis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VehHistoryRequestByChassis.Merge(m, src)
}
func (m *VehHistoryRequestByChassis) XXX_Size() int {
	return xxx_messageInfo_VehHistoryRequestByChassis.Size(m)
}
func (m *VehHistoryRequestByChassis) XXX_DiscardUnknown() {
	xxx_messageInfo_VehHistoryRequestByChassis.DiscardUnknown(m)
}

var xxx_messageInfo_VehHistoryRequestByChassis proto.InternalMessageInfo

func (m *VehHistoryRequestByChassis) GetChassisNumber() string {
	if m != nil {
		return m.ChassisNumber
	}
	return ""
}

type VehHistoryRequestByImei struct {
	Imei                 string   `protobuf:"bytes,1,opt,name=Imei,proto3" json:"Imei,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VehHistoryRequestByImei) Reset()         { *m = VehHistoryRequestByImei{} }
func (m *VehHistoryRequestByImei) String() string { return proto.CompactTextString(m) }
func (*VehHistoryRequestByImei) ProtoMessage()    {}
func (*VehHistoryRequestByImei) Descriptor() ([]byte, []int) {
	return fileDescriptor_a111f33d71f9ac4e, []int{2}
}

func (m *VehHistoryRequestByImei) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VehHistoryRequestByImei.Unmarshal(m, b)
}
func (m *VehHistoryRequestByImei) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VehHistoryRequestByImei.Marshal(b, m, deterministic)
}
func (m *VehHistoryRequestByImei) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VehHistoryRequestByImei.Merge(m, src)
}
func (m *VehHistoryRequestByImei) XXX_Size() int {
	return xxx_messageInfo_VehHistoryRequestByImei.Size(m)
}
func (m *VehHistoryRequestByImei) XXX_DiscardUnknown() {
	xxx_messageInfo_VehHistoryRequestByImei.DiscardUnknown(m)
}

var xxx_messageInfo_VehHistoryRequestByImei proto.InternalMessageInfo

func (m *VehHistoryRequestByImei) GetImei() string {
	if m != nil {
		return m.Imei
	}
	return ""
}

type VehHistoryRequestByDealerID struct {
	DealerID             string   `protobuf:"bytes,1,opt,name=DealerID,proto3" json:"DealerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VehHistoryRequestByDealerID) Reset()         { *m = VehHistoryRequestByDealerID{} }
func (m *VehHistoryRequestByDealerID) String() string { return proto.CompactTextString(m) }
func (*VehHistoryRequestByDealerID) ProtoMessage()    {}
func (*VehHistoryRequestByDealerID) Descriptor() ([]byte, []int) {
	return fileDescriptor_a111f33d71f9ac4e, []int{3}
}

func (m *VehHistoryRequestByDealerID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VehHistoryRequestByDealerID.Unmarshal(m, b)
}
func (m *VehHistoryRequestByDealerID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VehHistoryRequestByDealerID.Marshal(b, m, deterministic)
}
func (m *VehHistoryRequestByDealerID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VehHistoryRequestByDealerID.Merge(m, src)
}
func (m *VehHistoryRequestByDealerID) XXX_Size() int {
	return xxx_messageInfo_VehHistoryRequestByDealerID.Size(m)
}
func (m *VehHistoryRequestByDealerID) XXX_DiscardUnknown() {
	xxx_messageInfo_VehHistoryRequestByDealerID.DiscardUnknown(m)
}

var xxx_messageInfo_VehHistoryRequestByDealerID proto.InternalMessageInfo

func (m *VehHistoryRequestByDealerID) GetDealerID() string {
	if m != nil {
		return m.DealerID
	}
	return ""
}

type VehHistoryOutput struct {
	Status               int32      `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Data                 *anypb.Any `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *VehHistoryOutput) Reset()         { *m = VehHistoryOutput{} }
func (m *VehHistoryOutput) String() string { return proto.CompactTextString(m) }
func (*VehHistoryOutput) ProtoMessage()    {}
func (*VehHistoryOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_a111f33d71f9ac4e, []int{4}
}

func (m *VehHistoryOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VehHistoryOutput.Unmarshal(m, b)
}
func (m *VehHistoryOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VehHistoryOutput.Marshal(b, m, deterministic)
}
func (m *VehHistoryOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VehHistoryOutput.Merge(m, src)
}
func (m *VehHistoryOutput) XXX_Size() int {
	return xxx_messageInfo_VehHistoryOutput.Size(m)
}
func (m *VehHistoryOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_VehHistoryOutput.DiscardUnknown(m)
}

var xxx_messageInfo_VehHistoryOutput proto.InternalMessageInfo

func (m *VehHistoryOutput) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *VehHistoryOutput) GetData() *anypb.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*VehHistoryRequest)(nil), "packets.VehHistoryRequest")
	proto.RegisterType((*VehHistoryRequestByChassis)(nil), "packets.VehHistoryRequestByChassis")
	proto.RegisterType((*VehHistoryRequestByImei)(nil), "packets.VehHistoryRequestByImei")
	proto.RegisterType((*VehHistoryRequestByDealerID)(nil), "packets.VehHistoryRequestByDealerID")
	proto.RegisterType((*VehHistoryOutput)(nil), "packets.VehHistoryOutput")
}

func init() { proto.RegisterFile("vehicle_history.proto", fileDescriptor_a111f33d71f9ac4e) }

var fileDescriptor_a111f33d71f9ac4e = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xdf, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0x19, 0x1a, 0x83, 0xdd, 0x04, 0x02, 0xf3, 0x6b, 0x0d, 0x62, 0xaa, 0xb2, 0x21, 0xed,
	0x65, 0x99, 0x34, 0x9e, 0x10, 0xe2, 0xa1, 0x3f, 0x10, 0xf4, 0x61, 0x9b, 0x14, 0x50, 0x25, 0x40,
	0xa8, 0xb8, 0xc9, 0x51, 0x5b, 0xeb, 0xec, 0x10, 0x3b, 0x43, 0xf9, 0xdb, 0x79, 0x41, 0xd8, 0x6e,
	0xab, 0x25, 0x69, 0x26, 0x25, 0x3c, 0xc5, 0xb1, 0xef, 0x3e, 0x77, 0xfe, 0xde, 0x9d, 0xe1, 0xe9,
	0x15, 0x32, 0x1e, 0xcd, 0x71, 0xc2, 0xb8, 0xd2, 0x32, 0xcd, 0x83, 0x24, 0x95, 0x5a, 0x92, 0xbb,
	0x09, 0x8d, 0x2e, 0x50, 0x2b, 0xaf, 0x33, 0x93, 0x72, 0x36, 0xc7, 0x63, 0xb3, 0x3d, 0xcd, 0x7e,
	0x1e, 0x53, 0xe1, 0x6c, 0xfc, 0x77, 0xf0, 0x68, 0x8c, 0xec, 0xa3, 0xf5, 0x0b, 0xf1, 0x57, 0x86,
	0x4a, 0x93, 0x43, 0xd8, 0x1c, 0x52, 0x4d, 0x77, 0x37, 0xba, 0x1b, 0x87, 0x3b, 0x27, 0x4f, 0x02,
	0xeb, 0x1e, 0x2c, 0xdc, 0x83, 0x9e, 0xc8, 0x43, 0x63, 0xe1, 0xf7, 0xc1, 0x2b, 0xb9, 0xf7, 0xf3,
	0x01, 0xa3, 0x4a, 0x71, 0x45, 0x0e, 0xe0, 0xbe, 0x5b, 0x9e, 0x65, 0x97, 0x53, 0x4c, 0x0d, 0x70,
	0x3b, 0xbc, 0xbe, 0xe9, 0x1f, 0xc1, 0xf3, 0x0a, 0xc6, 0xe8, 0x12, 0x39, 0x21, 0xb0, 0xf9, 0xef,
	0xeb, 0xfc, 0xcc, 0xda, 0x7f, 0x03, 0x2f, 0x2a, 0xcc, 0x87, 0x48, 0xe7, 0x98, 0x8e, 0x86, 0xc4,
	0x83, 0x7b, 0x8b, 0xb5, 0x73, 0x5b, 0xfe, 0xfb, 0x9f, 0xe1, 0xe1, 0xca, 0xf5, 0x3c, 0xd3, 0x49,
	0xa6, 0xc9, 0x33, 0xd8, 0xfa, 0xa4, 0xa9, 0xce, 0x94, 0xb1, 0xbe, 0x13, 0xba, 0xbf, 0xa5, 0x06,
	0xb7, 0x6f, 0xd2, 0xe0, 0xe4, 0x0f, 0xc0, 0x83, 0xb1, 0x2d, 0x80, 0x43, 0x93, 0x53, 0x20, 0xd7,
	0x77, 0x26, 0x1f, 0x50, 0x13, 0x2f, 0x70, 0x05, 0x09, 0x4a, 0x17, 0xf0, 0x3a, 0x15, 0x67, 0x36,
	0x43, 0xff, 0x16, 0x09, 0x8d, 0x42, 0x05, 0xdc, 0x40, 0x66, 0xa2, 0x05, 0x73, 0x6c, 0x64, 0x2c,
	0x30, 0x57, 0xa5, 0x6b, 0xcc, 0xfd, 0x02, 0xbb, 0x55, 0x5c, 0x53, 0xce, 0xee, 0x7a, 0xa8, 0xb5,
	0xb8, 0x09, 0xbd, 0xb7, 0x46, 0x86, 0x7e, 0x7e, 0x2a, 0x85, 0x66, 0xcd, 0xb3, 0xfe, 0x06, 0xdd,
	0xb5, 0xe8, 0xf3, 0xdf, 0xc2, 0x76, 0x56, 0x53, 0x38, 0x85, 0x97, 0x05, 0xf8, 0x80, 0x61, 0x74,
	0xb1, 0x12, 0x7b, 0xbf, 0x4e, 0x17, 0x67, 0x54, 0x1f, 0x82, 0xc1, 0x7e, 0x39, 0x7f, 0x9b, 0x78,
	0xbc, 0x1c, 0x8e, 0x98, 0x1c, 0xd4, 0x05, 0x5a, 0x8c, 0x49, 0x7d, 0xa4, 0x14, 0x8e, 0xca, 0x91,
	0xc6, 0xc8, 0x7a, 0x91, 0xe6, 0x57, 0x38, 0x12, 0xd4, 0x7c, 0xff, 0x6f, 0xcc, 0x49, 0xa9, 0x57,
	0x07, 0x8c, 0x8a, 0x19, 0xba, 0x51, 0x6d, 0x2f, 0xdf, 0x0f, 0x78, 0x55, 0x19, 0xc0, 0x5c, 0x8a,
	0x6a, 0x2e, 0x85, 0x0b, 0xd5, 0xb8, 0x07, 0xca, 0x2f, 0x42, 0x2f, 0x8e, 0x9b, 0xe3, 0xce, 0xe0,
	0x71, 0x01, 0xf7, 0x3e, 0xe6, 0xba, 0xcd, 0xd4, 0xee, 0x55, 0x0a, 0x60, 0x5a, 0x48, 0x31, 0x9e,
	0x34, 0x47, 0x7f, 0x87, 0x4e, 0x01, 0x3d, 0x44, 0x6a, 0x75, 0xc5, 0xf6, 0xa5, 0xeb, 0xef, 0x7c,
	0xdd, 0x0e, 0xde, 0xba, 0xf3, 0xe9, 0x96, 0x79, 0x9e, 0x5f, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0x18, 0x14, 0x7c, 0x26, 0x11, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VehicleHistoryClient is the client API for VehicleHistory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VehicleHistoryClient interface {
	VehicleHistory_Get(ctx context.Context, in *VehHistoryRequest, opts ...grpc.CallOption) (*VehHistoryOutput, error)
	VehicleHistory_GetCount(ctx context.Context, in *VehHistoryRequest, opts ...grpc.CallOption) (*VehHistoryOutput, error)
	VehicleHistory_GetByChassis(ctx context.Context, in *VehHistoryRequest, opts ...grpc.CallOption) (*VehHistoryOutput, error)
	VehicleHistory_GetByImei(ctx context.Context, in *VehHistoryRequestByImei, opts ...grpc.CallOption) (*VehHistoryOutput, error)
	VehicleHistory_GetCountByMonth(ctx context.Context, in *VehHistoryRequest, opts ...grpc.CallOption) (*VehHistoryOutput, error)
	VehicleHistory_GetCountByOwnerID(ctx context.Context, in *VehHistoryRequest, opts ...grpc.CallOption) (*VehHistoryOutput, error)
	VehicleHistory_CheckByChassis(ctx context.Context, in *VehHistoryRequestByChassis, opts ...grpc.CallOption) (*VehHistoryOutput, error)
	VehicleHistory_GetOwnerIdByDealerId(ctx context.Context, in *VehHistoryRequestByDealerID, opts ...grpc.CallOption) (*VehHistoryOutput, error)
	VehicleHistory_GetVehActiveInactiveByDealerId(ctx context.Context, in *VehHistoryRequestByDealerID, opts ...grpc.CallOption) (*VehHistoryOutput, error)
	VehicleHistory_ChangeStatus(ctx context.Context, in *VehHistoryRequestByChassis, opts ...grpc.CallOption) (*VehHistoryOutput, error)
	VehicleHistory_ChangeActivationStatus(ctx context.Context, in *VehHistoryRequest, opts ...grpc.CallOption) (*VehHistoryOutput, error)
	VehicleHistory_Add(ctx context.Context, in *VehHistoryRequest, opts ...grpc.CallOption) (*VehHistoryOutput, error)
	VehicleHistory_Edit(ctx context.Context, in *VehHistoryRequest, opts ...grpc.CallOption) (*VehHistoryOutput, error)
	VehicleHistory_ChangeOwnership(ctx context.Context, in *VehHistoryRequest, opts ...grpc.CallOption) (*VehHistoryOutput, error)
	VehicleHistory_Deactivate(ctx context.Context, in *VehHistoryRequestByChassis, opts ...grpc.CallOption) (*VehHistoryOutput, error)
}

type vehicleHistoryClient struct {
	cc *grpc.ClientConn
}

func NewVehicleHistoryClient(cc *grpc.ClientConn) VehicleHistoryClient {
	return &vehicleHistoryClient{cc}
}

func (c *vehicleHistoryClient) VehicleHistory_Get(ctx context.Context, in *VehHistoryRequest, opts ...grpc.CallOption) (*VehHistoryOutput, error) {
	out := new(VehHistoryOutput)
	err := c.cc.Invoke(ctx, "/packets.VehicleHistory/VehicleHistory_Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleHistoryClient) VehicleHistory_GetCount(ctx context.Context, in *VehHistoryRequest, opts ...grpc.CallOption) (*VehHistoryOutput, error) {
	out := new(VehHistoryOutput)
	err := c.cc.Invoke(ctx, "/packets.VehicleHistory/VehicleHistory_GetCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleHistoryClient) VehicleHistory_GetByChassis(ctx context.Context, in *VehHistoryRequest, opts ...grpc.CallOption) (*VehHistoryOutput, error) {
	out := new(VehHistoryOutput)
	err := c.cc.Invoke(ctx, "/packets.VehicleHistory/VehicleHistory_GetByChassis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleHistoryClient) VehicleHistory_GetByImei(ctx context.Context, in *VehHistoryRequestByImei, opts ...grpc.CallOption) (*VehHistoryOutput, error) {
	out := new(VehHistoryOutput)
	err := c.cc.Invoke(ctx, "/packets.VehicleHistory/VehicleHistory_GetByImei", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleHistoryClient) VehicleHistory_GetCountByMonth(ctx context.Context, in *VehHistoryRequest, opts ...grpc.CallOption) (*VehHistoryOutput, error) {
	out := new(VehHistoryOutput)
	err := c.cc.Invoke(ctx, "/packets.VehicleHistory/VehicleHistory_GetCountByMonth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleHistoryClient) VehicleHistory_GetCountByOwnerID(ctx context.Context, in *VehHistoryRequest, opts ...grpc.CallOption) (*VehHistoryOutput, error) {
	out := new(VehHistoryOutput)
	err := c.cc.Invoke(ctx, "/packets.VehicleHistory/VehicleHistory_GetCountByOwnerID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleHistoryClient) VehicleHistory_CheckByChassis(ctx context.Context, in *VehHistoryRequestByChassis, opts ...grpc.CallOption) (*VehHistoryOutput, error) {
	out := new(VehHistoryOutput)
	err := c.cc.Invoke(ctx, "/packets.VehicleHistory/VehicleHistory_CheckByChassis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleHistoryClient) VehicleHistory_GetOwnerIdByDealerId(ctx context.Context, in *VehHistoryRequestByDealerID, opts ...grpc.CallOption) (*VehHistoryOutput, error) {
	out := new(VehHistoryOutput)
	err := c.cc.Invoke(ctx, "/packets.VehicleHistory/VehicleHistory_GetOwnerIdByDealerId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleHistoryClient) VehicleHistory_GetVehActiveInactiveByDealerId(ctx context.Context, in *VehHistoryRequestByDealerID, opts ...grpc.CallOption) (*VehHistoryOutput, error) {
	out := new(VehHistoryOutput)
	err := c.cc.Invoke(ctx, "/packets.VehicleHistory/VehicleHistory_GetVehActiveInactiveByDealerId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleHistoryClient) VehicleHistory_ChangeStatus(ctx context.Context, in *VehHistoryRequestByChassis, opts ...grpc.CallOption) (*VehHistoryOutput, error) {
	out := new(VehHistoryOutput)
	err := c.cc.Invoke(ctx, "/packets.VehicleHistory/VehicleHistory_ChangeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleHistoryClient) VehicleHistory_ChangeActivationStatus(ctx context.Context, in *VehHistoryRequest, opts ...grpc.CallOption) (*VehHistoryOutput, error) {
	out := new(VehHistoryOutput)
	err := c.cc.Invoke(ctx, "/packets.VehicleHistory/VehicleHistory_ChangeActivationStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleHistoryClient) VehicleHistory_Add(ctx context.Context, in *VehHistoryRequest, opts ...grpc.CallOption) (*VehHistoryOutput, error) {
	out := new(VehHistoryOutput)
	err := c.cc.Invoke(ctx, "/packets.VehicleHistory/VehicleHistory_Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleHistoryClient) VehicleHistory_Edit(ctx context.Context, in *VehHistoryRequest, opts ...grpc.CallOption) (*VehHistoryOutput, error) {
	out := new(VehHistoryOutput)
	err := c.cc.Invoke(ctx, "/packets.VehicleHistory/VehicleHistory_Edit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleHistoryClient) VehicleHistory_ChangeOwnership(ctx context.Context, in *VehHistoryRequest, opts ...grpc.CallOption) (*VehHistoryOutput, error) {
	out := new(VehHistoryOutput)
	err := c.cc.Invoke(ctx, "/packets.VehicleHistory/VehicleHistory_ChangeOwnership", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleHistoryClient) VehicleHistory_Deactivate(ctx context.Context, in *VehHistoryRequestByChassis, opts ...grpc.CallOption) (*VehHistoryOutput, error) {
	out := new(VehHistoryOutput)
	err := c.cc.Invoke(ctx, "/packets.VehicleHistory/VehicleHistory_Deactivate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VehicleHistoryServer is the server API for VehicleHistory service.
type VehicleHistoryServer interface {
	VehicleHistory_Get(context.Context, *VehHistoryRequest) (*VehHistoryOutput, error)
	VehicleHistory_GetCount(context.Context, *VehHistoryRequest) (*VehHistoryOutput, error)
	VehicleHistory_GetByChassis(context.Context, *VehHistoryRequest) (*VehHistoryOutput, error)
	VehicleHistory_GetByImei(context.Context, *VehHistoryRequestByImei) (*VehHistoryOutput, error)
	VehicleHistory_GetCountByMonth(context.Context, *VehHistoryRequest) (*VehHistoryOutput, error)
	VehicleHistory_GetCountByOwnerID(context.Context, *VehHistoryRequest) (*VehHistoryOutput, error)
	VehicleHistory_CheckByChassis(context.Context, *VehHistoryRequestByChassis) (*VehHistoryOutput, error)
	VehicleHistory_GetOwnerIdByDealerId(context.Context, *VehHistoryRequestByDealerID) (*VehHistoryOutput, error)
	VehicleHistory_GetVehActiveInactiveByDealerId(context.Context, *VehHistoryRequestByDealerID) (*VehHistoryOutput, error)
	VehicleHistory_ChangeStatus(context.Context, *VehHistoryRequestByChassis) (*VehHistoryOutput, error)
	VehicleHistory_ChangeActivationStatus(context.Context, *VehHistoryRequest) (*VehHistoryOutput, error)
	VehicleHistory_Add(context.Context, *VehHistoryRequest) (*VehHistoryOutput, error)
	VehicleHistory_Edit(context.Context, *VehHistoryRequest) (*VehHistoryOutput, error)
	VehicleHistory_ChangeOwnership(context.Context, *VehHistoryRequest) (*VehHistoryOutput, error)
	VehicleHistory_Deactivate(context.Context, *VehHistoryRequestByChassis) (*VehHistoryOutput, error)
}

// UnimplementedVehicleHistoryServer can be embedded to have forward compatible implementations.
type UnimplementedVehicleHistoryServer struct {
}

func (*UnimplementedVehicleHistoryServer) VehicleHistory_Get(ctx context.Context, req *VehHistoryRequest) (*VehHistoryOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VehicleHistory_Get not implemented")
}
func (*UnimplementedVehicleHistoryServer) VehicleHistory_GetCount(ctx context.Context, req *VehHistoryRequest) (*VehHistoryOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VehicleHistory_GetCount not implemented")
}
func (*UnimplementedVehicleHistoryServer) VehicleHistory_GetByChassis(ctx context.Context, req *VehHistoryRequest) (*VehHistoryOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VehicleHistory_GetByChassis not implemented")
}
func (*UnimplementedVehicleHistoryServer) VehicleHistory_GetByImei(ctx context.Context, req *VehHistoryRequestByImei) (*VehHistoryOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VehicleHistory_GetByImei not implemented")
}
func (*UnimplementedVehicleHistoryServer) VehicleHistory_GetCountByMonth(ctx context.Context, req *VehHistoryRequest) (*VehHistoryOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VehicleHistory_GetCountByMonth not implemented")
}
func (*UnimplementedVehicleHistoryServer) VehicleHistory_GetCountByOwnerID(ctx context.Context, req *VehHistoryRequest) (*VehHistoryOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VehicleHistory_GetCountByOwnerID not implemented")
}
func (*UnimplementedVehicleHistoryServer) VehicleHistory_CheckByChassis(ctx context.Context, req *VehHistoryRequestByChassis) (*VehHistoryOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VehicleHistory_CheckByChassis not implemented")
}
func (*UnimplementedVehicleHistoryServer) VehicleHistory_GetOwnerIdByDealerId(ctx context.Context, req *VehHistoryRequestByDealerID) (*VehHistoryOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VehicleHistory_GetOwnerIdByDealerId not implemented")
}
func (*UnimplementedVehicleHistoryServer) VehicleHistory_GetVehActiveInactiveByDealerId(ctx context.Context, req *VehHistoryRequestByDealerID) (*VehHistoryOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VehicleHistory_GetVehActiveInactiveByDealerId not implemented")
}
func (*UnimplementedVehicleHistoryServer) VehicleHistory_ChangeStatus(ctx context.Context, req *VehHistoryRequestByChassis) (*VehHistoryOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VehicleHistory_ChangeStatus not implemented")
}
func (*UnimplementedVehicleHistoryServer) VehicleHistory_ChangeActivationStatus(ctx context.Context, req *VehHistoryRequest) (*VehHistoryOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VehicleHistory_ChangeActivationStatus not implemented")
}
func (*UnimplementedVehicleHistoryServer) VehicleHistory_Add(ctx context.Context, req *VehHistoryRequest) (*VehHistoryOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VehicleHistory_Add not implemented")
}
func (*UnimplementedVehicleHistoryServer) VehicleHistory_Edit(ctx context.Context, req *VehHistoryRequest) (*VehHistoryOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VehicleHistory_Edit not implemented")
}
func (*UnimplementedVehicleHistoryServer) VehicleHistory_ChangeOwnership(ctx context.Context, req *VehHistoryRequest) (*VehHistoryOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VehicleHistory_ChangeOwnership not implemented")
}
func (*UnimplementedVehicleHistoryServer) VehicleHistory_Deactivate(ctx context.Context, req *VehHistoryRequestByChassis) (*VehHistoryOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VehicleHistory_Deactivate not implemented")
}

func RegisterVehicleHistoryServer(s *grpc.Server, srv VehicleHistoryServer) {
	s.RegisterService(&_VehicleHistory_serviceDesc, srv)
}

func _VehicleHistory_VehicleHistory_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VehHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleHistoryServer).VehicleHistory_Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/packets.VehicleHistory/VehicleHistory_Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleHistoryServer).VehicleHistory_Get(ctx, req.(*VehHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleHistory_VehicleHistory_GetCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VehHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleHistoryServer).VehicleHistory_GetCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/packets.VehicleHistory/VehicleHistory_GetCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleHistoryServer).VehicleHistory_GetCount(ctx, req.(*VehHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleHistory_VehicleHistory_GetByChassis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VehHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleHistoryServer).VehicleHistory_GetByChassis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/packets.VehicleHistory/VehicleHistory_GetByChassis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleHistoryServer).VehicleHistory_GetByChassis(ctx, req.(*VehHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleHistory_VehicleHistory_GetByImei_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VehHistoryRequestByImei)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleHistoryServer).VehicleHistory_GetByImei(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/packets.VehicleHistory/VehicleHistory_GetByImei",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleHistoryServer).VehicleHistory_GetByImei(ctx, req.(*VehHistoryRequestByImei))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleHistory_VehicleHistory_GetCountByMonth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VehHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleHistoryServer).VehicleHistory_GetCountByMonth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/packets.VehicleHistory/VehicleHistory_GetCountByMonth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleHistoryServer).VehicleHistory_GetCountByMonth(ctx, req.(*VehHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleHistory_VehicleHistory_GetCountByOwnerID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VehHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleHistoryServer).VehicleHistory_GetCountByOwnerID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/packets.VehicleHistory/VehicleHistory_GetCountByOwnerID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleHistoryServer).VehicleHistory_GetCountByOwnerID(ctx, req.(*VehHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleHistory_VehicleHistory_CheckByChassis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VehHistoryRequestByChassis)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleHistoryServer).VehicleHistory_CheckByChassis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/packets.VehicleHistory/VehicleHistory_CheckByChassis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleHistoryServer).VehicleHistory_CheckByChassis(ctx, req.(*VehHistoryRequestByChassis))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleHistory_VehicleHistory_GetOwnerIdByDealerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VehHistoryRequestByDealerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleHistoryServer).VehicleHistory_GetOwnerIdByDealerId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/packets.VehicleHistory/VehicleHistory_GetOwnerIdByDealerId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleHistoryServer).VehicleHistory_GetOwnerIdByDealerId(ctx, req.(*VehHistoryRequestByDealerID))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleHistory_VehicleHistory_GetVehActiveInactiveByDealerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VehHistoryRequestByDealerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleHistoryServer).VehicleHistory_GetVehActiveInactiveByDealerId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/packets.VehicleHistory/VehicleHistory_GetVehActiveInactiveByDealerId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleHistoryServer).VehicleHistory_GetVehActiveInactiveByDealerId(ctx, req.(*VehHistoryRequestByDealerID))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleHistory_VehicleHistory_ChangeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VehHistoryRequestByChassis)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleHistoryServer).VehicleHistory_ChangeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/packets.VehicleHistory/VehicleHistory_ChangeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleHistoryServer).VehicleHistory_ChangeStatus(ctx, req.(*VehHistoryRequestByChassis))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleHistory_VehicleHistory_ChangeActivationStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VehHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleHistoryServer).VehicleHistory_ChangeActivationStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/packets.VehicleHistory/VehicleHistory_ChangeActivationStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleHistoryServer).VehicleHistory_ChangeActivationStatus(ctx, req.(*VehHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleHistory_VehicleHistory_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VehHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleHistoryServer).VehicleHistory_Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/packets.VehicleHistory/VehicleHistory_Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleHistoryServer).VehicleHistory_Add(ctx, req.(*VehHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleHistory_VehicleHistory_Edit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VehHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleHistoryServer).VehicleHistory_Edit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/packets.VehicleHistory/VehicleHistory_Edit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleHistoryServer).VehicleHistory_Edit(ctx, req.(*VehHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleHistory_VehicleHistory_ChangeOwnership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VehHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleHistoryServer).VehicleHistory_ChangeOwnership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/packets.VehicleHistory/VehicleHistory_ChangeOwnership",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleHistoryServer).VehicleHistory_ChangeOwnership(ctx, req.(*VehHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleHistory_VehicleHistory_Deactivate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VehHistoryRequestByChassis)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleHistoryServer).VehicleHistory_Deactivate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/packets.VehicleHistory/VehicleHistory_Deactivate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleHistoryServer).VehicleHistory_Deactivate(ctx, req.(*VehHistoryRequestByChassis))
	}
	return interceptor(ctx, in, info, handler)
}

var _VehicleHistory_serviceDesc = grpc.ServiceDesc{
	ServiceName: "packets.VehicleHistory",
	HandlerType: (*VehicleHistoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VehicleHistory_Get",
			Handler:    _VehicleHistory_VehicleHistory_Get_Handler,
		},
		{
			MethodName: "VehicleHistory_GetCount",
			Handler:    _VehicleHistory_VehicleHistory_GetCount_Handler,
		},
		{
			MethodName: "VehicleHistory_GetByChassis",
			Handler:    _VehicleHistory_VehicleHistory_GetByChassis_Handler,
		},
		{
			MethodName: "VehicleHistory_GetByImei",
			Handler:    _VehicleHistory_VehicleHistory_GetByImei_Handler,
		},
		{
			MethodName: "VehicleHistory_GetCountByMonth",
			Handler:    _VehicleHistory_VehicleHistory_GetCountByMonth_Handler,
		},
		{
			MethodName: "VehicleHistory_GetCountByOwnerID",
			Handler:    _VehicleHistory_VehicleHistory_GetCountByOwnerID_Handler,
		},
		{
			MethodName: "VehicleHistory_CheckByChassis",
			Handler:    _VehicleHistory_VehicleHistory_CheckByChassis_Handler,
		},
		{
			MethodName: "VehicleHistory_GetOwnerIdByDealerId",
			Handler:    _VehicleHistory_VehicleHistory_GetOwnerIdByDealerId_Handler,
		},
		{
			MethodName: "VehicleHistory_GetVehActiveInactiveByDealerId",
			Handler:    _VehicleHistory_VehicleHistory_GetVehActiveInactiveByDealerId_Handler,
		},
		{
			MethodName: "VehicleHistory_ChangeStatus",
			Handler:    _VehicleHistory_VehicleHistory_ChangeStatus_Handler,
		},
		{
			MethodName: "VehicleHistory_ChangeActivationStatus",
			Handler:    _VehicleHistory_VehicleHistory_ChangeActivationStatus_Handler,
		},
		{
			MethodName: "VehicleHistory_Add",
			Handler:    _VehicleHistory_VehicleHistory_Add_Handler,
		},
		{
			MethodName: "VehicleHistory_Edit",
			Handler:    _VehicleHistory_VehicleHistory_Edit_Handler,
		},
		{
			MethodName: "VehicleHistory_ChangeOwnership",
			Handler:    _VehicleHistory_VehicleHistory_ChangeOwnership_Handler,
		},
		{
			MethodName: "VehicleHistory_Deactivate",
			Handler:    _VehicleHistory_VehicleHistory_Deactivate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vehicle_history.proto",
}
