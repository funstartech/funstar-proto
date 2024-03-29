// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: product.proto

package product

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

// 出售状态
type SaleStatus int32

const (
	SaleStatus_SALE_STATUS_UNKNOWN     SaleStatus = 0 // 未知
	SaleStatus_SALE_STATUS_AVAILABLE   SaleStatus = 1 // 现货
	SaleStatus_SALE_STATUS_UNAVAILABLE SaleStatus = 2 // 非现货
)

// Enum value maps for SaleStatus.
var (
	SaleStatus_name = map[int32]string{
		0: "SALE_STATUS_UNKNOWN",
		1: "SALE_STATUS_AVAILABLE",
		2: "SALE_STATUS_UNAVAILABLE",
	}
	SaleStatus_value = map[string]int32{
		"SALE_STATUS_UNKNOWN":     0,
		"SALE_STATUS_AVAILABLE":   1,
		"SALE_STATUS_UNAVAILABLE": 2,
	}
)

func (x SaleStatus) Enum() *SaleStatus {
	p := new(SaleStatus)
	*p = x
	return p
}

func (x SaleStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SaleStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_product_proto_enumTypes[0].Descriptor()
}

func (SaleStatus) Type() protoreflect.EnumType {
	return &file_product_proto_enumTypes[0]
}

func (x SaleStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SaleStatus.Descriptor instead.
func (SaleStatus) EnumDescriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

// 退货条款
type ReturnPolicy int32

const (
	ReturnPolicy_STATEMENT_UNKNOWN        ReturnPolicy = 0 // 未知
	ReturnPolicy_STATEMENT_SUPPORT_7DAY   ReturnPolicy = 1 // 支持7天无理由退货
	ReturnPolicy_STATEMENT_UNSUPPORT_7DAY ReturnPolicy = 2 // 不支持7天无理由退货
)

// Enum value maps for ReturnPolicy.
var (
	ReturnPolicy_name = map[int32]string{
		0: "STATEMENT_UNKNOWN",
		1: "STATEMENT_SUPPORT_7DAY",
		2: "STATEMENT_UNSUPPORT_7DAY",
	}
	ReturnPolicy_value = map[string]int32{
		"STATEMENT_UNKNOWN":        0,
		"STATEMENT_SUPPORT_7DAY":   1,
		"STATEMENT_UNSUPPORT_7DAY": 2,
	}
)

func (x ReturnPolicy) Enum() *ReturnPolicy {
	p := new(ReturnPolicy)
	*p = x
	return p
}

func (x ReturnPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReturnPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_product_proto_enumTypes[1].Descriptor()
}

func (ReturnPolicy) Type() protoreflect.EnumType {
	return &file_product_proto_enumTypes[1]
}

func (x ReturnPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReturnPolicy.Descriptor instead.
func (ReturnPolicy) EnumDescriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{1}
}

// 商品基础信息结构
type ProductBasic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid           string  `protobuf:"bytes,1,opt,name=pid,proto3" json:"pid,omitempty"`                                            // 商品id
	Cover         string  `protobuf:"bytes,2,opt,name=cover,proto3" json:"cover,omitempty"`                                        // 封面图
	Title         string  `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`                                        // 标题
	Desc          string  `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`                                          // 描述
	Price         int32   `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`                                       // 价格(单位：分)
	Brand         string  `protobuf:"bytes,6,opt,name=brand,proto3" json:"brand,omitempty"`                                        // 品牌
	Stock         uint32  `protobuf:"varint,7,opt,name=stock,proto3" json:"stock,omitempty"`                                       // 库存数量
	PackageWeight float32 `protobuf:"fixed32,8,opt,name=package_weight,json=packageWeight,proto3" json:"package_weight,omitempty"` // 包装权重
	SaleStatus    uint32  `protobuf:"varint,101,opt,name=sale_status,json=saleStatus,proto3" json:"sale_status,omitempty"`         // 出售状态 SaleStatus
	ReturnPolicy  uint32  `protobuf:"varint,102,opt,name=return_policy,json=returnPolicy,proto3" json:"return_policy,omitempty"`   // 退货条款 ReturnPolicy
}

func (x *ProductBasic) Reset() {
	*x = ProductBasic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductBasic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductBasic) ProtoMessage() {}

func (x *ProductBasic) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[0]
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
	return file_product_proto_rawDescGZIP(), []int{0}
}

func (x *ProductBasic) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
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

func (x *ProductBasic) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ProductBasic) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *ProductBasic) GetStock() uint32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *ProductBasic) GetPackageWeight() float32 {
	if x != nil {
		return x.PackageWeight
	}
	return 0
}

func (x *ProductBasic) GetSaleStatus() uint32 {
	if x != nil {
		return x.SaleStatus
	}
	return 0
}

func (x *ProductBasic) GetReturnPolicy() uint32 {
	if x != nil {
		return x.ReturnPolicy
	}
	return 0
}

// 商品详情结构
type ProductDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid      string   `protobuf:"bytes,1,opt,name=pid,proto3" json:"pid,omitempty"`                       // 商品id
	Pcode    string   `protobuf:"bytes,11,opt,name=pcode,proto3" json:"pcode,omitempty"`                  // 条形码
	Material string   `protobuf:"bytes,12,opt,name=material,proto3" json:"material,omitempty"`            // 材质
	Size     string   `protobuf:"bytes,13,opt,name=size,proto3" json:"size,omitempty"`                    // 尺寸
	UseAge   uint32   `protobuf:"varint,14,opt,name=use_age,json=useAge,proto3" json:"use_age,omitempty"` // 适用年龄
	Images   []string `protobuf:"bytes,15,rep,name=images,proto3" json:"images,omitempty"`                // 详情图片
}

func (x *ProductDetail) Reset() {
	*x = ProductDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductDetail) ProtoMessage() {}

func (x *ProductDetail) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductDetail.ProtoReflect.Descriptor instead.
func (*ProductDetail) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductDetail) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

func (x *ProductDetail) GetPcode() string {
	if x != nil {
		return x.Pcode
	}
	return ""
}

func (x *ProductDetail) GetMaterial() string {
	if x != nil {
		return x.Material
	}
	return ""
}

func (x *ProductDetail) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *ProductDetail) GetUseAge() uint32 {
	if x != nil {
		return x.UseAge
	}
	return 0
}

func (x *ProductDetail) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

type BatchGetProductBasicReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pids []string `protobuf:"bytes,1,rep,name=pids,proto3" json:"pids,omitempty"`
}

func (x *BatchGetProductBasicReq) Reset() {
	*x = BatchGetProductBasicReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetProductBasicReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetProductBasicReq) ProtoMessage() {}

func (x *BatchGetProductBasicReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetProductBasicReq.ProtoReflect.Descriptor instead.
func (*BatchGetProductBasicReq) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{2}
}

func (x *BatchGetProductBasicReq) GetPids() []string {
	if x != nil {
		return x.Pids
	}
	return nil
}

type BatchGetProductBasicRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*ProductBasic `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *BatchGetProductBasicRsp) Reset() {
	*x = BatchGetProductBasicRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetProductBasicRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetProductBasicRsp) ProtoMessage() {}

func (x *BatchGetProductBasicRsp) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetProductBasicRsp.ProtoReflect.Descriptor instead.
func (*BatchGetProductBasicRsp) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{3}
}

func (x *BatchGetProductBasicRsp) GetProducts() []*ProductBasic {
	if x != nil {
		return x.Products
	}
	return nil
}

type GetProductDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid string `protobuf:"bytes,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *GetProductDetailReq) Reset() {
	*x = GetProductDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductDetailReq) ProtoMessage() {}

func (x *GetProductDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductDetailReq.ProtoReflect.Descriptor instead.
func (*GetProductDetailReq) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{4}
}

func (x *GetProductDetailReq) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

type GetProductDetailRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Basic  *ProductBasic  `protobuf:"bytes,1,opt,name=basic,proto3" json:"basic,omitempty"`   // 基础信息
	Detail *ProductDetail `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"` // 详情
}

func (x *GetProductDetailRsp) Reset() {
	*x = GetProductDetailRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductDetailRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductDetailRsp) ProtoMessage() {}

func (x *GetProductDetailRsp) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductDetailRsp.ProtoReflect.Descriptor instead.
func (*GetProductDetailRsp) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{5}
}

func (x *GetProductDetailRsp) GetBasic() *ProductBasic {
	if x != nil {
		return x.Basic
	}
	return nil
}

func (x *GetProductDetailRsp) GetDetail() *ProductDetail {
	if x != nil {
		return x.Detail
	}
	return nil
}

type ProductStock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid   string `protobuf:"bytes,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Stock uint32 `protobuf:"varint,2,opt,name=stock,proto3" json:"stock,omitempty"`
}

func (x *ProductStock) Reset() {
	*x = ProductStock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductStock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductStock) ProtoMessage() {}

func (x *ProductStock) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductStock.ProtoReflect.Descriptor instead.
func (*ProductStock) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{6}
}

func (x *ProductStock) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

func (x *ProductStock) GetStock() uint32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

type GetProductStockReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pids []string `protobuf:"bytes,1,rep,name=pids,proto3" json:"pids,omitempty"`
}

func (x *GetProductStockReq) Reset() {
	*x = GetProductStockReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductStockReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductStockReq) ProtoMessage() {}

func (x *GetProductStockReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductStockReq.ProtoReflect.Descriptor instead.
func (*GetProductStockReq) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{7}
}

func (x *GetProductStockReq) GetPids() []string {
	if x != nil {
		return x.Pids
	}
	return nil
}

type GetProductStockRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stocks []*ProductStock `protobuf:"bytes,1,rep,name=stocks,proto3" json:"stocks,omitempty"`
}

func (x *GetProductStockRsp) Reset() {
	*x = GetProductStockRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductStockRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductStockRsp) ProtoMessage() {}

func (x *GetProductStockRsp) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductStockRsp.ProtoReflect.Descriptor instead.
func (*GetProductStockRsp) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{8}
}

func (x *GetProductStockRsp) GetStocks() []*ProductStock {
	if x != nil {
		return x.Stocks
	}
	return nil
}

type SetProductStockReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stocks []*ProductStock `protobuf:"bytes,1,rep,name=stocks,proto3" json:"stocks,omitempty"`
}

func (x *SetProductStockReq) Reset() {
	*x = SetProductStockReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetProductStockReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetProductStockReq) ProtoMessage() {}

func (x *SetProductStockReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetProductStockReq.ProtoReflect.Descriptor instead.
func (*SetProductStockReq) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{9}
}

func (x *SetProductStockReq) GetStocks() []*ProductStock {
	if x != nil {
		return x.Stocks
	}
	return nil
}

type SetProductStockRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetProductStockRsp) Reset() {
	*x = SetProductStockRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetProductStockRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetProductStockRsp) ProtoMessage() {}

func (x *SetProductStockRsp) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetProductStockRsp.ProtoReflect.Descriptor instead.
func (*SetProductStockRsp) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{10}
}

var File_product_proto protoreflect.FileDescriptor

var file_product_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x16, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x8f, 0x02, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x42, 0x61, 0x73, 0x69, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x25, 0x0a, 0x0e,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x57, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x72, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x98, 0x01, 0x0a, 0x0d, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x70,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x41, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x22, 0x2d, 0x0a, 0x17, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x65, 0x71, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x69, 0x64, 0x73, 0x22, 0x5b, 0x0a, 0x17, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x73, 0x70, 0x12, 0x40,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x22, 0x27, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x90, 0x01, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x73,
	0x70, 0x12, 0x3a, 0x0a, 0x05, 0x62, 0x61, 0x73, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x05, 0x62, 0x61, 0x73, 0x69, 0x63, 0x12, 0x3d, 0x0a,
	0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x36, 0x0a, 0x0c,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x10, 0x0a, 0x03,
	0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x22, 0x28, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x70, 0x69, 0x64, 0x73, 0x22, 0x52,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x52, 0x73, 0x70, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x73, 0x22, 0x52, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74,
	0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x06,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x22, 0x14, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x73, 0x70, 0x2a, 0x5d, 0x0a, 0x0a,
	0x53, 0x61, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x41,
	0x4c, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x41, 0x4c, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x1b,
	0x0a, 0x17, 0x53, 0x41, 0x4c, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e,
	0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x2a, 0x5f, 0x0a, 0x0c, 0x52,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x15, 0x0a, 0x11, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x54, 0x41, 0x54, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x37, 0x44, 0x41, 0x59, 0x10, 0x01, 0x12, 0x1c,
	0x0a, 0x18, 0x53, 0x54, 0x41, 0x54, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x55,
	0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x37, 0x44, 0x41, 0x59, 0x10, 0x02, 0x32, 0xca, 0x03, 0x0a,
	0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x76, 0x72, 0x12, 0x78, 0x0a, 0x14, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61,
	0x73, 0x69, 0x63, 0x12, 0x2f, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x52, 0x65, 0x71, 0x1a, 0x2f, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x52, 0x73, 0x70, 0x12, 0x6c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x2b, 0x2e, 0x66, 0x75, 0x6e, 0x73,
	0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x2b, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x73, 0x70, 0x12, 0x69, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x2a, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x1a, 0x2a, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x73, 0x70, 0x12, 0x69,
	0x0a, 0x0f, 0x53, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x12, 0x2a, 0x2e, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x2a, 0x2e,
	0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x73, 0x70, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x65, 0x63, 0x68, 0x2f, 0x66, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_proto_rawDescOnce sync.Once
	file_product_proto_rawDescData = file_product_proto_rawDesc
)

func file_product_proto_rawDescGZIP() []byte {
	file_product_proto_rawDescOnce.Do(func() {
		file_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_proto_rawDescData)
	})
	return file_product_proto_rawDescData
}

var file_product_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_product_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_product_proto_goTypes = []interface{}{
	(SaleStatus)(0),                 // 0: funstar.server.product.SaleStatus
	(ReturnPolicy)(0),               // 1: funstar.server.product.ReturnPolicy
	(*ProductBasic)(nil),            // 2: funstar.server.product.ProductBasic
	(*ProductDetail)(nil),           // 3: funstar.server.product.ProductDetail
	(*BatchGetProductBasicReq)(nil), // 4: funstar.server.product.BatchGetProductBasicReq
	(*BatchGetProductBasicRsp)(nil), // 5: funstar.server.product.BatchGetProductBasicRsp
	(*GetProductDetailReq)(nil),     // 6: funstar.server.product.GetProductDetailReq
	(*GetProductDetailRsp)(nil),     // 7: funstar.server.product.GetProductDetailRsp
	(*ProductStock)(nil),            // 8: funstar.server.product.ProductStock
	(*GetProductStockReq)(nil),      // 9: funstar.server.product.GetProductStockReq
	(*GetProductStockRsp)(nil),      // 10: funstar.server.product.GetProductStockRsp
	(*SetProductStockReq)(nil),      // 11: funstar.server.product.SetProductStockReq
	(*SetProductStockRsp)(nil),      // 12: funstar.server.product.SetProductStockRsp
}
var file_product_proto_depIdxs = []int32{
	2,  // 0: funstar.server.product.BatchGetProductBasicRsp.products:type_name -> funstar.server.product.ProductBasic
	2,  // 1: funstar.server.product.GetProductDetailRsp.basic:type_name -> funstar.server.product.ProductBasic
	3,  // 2: funstar.server.product.GetProductDetailRsp.detail:type_name -> funstar.server.product.ProductDetail
	8,  // 3: funstar.server.product.GetProductStockRsp.stocks:type_name -> funstar.server.product.ProductStock
	8,  // 4: funstar.server.product.SetProductStockReq.stocks:type_name -> funstar.server.product.ProductStock
	4,  // 5: funstar.server.product.ProductSvr.BatchGetProductBasic:input_type -> funstar.server.product.BatchGetProductBasicReq
	6,  // 6: funstar.server.product.ProductSvr.GetProductDetail:input_type -> funstar.server.product.GetProductDetailReq
	9,  // 7: funstar.server.product.ProductSvr.GetProductStock:input_type -> funstar.server.product.GetProductStockReq
	11, // 8: funstar.server.product.ProductSvr.SetProductStock:input_type -> funstar.server.product.SetProductStockReq
	5,  // 9: funstar.server.product.ProductSvr.BatchGetProductBasic:output_type -> funstar.server.product.BatchGetProductBasicRsp
	7,  // 10: funstar.server.product.ProductSvr.GetProductDetail:output_type -> funstar.server.product.GetProductDetailRsp
	10, // 11: funstar.server.product.ProductSvr.GetProductStock:output_type -> funstar.server.product.GetProductStockRsp
	12, // 12: funstar.server.product.ProductSvr.SetProductStock:output_type -> funstar.server.product.SetProductStockRsp
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_product_proto_init() }
func file_product_proto_init() {
	if File_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductDetail); i {
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
		file_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetProductBasicReq); i {
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
		file_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetProductBasicRsp); i {
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
		file_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductDetailReq); i {
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
		file_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductDetailRsp); i {
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
		file_product_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductStock); i {
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
		file_product_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductStockReq); i {
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
		file_product_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductStockRsp); i {
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
		file_product_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetProductStockReq); i {
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
		file_product_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetProductStockRsp); i {
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
			RawDescriptor: file_product_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_proto_goTypes,
		DependencyIndexes: file_product_proto_depIdxs,
		EnumInfos:         file_product_proto_enumTypes,
		MessageInfos:      file_product_proto_msgTypes,
	}.Build()
	File_product_proto = out.File
	file_product_proto_rawDesc = nil
	file_product_proto_goTypes = nil
	file_product_proto_depIdxs = nil
}
