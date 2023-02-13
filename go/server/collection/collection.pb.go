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

	Cid          string `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`                                        // 赏id
	Cover        string `protobuf:"bytes,2,opt,name=cover,proto3" json:"cover,omitempty"`                                    // 封面图
	Title        string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`                                    // 标题
	Desc         string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`                                      // 描述
	Price        uint32 `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`                                   // 价格
	ProductCount uint32 `protobuf:"varint,6,opt,name=product_count,json=productCount,proto3" json:"product_count,omitempty"` // 商品数量
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
		mi := &file_collection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionListReq) ProtoMessage() {}

func (x *GetCollectionListReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetCollectionListReq.ProtoReflect.Descriptor instead.
func (*GetCollectionListReq) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{1}
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
		mi := &file_collection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionListRsp) ProtoMessage() {}

func (x *GetCollectionListRsp) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetCollectionListRsp.ProtoReflect.Descriptor instead.
func (*GetCollectionListRsp) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{2}
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
		mi := &file_collection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionDetailReq) ProtoMessage() {}

func (x *GetCollectionDetailReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetCollectionDetailReq.ProtoReflect.Descriptor instead.
func (*GetCollectionDetailReq) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{3}
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

	CollectionBasic    *CollectionBasic        `protobuf:"bytes,1,opt,name=collection_basic,json=collectionBasic,proto3" json:"collection_basic,omitempty"`
	Products           []*product.ProductBasic `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`                                                                                                                                          // 商品列表
	LotteryProbability map[uint32]float32      `protobuf:"bytes,3,rep,name=lottery_probability,json=lotteryProbability,proto3" json:"lottery_probability,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"` // 抽奖概率 key:ProductLevel
}

func (x *GetCollectionDetailRsp) Reset() {
	*x = GetCollectionDetailRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionDetailRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionDetailRsp) ProtoMessage() {}

func (x *GetCollectionDetailRsp) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetCollectionDetailRsp.ProtoReflect.Descriptor instead.
func (*GetCollectionDetailRsp) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{4}
}

func (x *GetCollectionDetailRsp) GetCollectionBasic() *CollectionBasic {
	if x != nil {
		return x.CollectionBasic
	}
	return nil
}

func (x *GetCollectionDetailRsp) GetProducts() []*product.ProductBasic {
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

type BatchGetCollectionBasicReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid []string `protobuf:"bytes,1,rep,name=cid,proto3" json:"cid,omitempty"`
}

func (x *BatchGetCollectionBasicReq) Reset() {
	*x = BatchGetCollectionBasicReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetCollectionBasicReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetCollectionBasicReq) ProtoMessage() {}

func (x *BatchGetCollectionBasicReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use BatchGetCollectionBasicReq.ProtoReflect.Descriptor instead.
func (*BatchGetCollectionBasicReq) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{5}
}

func (x *BatchGetCollectionBasicReq) GetCid() []string {
	if x != nil {
		return x.Cid
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
		mi := &file_collection_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetCollectionBasicRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetCollectionBasicRsp) ProtoMessage() {}

func (x *BatchGetCollectionBasicRsp) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use BatchGetCollectionBasicRsp.ProtoReflect.Descriptor instead.
func (*BatchGetCollectionBasicRsp) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{6}
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
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a,
	0x0f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x73, 0x69, 0x63,
	0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65,
	0x73, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2a, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x64, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73,
	0x70, 0x12, 0x4c, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x2a, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x22, 0xf4, 0x02, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x73, 0x70, 0x12, 0x55, 0x0a, 0x10, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x0f, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x12, 0x40, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12,
	0x7a, 0x0a, 0x13, 0x6c, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x66,
	0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x73, 0x70, 0x2e,
	0x4c, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x6c, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x79,
	0x50, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x1a, 0x45, 0x0a, 0x17, 0x4c,
	0x6f, 0x74, 0x74, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x2e, 0x0a, 0x1a, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x65, 0x71,
	0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x63,
	0x69, 0x64, 0x22, 0x6a, 0x0a, 0x1a, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x73, 0x70,
	0x12, 0x4c, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x8d,
	0x03, 0x0a, 0x0d, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x76, 0x72,
	0x12, 0x75, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x2f, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x87, 0x01, 0x0a, 0x17, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61,
	0x73, 0x69, 0x63, 0x12, 0x35, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x65, 0x71, 0x1a, 0x35, 0x2e, 0x66, 0x75, 0x6e,
	0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x73,
	0x70, 0x12, 0x7b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x31, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74,
	0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x31, 0x2e, 0x66, 0x75,
	0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x73, 0x70, 0x42, 0x3b,
	0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x75, 0x6e,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
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

var file_collection_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_collection_proto_goTypes = []interface{}{
	(*CollectionBasic)(nil),            // 0: funstar.server.collection.CollectionBasic
	(*GetCollectionListReq)(nil),       // 1: funstar.server.collection.GetCollectionListReq
	(*GetCollectionListRsp)(nil),       // 2: funstar.server.collection.GetCollectionListRsp
	(*GetCollectionDetailReq)(nil),     // 3: funstar.server.collection.GetCollectionDetailReq
	(*GetCollectionDetailRsp)(nil),     // 4: funstar.server.collection.GetCollectionDetailRsp
	(*BatchGetCollectionBasicReq)(nil), // 5: funstar.server.collection.BatchGetCollectionBasicReq
	(*BatchGetCollectionBasicRsp)(nil), // 6: funstar.server.collection.BatchGetCollectionBasicRsp
	nil,                                // 7: funstar.server.collection.GetCollectionDetailRsp.LotteryProbabilityEntry
	(*product.ProductBasic)(nil),       // 8: funstar.server.product.ProductBasic
}
var file_collection_proto_depIdxs = []int32{
	0, // 0: funstar.server.collection.GetCollectionListRsp.collections:type_name -> funstar.server.collection.CollectionBasic
	0, // 1: funstar.server.collection.GetCollectionDetailRsp.collection_basic:type_name -> funstar.server.collection.CollectionBasic
	8, // 2: funstar.server.collection.GetCollectionDetailRsp.products:type_name -> funstar.server.product.ProductBasic
	7, // 3: funstar.server.collection.GetCollectionDetailRsp.lottery_probability:type_name -> funstar.server.collection.GetCollectionDetailRsp.LotteryProbabilityEntry
	0, // 4: funstar.server.collection.BatchGetCollectionBasicRsp.collections:type_name -> funstar.server.collection.CollectionBasic
	1, // 5: funstar.server.collection.CollectionSvr.GetCollectionList:input_type -> funstar.server.collection.GetCollectionListReq
	5, // 6: funstar.server.collection.CollectionSvr.BatchGetCollectionBasic:input_type -> funstar.server.collection.BatchGetCollectionBasicReq
	3, // 7: funstar.server.collection.CollectionSvr.GetCollectionDetail:input_type -> funstar.server.collection.GetCollectionDetailReq
	2, // 8: funstar.server.collection.CollectionSvr.GetCollectionList:output_type -> funstar.server.collection.GetCollectionListRsp
	6, // 9: funstar.server.collection.CollectionSvr.BatchGetCollectionBasic:output_type -> funstar.server.collection.BatchGetCollectionBasicRsp
	4, // 10: funstar.server.collection.CollectionSvr.GetCollectionDetail:output_type -> funstar.server.collection.GetCollectionDetailRsp
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
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
		file_collection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_collection_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_collection_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_collection_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_collection_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
			NumMessages:   8,
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
