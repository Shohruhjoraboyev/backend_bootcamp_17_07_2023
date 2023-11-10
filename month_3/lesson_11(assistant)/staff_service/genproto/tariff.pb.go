// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: tariff.proto

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

type CreateTariffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TariffType    string  `protobuf:"bytes,2,opt,name=tariff_type,json=tariffType,proto3" json:"tariff_type,omitempty"`
	AmountForCash float64 `protobuf:"fixed64,3,opt,name=amount_for_cash,json=amountForCash,proto3" json:"amount_for_cash,omitempty"`
	AmountForCard float64 `protobuf:"fixed64,4,opt,name=amount_for_card,json=amountForCard,proto3" json:"amount_for_card,omitempty"`
}

func (x *CreateTariffRequest) Reset() {
	*x = CreateTariffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tariff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTariffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTariffRequest) ProtoMessage() {}

func (x *CreateTariffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tariff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTariffRequest.ProtoReflect.Descriptor instead.
func (*CreateTariffRequest) Descriptor() ([]byte, []int) {
	return file_tariff_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTariffRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateTariffRequest) GetTariffType() string {
	if x != nil {
		return x.TariffType
	}
	return ""
}

func (x *CreateTariffRequest) GetAmountForCash() float64 {
	if x != nil {
		return x.AmountForCash
	}
	return 0
}

func (x *CreateTariffRequest) GetAmountForCard() float64 {
	if x != nil {
		return x.AmountForCard
	}
	return 0
}

type Tariff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TariffType    string  `protobuf:"bytes,3,opt,name=tariff_type,json=tariffType,proto3" json:"tariff_type,omitempty"`
	AmountForCash float64 `protobuf:"fixed64,4,opt,name=amount_for_cash,json=amountForCash,proto3" json:"amount_for_cash,omitempty"`
	AmountForCard float64 `protobuf:"fixed64,5,opt,name=amount_for_card,json=amountForCard,proto3" json:"amount_for_card,omitempty"`
	CreatedAt     string  `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string  `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Tariff) Reset() {
	*x = Tariff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tariff_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tariff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tariff) ProtoMessage() {}

func (x *Tariff) ProtoReflect() protoreflect.Message {
	mi := &file_tariff_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tariff.ProtoReflect.Descriptor instead.
func (*Tariff) Descriptor() ([]byte, []int) {
	return file_tariff_proto_rawDescGZIP(), []int{1}
}

func (x *Tariff) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tariff) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tariff) GetTariffType() string {
	if x != nil {
		return x.TariffType
	}
	return ""
}

func (x *Tariff) GetAmountForCash() float64 {
	if x != nil {
		return x.AmountForCash
	}
	return 0
}

func (x *Tariff) GetAmountForCard() float64 {
	if x != nil {
		return x.AmountForCard
	}
	return 0
}

func (x *Tariff) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Tariff) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CreateTariffResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateTariffResponse) Reset() {
	*x = CreateTariffResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tariff_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTariffResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTariffResponse) ProtoMessage() {}

func (x *CreateTariffResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tariff_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTariffResponse.ProtoReflect.Descriptor instead.
func (*CreateTariffResponse) Descriptor() ([]byte, []int) {
	return file_tariff_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTariffResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTariffResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tariff *Tariff `protobuf:"bytes,1,opt,name=Tariff,proto3" json:"Tariff,omitempty"`
}

func (x *GetTariffResponse) Reset() {
	*x = GetTariffResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tariff_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTariffResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTariffResponse) ProtoMessage() {}

func (x *GetTariffResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tariff_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTariffResponse.ProtoReflect.Descriptor instead.
func (*GetTariffResponse) Descriptor() ([]byte, []int) {
	return file_tariff_proto_rawDescGZIP(), []int{3}
}

func (x *GetTariffResponse) GetTariff() *Tariff {
	if x != nil {
		return x.Tariff
	}
	return nil
}

type UpdateTariffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TariffType    string  `protobuf:"bytes,3,opt,name=tariff_type,json=tariffType,proto3" json:"tariff_type,omitempty"`
	AmountForCash float64 `protobuf:"fixed64,4,opt,name=amount_for_cash,json=amountForCash,proto3" json:"amount_for_cash,omitempty"`
	AmountForCard float64 `protobuf:"fixed64,5,opt,name=amount_for_card,json=amountForCard,proto3" json:"amount_for_card,omitempty"`
}

func (x *UpdateTariffRequest) Reset() {
	*x = UpdateTariffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tariff_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTariffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTariffRequest) ProtoMessage() {}

func (x *UpdateTariffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tariff_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTariffRequest.ProtoReflect.Descriptor instead.
func (*UpdateTariffRequest) Descriptor() ([]byte, []int) {
	return file_tariff_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTariffRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTariffRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateTariffRequest) GetTariffType() string {
	if x != nil {
		return x.TariffType
	}
	return ""
}

func (x *UpdateTariffRequest) GetAmountForCash() float64 {
	if x != nil {
		return x.AmountForCash
	}
	return 0
}

func (x *UpdateTariffRequest) GetAmountForCard() float64 {
	if x != nil {
		return x.AmountForCard
	}
	return 0
}

type ListTariffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *ListTariffRequest) Reset() {
	*x = ListTariffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tariff_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTariffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTariffRequest) ProtoMessage() {}

func (x *ListTariffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tariff_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTariffRequest.ProtoReflect.Descriptor instead.
func (*ListTariffRequest) Descriptor() ([]byte, []int) {
	return file_tariff_proto_rawDescGZIP(), []int{5}
}

func (x *ListTariffRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListTariffRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListTariffRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type ListTariffResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tariffs []*Tariff `protobuf:"bytes,1,rep,name=Tariffs,proto3" json:"Tariffs,omitempty"`
	Count   int32     `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ListTariffResponse) Reset() {
	*x = ListTariffResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tariff_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTariffResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTariffResponse) ProtoMessage() {}

func (x *ListTariffResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tariff_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTariffResponse.ProtoReflect.Descriptor instead.
func (*ListTariffResponse) Descriptor() ([]byte, []int) {
	return file_tariff_proto_rawDescGZIP(), []int{6}
}

func (x *ListTariffResponse) GetTariffs() []*Tariff {
	if x != nil {
		return x.Tariffs
	}
	return nil
}

func (x *ListTariffResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_tariff_proto protoreflect.FileDescriptor

var file_tariff_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0b, 0x73,
	0x74, 0x61, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x72,
	0x69, 0x66, 0x66, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x63, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0d, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f, 0x72, 0x43, 0x61, 0x73, 0x68, 0x12,
	0x26, 0x0a, 0x0f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x63, 0x61,
	0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x46, 0x6f, 0x72, 0x43, 0x61, 0x72, 0x64, 0x22, 0xdb, 0x01, 0x0a, 0x06, 0x54, 0x61, 0x72, 0x69,
	0x66, 0x66, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x72,
	0x69, 0x66, 0x66, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x63, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0d, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f, 0x72, 0x43, 0x61, 0x73, 0x68, 0x12,
	0x26, 0x0a, 0x0f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x63, 0x61,
	0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x46, 0x6f, 0x72, 0x43, 0x61, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x26, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x42, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x06, 0x54, 0x61, 0x72, 0x69, 0x66,
	0x66, 0x22, 0xaa, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x72, 0x69,
	0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26,
	0x0a, 0x0f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x63, 0x61, 0x73,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x46,
	0x6f, 0x72, 0x43, 0x61, 0x73, 0x68, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0d, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f, 0x72, 0x43, 0x61, 0x72, 0x64, 0x22, 0x59,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x5b, 0x0a, 0x12, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2f, 0x0a, 0x07, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x07, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x80, 0x03, 0x0a, 0x0d, 0x54, 0x61, 0x72, 0x69, 0x66,
	0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x22, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x72,
	0x69, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x73, 0x74, 0x61,
	0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73,
	0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x47, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x73, 0x74,
	0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tariff_proto_rawDescOnce sync.Once
	file_tariff_proto_rawDescData = file_tariff_proto_rawDesc
)

func file_tariff_proto_rawDescGZIP() []byte {
	file_tariff_proto_rawDescOnce.Do(func() {
		file_tariff_proto_rawDescData = protoimpl.X.CompressGZIP(file_tariff_proto_rawDescData)
	})
	return file_tariff_proto_rawDescData
}

var file_tariff_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_tariff_proto_goTypes = []interface{}{
	(*CreateTariffRequest)(nil),  // 0: staff_service.CreateTariffRequest
	(*Tariff)(nil),               // 1: staff_service.Tariff
	(*CreateTariffResponse)(nil), // 2: staff_service.CreateTariffResponse
	(*GetTariffResponse)(nil),    // 3: staff_service.GetTariffResponse
	(*UpdateTariffRequest)(nil),  // 4: staff_service.UpdateTariffRequest
	(*ListTariffRequest)(nil),    // 5: staff_service.ListTariffRequest
	(*ListTariffResponse)(nil),   // 6: staff_service.ListTariffResponse
	(*IdRequest)(nil),            // 7: staff_service.IdRequest
	(*Response)(nil),             // 8: staff_service.Response
}
var file_tariff_proto_depIdxs = []int32{
	1, // 0: staff_service.GetTariffResponse.Tariff:type_name -> staff_service.Tariff
	1, // 1: staff_service.ListTariffResponse.Tariffs:type_name -> staff_service.Tariff
	0, // 2: staff_service.TariffService.Create:input_type -> staff_service.CreateTariffRequest
	7, // 3: staff_service.TariffService.Get:input_type -> staff_service.IdRequest
	5, // 4: staff_service.TariffService.List:input_type -> staff_service.ListTariffRequest
	4, // 5: staff_service.TariffService.Update:input_type -> staff_service.UpdateTariffRequest
	7, // 6: staff_service.TariffService.Delete:input_type -> staff_service.IdRequest
	2, // 7: staff_service.TariffService.Create:output_type -> staff_service.CreateTariffResponse
	3, // 8: staff_service.TariffService.Get:output_type -> staff_service.GetTariffResponse
	6, // 9: staff_service.TariffService.List:output_type -> staff_service.ListTariffResponse
	8, // 10: staff_service.TariffService.Update:output_type -> staff_service.Response
	8, // 11: staff_service.TariffService.Delete:output_type -> staff_service.Response
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tariff_proto_init() }
func file_tariff_proto_init() {
	if File_tariff_proto != nil {
		return
	}
	file_staff_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tariff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTariffRequest); i {
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
		file_tariff_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tariff); i {
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
		file_tariff_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTariffResponse); i {
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
		file_tariff_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTariffResponse); i {
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
		file_tariff_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTariffRequest); i {
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
		file_tariff_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTariffRequest); i {
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
		file_tariff_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTariffResponse); i {
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
			RawDescriptor: file_tariff_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tariff_proto_goTypes,
		DependencyIndexes: file_tariff_proto_depIdxs,
		MessageInfos:      file_tariff_proto_msgTypes,
	}.Build()
	File_tariff_proto = out.File
	file_tariff_proto_rawDesc = nil
	file_tariff_proto_goTypes = nil
	file_tariff_proto_depIdxs = nil
}
