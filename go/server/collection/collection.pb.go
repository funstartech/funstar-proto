// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: collection.proto

package collection

import (
	order "github.com/funstartech/funstar-proto/go/server/order"
	product "github.com/funstartech/funstar-proto/go/server/product"
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
type CollectionProductLevel int32

const (
	CollectionProductLevel_PRODUCT_LEVEL_UNKNOWN CollectionProductLevel = 0 // 未知
	CollectionProductLevel_PRODUCT_LEVEL_L1      CollectionProductLevel = 1 // 普通款
	CollectionProductLevel_PRODUCT_LEVEL_L2      CollectionProductLevel = 2 // 稀有款
	CollectionProductLevel_PRODUCT_LEVEL_L3      CollectionProductLevel = 3 // 欧皇款
	CollectionProductLevel_PRODUCT_LEVEL_L4      CollectionProductLevel = 4 // 超神款
)

// Enum value maps for CollectionProductLevel.
var (
	CollectionProductLevel_name = map[int32]string{
		0: "PRODUCT_LEVEL_UNKNOWN",
		1: "PRODUCT_LEVEL_L1",
		2: "PRODUCT_LEVEL_L2",
		3: "PRODUCT_LEVEL_L3",
		4: "PRODUCT_LEVEL_L4",
	}
	CollectionProductLevel_value = map[string]int32{
		"PRODUCT_LEVEL_UNKNOWN": 0,
		"PRODUCT_LEVEL_L1":      1,
		"PRODUCT_LEVEL_L2":      2,
		"PRODUCT_LEVEL_L3":      3,
		"PRODUCT_LEVEL_L4":      4,
	}
)

func (x CollectionProductLevel) Enum() *CollectionProductLevel {
	p := new(CollectionProductLevel)
	*p = x
	return p
}

func (x CollectionProductLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CollectionProductLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_collection_proto_enumTypes[0].Descriptor()
}

func (CollectionProductLevel) Type() protoreflect.EnumType {
	return &file_collection_proto_enumTypes[0]
}

func (x CollectionProductLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CollectionProductLevel.Descriptor instead.
func (CollectionProductLevel) EnumDescriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{0}
}

// 赏基础信息结构
type CollectionBasic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid           string `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`                                           // 赏id
	Cover         string `protobuf:"bytes,2,opt,name=cover,proto3" json:"cover,omitempty"`                                       // 封面图
	Title         string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`                                       // 标题
	Desc          string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`                                         // 描述
	Price         int32  `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`                                      // 价格(单位：分)
	DiscountPrice int32  `protobuf:"varint,6,opt,name=discount_price,json=discountPrice,proto3" json:"discount_price,omitempty"` // 折扣价(单位：分)
	ProductCount  uint32 `protobuf:"varint,7,opt,name=product_count,json=productCount,proto3" json:"product_count,omitempty"`    // 商品数量
}

func (x *CollectionBasic) Reset() {
	*x = CollectionBasic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionBasic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionBasic) ProtoMessage() {}

func (x *CollectionBasic) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[0]
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
	return file_collection_proto_rawDescGZIP(), []int{0}
}

func (x *CollectionBasic) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
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

func (x *CollectionBasic) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CollectionBasic) GetDiscountPrice() int32 {
	if x != nil {
		return x.DiscountPrice
	}
	return 0
}

func (x *CollectionBasic) GetProductCount() uint32 {
	if x != nil {
		return x.ProductCount
	}
	return 0
}

type CollectionDetailProducts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level       uint32                  `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"` // CollectionProductLevel
	Probability float32                 `protobuf:"fixed32,2,opt,name=probability,proto3" json:"probability,omitempty"`
	Products    []*product.ProductBasic `protobuf:"bytes,3,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *CollectionDetailProducts) Reset() {
	*x = CollectionDetailProducts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionDetailProducts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionDetailProducts) ProtoMessage() {}

func (x *CollectionDetailProducts) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionDetailProducts.ProtoReflect.Descriptor instead.
func (*CollectionDetailProducts) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{1}
}

func (x *CollectionDetailProducts) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *CollectionDetailProducts) GetProbability() float32 {
	if x != nil {
		return x.Probability
	}
	return 0
}

func (x *CollectionDetailProducts) GetProducts() []*product.ProductBasic {
	if x != nil {
		return x.Products
	}
	return nil
}

// 赏详情结构
type CollectionDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid            string                      `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`                                             // 赏id
	DetailProducts []*CollectionDetailProducts `protobuf:"bytes,2,rep,name=detail_products,json=detailProducts,proto3" json:"detail_products,omitempty"` // 商品列表
}

func (x *CollectionDetail) Reset() {
	*x = CollectionDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionDetail) ProtoMessage() {}

func (x *CollectionDetail) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionDetail.ProtoReflect.Descriptor instead.
func (*CollectionDetail) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{2}
}

func (x *CollectionDetail) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (x *CollectionDetail) GetDetailProducts() []*CollectionDetailProducts {
	if x != nil {
		return x.DetailProducts
	}
	return nil
}

type GetCollectionListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cookie []byte `protobuf:"bytes,1,opt,name=cookie,proto3" json:"cookie,omitempty"`
}

func (x *GetCollectionListReq) Reset() {
	*x = GetCollectionListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionListReq) ProtoMessage() {}

func (x *GetCollectionListReq) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[3]
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
	return file_collection_proto_rawDescGZIP(), []int{3}
}

func (x *GetCollectionListReq) GetCookie() []byte {
	if x != nil {
		return x.Cookie
	}
	return nil
}

type GetCollectionListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collections []*CollectionBasic `protobuf:"bytes,1,rep,name=collections,proto3" json:"collections,omitempty"`
	Cookie      []byte             `protobuf:"bytes,2,opt,name=cookie,proto3" json:"cookie,omitempty"`
	IsEnd       bool               `protobuf:"varint,3,opt,name=is_end,json=isEnd,proto3" json:"is_end,omitempty"`
}

func (x *GetCollectionListRsp) Reset() {
	*x = GetCollectionListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionListRsp) ProtoMessage() {}

func (x *GetCollectionListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[4]
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
	return file_collection_proto_rawDescGZIP(), []int{4}
}

func (x *GetCollectionListRsp) GetCollections() []*CollectionBasic {
	if x != nil {
		return x.Collections
	}
	return nil
}

func (x *GetCollectionListRsp) GetCookie() []byte {
	if x != nil {
		return x.Cookie
	}
	return nil
}

func (x *GetCollectionListRsp) GetIsEnd() bool {
	if x != nil {
		return x.IsEnd
	}
	return false
}

type GetCollectionDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid string `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
}

func (x *GetCollectionDetailReq) Reset() {
	*x = GetCollectionDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionDetailReq) ProtoMessage() {}

func (x *GetCollectionDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[5]
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
	return file_collection_proto_rawDescGZIP(), []int{5}
}

func (x *GetCollectionDetailReq) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

type GetCollectionDetailRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Basic  *CollectionBasic  `protobuf:"bytes,1,opt,name=basic,proto3" json:"basic,omitempty"`   // 基础信息
	Detail *CollectionDetail `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"` // 详情
}

func (x *GetCollectionDetailRsp) Reset() {
	*x = GetCollectionDetailRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionDetailRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionDetailRsp) ProtoMessage() {}

func (x *GetCollectionDetailRsp) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[6]
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
	return file_collection_proto_rawDescGZIP(), []int{6}
}

func (x *GetCollectionDetailRsp) GetBasic() *CollectionBasic {
	if x != nil {
		return x.Basic
	}
	return nil
}

func (x *GetCollectionDetailRsp) GetDetail() *CollectionDetail {
	if x != nil {
		return x.Detail
	}
	return nil
}

type BatchGetCollectionBasicReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cids []string `protobuf:"bytes,1,rep,name=cids,proto3" json:"cids,omitempty"`
}

func (x *BatchGetCollectionBasicReq) Reset() {
	*x = BatchGetCollectionBasicReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetCollectionBasicReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetCollectionBasicReq) ProtoMessage() {}

func (x *BatchGetCollectionBasicReq) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetCollectionBasicReq.ProtoReflect.Descriptor instead.
func (*BatchGetCollectionBasicReq) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{7}
}

func (x *BatchGetCollectionBasicReq) GetCids() []string {
	if x != nil {
		return x.Cids
	}
	return nil
}

type BatchGetCollectionBasicRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collections []*CollectionBasic `protobuf:"bytes,1,rep,name=collections,proto3" json:"collections,omitempty"`
}

func (x *BatchGetCollectionBasicRsp) Reset() {
	*x = BatchGetCollectionBasicRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetCollectionBasicRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetCollectionBasicRsp) ProtoMessage() {}

func (x *BatchGetCollectionBasicRsp) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetCollectionBasicRsp.ProtoReflect.Descriptor instead.
func (*BatchGetCollectionBasicRsp) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{8}
}

func (x *BatchGetCollectionBasicRsp) GetCollections() []*CollectionBasic {
	if x != nil {
		return x.Collections
	}
	return nil
}

type LotteryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collections []*order.Collection `protobuf:"bytes,1,rep,name=collections,proto3" json:"collections,omitempty"`
}

func (x *LotteryReq) Reset() {
	*x = LotteryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LotteryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LotteryReq) ProtoMessage() {}

func (x *LotteryReq) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LotteryReq.ProtoReflect.Descriptor instead.
func (*LotteryReq) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{9}
}

func (x *LotteryReq) GetCollections() []*order.Collection {
	if x != nil {
		return x.Collections
	}
	return nil
}

type LotteryRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collections []*order.Collection `protobuf:"bytes,1,rep,name=collections,proto3" json:"collections,omitempty"`
}

func (x *LotteryRsp) Reset() {
	*x = LotteryRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LotteryRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LotteryRsp) ProtoMessage() {}

func (x *LotteryRsp) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LotteryRsp.ProtoReflect.Descriptor instead.
func (*LotteryRsp) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{10}
}

func (x *LotteryRsp) GetCollections() []*order.Collection {
	if x != nil {
		return x.Collections
	}
	return nil
}

var File_collection_proto protoreflect.FileDescriptor

var file_collection_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x19, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0d, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x0f, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x12, 0x10, 0x0a,
	0x03, 0x63, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x94, 0x01, 0x0a, 0x18, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x40, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74,
	0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a,
	0x03, 0x63, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12,
	0x5c, 0x0a, 0x0f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74,
	0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x0e, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x2e, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x22, 0x93, 0x01,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x4c, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x66, 0x75,
	0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x15, 0x0a, 0x06,
	0x69, 0x73, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73,
	0x45, 0x6e, 0x64, 0x22, 0x2a, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a,
	0x03, 0x63, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x22,
	0x9f, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x73, 0x70, 0x12, 0x40, 0x0a, 0x05, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x66, 0x75, 0x6e, 0x73,
	0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x05, 0x62, 0x61, 0x73, 0x69, 0x63, 0x12, 0x43, 0x0a, 0x06,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x66,
	0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x22, 0x30, 0x0a, 0x1a, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x65, 0x71, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x69, 0x64, 0x73, 0x22, 0x6a, 0x0a, 0x1a, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x73,
	0x70, 0x12, 0x4c, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x50, 0x0a, 0x0a, 0x4c, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12, 0x42, 0x0a,
	0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x50, 0x0a, 0x0a, 0x4c, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x79, 0x52, 0x73, 0x70, 0x12,
	0x42, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2a, 0x8b, 0x01, 0x0a, 0x16, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x19,
	0x0a, 0x15, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x52, 0x4f,
	0x44, 0x55, 0x43, 0x54, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x4c, 0x31, 0x10, 0x01, 0x12,
	0x14, 0x0a, 0x10, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c,
	0x5f, 0x4c, 0x32, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54,
	0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x4c, 0x33, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x50,
	0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x4c, 0x34, 0x10,
	0x04, 0x32, 0xe6, 0x03, 0x0a, 0x0d, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x76, 0x72, 0x12, 0x75, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74,
	0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x2f, 0x2e, 0x66, 0x75, 0x6e, 0x73,
	0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x87, 0x01, 0x0a, 0x17, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x12, 0x35, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x65, 0x71, 0x1a, 0x35, 0x2e,
	0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x52, 0x73, 0x70, 0x12, 0x7b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x31, 0x2e, 0x66, 0x75,
	0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x31,
	0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x73,
	0x70, 0x12, 0x57, 0x0a, 0x07, 0x4c, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x79, 0x12, 0x25, 0x2e, 0x66,
	0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x4c, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x79, 0x52, 0x73, 0x70, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x65, 0x63, 0x68, 0x2f, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_collection_proto_rawDescOnce sync.Once
	file_collection_proto_rawDescData = file_collection_proto_rawDesc
)

func file_collection_proto_rawDescGZIP() []byte {
	file_collection_proto_rawDescOnce.Do(func() {
		file_collection_proto_rawDescData = protoimpl.X.CompressGZIP(file_collection_proto_rawDescData)
	})
	return file_collection_proto_rawDescData
}

var file_collection_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_collection_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_collection_proto_goTypes = []interface{}{
	(CollectionProductLevel)(0),        // 0: funstar.server.collection.CollectionProductLevel
	(*CollectionBasic)(nil),            // 1: funstar.server.collection.CollectionBasic
	(*CollectionDetailProducts)(nil),   // 2: funstar.server.collection.CollectionDetailProducts
	(*CollectionDetail)(nil),           // 3: funstar.server.collection.CollectionDetail
	(*GetCollectionListReq)(nil),       // 4: funstar.server.collection.GetCollectionListReq
	(*GetCollectionListRsp)(nil),       // 5: funstar.server.collection.GetCollectionListRsp
	(*GetCollectionDetailReq)(nil),     // 6: funstar.server.collection.GetCollectionDetailReq
	(*GetCollectionDetailRsp)(nil),     // 7: funstar.server.collection.GetCollectionDetailRsp
	(*BatchGetCollectionBasicReq)(nil), // 8: funstar.server.collection.BatchGetCollectionBasicReq
	(*BatchGetCollectionBasicRsp)(nil), // 9: funstar.server.collection.BatchGetCollectionBasicRsp
	(*LotteryReq)(nil),                 // 10: funstar.server.collection.LotteryReq
	(*LotteryRsp)(nil),                 // 11: funstar.server.collection.LotteryRsp
	(*product.ProductBasic)(nil),       // 12: funstar.server.product.ProductBasic
	(*order.Collection)(nil),           // 13: funstar.server.order.Collection
}
var file_collection_proto_depIdxs = []int32{
	12, // 0: funstar.server.collection.CollectionDetailProducts.products:type_name -> funstar.server.product.ProductBasic
	2,  // 1: funstar.server.collection.CollectionDetail.detail_products:type_name -> funstar.server.collection.CollectionDetailProducts
	1,  // 2: funstar.server.collection.GetCollectionListRsp.collections:type_name -> funstar.server.collection.CollectionBasic
	1,  // 3: funstar.server.collection.GetCollectionDetailRsp.basic:type_name -> funstar.server.collection.CollectionBasic
	3,  // 4: funstar.server.collection.GetCollectionDetailRsp.detail:type_name -> funstar.server.collection.CollectionDetail
	1,  // 5: funstar.server.collection.BatchGetCollectionBasicRsp.collections:type_name -> funstar.server.collection.CollectionBasic
	13, // 6: funstar.server.collection.LotteryReq.collections:type_name -> funstar.server.order.Collection
	13, // 7: funstar.server.collection.LotteryRsp.collections:type_name -> funstar.server.order.Collection
	4,  // 8: funstar.server.collection.CollectionSvr.GetCollectionList:input_type -> funstar.server.collection.GetCollectionListReq
	8,  // 9: funstar.server.collection.CollectionSvr.BatchGetCollectionBasic:input_type -> funstar.server.collection.BatchGetCollectionBasicReq
	6,  // 10: funstar.server.collection.CollectionSvr.GetCollectionDetail:input_type -> funstar.server.collection.GetCollectionDetailReq
	10, // 11: funstar.server.collection.CollectionSvr.Lottery:input_type -> funstar.server.collection.LotteryReq
	5,  // 12: funstar.server.collection.CollectionSvr.GetCollectionList:output_type -> funstar.server.collection.GetCollectionListRsp
	9,  // 13: funstar.server.collection.CollectionSvr.BatchGetCollectionBasic:output_type -> funstar.server.collection.BatchGetCollectionBasicRsp
	7,  // 14: funstar.server.collection.CollectionSvr.GetCollectionDetail:output_type -> funstar.server.collection.GetCollectionDetailRsp
	11, // 15: funstar.server.collection.CollectionSvr.Lottery:output_type -> funstar.server.collection.LotteryRsp
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_collection_proto_init() }
func file_collection_proto_init() {
	if File_collection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_collection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_collection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectionDetailProducts); i {
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
		file_collection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectionDetail); i {
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
		file_collection_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_collection_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_collection_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_collection_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_collection_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetCollectionBasicReq); i {
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
		file_collection_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetCollectionBasicRsp); i {
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
		file_collection_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LotteryReq); i {
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
		file_collection_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LotteryRsp); i {
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
			RawDescriptor: file_collection_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_collection_proto_goTypes,
		DependencyIndexes: file_collection_proto_depIdxs,
		EnumInfos:         file_collection_proto_enumTypes,
		MessageInfos:      file_collection_proto_msgTypes,
	}.Build()
	File_collection_proto = out.File
	file_collection_proto_rawDesc = nil
	file_collection_proto_goTypes = nil
	file_collection_proto_depIdxs = nil
}
