// Code generated by protoc-gen-go. DO NOT EDIT.
// source: request.proto

package request_repository

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Request struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UserUuid             string               `protobuf:"bytes,3,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	UserIp               string               `protobuf:"bytes,4,opt,name=user_ip,json=userIp,proto3" json:"user_ip,omitempty"`
	Url                  string               `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	Data                 string               `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Request) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Request) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *Request) GetUserIp() string {
	if m != nil {
		return m.UserIp
	}
	return ""
}

func (m *Request) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Request) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type CreateRequestParams struct {
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UserUuid             string               `protobuf:"bytes,2,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	UserIp               string               `protobuf:"bytes,3,opt,name=user_ip,json=userIp,proto3" json:"user_ip,omitempty"`
	Url                  string               `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Data                 string               `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreateRequestParams) Reset()         { *m = CreateRequestParams{} }
func (m *CreateRequestParams) String() string { return proto.CompactTextString(m) }
func (*CreateRequestParams) ProtoMessage()    {}
func (*CreateRequestParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{1}
}

func (m *CreateRequestParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequestParams.Unmarshal(m, b)
}
func (m *CreateRequestParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequestParams.Marshal(b, m, deterministic)
}
func (m *CreateRequestParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequestParams.Merge(m, src)
}
func (m *CreateRequestParams) XXX_Size() int {
	return xxx_messageInfo_CreateRequestParams.Size(m)
}
func (m *CreateRequestParams) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequestParams.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequestParams proto.InternalMessageInfo

func (m *CreateRequestParams) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *CreateRequestParams) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *CreateRequestParams) GetUserIp() string {
	if m != nil {
		return m.UserIp
	}
	return ""
}

func (m *CreateRequestParams) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *CreateRequestParams) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type ListRequestsParams struct {
	Limit                int32    `protobuf:"varint,1,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=Offset,proto3" json:"Offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequestsParams) Reset()         { *m = ListRequestsParams{} }
func (m *ListRequestsParams) String() string { return proto.CompactTextString(m) }
func (*ListRequestsParams) ProtoMessage()    {}
func (*ListRequestsParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{2}
}

func (m *ListRequestsParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequestsParams.Unmarshal(m, b)
}
func (m *ListRequestsParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequestsParams.Marshal(b, m, deterministic)
}
func (m *ListRequestsParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequestsParams.Merge(m, src)
}
func (m *ListRequestsParams) XXX_Size() int {
	return xxx_messageInfo_ListRequestsParams.Size(m)
}
func (m *ListRequestsParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequestsParams.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequestsParams proto.InternalMessageInfo

func (m *ListRequestsParams) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListRequestsParams) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ListRequestsResponse struct {
	Requests             []*Request `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
	Total                int32      `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListRequestsResponse) Reset()         { *m = ListRequestsResponse{} }
func (m *ListRequestsResponse) String() string { return proto.CompactTextString(m) }
func (*ListRequestsResponse) ProtoMessage()    {}
func (*ListRequestsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{3}
}

func (m *ListRequestsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequestsResponse.Unmarshal(m, b)
}
func (m *ListRequestsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequestsResponse.Marshal(b, m, deterministic)
}
func (m *ListRequestsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequestsResponse.Merge(m, src)
}
func (m *ListRequestsResponse) XXX_Size() int {
	return xxx_messageInfo_ListRequestsResponse.Size(m)
}
func (m *ListRequestsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequestsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequestsResponse proto.InternalMessageInfo

func (m *ListRequestsResponse) GetRequests() []*Request {
	if m != nil {
		return m.Requests
	}
	return nil
}

func (m *ListRequestsResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "request.repository.Request")
	proto.RegisterType((*CreateRequestParams)(nil), "request.repository.CreateRequestParams")
	proto.RegisterType((*ListRequestsParams)(nil), "request.repository.ListRequestsParams")
	proto.RegisterType((*ListRequestsResponse)(nil), "request.repository.ListRequestsResponse")
}

func init() { proto.RegisterFile("request.proto", fileDescriptor_7f73548e33e655fe) }

var fileDescriptor_7f73548e33e655fe = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x92, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x95, 0xa6, 0x49, 0xdb, 0x43, 0x20, 0x74, 0x54, 0x10, 0xb5, 0x03, 0xa8, 0x13, 0x53,
	0x2a, 0xc1, 0x80, 0x18, 0x81, 0x09, 0xa9, 0x12, 0xc8, 0x82, 0xb9, 0x72, 0xa9, 0x5b, 0x59, 0x6a,
	0x70, 0xb0, 0xcf, 0x03, 0xef, 0xc4, 0xc0, 0x23, 0xe2, 0xd8, 0x0e, 0xaa, 0x28, 0x2c, 0x6c, 0xf7,
	0xdd, 0x9f, 0xef, 0x7e, 0xb9, 0x18, 0xf6, 0xb5, 0x78, 0xb3, 0xc2, 0x50, 0x59, 0x6b, 0x45, 0x0a,
	0xb1, 0x95, 0x5a, 0xd4, 0xca, 0x48, 0x52, 0xfa, 0x7d, 0x74, 0xba, 0x56, 0x6a, 0xbd, 0x11, 0x53,
	0xdf, 0xb1, 0xb0, 0xab, 0x29, 0xc9, 0xca, 0xb5, 0xf0, 0xaa, 0x0e, 0x43, 0x93, 0xcf, 0x04, 0x7a,
	0x2c, 0xcc, 0xe1, 0x01, 0x74, 0xe4, 0xb2, 0x48, 0xce, 0x92, 0xf3, 0x8c, 0xb9, 0x08, 0xaf, 0x01,
	0x5e, 0xb4, 0xe0, 0x24, 0x96, 0x73, 0x4e, 0x45, 0xc7, 0xe5, 0xf7, 0x2e, 0x46, 0x65, 0x70, 0x2c,
	0x5b, 0xc7, 0xf2, 0xa9, 0x75, 0x64, 0x83, 0xd8, 0x7d, 0x43, 0x38, 0x86, 0x81, 0x35, 0x42, 0xcf,
	0xad, 0x75, 0x8e, 0xa9, 0x9b, 0x1c, 0xb0, 0x7e, 0x93, 0x78, 0x76, 0x1a, 0x4f, 0xa0, 0xe7, 0x8b,
	0xb2, 0x2e, 0xba, 0xbe, 0x94, 0x37, 0xf2, 0xbe, 0xc6, 0x43, 0x48, 0xad, 0xde, 0x14, 0x99, 0x4f,
	0x36, 0x21, 0x22, 0x74, 0x97, 0x9c, 0x78, 0x91, 0xfb, 0x94, 0x8f, 0x27, 0x1f, 0x09, 0x1c, 0xdd,
	0xf9, 0x4d, 0x11, 0xfc, 0x91, 0x6b, 0x5e, 0x99, 0x1f, 0xb8, 0xc9, 0xbf, 0x71, 0x3b, 0x7f, 0xe3,
	0xa6, 0xbf, 0xe1, 0x76, 0x77, 0x71, 0xb3, 0x2d, 0xdc, 0x5b, 0xc0, 0x99, 0x34, 0x14, 0x59, 0x4d,
	0x84, 0x1d, 0x42, 0x36, 0x93, 0x95, 0xa4, 0x78, 0xee, 0x20, 0xf0, 0x18, 0xf2, 0x87, 0xd5, 0xca,
	0x88, 0x70, 0xed, 0x8c, 0x45, 0x35, 0x11, 0x30, 0xdc, 0xf6, 0x60, 0xc2, 0xd4, 0xea, 0xd5, 0x08,
	0xbc, 0x82, 0x7e, 0xfc, 0xe9, 0xc6, 0x19, 0xa5, 0xee, 0x83, 0xc7, 0xe5, 0xee, 0x2b, 0x28, 0xe3,
	0x1c, 0xfb, 0x6e, 0x6e, 0xd6, 0x93, 0x22, 0xbe, 0x89, 0x7b, 0x82, 0x58, 0xe4, 0xfe, 0x4a, 0x97,
	0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x24, 0x68, 0x65, 0x59, 0x02, 0x00, 0x00,
}
