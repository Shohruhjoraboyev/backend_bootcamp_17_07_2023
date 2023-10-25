// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: staff.proto

package staff_service

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

type CreateStaffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BranchId  string `protobuf:"bytes,1,opt,name=branchId,proto3" json:"branchId,omitempty"`
	TariffId  string `protobuf:"bytes,2,opt,name=tariffId,proto3" json:"tariffId,omitempty"`
	StaffType string `protobuf:"bytes,3,opt,name=staffType,proto3" json:"staffType,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Login     string `protobuf:"bytes,5,opt,name=login,proto3" json:"login,omitempty"`
	Password  string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	Phone     string `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *CreateStaffRequest) Reset() {
	*x = CreateStaffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStaffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStaffRequest) ProtoMessage() {}

func (x *CreateStaffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStaffRequest.ProtoReflect.Descriptor instead.
func (*CreateStaffRequest) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{0}
}

func (x *CreateStaffRequest) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *CreateStaffRequest) GetTariffId() string {
	if x != nil {
		return x.TariffId
	}
	return ""
}

func (x *CreateStaffRequest) GetStaffType() string {
	if x != nil {
		return x.StaffType
	}
	return ""
}

func (x *CreateStaffRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateStaffRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *CreateStaffRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateStaffRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type Staff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BranchId  string  `protobuf:"bytes,2,opt,name=branchId,proto3" json:"branchId,omitempty"`
	TariffId  string  `protobuf:"bytes,3,opt,name=tariffId,proto3" json:"tariffId,omitempty"`
	StaffType string  `protobuf:"bytes,4,opt,name=staffType,proto3" json:"staffType,omitempty"`
	Name      string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Balance   float64 `protobuf:"fixed64,6,opt,name=balance,proto3" json:"balance,omitempty"`
	Login     string  `protobuf:"bytes,7,opt,name=login,proto3" json:"login,omitempty"`
	Password  string  `protobuf:"bytes,8,opt,name=password,proto3" json:"password,omitempty"`
	Phone     string  `protobuf:"bytes,9,opt,name=phone,proto3" json:"phone,omitempty"`
	CreatedAt string  `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string  `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Staff) Reset() {
	*x = Staff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Staff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Staff) ProtoMessage() {}

func (x *Staff) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Staff.ProtoReflect.Descriptor instead.
func (*Staff) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{1}
}

func (x *Staff) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Staff) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *Staff) GetTariffId() string {
	if x != nil {
		return x.TariffId
	}
	return ""
}

func (x *Staff) GetStaffType() string {
	if x != nil {
		return x.StaffType
	}
	return ""
}

func (x *Staff) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Staff) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *Staff) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *Staff) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Staff) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Staff) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Staff) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type IdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdResponse) Reset() {
	*x = IdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdResponse) ProtoMessage() {}

func (x *IdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdResponse.ProtoReflect.Descriptor instead.
func (*IdResponse) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{2}
}

func (x *IdResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateStaffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BranchId  string  `protobuf:"bytes,2,opt,name=branchId,proto3" json:"branchId,omitempty"`
	TariffId  string  `protobuf:"bytes,3,opt,name=tariffId,proto3" json:"tariffId,omitempty"`
	StaffType string  `protobuf:"bytes,4,opt,name=staffType,proto3" json:"staffType,omitempty"`
	Name      string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Balance   float64 `protobuf:"fixed64,6,opt,name=balance,proto3" json:"balance,omitempty"`
	Login     string  `protobuf:"bytes,7,opt,name=login,proto3" json:"login,omitempty"`
	Password  string  `protobuf:"bytes,8,opt,name=password,proto3" json:"password,omitempty"`
	Phone     string  `protobuf:"bytes,9,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *UpdateStaffRequest) Reset() {
	*x = UpdateStaffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStaffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStaffRequest) ProtoMessage() {}

func (x *UpdateStaffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStaffRequest.ProtoReflect.Descriptor instead.
func (*UpdateStaffRequest) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateStaffRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateStaffRequest) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *UpdateStaffRequest) GetTariffId() string {
	if x != nil {
		return x.TariffId
	}
	return ""
}

func (x *UpdateStaffRequest) GetStaffType() string {
	if x != nil {
		return x.StaffType
	}
	return ""
}

func (x *UpdateStaffRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateStaffRequest) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *UpdateStaffRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *UpdateStaffRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UpdateStaffRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type GetStaffResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Staff *Staff `protobuf:"bytes,1,opt,name=staff,proto3" json:"staff,omitempty"`
}

func (x *GetStaffResponse) Reset() {
	*x = GetStaffResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStaffResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStaffResponse) ProtoMessage() {}

func (x *GetStaffResponse) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStaffResponse.ProtoReflect.Descriptor instead.
func (*GetStaffResponse) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{4}
}

func (x *GetStaffResponse) GetStaff() *Staff {
	if x != nil {
		return x.Staff
	}
	return nil
}

type GetAllStaffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit       int32   `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset      int32   `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Search      string  `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
	BalanceFrom float64 `protobuf:"fixed64,4,opt,name=balanceFrom,proto3" json:"balanceFrom,omitempty"`
	BalanceTo   float64 `protobuf:"fixed64,5,opt,name=balanceTo,proto3" json:"balanceTo,omitempty"`
}

func (x *GetAllStaffRequest) Reset() {
	*x = GetAllStaffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllStaffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllStaffRequest) ProtoMessage() {}

func (x *GetAllStaffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllStaffRequest.ProtoReflect.Descriptor instead.
func (*GetAllStaffRequest) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllStaffRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllStaffRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAllStaffRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *GetAllStaffRequest) GetBalanceFrom() float64 {
	if x != nil {
		return x.BalanceFrom
	}
	return 0
}

func (x *GetAllStaffRequest) GetBalanceTo() float64 {
	if x != nil {
		return x.BalanceTo
	}
	return 0
}

type GetAllStaffResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Staffs []*Staff `protobuf:"bytes,1,rep,name=staffs,proto3" json:"staffs,omitempty"`
	Count  int32    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllStaffResponse) Reset() {
	*x = GetAllStaffResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllStaffResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllStaffResponse) ProtoMessage() {}

func (x *GetAllStaffResponse) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllStaffResponse.ProtoReflect.Descriptor instead.
func (*GetAllStaffResponse) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllStaffResponse) GetStaffs() []*Staff {
	if x != nil {
		return x.Staffs
	}
	return nil
}

func (x *GetAllStaffResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{7}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type IdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdRequest) Reset() {
	*x = IdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdRequest) ProtoMessage() {}

func (x *IdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdRequest.ProtoReflect.Descriptor instead.
func (*IdRequest) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{8}
}

func (x *IdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_staff_proto protoreflect.FileDescriptor

var file_staff_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73,
	0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xc6, 0x01, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x74, 0x61, 0x66, 0x66, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x66, 0x66, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0xa1, 0x02, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x66, 0x66, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74,
	0x61, 0x72, 0x69, 0x66, 0x66, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x61, 0x72, 0x69, 0x66, 0x66, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x66, 0x66,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x1c, 0x0a, 0x0a, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xf0, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x66, 0x66, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x66, 0x66,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x3e, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x66, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74,
	0x61, 0x66, 0x66, 0x52, 0x05, 0x73, 0x74, 0x61, 0x66, 0x66, 0x22, 0x9a, 0x01, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x54, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x6f, 0x22, 0x59, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x66, 0x66, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53,
	0x74, 0x61, 0x66, 0x66, 0x52, 0x06, 0x73, 0x74, 0x61, 0x66, 0x66, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x24, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1b, 0x0a, 0x09, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xf6, 0x02, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x66, 0x66, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x21, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x42, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x21,
	0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x21, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3d, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x18,
	0x5a, 0x16, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x66, 0x66,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_staff_proto_rawDescOnce sync.Once
	file_staff_proto_rawDescData = file_staff_proto_rawDesc
)

func file_staff_proto_rawDescGZIP() []byte {
	file_staff_proto_rawDescOnce.Do(func() {
		file_staff_proto_rawDescData = protoimpl.X.CompressGZIP(file_staff_proto_rawDescData)
	})
	return file_staff_proto_rawDescData
}

var file_staff_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_staff_proto_goTypes = []interface{}{
	(*CreateStaffRequest)(nil),  // 0: staff_service.CreateStaffRequest
	(*Staff)(nil),               // 1: staff_service.Staff
	(*IdResponse)(nil),          // 2: staff_service.IdResponse
	(*UpdateStaffRequest)(nil),  // 3: staff_service.UpdateStaffRequest
	(*GetStaffResponse)(nil),    // 4: staff_service.GetStaffResponse
	(*GetAllStaffRequest)(nil),  // 5: staff_service.GetAllStaffRequest
	(*GetAllStaffResponse)(nil), // 6: staff_service.GetAllStaffResponse
	(*Response)(nil),            // 7: staff_service.Response
	(*IdRequest)(nil),           // 8: staff_service.IdRequest
}
var file_staff_proto_depIdxs = []int32{
	1, // 0: staff_service.GetStaffResponse.staff:type_name -> staff_service.Staff
	1, // 1: staff_service.GetAllStaffResponse.staffs:type_name -> staff_service.Staff
	0, // 2: staff_service.StaffService.Create:input_type -> staff_service.CreateStaffRequest
	8, // 3: staff_service.StaffService.Get:input_type -> staff_service.IdRequest
	5, // 4: staff_service.StaffService.GetAll:input_type -> staff_service.GetAllStaffRequest
	3, // 5: staff_service.StaffService.Update:input_type -> staff_service.UpdateStaffRequest
	8, // 6: staff_service.StaffService.Delete:input_type -> staff_service.IdRequest
	2, // 7: staff_service.StaffService.Create:output_type -> staff_service.IdResponse
	4, // 8: staff_service.StaffService.Get:output_type -> staff_service.GetStaffResponse
	6, // 9: staff_service.StaffService.GetAll:output_type -> staff_service.GetAllStaffResponse
	7, // 10: staff_service.StaffService.Update:output_type -> staff_service.Response
	7, // 11: staff_service.StaffService.Delete:output_type -> staff_service.Response
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_staff_proto_init() }
func file_staff_proto_init() {
	if File_staff_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_staff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStaffRequest); i {
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
		file_staff_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Staff); i {
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
		file_staff_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdResponse); i {
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
		file_staff_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStaffRequest); i {
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
		file_staff_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStaffResponse); i {
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
		file_staff_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllStaffRequest); i {
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
		file_staff_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllStaffResponse); i {
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
		file_staff_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_staff_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdRequest); i {
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
			RawDescriptor: file_staff_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_staff_proto_goTypes,
		DependencyIndexes: file_staff_proto_depIdxs,
		MessageInfos:      file_staff_proto_msgTypes,
	}.Build()
	File_staff_proto = out.File
	file_staff_proto_rawDesc = nil
	file_staff_proto_goTypes = nil
	file_staff_proto_depIdxs = nil
}
