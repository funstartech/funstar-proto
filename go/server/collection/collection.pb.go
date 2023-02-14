// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: collection.proto

package collection

import (
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

// 赏基础信息结构
type CollectionBasic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid           string  `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`                                            // 赏id
	Cover         string  `protobuf:"bytes,2,opt,name=cover,proto3" json:"cover,omitempty"`                                        // 封面图
	Title         string  `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`                                        // 标题
	Desc          string  `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`                                          // 描述
	Price         float32 `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`                                      // 价格
	DiscountPrice float32 `protobuf:"fixed32,6,opt,name=discount_price,json=discountPrice,proto3" json:"discount_price,omitempty"` // 折扣价
	ProductCount  uint32  `protobuf:"varint,7,opt,name=product_count,json=productCount,proto3" json:"product_count,omitempty"`     // 商品数量
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

func (x *CollectionBasic) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CollectionBasic) GetDiscountPrice() float32 {
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

// 赏详情结构
type CollectionDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid                string                  `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`                                                                                                                                                    // 赏id
	LotteryProbability map[uint32]float32      `protobuf:"bytes,2,rep,name=lottery_probability,json=lotteryProbability,proto3" json:"lottery_probability,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"` // 抽奖概率 key:ProductLevel
	Products           []*product.ProductBasic `protobuf:"bytes,3,rep,name=products,proto3" json:"products,omitempty"`                                                                                                                                          // 商品列表
}

func (x *CollectionDetail) Reset() {
	*x = CollectionDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionDetail) ProtoMessage() {}

func (x *CollectionDetail) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CollectionDetail.ProtoReflect.Descriptor instead.
func (*CollectionDetail) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{1}
}

func (x *CollectionDetail) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (x *CollectionDetail) GetLotteryProbability() map[uint32]float32 {
	if x != nil {
		return x.LotteryProbability
	}
	return nil
}

func (x *CollectionDetail) GetProducts() []*product.ProductBasic {
	if x != nil {
		return x.Products
	}
	return nil
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
		mi := &file_collection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionListReq) ProtoMessage() {}

func (x *GetCollectionListReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetCollectionListReq.ProtoReflect.Descriptor instead.
func (*GetCollectionListReq) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{2}
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
		mi := &file_collection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionListRsp) ProtoMessage() {}

func (x *GetCollectionListRsp) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetCollectionListRsp.ProtoReflect.Descriptor instead.
func (*GetCollectionListRsp) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{3}
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

	Cid string `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
}

func (x *GetCollectionDetailReq) Reset() {
	*x = GetCollectionDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionDetailReq) ProtoMessage() {}

func (x *GetCollectionDetailReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetCollectionDetailReq.ProtoReflect.Descriptor instead.
func (*GetCollectionDetailReq) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{4}
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
		mi := &file_collection_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionDetailRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionDetailRsp) ProtoMessage() {}

func (x *GetCollectionDetailRsp) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetCollectionDetailRsp.ProtoReflect.Descriptor instead.
func (*GetCollectionDetailRsp) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{5}
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
		mi := &file_collection_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetCollectionBasicReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetCollectionBasicReq) ProtoMessage() {}

func (x *BatchGetCollectionBasicReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use BatchGetCollectionBasicReq.ProtoReflect.Descriptor instead.
func (*BatchGetCollectionBasicReq) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{6}
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
		mi := &file_collection_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetCollectionBasicRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetCollectionBasicRsp) ProtoMessage() {}

func (x *BatchGetCollectionBasicRsp) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use BatchGetCollectionBasicRsp.ProtoReflect.Descriptor instead.
func (*BatchGetCollectionBasicRsp) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{7}
}

func (x *BatchGetCollectionBasicRsp) GetCollections() []*CollectionBasic {
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
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x01, 0x0a,
	0x0f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x73, 0x69, 0x63,
	0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65,
	0x73, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0xa3, 0x02, 0x0a, 0x10, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x74, 0x0a, 0x13, 0x6c,
	0x6f, 0x74, 0x74, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74,
	0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x4c, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f,
	0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x6c,
	0x6f, 0x74, 0x74, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x12, 0x40, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x1a, 0x45, 0x0a, 0x17, 0x4c, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x79, 0x50, 0x72,
	0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2a, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x64, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x4c,
	0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52,
	0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x2a, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x22, 0x9f, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x73, 0x70, 0x12, 0x40, 0x0a, 0x05, 0x62, 0x61, 0x73, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x05,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x12, 0x43, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x30, 0x0a, 0x1a, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x64, 0x73, 0x22, 0x6a, 0x0a, 0x1a,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x73, 0x70, 0x12, 0x4c, 0x0a, 0x0b, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x0b, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x8d, 0x03, 0x0a, 0x0d, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x76, 0x72, 0x12, 0x75, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x2f, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x2f, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73,
	0x70, 0x12, 0x87, 0x01, 0x0a, 0x17, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x12, 0x35, 0x2e,
	0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x52, 0x65, 0x71, 0x1a, 0x35, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x73, 0x70, 0x12, 0x7b, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0x31, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x31, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x73, 0x70, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65,
	0x63, 0x68, 0x2f, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_collection_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_collection_proto_goTypes = []interface{}{
	(*CollectionBasic)(nil),            // 0: funstar.server.collection.CollectionBasic
	(*CollectionDetail)(nil),           // 1: funstar.server.collection.CollectionDetail
	(*GetCollectionListReq)(nil),       // 2: funstar.server.collection.GetCollectionListReq
	(*GetCollectionListRsp)(nil),       // 3: funstar.server.collection.GetCollectionListRsp
	(*GetCollectionDetailReq)(nil),     // 4: funstar.server.collection.GetCollectionDetailReq
	(*GetCollectionDetailRsp)(nil),     // 5: funstar.server.collection.GetCollectionDetailRsp
	(*BatchGetCollectionBasicReq)(nil), // 6: funstar.server.collection.BatchGetCollectionBasicReq
	(*BatchGetCollectionBasicRsp)(nil), // 7: funstar.server.collection.BatchGetCollectionBasicRsp
	nil,                                // 8: funstar.server.collection.CollectionDetail.LotteryProbabilityEntry
	(*product.ProductBasic)(nil),       // 9: funstar.server.product.ProductBasic
}
var file_collection_proto_depIdxs = []int32{
	8, // 0: funstar.server.collection.CollectionDetail.lottery_probability:type_name -> funstar.server.collection.CollectionDetail.LotteryProbabilityEntry
	9, // 1: funstar.server.collection.CollectionDetail.products:type_name -> funstar.server.product.ProductBasic
	0, // 2: funstar.server.collection.GetCollectionListRsp.collections:type_name -> funstar.server.collection.CollectionBasic
	0, // 3: funstar.server.collection.GetCollectionDetailRsp.basic:type_name -> funstar.server.collection.CollectionBasic
	1, // 4: funstar.server.collection.GetCollectionDetailRsp.detail:type_name -> funstar.server.collection.CollectionDetail
	0, // 5: funstar.server.collection.BatchGetCollectionBasicRsp.collections:type_name -> funstar.server.collection.CollectionBasic
	2, // 6: funstar.server.collection.CollectionSvr.GetCollectionList:input_type -> funstar.server.collection.GetCollectionListReq
	6, // 7: funstar.server.collection.CollectionSvr.BatchGetCollectionBasic:input_type -> funstar.server.collection.BatchGetCollectionBasicReq
	4, // 8: funstar.server.collection.CollectionSvr.GetCollectionDetail:input_type -> funstar.server.collection.GetCollectionDetailReq
	3, // 9: funstar.server.collection.CollectionSvr.GetCollectionList:output_type -> funstar.server.collection.GetCollectionListRsp
	7, // 10: funstar.server.collection.CollectionSvr.BatchGetCollectionBasic:output_type -> funstar.server.collection.BatchGetCollectionBasicRsp
	5, // 11: funstar.server.collection.CollectionSvr.GetCollectionDetail:output_type -> funstar.server.collection.GetCollectionDetailRsp
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
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
		file_collection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_collection_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_collection_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_collection_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_collection_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_collection_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_collection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_collection_proto_goTypes,
		DependencyIndexes: file_collection_proto_depIdxs,
		MessageInfos:      file_collection_proto_msgTypes,
	}.Build()
	File_collection_proto = out.File
	file_collection_proto_rawDesc = nil
	file_collection_proto_goTypes = nil
	file_collection_proto_depIdxs = nil
}
