// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: proto/item.proto

package item

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

// 商品等级
type ProductLevel int32

const (
	ProductLevel_PRODUCT_LEVEL_UNKNOWN ProductLevel = 0 // 未知等级
	ProductLevel_PRODUCT_LEVEL_L1      ProductLevel = 1 // 普通款
	ProductLevel_PRODUCT_LEVEL_L2      ProductLevel = 2 // 稀有款
	ProductLevel_PRODUCT_LEVEL_L3      ProductLevel = 3 // 欧皇款
	ProductLevel_PRODUCT_LEVEL_L4      ProductLevel = 4 // 超神款
)

// Enum value maps for ProductLevel.
var (
	ProductLevel_name = map[int32]string{
		0: "PRODUCT_LEVEL_UNKNOWN",
		1: "PRODUCT_LEVEL_L1",
		2: "PRODUCT_LEVEL_L2",
		3: "PRODUCT_LEVEL_L3",
		4: "PRODUCT_LEVEL_L4",
	}
	ProductLevel_value = map[string]int32{
		"PRODUCT_LEVEL_UNKNOWN": 0,
		"PRODUCT_LEVEL_L1":      1,
		"PRODUCT_LEVEL_L2":      2,
		"PRODUCT_LEVEL_L3":      3,
		"PRODUCT_LEVEL_L4":      4,
	}
)

func (x ProductLevel) Enum() *ProductLevel {
	p := new(ProductLevel)
	*p = x
	return p
}

func (x ProductLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProductLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_item_proto_enumTypes[0].Descriptor()
}

func (ProductLevel) Type() protoreflect.EnumType {
	return &file_proto_item_proto_enumTypes[0]
}

func (x ProductLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProductLevel.Descriptor instead.
func (ProductLevel) EnumDescriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{0}
}

// 商品结构
type ProductBasic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid          uint32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`                                       // 商品id
	Cover        string `protobuf:"bytes,2,opt,name=cover,proto3" json:"cover,omitempty"`                                    // 封面图
	Title        string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`                                    // 标题
	Desc         string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`                                      // 描述
	Price        uint32 `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`                                   // 价格
	ProductLevel uint32 `protobuf:"varint,6,opt,name=product_level,json=productLevel,proto3" json:"product_level,omitempty"` // 商品等级 ProductLevel
}

func (x *ProductBasic) Reset() {
	*x = ProductBasic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductBasic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductBasic) ProtoMessage() {}

func (x *ProductBasic) ProtoReflect() protoreflect.Message {
	mi := &file_proto_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductBasic.ProtoReflect.Descriptor instead.
func (*ProductBasic) Descriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{0}
}

func (x *ProductBasic) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *ProductBasic) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *ProductBasic) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ProductBasic) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *ProductBasic) GetPrice() uint32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ProductBasic) GetProductLevel() uint32 {
	if x != nil {
		return x.ProductLevel
	}
	return 0
}

// 赏结构
type CollectionBasic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid          uint32 `protobuf:"varint,1,opt,name=cid,proto3" json:"cid,omitempty"`                                       // 赏id
	Cover        string `protobuf:"bytes,2,opt,name=cover,proto3" json:"cover,omitempty"`                                    // 封面图
	Title        string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`                                    // 标题
	Desc         string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`                                      // 描述
	Price        uint32 `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`                                   // 价格
	ProductCount uint32 `protobuf:"varint,6,opt,name=product_count,json=productCount,proto3" json:"product_count,omitempty"` // 商品数量
}

func (x *CollectionBasic) Reset() {
	*x = CollectionBasic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionBasic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionBasic) ProtoMessage() {}

func (x *CollectionBasic) ProtoReflect() protoreflect.Message {
	mi := &file_proto_item_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionBasic.ProtoReflect.Descriptor instead.
func (*CollectionBasic) Descriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{1}
}

func (x *CollectionBasic) GetCid() uint32 {
	if x != nil {
		return x.Cid
	}
	return 0
}

func (x *CollectionBasic) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *CollectionBasic) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CollectionBasic) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *CollectionBasic) GetPrice() uint32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CollectionBasic) GetProductCount() uint32 {
	if x != nil {
		return x.ProductCount
	}
	return 0
}

type GetCollectionListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page uint32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *GetCollectionListReq) Reset() {
	*x = GetCollectionListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_item_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionListReq) ProtoMessage() {}

func (x *GetCollectionListReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_item_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectionListReq.ProtoReflect.Descriptor instead.
func (*GetCollectionListReq) Descriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{2}
}

func (x *GetCollectionListReq) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type GetCollectionListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collections []*CollectionBasic `protobuf:"bytes,1,rep,name=collections,proto3" json:"collections,omitempty"`
}

func (x *GetCollectionListRsp) Reset() {
	*x = GetCollectionListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_item_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionListRsp) ProtoMessage() {}

func (x *GetCollectionListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_item_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectionListRsp.ProtoReflect.Descriptor instead.
func (*GetCollectionListRsp) Descriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{3}
}

func (x *GetCollectionListRsp) GetCollections() []*CollectionBasic {
	if x != nil {
		return x.Collections
	}
	return nil
}

type GetCollectionDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid uint32 `protobuf:"varint,1,opt,name=cid,proto3" json:"cid,omitempty"`
}

func (x *GetCollectionDetailReq) Reset() {
	*x = GetCollectionDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_item_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionDetailReq) ProtoMessage() {}

func (x *GetCollectionDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_item_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectionDetailReq.ProtoReflect.Descriptor instead.
func (*GetCollectionDetailReq) Descriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{4}
}

func (x *GetCollectionDetailReq) GetCid() uint32 {
	if x != nil {
		return x.Cid
	}
	return 0
}

type GetCollectionDetailRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollectionBasic    *CollectionBasic   `protobuf:"bytes,1,opt,name=collection_basic,json=collectionBasic,proto3" json:"collection_basic,omitempty"`
	Products           []*ProductBasic    `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`                                                                                                                                          // 商品列表
	LotteryProbability map[uint32]float32 `protobuf:"bytes,3,rep,name=lottery_probability,json=lotteryProbability,proto3" json:"lottery_probability,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"` // 抽奖概率 key:ProductLevel
}

func (x *GetCollectionDetailRsp) Reset() {
	*x = GetCollectionDetailRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_item_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionDetailRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionDetailRsp) ProtoMessage() {}

func (x *GetCollectionDetailRsp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_item_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectionDetailRsp.ProtoReflect.Descriptor instead.
func (*GetCollectionDetailRsp) Descriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{5}
}

func (x *GetCollectionDetailRsp) GetCollectionBasic() *CollectionBasic {
	if x != nil {
		return x.CollectionBasic
	}
	return nil
}

func (x *GetCollectionDetailRsp) GetProducts() []*ProductBasic {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *GetCollectionDetailRsp) GetLotteryProbability() map[uint32]float32 {
	if x != nil {
		return x.LotteryProbability
	}
	return nil
}

var File_proto_item_proto protoreflect.FileDescriptor

var file_proto_item_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x9b, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x42, 0x61, 0x73, 0x69, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x9e, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x22, 0x5e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x46, 0x0a, 0x0b, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x2a, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03,
	0x63, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x63, 0x69, 0x64, 0x22, 0xe5,
	0x02, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x73, 0x70, 0x12, 0x4f, 0x0a, 0x10, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x12, 0x3d, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66,
	0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x69, 0x74,
	0x65, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x74, 0x0a, 0x13, 0x6c, 0x6f, 0x74,
	0x74, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x73, 0x70, 0x2e, 0x4c, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x62, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x6c, 0x6f, 0x74,
	0x74, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x1a,
	0x45, 0x0a, 0x17, 0x4c, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x62, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x81, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x52, 0x4f, 0x44, 0x55,
	0x43, 0x54, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x4c, 0x45,
	0x56, 0x45, 0x4c, 0x5f, 0x4c, 0x31, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x52, 0x4f, 0x44,
	0x55, 0x43, 0x54, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x4c, 0x32, 0x10, 0x02, 0x12, 0x14,
	0x0a, 0x10, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f,
	0x4c, 0x33, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f,
	0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x4c, 0x34, 0x10, 0x04, 0x32, 0xeb, 0x01, 0x0a, 0x0d, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x76, 0x72, 0x12, 0x69, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x29, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x29, 0x2e, 0x66,
	0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x69, 0x74,
	0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x6f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x2b,
	0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x69, 0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x2b, 0x2e, 0x66, 0x75,
	0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x69, 0x74, 0x65,
	0x6d, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x73, 0x70, 0x42, 0x0d, 0x5a, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_item_proto_rawDescOnce sync.Once
	file_proto_item_proto_rawDescData = file_proto_item_proto_rawDesc
)

func file_proto_item_proto_rawDescGZIP() []byte {
	file_proto_item_proto_rawDescOnce.Do(func() {
		file_proto_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_item_proto_rawDescData)
	})
	return file_proto_item_proto_rawDescData
}

var file_proto_item_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_item_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_item_proto_goTypes = []interface{}{
	(ProductLevel)(0),              // 0: funstar.server.item.ProductLevel
	(*ProductBasic)(nil),           // 1: funstar.server.item.ProductBasic
	(*CollectionBasic)(nil),        // 2: funstar.server.item.CollectionBasic
	(*GetCollectionListReq)(nil),   // 3: funstar.server.item.GetCollectionListReq
	(*GetCollectionListRsp)(nil),   // 4: funstar.server.item.GetCollectionListRsp
	(*GetCollectionDetailReq)(nil), // 5: funstar.server.item.GetCollectionDetailReq
	(*GetCollectionDetailRsp)(nil), // 6: funstar.server.item.GetCollectionDetailRsp
	nil,                            // 7: funstar.server.item.GetCollectionDetailRsp.LotteryProbabilityEntry
}
var file_proto_item_proto_depIdxs = []int32{
	2, // 0: funstar.server.item.GetCollectionListRsp.collections:type_name -> funstar.server.item.CollectionBasic
	2, // 1: funstar.server.item.GetCollectionDetailRsp.collection_basic:type_name -> funstar.server.item.CollectionBasic
	1, // 2: funstar.server.item.GetCollectionDetailRsp.products:type_name -> funstar.server.item.ProductBasic
	7, // 3: funstar.server.item.GetCollectionDetailRsp.lottery_probability:type_name -> funstar.server.item.GetCollectionDetailRsp.LotteryProbabilityEntry
	3, // 4: funstar.server.item.CollectionSvr.GetCollectionList:input_type -> funstar.server.item.GetCollectionListReq
	5, // 5: funstar.server.item.CollectionSvr.GetCollectionDetail:input_type -> funstar.server.item.GetCollectionDetailReq
	4, // 6: funstar.server.item.CollectionSvr.GetCollectionList:output_type -> funstar.server.item.GetCollectionListRsp
	6, // 7: funstar.server.item.CollectionSvr.GetCollectionDetail:output_type -> funstar.server.item.GetCollectionDetailRsp
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_item_proto_init() }
func file_proto_item_proto_init() {
	if File_proto_item_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductBasic); i {
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
		file_proto_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectionBasic); i {
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
		file_proto_item_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCollectionListReq); i {
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
		file_proto_item_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCollectionListRsp); i {
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
		file_proto_item_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCollectionDetailReq); i {
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
		file_proto_item_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCollectionDetailRsp); i {
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
			RawDescriptor: file_proto_item_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_item_proto_goTypes,
		DependencyIndexes: file_proto_item_proto_depIdxs,
		EnumInfos:         file_proto_item_proto_enumTypes,
		MessageInfos:      file_proto_item_proto_msgTypes,
	}.Build()
	File_proto_item_proto = out.File
	file_proto_item_proto_rawDesc = nil
	file_proto_item_proto_goTypes = nil
	file_proto_item_proto_depIdxs = nil
}
