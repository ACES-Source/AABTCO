// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/openbank/openbank/v1/branch/all.proto

package branch

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	types "github.com/openbank/openbank/v1/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Branch holds all details about a branch.
type Branch struct {
	// ID is the unique identifier of the branch.
	ID string `protobuf:"bytes,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	// BankID is an identifier for the bank the branch is associated with.
	BankID string `protobuf:"bytes,2,opt,name=BankID,json=bank_id,proto3" json:"BankID,omitempty"`
	// Name is the branch name.
	Name string `protobuf:"bytes,3,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	// PhoneNumber is the branch's phone number.
	PhoneNumber string `protobuf:"bytes,4,opt,name=PhoneNumber,json=phone_number,proto3" json:"PhoneNumber,omitempty"`
	// Address is the branch's address.
	Address *types.Address `protobuf:"bytes,5,opt,name=Address,json=address,proto3" json:"Address,omitempty"`
	// Location is the branch's longitude and latitude.
	Location *types.Location `protobuf:"bytes,6,opt,name=Location,json=location,proto3" json:"Location,omitempty"`
	// Description is the branch's description.
	Description string `protobuf:"bytes,7,opt,name=Description,json=description,proto3" json:"Description,omitempty"`
	// Metadata is the branch's metadata.
	Metadata             string   `protobuf:"bytes,8,opt,name=Metadata,json=metadata,proto3" json:"Metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Branch) Reset()         { *m = Branch{} }
func (m *Branch) String() string { return proto.CompactTextString(m) }
func (*Branch) ProtoMessage()    {}
func (*Branch) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c1d5fc3994e55fa, []int{0}
}

func (m *Branch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Branch.Unmarshal(m, b)
}
func (m *Branch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Branch.Marshal(b, m, deterministic)
}
func (m *Branch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Branch.Merge(m, src)
}
func (m *Branch) XXX_Size() int {
	return xxx_messageInfo_Branch.Size(m)
}
func (m *Branch) XXX_DiscardUnknown() {
	xxx_messageInfo_Branch.DiscardUnknown(m)
}

var xxx_messageInfo_Branch proto.InternalMessageInfo

func (m *Branch) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Branch) GetBankID() string {
	if m != nil {
		return m.BankID
	}
	return ""
}

func (m *Branch) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Branch) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *Branch) GetAddress() *types.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Branch) GetLocation() *types.Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Branch) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Branch) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

// CreateBranchRequest is a request envelope to create a branch.
type CreateBranchRequest struct {
	// BankID is an identifier for the bank the branch is associated with.
	BankID string `protobuf:"bytes,1,opt,name=BankID,json=bank_id,proto3" json:"BankID,omitempty"`
	// Name is the branch name.
	Name string `protobuf:"bytes,2,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	// PhoneNumber is the branch's phone number.
	PhoneNumber string `protobuf:"bytes,3,opt,name=PhoneNumber,json=phone_number,proto3" json:"PhoneNumber,omitempty"`
	// Address is the branch's address.
	Address *types.Address `protobuf:"bytes,4,opt,name=Address,json=address,proto3" json:"Address,omitempty"`
	// Location is the branch's longitude and latitude.
	Location *types.Location `protobuf:"bytes,5,opt,name=Location,json=location,proto3" json:"Location,omitempty"`
	// Description is the branch's description.
	Description string `protobuf:"bytes,6,opt,name=Description,json=description,proto3" json:"Description,omitempty"`
	// Metadata is the branch's metadata.
	Metadata             string   `protobuf:"bytes,7,opt,name=Metadata,json=metadata,proto3" json:"Metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBranchRequest) Reset()         { *m = CreateBranchRequest{} }
func (m *CreateBranchRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBranchRequest) ProtoMessage()    {}
func (*CreateBranchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c1d5fc3994e55fa, []int{1}
}

func (m *CreateBranchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBranchRequest.Unmarshal(m, b)
}
func (m *CreateBranchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBranchRequest.Marshal(b, m, deterministic)
}
func (m *CreateBranchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBranchRequest.Merge(m, src)
}
func (m *CreateBranchRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBranchRequest.Size(m)
}
func (m *CreateBranchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBranchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBranchRequest proto.InternalMessageInfo

func (m *CreateBranchRequest) GetBankID() string {
	if m != nil {
		return m.BankID
	}
	return ""
}

func (m *CreateBranchRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateBranchRequest) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *CreateBranchRequest) GetAddress() *types.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *CreateBranchRequest) GetLocation() *types.Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *CreateBranchRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateBranchRequest) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

// CreateBranchResponse is a response envelope for creating a branch.
type CreateBranchResponse struct {
	// BranchID is the branch unique identifier.
	BranchID             string   `protobuf:"bytes,1,opt,name=BranchID,json=branch_id,proto3" json:"BranchID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBranchResponse) Reset()         { *m = CreateBranchResponse{} }
func (m *CreateBranchResponse) String() string { return proto.CompactTextString(m) }
func (*CreateBranchResponse) ProtoMessage()    {}
func (*CreateBranchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c1d5fc3994e55fa, []int{2}
}

func (m *CreateBranchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBranchResponse.Unmarshal(m, b)
}
func (m *CreateBranchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBranchResponse.Marshal(b, m, deterministic)
}
func (m *CreateBranchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBranchResponse.Merge(m, src)
}
func (m *CreateBranchResponse) XXX_Size() int {
	return xxx_messageInfo_CreateBranchResponse.Size(m)
}
func (m *CreateBranchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBranchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBranchResponse proto.InternalMessageInfo

func (m *CreateBranchResponse) GetBranchID() string {
	if m != nil {
		return m.BranchID
	}
	return ""
}

// UpdateBranchRequest is the request envelope to update a branch.
type UpdateBranchRequest struct {
	// BranchID is the branch unique identifier.
	BranchID string `protobuf:"bytes,1,opt,name=BranchID,json=branch_id,proto3" json:"BranchID,omitempty"`
	// Name is the branch name.
	Name string `protobuf:"bytes,2,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	// PhoneNumber is the branch phone number.
	PhoneNumber string `protobuf:"bytes,3,opt,name=PhoneNumber,json=phone_number,proto3" json:"PhoneNumber,omitempty"`
	// Address is the branch address details.
	Address *types.Address `protobuf:"bytes,4,opt,name=Address,json=address,proto3" json:"Address,omitempty"`
	// Location is the branch longitude and latitude.
	Location *types.Location `protobuf:"bytes,5,opt,name=Location,json=location,proto3" json:"Location,omitempty"`
	// Description is the branch description.
	Description string `protobuf:"bytes,6,opt,name=Description,json=description,proto3" json:"Description,omitempty"`
	// Metadata is the branch metadata.
	Metadata             string   `protobuf:"bytes,7,opt,name=Metadata,json=metadata,proto3" json:"Metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBranchRequest) Reset()         { *m = UpdateBranchRequest{} }
func (m *UpdateBranchRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateBranchRequest) ProtoMessage()    {}
func (*UpdateBranchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c1d5fc3994e55fa, []int{3}
}

func (m *UpdateBranchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBranchRequest.Unmarshal(m, b)
}
func (m *UpdateBranchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBranchRequest.Marshal(b, m, deterministic)
}
func (m *UpdateBranchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBranchRequest.Merge(m, src)
}
func (m *UpdateBranchRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateBranchRequest.Size(m)
}
func (m *UpdateBranchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBranchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBranchRequest proto.InternalMessageInfo

func (m *UpdateBranchRequest) GetBranchID() string {
	if m != nil {
		return m.BranchID
	}
	return ""
}

func (m *UpdateBranchRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateBranchRequest) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *UpdateBranchRequest) GetAddress() *types.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *UpdateBranchRequest) GetLocation() *types.Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *UpdateBranchRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateBranchRequest) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

// DeleteBranchRequest is the request envelope to delete a branch.
type DeleteBranchRequest struct {
	// BranchID is the branch unique identifier.
	BranchID             string   `protobuf:"bytes,1,opt,name=BranchID,json=branch_id,proto3" json:"BranchID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBranchRequest) Reset()         { *m = DeleteBranchRequest{} }
func (m *DeleteBranchRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteBranchRequest) ProtoMessage()    {}
func (*DeleteBranchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c1d5fc3994e55fa, []int{4}
}

func (m *DeleteBranchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBranchRequest.Unmarshal(m, b)
}
func (m *DeleteBranchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBranchRequest.Marshal(b, m, deterministic)
}
func (m *DeleteBranchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBranchRequest.Merge(m, src)
}
func (m *DeleteBranchRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteBranchRequest.Size(m)
}
func (m *DeleteBranchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBranchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBranchRequest proto.InternalMessageInfo

func (m *DeleteBranchRequest) GetBranchID() string {
	if m != nil {
		return m.BranchID
	}
	return ""
}

// GetBranchRequest is the request envelope to get the branch data.
type GetBranchRequest struct {
	// BranchID is the branch unique identifier.
	BranchID             string   `protobuf:"bytes,1,opt,name=BranchID,json=branch_id,proto3" json:"BranchID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBranchRequest) Reset()         { *m = GetBranchRequest{} }
func (m *GetBranchRequest) String() string { return proto.CompactTextString(m) }
func (*GetBranchRequest) ProtoMessage()    {}
func (*GetBranchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c1d5fc3994e55fa, []int{5}
}

func (m *GetBranchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBranchRequest.Unmarshal(m, b)
}
func (m *GetBranchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBranchRequest.Marshal(b, m, deterministic)
}
func (m *GetBranchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBranchRequest.Merge(m, src)
}
func (m *GetBranchRequest) XXX_Size() int {
	return xxx_messageInfo_GetBranchRequest.Size(m)
}
func (m *GetBranchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBranchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBranchRequest proto.InternalMessageInfo

func (m *GetBranchRequest) GetBranchID() string {
	if m != nil {
		return m.BranchID
	}
	return ""
}

// GetBranchesResponse is the response for GetBranches
type GetBranchesResponse struct {
	// Result is the list of the branch.
	Result               []*Branch `protobuf:"bytes,1,rep,name=Result,json=result,proto3" json:"Result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetBranchesResponse) Reset()         { *m = GetBranchesResponse{} }
func (m *GetBranchesResponse) String() string { return proto.CompactTextString(m) }
func (*GetBranchesResponse) ProtoMessage()    {}
func (*GetBranchesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c1d5fc3994e55fa, []int{6}
}

func (m *GetBranchesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBranchesResponse.Unmarshal(m, b)
}
func (m *GetBranchesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBranchesResponse.Marshal(b, m, deterministic)
}
func (m *GetBranchesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBranchesResponse.Merge(m, src)
}
func (m *GetBranchesResponse) XXX_Size() int {
	return xxx_messageInfo_GetBranchesResponse.Size(m)
}
func (m *GetBranchesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBranchesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBranchesResponse proto.InternalMessageInfo

func (m *GetBranchesResponse) GetResult() []*Branch {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*Branch)(nil), "branch.Branch")
	proto.RegisterType((*CreateBranchRequest)(nil), "branch.CreateBranchRequest")
	proto.RegisterType((*CreateBranchResponse)(nil), "branch.CreateBranchResponse")
	proto.RegisterType((*UpdateBranchRequest)(nil), "branch.UpdateBranchRequest")
	proto.RegisterType((*DeleteBranchRequest)(nil), "branch.DeleteBranchRequest")
	proto.RegisterType((*GetBranchRequest)(nil), "branch.GetBranchRequest")
	proto.RegisterType((*GetBranchesResponse)(nil), "branch.GetBranchesResponse")
}

func init() {
	proto.RegisterFile("github.com/openbank/openbank/v1/branch/all.proto", fileDescriptor_0c1d5fc3994e55fa)
}

var fileDescriptor_0c1d5fc3994e55fa = []byte{
	// 1390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x97, 0x5f, 0x6f, 0x1b, 0xc5,
	0x16, 0xc0, 0x77, 0xd7, 0x89, 0xe3, 0x4c, 0x72, 0xdb, 0x74, 0x72, 0x75, 0xaf, 0xbb, 0x8d, 0x72,
	0x47, 0xa9, 0xee, 0x4d, 0x6e, 0x6e, 0xbb, 0x76, 0xdc, 0x5c, 0x10, 0x41, 0x08, 0x9c, 0xa4, 0x0d,
	0x89, 0xda, 0x62, 0xdc, 0x82, 0x50, 0x5f, 0xa2, 0xf1, 0xee, 0xb1, 0xbd, 0x74, 0x3d, 0xb3, 0xcc,
	0xec, 0x26, 0x0d, 0x08, 0xa9, 0xe2, 0x29, 0x8f, 0x55, 0x78, 0x42, 0x20, 0xc4, 0x17, 0xe0, 0x9d,
	0x0f, 0xc0, 0x33, 0x42, 0x42, 0xa0, 0x8a, 0xd7, 0x00, 0xe2, 0x91, 0x47, 0x9e, 0x10, 0xda, 0x99,
	0xf5, 0xda, 0x89, 0x9d, 0x26, 0xb4, 0x3c, 0xf2, 0x94, 0xf8, 0xfc, 0x9b, 0x73, 0x7e, 0x73, 0xce,
	0x19, 0x1b, 0x95, 0x5b, 0x7e, 0xd4, 0x8e, 0x1b, 0x8e, 0xcb, 0x3b, 0x25, 0x1e, 0x02, 0x6b, 0x50,
	0x76, 0xbf, 0xf7, 0xcf, 0xce, 0x52, 0xa9, 0x21, 0x28, 0x73, 0xdb, 0x25, 0x1a, 0x04, 0x4e, 0x28,
	0x78, 0xc4, 0x71, 0x5e, 0x4b, 0xec, 0x99, 0x16, 0xe7, 0xad, 0x00, 0x4a, 0x34, 0xf4, 0x4b, 0x94,
	0x31, 0x1e, 0xd1, 0xc8, 0xe7, 0x4c, 0x6a, 0x2b, 0xfb, 0x8a, 0xfa, 0xe3, 0x5e, 0x6d, 0x01, 0xbb,
	0x2a, 0x77, 0x69, 0xab, 0x05, 0xa2, 0xc4, 0x43, 0x65, 0x31, 0xc4, 0xfa, 0x52, 0x1a, 0x4b, 0x7d,
	0x6a, 0xc4, 0xcd, 0x12, 0x74, 0xc2, 0x68, 0x2f, 0x55, 0x96, 0x4e, 0x4b, 0x31, 0xda, 0x0b, 0x41,
	0xf6, 0x32, 0x9c, 0xfb, 0xd1, 0x42, 0xf9, 0x55, 0x95, 0x24, 0xb6, 0x91, 0xb5, 0xb9, 0x5e, 0x34,
	0x89, 0xb9, 0x30, 0xbe, 0x8a, 0x0a, 0x46, 0xd1, 0x58, 0x30, 0xca, 0x46, 0xcd, 0xa8, 0x5b, 0xbe,
	0x87, 0x2f, 0xa3, 0xfc, 0x2a, 0x65, 0xf7, 0x37, 0xd7, 0x8b, 0xd6, 0x80, 0x7e, 0x2c, 0x89, 0xbd,
	0xed, 0x7b, 0x78, 0x16, 0x8d, 0xdc, 0xa6, 0x1d, 0x28, 0xe6, 0x06, 0x4c, 0x46, 0x18, 0xed, 0x00,
	0xbe, 0x8a, 0x26, 0x6a, 0x6d, 0xce, 0xe0, 0x76, 0xdc, 0x69, 0x80, 0x28, 0x8e, 0x0c, 0x98, 0x4d,
	0x86, 0x89, 0x7a, 0x9b, 0x29, 0x3d, 0x5e, 0x46, 0x63, 0x55, 0xcf, 0x13, 0x20, 0x65, 0x71, 0x94,
	0x98, 0x0b, 0x13, 0x95, 0x73, 0x8e, 0xca, 0xde, 0x49, 0xa5, 0x47, 0x93, 0xa0, 0x5a, 0x88, 0x9f,
	0x47, 0x85, 0x9b, 0xdc, 0x55, 0xc4, 0x8a, 0x79, 0xe5, 0x76, 0x3e, 0x75, 0xeb, 0x8a, 0x8f, 0xf8,
	0x15, 0x82, 0x54, 0x8a, 0xaf, 0xa0, 0x89, 0x75, 0x90, 0xae, 0xf0, 0x15, 0xf9, 0xe2, 0xd8, 0x40,
	0x76, 0x13, 0x5e, 0x4f, 0x8d, 0xff, 0x83, 0x0a, 0xb7, 0x20, 0xa2, 0x1e, 0x8d, 0x68, 0xb1, 0x30,
	0x60, 0x5a, 0xe8, 0xa4, 0xba, 0x95, 0x7c, 0xc1, 0x98, 0x32, 0x8a, 0xc6, 0xdc, 0x77, 0x16, 0x9a,
	0x5e, 0x13, 0x40, 0x23, 0xd0, 0xb4, 0xeb, 0xf0, 0x4e, 0x0c, 0x32, 0xea, 0x03, 0x6b, 0x9e, 0x0e,
	0xd6, 0x3a, 0x1b, 0xd8, 0xdc, 0xd9, 0xc1, 0x8e, 0x3c, 0x1d, 0xd8, 0xd1, 0x67, 0x00, 0x9b, 0x3f,
	0x3b, 0xd8, 0xb1, 0x33, 0x80, 0xdd, 0x40, 0x7f, 0x3f, 0xca, 0x55, 0x86, 0x9c, 0x49, 0xc0, 0xf3,
	0xa8, 0xa0, 0x25, 0x43, 0xd1, 0x8e, 0xeb, 0xc1, 0xdc, 0xf6, 0xbd, 0x2c, 0xd0, 0xf7, 0x16, 0x9a,
	0x7e, 0x23, 0xf4, 0x06, 0x6e, 0xe8, 0xac, 0x81, 0xfe, 0xba, 0xa5, 0x27, 0xdc, 0xd2, 0x0d, 0x34,
	0xbd, 0x0e, 0x01, 0x3c, 0x2d, 0xdb, 0x2c, 0xce, 0x1a, 0x9a, 0xda, 0x80, 0xe8, 0x19, 0x83, 0xbc,
	0x8e, 0xa6, 0xb3, 0x20, 0x20, 0xb3, 0x8e, 0xa9, 0xa0, 0x7c, 0x1d, 0x64, 0x1c, 0x44, 0x45, 0x93,
	0xe4, 0x14, 0x6f, 0xed, 0xea, 0x68, 0xcb, 0x23, 0x51, 0xf3, 0x42, 0x59, 0x76, 0x43, 0x56, 0xbe,
	0x9c, 0x40, 0x7f, 0xd3, 0x66, 0x77, 0x40, 0xec, 0xf8, 0x2e, 0xe0, 0x6f, 0x2d, 0x34, 0x9e, 0x9d,
	0x82, 0x8b, 0xdd, 0x58, 0xc7, 0xb3, 0xb7, 0x8f, 0x9d, 0x32, 0xf7, 0x91, 0x75, 0x50, 0xfd, 0xc5,
	0xcc, 0x96, 0xf2, 0xa5, 0x3a, 0x44, 0xc2, 0x87, 0x1d, 0x20, 0xda, 0x8c, 0xf8, 0xac, 0xc9, 0x45,
	0x47, 0xdd, 0x99, 0xfd, 0x42, 0xa6, 0xec, 0x93, 0x12, 0xda, 0xe0, 0x71, 0x44, 0xa2, 0x76, 0xe6,
	0x22, 0x43, 0x70, 0xfd, 0xa6, 0x0f, 0x1e, 0x69, 0xec, 0x29, 0xf9, 0xe6, 0xfa, 0xd6, 0x06, 0xca,
	0x55, 0xca, 0x65, 0xfc, 0x0a, 0x9a, 0x4d, 0x13, 0x21, 0xf0, 0x00, 0xdc, 0x38, 0x02, 0x8f, 0xc8,
	0xd8, 0x75, 0x41, 0xca, 0x66, 0x1c, 0x04, 0x7b, 0x0e, 0x9e, 0x45, 0x33, 0xb6, 0x7d, 0xb9, 0xe4,
	0x41, 0xd3, 0x67, 0xbe, 0x7e, 0x93, 0x74, 0x58, 0x9d, 0xe0, 0xd6, 0x12, 0xca, 0x2d, 0x97, 0x97,
	0xf1, 0x22, 0x5a, 0xa8, 0x43, 0x14, 0x0b, 0x06, 0x1e, 0xd9, 0x6d, 0x03, 0x53, 0xe7, 0x08, 0x90,
	0x3c, 0x16, 0x2e, 0x10, 0x5f, 0x12, 0xc6, 0x23, 0xd2, 0xe4, 0x31, 0xf3, 0x9c, 0x06, 0x46, 0x53,
	0x28, 0xff, 0x5a, 0x35, 0x8e, 0xda, 0x15, 0x9c, 0x47, 0x23, 0x02, 0xa8, 0xf7, 0xc1, 0x37, 0x3f,
	0x7c, 0x68, 0x5d, 0xc4, 0xff, 0xec, 0xbd, 0xa2, 0x20, 0x4b, 0xef, 0x75, 0xaf, 0xf5, 0xfd, 0x7d,
	0xcb, 0x78, 0x64, 0x29, 0xf6, 0xf8, 0xd0, 0x42, 0x13, 0x7d, 0xb7, 0x87, 0xff, 0xe1, 0xe8, 0xf7,
	0xd0, 0xe9, 0xbe, 0x87, 0xce, 0xf5, 0xe4, 0x3d, 0xb4, 0x2f, 0x0d, 0x10, 0xef, 0x5d, 0xf5, 0xdc,
	0x27, 0xd6, 0x41, 0xf5, 0xb7, 0x1e, 0xe4, 0x7f, 0x65, 0x1c, 0x69, 0x10, 0x10, 0xba, 0x43, 0xfd,
	0x80, 0x36, 0x82, 0x2e, 0x3f, 0x90, 0xf6, 0x73, 0x43, 0x41, 0x0b, 0x68, 0x51, 0xe1, 0xf9, 0xac,
	0x75, 0x82, 0x9b, 0xb3, 0x75, 0x57, 0x53, 0xbe, 0x75, 0x2a, 0xe5, 0xff, 0xa1, 0xff, 0xda, 0xf3,
	0xc3, 0x28, 0x0f, 0xc9, 0xfe, 0xcf, 0x44, 0x7e, 0x0e, 0x4f, 0xf6, 0x23, 0xd7, 0x9c, 0xcb, 0x06,
	0xfe, 0xd4, 0x42, 0x93, 0xfd, 0x6b, 0x15, 0x67, 0x38, 0x87, 0x3c, 0x62, 0xf6, 0xcc, 0x70, 0x65,
	0x0a, 0xfb, 0x2b, 0xf3, 0xa0, 0xfa, 0x79, 0x0f, 0xf6, 0x79, 0x6d, 0x44, 0x68, 0xca, 0xc9, 0x5e,
	0xd0, 0x02, 0x49, 0x28, 0x61, 0xb0, 0xdb, 0x6d, 0x5a, 0xca, 0x3c, 0x22, 0x54, 0x79, 0x92, 0xf8,
	0x91, 0x24, 0xbe, 0xe7, 0x6c, 0xdd, 0x49, 0x70, 0x2e, 0xe1, 0x9b, 0x68, 0x46, 0xc7, 0x22, 0xae,
	0xf2, 0x3c, 0x0e, 0xf3, 0x0a, 0x5a, 0xb4, 0x17, 0x86, 0xc1, 0x1c, 0x96, 0x5e, 0x63, 0x1a, 0x5d,
	0xc8, 0xd0, 0x8c, 0xa1, 0xd1, 0x5d, 0xe1, 0x47, 0xa0, 0xd8, 0x5c, 0x58, 0x31, 0x17, 0xe7, 0x86,
	0xe0, 0x51, 0x6d, 0xf8, 0xb3, 0x89, 0x26, 0xfb, 0x5f, 0x8b, 0x1e, 0xa0, 0x21, 0x6f, 0x88, 0x7d,
	0x42, 0x93, 0xce, 0x7d, 0x66, 0x1e, 0x54, 0xe3, 0x1e, 0x19, 0xed, 0xda, 0x23, 0x33, 0xab, 0x05,
	0x32, 0x93, 0xcc, 0xcb, 0xfe, 0x06, 0xdc, 0xfa, 0x77, 0xc2, 0x63, 0x39, 0x19, 0xd1, 0x94, 0x47,
	0x3f, 0x07, 0x12, 0x2b, 0x67, 0xcf, 0x39, 0xb9, 0xc2, 0x99, 0x15, 0x73, 0xd1, 0x3e, 0x6d, 0xe6,
	0xf0, 0xa1, 0x89, 0x26, 0xfb, 0x97, 0x77, 0xaf, 0xd4, 0x21, 0x2b, 0xfd, 0xc4, 0x52, 0x3f, 0x36,
	0x0f, 0xaa, 0xb2, 0x57, 0xaa, 0x76, 0xed, 0x95, 0x3a, 0x53, 0x03, 0xd1, 0xa1, 0x0c, 0x58, 0x14,
	0xec, 0x11, 0xef, 0xa8, 0xd2, 0x39, 0xa5, 0x50, 0x6d, 0xfe, 0xa4, 0x42, 0x2f, 0x2e, 0x9e, 0x56,
	0xa5, 0x9d, 0xdb, 0xb7, 0x8c, 0xd5, 0xfd, 0xfc, 0x41, 0xf5, 0x70, 0x14, 0xe5, 0x2a, 0x4e, 0x19,
	0xd7, 0x10, 0x4a, 0x0f, 0xab, 0xd6, 0x36, 0xf1, 0x8b, 0x35, 0xc1, 0x77, 0x7c, 0x0f, 0x64, 0xda,
	0x73, 0x69, 0x7f, 0x52, 0x8f, 0xf0, 0x10, 0x84, 0xfe, 0x82, 0x4e, 0x38, 0xeb, 0x5f, 0xbb, 0xdd,
	0x51, 0x74, 0x2a, 0xa3, 0x4b, 0x4e, 0xd9, 0x29, 0x2f, 0x9a, 0x56, 0x65, 0x8a, 0x86, 0x61, 0xe0,
	0xeb, 0x27, 0xb7, 0xf4, 0xb6, 0xe4, 0x6c, 0x65, 0x40, 0x52, 0xdf, 0x4e, 0xc6, 0xbb, 0x8c, 0xdf,
	0x42, 0x6f, 0x0e, 0x1b, 0x6f, 0xbd, 0x46, 0x1a, 0xdc, 0xdb, 0x4b, 0x46, 0xbc, 0x43, 0x03, 0xdd,
	0x0e, 0xc9, 0x14, 0x70, 0x41, 0x3c, 0x0e, 0x7a, 0xee, 0x3b, 0x34, 0x72, 0xdb, 0xca, 0x05, 0x1e,
	0x84, 0xe0, 0x26, 0xea, 0xd4, 0xd7, 0xa9, 0xdf, 0x4c, 0x0e, 0x58, 0xc2, 0xd7, 0xd1, 0xda, 0xc9,
	0x07, 0x64, 0x81, 0x5c, 0xce, 0x22, 0xea, 0x33, 0xa9, 0xb4, 0xb1, 0x04, 0x31, 0xaf, 0x10, 0x78,
	0xc0, 0x22, 0x9f, 0x06, 0xd2, 0xa9, 0xd7, 0x92, 0x68, 0xd7, 0xf0, 0x26, 0xda, 0x18, 0x8c, 0x96,
	0xd8, 0xf7, 0x42, 0xb5, 0xe9, 0x0e, 0x90, 0x10, 0x44, 0xc7, 0x97, 0x32, 0xd9, 0x9f, 0x11, 0x27,
	0x54, 0x5d, 0xe3, 0x91, 0xcd, 0xe5, 0xd4, 0xff, 0xf8, 0x7e, 0xab, 0xdf, 0x40, 0xb9, 0xff, 0x97,
	0xcb, 0xf8, 0x65, 0xf4, 0xd2, 0x51, 0x17, 0xca, 0x48, 0xcc, 0x32, 0x02, 0x20, 0x04, 0x17, 0x84,
	0xbb, 0x6e, 0x2c, 0x12, 0x5c, 0x3a, 0xa2, 0x04, 0xb1, 0x03, 0x82, 0x48, 0xdf, 0x03, 0xe7, 0xde,
	0x4f, 0x26, 0x3a, 0x34, 0xb3, 0x1e, 0x7a, 0x6c, 0x16, 0x72, 0xf8, 0xa1, 0x59, 0x4d, 0x93, 0xe4,
	0xc7, 0x6f, 0x56, 0x26, 0x59, 0x08, 0x90, 0x91, 0xf0, 0x55, 0xfc, 0xa4, 0xa0, 0x38, 0x6a, 0x27,
	0x68, 0x5c, 0xb5, 0x97, 0x92, 0xfa, 0xa5, 0x43, 0xee, 0xb6, 0xa1, 0x5f, 0x91, 0xd4, 0x1e, 0x0a,
	0xae, 0x82, 0x36, 0x79, 0x10, 0xf0, 0x5d, 0x4d, 0x20, 0x39, 0x94, 0x0b, 0xff, 0x5d, 0x6d, 0xb1,
	0xc6, 0x3d, 0x20, 0xcd, 0x80, 0xef, 0x3a, 0x0b, 0x23, 0x95, 0x42, 0xd2, 0xbe, 0x49, 0x88, 0x95,
	0x71, 0xf5, 0x2b, 0x8e, 0xdf, 0x07, 0xb6, 0xba, 0x82, 0x66, 0xd2, 0x26, 0xc7, 0xd3, 0x1b, 0x82,
	0xb2, 0x48, 0x12, 0xf5, 0x29, 0xa5, 0x8a, 0x6c, 0xbd, 0xe8, 0x31, 0x4e, 0x95, 0xaa, 0x5d, 0xb5,
	0xee, 0x55, 0xb3, 0x66, 0xdc, 0x4b, 0x7f, 0x9d, 0x3e, 0x34, 0x8d, 0x7d, 0xd3, 0x78, 0x64, 0x1a,
	0x5f, 0x98, 0xc6, 0x63, 0xd3, 0xf8, 0xd5, 0x34, 0xbe, 0xb6, 0x8c, 0x46, 0x5e, 0xcd, 0xef, 0xb5,
	0xdf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x54, 0x7d, 0x82, 0x1e, 0xf5, 0x0e, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BranchServiceClient is the client API for BranchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BranchServiceClient interface {
	// GetBranch retrieves the details of a specific branch information, selected by its ID.
	GetBranch(ctx context.Context, in *GetBranchRequest, opts ...grpc.CallOption) (*Branch, error)
	// GetBranches get all the available branch.
	GetBranches(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetBranchesResponse, error)
	// CreateBranch creates a new branch and returns its id.
	CreateBranch(ctx context.Context, in *CreateBranchRequest, opts ...grpc.CallOption) (*CreateBranchResponse, error)
	// UpdateBranch updates a branch.
	UpdateBranch(ctx context.Context, in *UpdateBranchRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// DeleteBranch deletes a branch.
	DeleteBranch(ctx context.Context, in *DeleteBranchRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type branchServiceClient struct {
	cc *grpc.ClientConn
}

func NewBranchServiceClient(cc *grpc.ClientConn) BranchServiceClient {
	return &branchServiceClient{cc}
}

func (c *branchServiceClient) GetBranch(ctx context.Context, in *GetBranchRequest, opts ...grpc.CallOption) (*Branch, error) {
	out := new(Branch)
	err := c.cc.Invoke(ctx, "/branch.BranchService/GetBranch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchServiceClient) GetBranches(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetBranchesResponse, error) {
	out := new(GetBranchesResponse)
	err := c.cc.Invoke(ctx, "/branch.BranchService/GetBranches", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchServiceClient) CreateBranch(ctx context.Context, in *CreateBranchRequest, opts ...grpc.CallOption) (*CreateBranchResponse, error) {
	out := new(CreateBranchResponse)
	err := c.cc.Invoke(ctx, "/branch.BranchService/CreateBranch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchServiceClient) UpdateBranch(ctx context.Context, in *UpdateBranchRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/branch.BranchService/UpdateBranch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchServiceClient) DeleteBranch(ctx context.Context, in *DeleteBranchRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/branch.BranchService/DeleteBranch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BranchServiceServer is the server API for BranchService service.
type BranchServiceServer interface {
	// GetBranch retrieves the details of a specific branch information, selected by its ID.
	GetBranch(context.Context, *GetBranchRequest) (*Branch, error)
	// GetBranches get all the available branch.
	GetBranches(context.Context, *empty.Empty) (*GetBranchesResponse, error)
	// CreateBranch creates a new branch and returns its id.
	CreateBranch(context.Context, *CreateBranchRequest) (*CreateBranchResponse, error)
	// UpdateBranch updates a branch.
	UpdateBranch(context.Context, *UpdateBranchRequest) (*empty.Empty, error)
	// DeleteBranch deletes a branch.
	DeleteBranch(context.Context, *DeleteBranchRequest) (*empty.Empty, error)
}

// UnimplementedBranchServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBranchServiceServer struct {
}

func (*UnimplementedBranchServiceServer) GetBranch(ctx context.Context, req *GetBranchRequest) (*Branch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBranch not implemented")
}
func (*UnimplementedBranchServiceServer) GetBranches(ctx context.Context, req *empty.Empty) (*GetBranchesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBranches not implemented")
}
func (*UnimplementedBranchServiceServer) CreateBranch(ctx context.Context, req *CreateBranchRequest) (*CreateBranchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBranch not implemented")
}
func (*UnimplementedBranchServiceServer) UpdateBranch(ctx context.Context, req *UpdateBranchRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBranch not implemented")
}
func (*UnimplementedBranchServiceServer) DeleteBranch(ctx context.Context, req *DeleteBranchRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBranch not implemented")
}

func RegisterBranchServiceServer(s *grpc.Server, srv BranchServiceServer) {
	s.RegisterService(&_BranchService_serviceDesc, srv)
}

func _BranchService_GetBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).GetBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/branch.BranchService/GetBranch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).GetBranch(ctx, req.(*GetBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchService_GetBranches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).GetBranches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/branch.BranchService/GetBranches",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).GetBranches(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchService_CreateBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).CreateBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/branch.BranchService/CreateBranch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).CreateBranch(ctx, req.(*CreateBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchService_UpdateBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).UpdateBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/branch.BranchService/UpdateBranch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).UpdateBranch(ctx, req.(*UpdateBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchService_DeleteBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).DeleteBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/branch.BranchService/DeleteBranch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).DeleteBranch(ctx, req.(*DeleteBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BranchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "branch.BranchService",
	HandlerType: (*BranchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBranch",
			Handler:    _BranchService_GetBranch_Handler,
		},
		{
			MethodName: "GetBranches",
			Handler:    _BranchService_GetBranches_Handler,
		},
		{
			MethodName: "CreateBranch",
			Handler:    _BranchService_CreateBranch_Handler,
		},
		{
			MethodName: "UpdateBranch",
			Handler:    _BranchService_UpdateBranch_Handler,
		},
		{
			MethodName: "DeleteBranch",
			Handler:    _BranchService_DeleteBranch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/openbank/openbank/v1/branch/all.proto",
}
