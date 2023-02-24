// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: deliver.proto

package deliver

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

// 定义发送请求信息
type SimpleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 定义发送的参数
	// 参数类型 参数名 标识号(不可重复)
	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SimpleRequest) Reset() {
	*x = SimpleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deliver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleRequest) ProtoMessage() {}

func (x *SimpleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deliver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleRequest.ProtoReflect.Descriptor instead.
func (*SimpleRequest) Descriptor() ([]byte, []int) {
	return file_deliver_proto_rawDescGZIP(), []int{0}
}

func (x *SimpleRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type BarData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Insid              string  `protobuf:"bytes,1,opt,name=Insid,proto3" json:"Insid,omitempty"`
	TsOpen             int64   `protobuf:"varint,2,opt,name=Ts_open,json=TsOpen,proto3" json:"Ts_open,omitempty"`
	OpenPrice          float32 `protobuf:"fixed32,3,opt,name=Open_price,json=OpenPrice,proto3" json:"Open_price,omitempty"`
	HighPrice          float32 `protobuf:"fixed32,4,opt,name=High_price,json=HighPrice,proto3" json:"High_price,omitempty"`
	LowPrice           float32 `protobuf:"fixed32,5,opt,name=Low_price,json=LowPrice,proto3" json:"Low_price,omitempty"`
	ClosePrice         float32 `protobuf:"fixed32,6,opt,name=Close_price,json=ClosePrice,proto3" json:"Close_price,omitempty"`
	Vol                float32 `protobuf:"fixed32,7,opt,name=Vol,proto3" json:"Vol,omitempty"`
	VolCcy             float32 `protobuf:"fixed32,8,opt,name=VolCcy,proto3" json:"VolCcy,omitempty"`
	VolCcyQuote        float32 `protobuf:"fixed32,9,opt,name=VolCcyQuote,proto3" json:"VolCcyQuote,omitempty"`
	Oi                 float32 `protobuf:"fixed32,10,opt,name=Oi,proto3" json:"Oi,omitempty"`
	OiCcy              float32 `protobuf:"fixed32,11,opt,name=OiCcy,proto3" json:"OiCcy,omitempty"`
	TsOi               int64   `protobuf:"varint,12,opt,name=Ts_oi,json=TsOi,proto3" json:"Ts_oi,omitempty"`
	FundingRate        float32 `protobuf:"fixed32,13,opt,name=FundingRate,proto3" json:"FundingRate,omitempty"`
	NextFundingRate    float32 `protobuf:"fixed32,14,opt,name=NextFundingRate,proto3" json:"NextFundingRate,omitempty"`
	Ts_FundingRate     int64   `protobuf:"varint,15,opt,name=Ts_FundingRate,json=TsFundingRate,proto3" json:"Ts_FundingRate,omitempty"`
	TS_NextFundingRate int64   `protobuf:"varint,16,opt,name=TS_NextFundingRate,json=TSNextFundingRate,proto3" json:"TS_NextFundingRate,omitempty"`
}

func (x *BarData) Reset() {
	*x = BarData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deliver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BarData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BarData) ProtoMessage() {}

func (x *BarData) ProtoReflect() protoreflect.Message {
	mi := &file_deliver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BarData.ProtoReflect.Descriptor instead.
func (*BarData) Descriptor() ([]byte, []int) {
	return file_deliver_proto_rawDescGZIP(), []int{1}
}

func (x *BarData) GetInsid() string {
	if x != nil {
		return x.Insid
	}
	return ""
}

func (x *BarData) GetTsOpen() int64 {
	if x != nil {
		return x.TsOpen
	}
	return 0
}

func (x *BarData) GetOpenPrice() float32 {
	if x != nil {
		return x.OpenPrice
	}
	return 0
}

func (x *BarData) GetHighPrice() float32 {
	if x != nil {
		return x.HighPrice
	}
	return 0
}

func (x *BarData) GetLowPrice() float32 {
	if x != nil {
		return x.LowPrice
	}
	return 0
}

func (x *BarData) GetClosePrice() float32 {
	if x != nil {
		return x.ClosePrice
	}
	return 0
}

func (x *BarData) GetVol() float32 {
	if x != nil {
		return x.Vol
	}
	return 0
}

func (x *BarData) GetVolCcy() float32 {
	if x != nil {
		return x.VolCcy
	}
	return 0
}

func (x *BarData) GetVolCcyQuote() float32 {
	if x != nil {
		return x.VolCcyQuote
	}
	return 0
}

func (x *BarData) GetOi() float32 {
	if x != nil {
		return x.Oi
	}
	return 0
}

func (x *BarData) GetOiCcy() float32 {
	if x != nil {
		return x.OiCcy
	}
	return 0
}

func (x *BarData) GetTsOi() int64 {
	if x != nil {
		return x.TsOi
	}
	return 0
}

func (x *BarData) GetFundingRate() float32 {
	if x != nil {
		return x.FundingRate
	}
	return 0
}

func (x *BarData) GetNextFundingRate() float32 {
	if x != nil {
		return x.NextFundingRate
	}
	return 0
}

func (x *BarData) GetTs_FundingRate() int64 {
	if x != nil {
		return x.Ts_FundingRate
	}
	return 0
}

func (x *BarData) GetTS_NextFundingRate() int64 {
	if x != nil {
		return x.TS_NextFundingRate
	}
	return 0
}

type TickData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Insid      string  `protobuf:"bytes,1,opt,name=Insid,proto3" json:"Insid,omitempty"`
	Ts_Price   int64   `protobuf:"varint,2,opt,name=Ts_Price,json=TsPrice,proto3" json:"Ts_Price,omitempty"`
	Ask1Price  float32 `protobuf:"fixed32,3,opt,name=Ask1_price,json=Ask1Price,proto3" json:"Ask1_price,omitempty"`
	Bid1Price  float32 `protobuf:"fixed32,4,opt,name=Bid1_price,json=Bid1Price,proto3" json:"Bid1_price,omitempty"`
	Ask1Volumn float32 `protobuf:"fixed32,5,opt,name=Ask1_volumn,json=Ask1Volumn,proto3" json:"Ask1_volumn,omitempty"`
	Bid1Volumn float32 `protobuf:"fixed32,6,opt,name=Bid1_volumn,json=Bid1Volumn,proto3" json:"Bid1_volumn,omitempty"`
}

func (x *TickData) Reset() {
	*x = TickData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deliver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TickData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TickData) ProtoMessage() {}

func (x *TickData) ProtoReflect() protoreflect.Message {
	mi := &file_deliver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TickData.ProtoReflect.Descriptor instead.
func (*TickData) Descriptor() ([]byte, []int) {
	return file_deliver_proto_rawDescGZIP(), []int{2}
}

func (x *TickData) GetInsid() string {
	if x != nil {
		return x.Insid
	}
	return ""
}

func (x *TickData) GetTs_Price() int64 {
	if x != nil {
		return x.Ts_Price
	}
	return 0
}

func (x *TickData) GetAsk1Price() float32 {
	if x != nil {
		return x.Ask1Price
	}
	return 0
}

func (x *TickData) GetBid1Price() float32 {
	if x != nil {
		return x.Bid1Price
	}
	return 0
}

func (x *TickData) GetAsk1Volumn() float32 {
	if x != nil {
		return x.Ask1Volumn
	}
	return 0
}

func (x *TickData) GetBid1Volumn() float32 {
	if x != nil {
		return x.Bid1Volumn
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseMe string `protobuf:"bytes,1,opt,name=response_me,json=responseMe,proto3" json:"response_me,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deliver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_deliver_proto_msgTypes[3]
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
	return file_deliver_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetResponseMe() string {
	if x != nil {
		return x.ResponseMe
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InsId           string `protobuf:"bytes,1,opt,name=insId,proto3" json:"insId,omitempty"`                      // 产品ID   必填
	TdMode          string `protobuf:"bytes,2,opt,name=tdMode,proto3" json:"tdMode,omitempty"`                    // 交易模式 必填
	Ccy             string `protobuf:"bytes,3,opt,name=ccy,proto3" json:"ccy,omitempty"`                          // 保证金币种 非必填
	ClOrdId         string `protobuf:"bytes,4,opt,name=clOrdId,proto3" json:"clOrdId,omitempty"`                  // 客户自定义订单ID 非必填
	Tag             string `protobuf:"bytes,5,opt,name=tag,proto3" json:"tag,omitempty"`                          // 订单标签 非必填
	Side            string `protobuf:"bytes,6,opt,name=side,proto3" json:"side,omitempty"`                        // 订单方向 必填 buy：买， sell：卖
	PosSide         string `protobuf:"bytes,7,opt,name=posSide,proto3" json:"posSide,omitempty"`                  // 持仓方向 在双向持仓模式下必填，且仅可选择 long 或 short。 仅适用交割、永续。
	OrdType         string `protobuf:"bytes,8,opt,name=ordType,proto3" json:"ordType,omitempty"`                  // 订单类型 必填 // market：市价单，limit：限价单，post_only：只做maker单，fok：全部成交或立即取消，ioc：立即成交并取消剩余，optimal_limit_ioc：市价委托立即成交并取消剩余（仅适用交割、永续）
	Sz              string `protobuf:"bytes,9,opt,name=sz,proto3" json:"sz,omitempty"`                            // 委托数量 必填
	Px              string `protobuf:"bytes,10,opt,name=px,proto3" json:"px,omitempty"`                           // 委托价格，仅适用于limit、post_only、fok、ioc类型的订单 非必填
	ReduceOnly      bool   `protobuf:"varint,11,opt,name=reduceOnly,proto3" json:"reduceOnly,omitempty"`          // 是否只减仓，true 或 false，默认false 非必填
	TgtCcy          string `protobuf:"bytes,12,opt,name=tgtCcy,proto3" json:"tgtCcy,omitempty"`                   // 市价单委托数量sz的单位，仅适用于币币市价订单 非必填
	BanAmend        bool   `protobuf:"varint,13,opt,name=banAmend,proto3" json:"banAmend,omitempty"`              // 是否禁止币币市价改单，true 或 false，默认false 非必填
	TpTriggerPx     string `protobuf:"bytes,14,opt,name=tpTriggerPx,proto3" json:"tpTriggerPx,omitempty"`         // 止盈触发价，如果填写此参数，必须填写 止盈委托价 非必填
	TpOrdPx         string `protobuf:"bytes,15,opt,name=tpOrdPx,proto3" json:"tpOrdPx,omitempty"`                 // 止盈委托价，如果填写此参数，必须填写 止盈触发价 非必填（PS:委托价格为-1时，执行市价止盈）
	SlTriggerPx     string `protobuf:"bytes,16,opt,name=slTriggerPx,proto3" json:"slTriggerPx,omitempty"`         // 止损触发价，如果填写此参数，必须填写 止损委托价 非必填
	SlOrdPx         string `protobuf:"bytes,17,opt,name=slOrdPx,proto3" json:"slOrdPx,omitempty"`                 // 止损委托价，如果填写此参数，必须填写 止损触发价 非必填 （PS:委托价格为-1时，执行市价止损）
	TpTriggerPxType string `protobuf:"bytes,18,opt,name=tpTriggerPxType,proto3" json:"tpTriggerPxType,omitempty"` // 止盈触发类型 last：最新价格 index：指数价格 mark：标记价格 默认为last
	SlTriggerPxType string `protobuf:"bytes,19,opt,name=slTriggerPxType,proto3" json:"slTriggerPxType,omitempty"` // 止损触发价类型 last：最新价格 index：指数价格 mark：标记价格 默认为last
	QuickMgnType    string `protobuf:"bytes,20,opt,name=quickMgnType,proto3" json:"quickMgnType,omitempty"`       // 一键借币类型，仅适用于杠杆逐仓的一键借币模式：manual：手动，auto_borrow： 自动借币，auto_repay： 自动还币
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deliver_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_deliver_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_deliver_proto_rawDescGZIP(), []int{4}
}

func (x *Order) GetInsId() string {
	if x != nil {
		return x.InsId
	}
	return ""
}

func (x *Order) GetTdMode() string {
	if x != nil {
		return x.TdMode
	}
	return ""
}

func (x *Order) GetCcy() string {
	if x != nil {
		return x.Ccy
	}
	return ""
}

func (x *Order) GetClOrdId() string {
	if x != nil {
		return x.ClOrdId
	}
	return ""
}

func (x *Order) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Order) GetSide() string {
	if x != nil {
		return x.Side
	}
	return ""
}

func (x *Order) GetPosSide() string {
	if x != nil {
		return x.PosSide
	}
	return ""
}

func (x *Order) GetOrdType() string {
	if x != nil {
		return x.OrdType
	}
	return ""
}

func (x *Order) GetSz() string {
	if x != nil {
		return x.Sz
	}
	return ""
}

func (x *Order) GetPx() string {
	if x != nil {
		return x.Px
	}
	return ""
}

func (x *Order) GetReduceOnly() bool {
	if x != nil {
		return x.ReduceOnly
	}
	return false
}

func (x *Order) GetTgtCcy() string {
	if x != nil {
		return x.TgtCcy
	}
	return ""
}

func (x *Order) GetBanAmend() bool {
	if x != nil {
		return x.BanAmend
	}
	return false
}

func (x *Order) GetTpTriggerPx() string {
	if x != nil {
		return x.TpTriggerPx
	}
	return ""
}

func (x *Order) GetTpOrdPx() string {
	if x != nil {
		return x.TpOrdPx
	}
	return ""
}

func (x *Order) GetSlTriggerPx() string {
	if x != nil {
		return x.SlTriggerPx
	}
	return ""
}

func (x *Order) GetSlOrdPx() string {
	if x != nil {
		return x.SlOrdPx
	}
	return ""
}

func (x *Order) GetTpTriggerPxType() string {
	if x != nil {
		return x.TpTriggerPxType
	}
	return ""
}

func (x *Order) GetSlTriggerPxType() string {
	if x != nil {
		return x.SlTriggerPxType
	}
	return ""
}

func (x *Order) GetQuickMgnType() string {
	if x != nil {
		return x.QuickMgnType
	}
	return ""
}

var File_deliver_proto protoreflect.FileDescriptor

var file_deliver_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x23, 0x0a, 0x0d, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0xdd, 0x03, 0x0a, 0x07, 0x42, 0x61, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x49, 0x6e, 0x73, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x54, 0x73, 0x5f, 0x6f, 0x70, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x54, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x4f, 0x70, 0x65, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x09, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x48, 0x69, 0x67, 0x68, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x09, 0x48, 0x69, 0x67, 0x68, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x4c, 0x6f, 0x77, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x08, 0x4c, 0x6f, 0x77, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0a, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x56,
	0x6f, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x56, 0x6f, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x56, 0x6f, 0x6c, 0x43, 0x63, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x56,
	0x6f, 0x6c, 0x43, 0x63, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x56, 0x6f, 0x6c, 0x43, 0x63, 0x79, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x56, 0x6f, 0x6c, 0x43,
	0x63, 0x79, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x69, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x02, 0x4f, 0x69, 0x12, 0x14, 0x0a, 0x05, 0x4f, 0x69, 0x43, 0x63, 0x79,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x4f, 0x69, 0x43, 0x63, 0x79, 0x12, 0x13, 0x0a,
	0x05, 0x54, 0x73, 0x5f, 0x6f, 0x69, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x54, 0x73,
	0x4f, 0x69, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74,
	0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x46, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x52, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x4e, 0x65, 0x78, 0x74, 0x46, 0x75, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x4e,
	0x65, 0x78, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x54, 0x73, 0x5f, 0x46, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x65,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x54, 0x73, 0x46, 0x75, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x54, 0x53, 0x5f, 0x4e, 0x65, 0x78, 0x74,
	0x46, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x11, 0x54, 0x53, 0x4e, 0x65, 0x78, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x52, 0x61, 0x74, 0x65, 0x22, 0xbb, 0x01, 0x0a, 0x08, 0x54, 0x69, 0x63, 0x6b, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x49, 0x6e, 0x73, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x54, 0x73, 0x5f, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x54, 0x73, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x41, 0x73, 0x6b, 0x31, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x41, 0x73, 0x6b, 0x31, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x42, 0x69, 0x64, 0x31, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x42, 0x69, 0x64, 0x31, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x41, 0x73, 0x6b, 0x31, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x41, 0x73, 0x6b, 0x31, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x42, 0x69, 0x64, 0x31, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x42, 0x69, 0x64, 0x31, 0x56, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x22, 0x2b, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x22,
	0x9f, 0x04, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x73,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x73, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x63, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x63, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x4f,
	0x72, 0x64, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x4f, 0x72,
	0x64, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x64, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6f, 0x73,
	0x53, 0x69, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x73, 0x53,
	0x69, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x73, 0x7a, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x73, 0x7a, 0x12, 0x0e, 0x0a,
	0x02, 0x70, 0x78, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x70, 0x78, 0x12, 0x1e, 0x0a,
	0x0a, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x67, 0x74, 0x43, 0x63, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x67, 0x74, 0x43, 0x63, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x61, 0x6e, 0x41, 0x6d, 0x65, 0x6e,
	0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x41, 0x6d, 0x65, 0x6e,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x70, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x50, 0x78,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x70, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x50, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x70, 0x4f, 0x72, 0x64, 0x50, 0x78, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x70, 0x4f, 0x72, 0x64, 0x50, 0x78, 0x12, 0x20, 0x0a,
	0x0b, 0x73, 0x6c, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x50, 0x78, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x6c, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x50, 0x78, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x6c, 0x4f, 0x72, 0x64, 0x50, 0x78, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x6c, 0x4f, 0x72, 0x64, 0x50, 0x78, 0x12, 0x28, 0x0a, 0x0f, 0x74, 0x70, 0x54,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x50, 0x78, 0x54, 0x79, 0x70, 0x65, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x74, 0x70, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x50, 0x78, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x6c, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x50, 0x78, 0x54, 0x79, 0x70, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x6c,
	0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x50, 0x78, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x4d, 0x67, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x4d, 0x67, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x32, 0x3b, 0x0a, 0x0f, 0x42, 0x61, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x42, 0x61, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x08, 0x2e, 0x42, 0x61, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x3e,
	0x0a, 0x10, 0x54, 0x69, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x54, 0x69, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x09, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x44, 0x61, 0x74,
	0x61, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x34,
	0x0a, 0x0c, 0x4f, 0x72, 0x65, 0x72, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x24,
	0x0a, 0x0d, 0x4f, 0x72, 0x65, 0x72, 0x52, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12,
	0x06, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deliver_proto_rawDescOnce sync.Once
	file_deliver_proto_rawDescData = file_deliver_proto_rawDesc
)

func file_deliver_proto_rawDescGZIP() []byte {
	file_deliver_proto_rawDescOnce.Do(func() {
		file_deliver_proto_rawDescData = protoimpl.X.CompressGZIP(file_deliver_proto_rawDescData)
	})
	return file_deliver_proto_rawDescData
}

var file_deliver_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_deliver_proto_goTypes = []interface{}{
	(*SimpleRequest)(nil), // 0: SimpleRequest
	(*BarData)(nil),       // 1: BarData
	(*TickData)(nil),      // 2: TickData
	(*Response)(nil),      // 3: Response
	(*Order)(nil),         // 4: Order
}
var file_deliver_proto_depIdxs = []int32{
	1, // 0: BarDataReceiver.BarDataReceiver:input_type -> BarData
	2, // 1: TickDataReceiver.TickDataReceiver:input_type -> TickData
	4, // 2: OrerReceiver.OrerRReceiver:input_type -> Order
	3, // 3: BarDataReceiver.BarDataReceiver:output_type -> Response
	3, // 4: TickDataReceiver.TickDataReceiver:output_type -> Response
	3, // 5: OrerReceiver.OrerRReceiver:output_type -> Response
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_deliver_proto_init() }
func file_deliver_proto_init() {
	if File_deliver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_deliver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleRequest); i {
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
		file_deliver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BarData); i {
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
		file_deliver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TickData); i {
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
		file_deliver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_deliver_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
			RawDescriptor: file_deliver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_deliver_proto_goTypes,
		DependencyIndexes: file_deliver_proto_depIdxs,
		MessageInfos:      file_deliver_proto_msgTypes,
	}.Build()
	File_deliver_proto = out.File
	file_deliver_proto_rawDesc = nil
	file_deliver_proto_goTypes = nil
	file_deliver_proto_depIdxs = nil
}
