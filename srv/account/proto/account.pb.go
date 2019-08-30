// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv/account/proto/account.proto

package account

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

// 账户级别
type AccountBase_Level int32

const (
	AccountBase_LevelDefault AccountBase_Level = 0
	AccountBase_LevelOne     AccountBase_Level = 1
	AccountBase_LevelTwo     AccountBase_Level = 2
	AccountBase_LevelThree   AccountBase_Level = 3
)

var AccountBase_Level_name = map[int32]string{
	0: "LevelDefault",
	1: "LevelOne",
	2: "LevelTwo",
	3: "LevelThree",
}

var AccountBase_Level_value = map[string]int32{
	"LevelDefault": 0,
	"LevelOne":     1,
	"LevelTwo":     2,
	"LevelThree":   3,
}

func (x AccountBase_Level) String() string {
	return proto.EnumName(AccountBase_Level_name, int32(x))
}

func (AccountBase_Level) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_89ed2ef530ab98be, []int{0, 0}
}

// 账户状态
type AccountBase_Status int32

const (
	AccountBase_StatusDefault AccountBase_Status = 0
	// 正常
	AccountBase_StatusNormal AccountBase_Status = 1
	// 冻结
	AccountBase_StatusFrozen AccountBase_Status = 2
)

var AccountBase_Status_name = map[int32]string{
	0: "StatusDefault",
	1: "StatusNormal",
	2: "StatusFrozen",
}

var AccountBase_Status_value = map[string]int32{
	"StatusDefault": 0,
	"StatusNormal":  1,
	"StatusFrozen":  2,
}

func (x AccountBase_Status) String() string {
	return proto.EnumName(AccountBase_Status_name, int32(x))
}

func (AccountBase_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_89ed2ef530ab98be, []int{0, 1}
}

// 账户基本信息
type AccountBase struct {
	// @inject_tag: db:"id"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" db:"id"`
	// @inject_tag: db:"user_id" valid:"required~用户id必须存在"
	UserId int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" db:"user_id" valid:"required~用户id必须存在"`
	// @inject_tag: db:"account_level" valid:"required~账户级别必须存在"
	AccountLevel AccountBase_Level `protobuf:"varint,3,opt,name=account_level,json=accountLevel,proto3,enum=account.AccountBase_Level" json:"account_level,omitempty" db:"account_level" valid:"required~账户级别必须存在"`
	// @inject_tag: db:"balance"
	Balance float32 `protobuf:"fixed32,4,opt,name=balance,proto3" json:"balance,omitempty" db:"balance"`
	// @inject_tag: db:"account_status" valid:"required~账户状态必须存在"
	AccountStatus        AccountBase_Status `protobuf:"varint,5,opt,name=account_status,json=accountStatus,proto3,enum=account.AccountBase_Status" json:"account_status,omitempty" db:"account_status" valid:"required~账户状态必须存在"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AccountBase) Reset()         { *m = AccountBase{} }
func (m *AccountBase) String() string { return proto.CompactTextString(m) }
func (*AccountBase) ProtoMessage()    {}
func (*AccountBase) Descriptor() ([]byte, []int) {
	return fileDescriptor_89ed2ef530ab98be, []int{0}
}

func (m *AccountBase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountBase.Unmarshal(m, b)
}
func (m *AccountBase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountBase.Marshal(b, m, deterministic)
}
func (m *AccountBase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountBase.Merge(m, src)
}
func (m *AccountBase) XXX_Size() int {
	return xxx_messageInfo_AccountBase.Size(m)
}
func (m *AccountBase) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountBase.DiscardUnknown(m)
}

var xxx_messageInfo_AccountBase proto.InternalMessageInfo

func (m *AccountBase) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AccountBase) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *AccountBase) GetAccountLevel() AccountBase_Level {
	if m != nil {
		return m.AccountLevel
	}
	return AccountBase_LevelDefault
}

func (m *AccountBase) GetBalance() float32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *AccountBase) GetAccountStatus() AccountBase_Status {
	if m != nil {
		return m.AccountStatus
	}
	return AccountBase_StatusDefault
}

// 账户Id
type AccountId struct {
	// 账户id
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountId) Reset()         { *m = AccountId{} }
func (m *AccountId) String() string { return proto.CompactTextString(m) }
func (*AccountId) ProtoMessage()    {}
func (*AccountId) Descriptor() ([]byte, []int) {
	return fileDescriptor_89ed2ef530ab98be, []int{1}
}

func (m *AccountId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountId.Unmarshal(m, b)
}
func (m *AccountId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountId.Marshal(b, m, deterministic)
}
func (m *AccountId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountId.Merge(m, src)
}
func (m *AccountId) XXX_Size() int {
	return xxx_messageInfo_AccountId.Size(m)
}
func (m *AccountId) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountId.DiscardUnknown(m)
}

var xxx_messageInfo_AccountId proto.InternalMessageInfo

func (m *AccountId) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// 账户更新请求
type AccountUpdateReq struct {
	// 账户id
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 账户余额
	Balance              float32  `protobuf:"fixed32,2,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountUpdateReq) Reset()         { *m = AccountUpdateReq{} }
func (m *AccountUpdateReq) String() string { return proto.CompactTextString(m) }
func (*AccountUpdateReq) ProtoMessage()    {}
func (*AccountUpdateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_89ed2ef530ab98be, []int{2}
}

func (m *AccountUpdateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountUpdateReq.Unmarshal(m, b)
}
func (m *AccountUpdateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountUpdateReq.Marshal(b, m, deterministic)
}
func (m *AccountUpdateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountUpdateReq.Merge(m, src)
}
func (m *AccountUpdateReq) XXX_Size() int {
	return xxx_messageInfo_AccountUpdateReq.Size(m)
}
func (m *AccountUpdateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountUpdateReq.DiscardUnknown(m)
}

var xxx_messageInfo_AccountUpdateReq proto.InternalMessageInfo

func (m *AccountUpdateReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AccountUpdateReq) GetBalance() float32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

// 账户全部
type AccountAll struct {
	// 账户信息
	All []*AccountBase `protobuf:"bytes,1,rep,name=all,proto3" json:"all,omitempty"`
	// 页数
	Page                 *Page    `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountAll) Reset()         { *m = AccountAll{} }
func (m *AccountAll) String() string { return proto.CompactTextString(m) }
func (*AccountAll) ProtoMessage()    {}
func (*AccountAll) Descriptor() ([]byte, []int) {
	return fileDescriptor_89ed2ef530ab98be, []int{3}
}

func (m *AccountAll) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountAll.Unmarshal(m, b)
}
func (m *AccountAll) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountAll.Marshal(b, m, deterministic)
}
func (m *AccountAll) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountAll.Merge(m, src)
}
func (m *AccountAll) XXX_Size() int {
	return xxx_messageInfo_AccountAll.Size(m)
}
func (m *AccountAll) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountAll.DiscardUnknown(m)
}

var xxx_messageInfo_AccountAll proto.InternalMessageInfo

func (m *AccountAll) GetAll() []*AccountBase {
	if m != nil {
		return m.All
	}
	return nil
}

func (m *AccountAll) GetPage() *Page {
	if m != nil {
		return m.Page
	}
	return nil
}

// 空消息
type Null struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Null) Reset()         { *m = Null{} }
func (m *Null) String() string { return proto.CompactTextString(m) }
func (*Null) ProtoMessage()    {}
func (*Null) Descriptor() ([]byte, []int) {
	return fileDescriptor_89ed2ef530ab98be, []int{4}
}

func (m *Null) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Null.Unmarshal(m, b)
}
func (m *Null) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Null.Marshal(b, m, deterministic)
}
func (m *Null) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Null.Merge(m, src)
}
func (m *Null) XXX_Size() int {
	return xxx_messageInfo_Null.Size(m)
}
func (m *Null) XXX_DiscardUnknown() {
	xxx_messageInfo_Null.DiscardUnknown(m)
}

var xxx_messageInfo_Null proto.InternalMessageInfo

// 分页
type Page struct {
	// 页
	PageIndex int64 `protobuf:"varint,1,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	// 每页大小
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// 总页数
	PageTotal int64 `protobuf:"varint,3,opt,name=page_total,json=pageTotal,proto3" json:"page_total,omitempty"`
	// 条数
	Count int64 `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	// 总条数
	Total                int64    `protobuf:"varint,5,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Page) Reset()         { *m = Page{} }
func (m *Page) String() string { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()    {}
func (*Page) Descriptor() ([]byte, []int) {
	return fileDescriptor_89ed2ef530ab98be, []int{5}
}

func (m *Page) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Page.Unmarshal(m, b)
}
func (m *Page) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Page.Marshal(b, m, deterministic)
}
func (m *Page) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Page.Merge(m, src)
}
func (m *Page) XXX_Size() int {
	return xxx_messageInfo_Page.Size(m)
}
func (m *Page) XXX_DiscardUnknown() {
	xxx_messageInfo_Page.DiscardUnknown(m)
}

var xxx_messageInfo_Page proto.InternalMessageInfo

func (m *Page) GetPageIndex() int64 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

func (m *Page) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *Page) GetPageTotal() int64 {
	if m != nil {
		return m.PageTotal
	}
	return 0
}

func (m *Page) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Page) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterEnum("account.AccountBase_Level", AccountBase_Level_name, AccountBase_Level_value)
	proto.RegisterEnum("account.AccountBase_Status", AccountBase_Status_name, AccountBase_Status_value)
	proto.RegisterType((*AccountBase)(nil), "account.AccountBase")
	proto.RegisterType((*AccountId)(nil), "account.AccountId")
	proto.RegisterType((*AccountUpdateReq)(nil), "account.AccountUpdateReq")
	proto.RegisterType((*AccountAll)(nil), "account.AccountAll")
	proto.RegisterType((*Null)(nil), "account.Null")
	proto.RegisterType((*Page)(nil), "account.Page")
}

func init() { proto.RegisterFile("srv/account/proto/account.proto", fileDescriptor_89ed2ef530ab98be) }

var fileDescriptor_89ed2ef530ab98be = []byte{
	// 830 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcf, 0x8f, 0xdb, 0x44,
	0x14, 0xae, 0xed, 0xdd, 0x6c, 0xf6, 0xe5, 0xc7, 0xba, 0xd3, 0x0a, 0x19, 0x2f, 0xab, 0x0d, 0x96,
	0xa8, 0xa2, 0xa8, 0xb1, 0xb3, 0x29, 0xd2, 0x2a, 0xb9, 0xb0, 0x8e, 0x28, 0x52, 0x54, 0xda, 0x2d,
	0xde, 0xae, 0x40, 0x08, 0x29, 0x9a, 0xc4, 0x53, 0xc7, 0xc8, 0xb1, 0x8d, 0x3d, 0xde, 0x36, 0x8b,
	0x90, 0xf8, 0x71, 0xa9, 0x90, 0x90, 0x50, 0xe8, 0xa5, 0x87, 0x1e, 0xe0, 0x00, 0x48, 0x48, 0x1c,
	0x10, 0x02, 0x09, 0x89, 0xbf, 0x81, 0x43, 0xff, 0x85, 0x92, 0xd2, 0xff, 0x02, 0x79, 0xc6, 0x49,
	0x36, 0xdd, 0x20, 0x4e, 0x9e, 0xef, 0x7b, 0xdf, 0x7c, 0x6f, 0xde, 0x7b, 0x33, 0x86, 0xdd, 0x38,
	0x3a, 0x31, 0xf0, 0x60, 0x10, 0x24, 0x3e, 0x35, 0xc2, 0x28, 0xa0, 0xc1, 0x0c, 0xe9, 0x0c, 0xa1,
	0x8d, 0x0c, 0xaa, 0xaf, 0x38, 0x41, 0xe0, 0x78, 0xc4, 0xc0, 0xa1, 0x6b, 0x60, 0xdf, 0x0f, 0x28,
	0xa6, 0x6e, 0xe0, 0xc7, 0x5c, 0xa6, 0x6e, 0x67, 0x51, 0x86, 0xfa, 0xc9, 0x5d, 0x83, 0x8c, 0x42,
	0x3a, 0xce, 0x82, 0x57, 0xd9, 0x67, 0x50, 0x77, 0x88, 0x5f, 0x8f, 0xef, 0x61, 0xc7, 0x21, 0x91,
	0x11, 0x84, 0x6c, 0xfb, 0x79, 0x2b, 0xed, 0x67, 0x09, 0x0a, 0x26, 0x4f, 0xda, 0xc1, 0x31, 0x41,
	0x3b, 0x20, 0xba, 0xb6, 0x22, 0x54, 0x84, 0xaa, 0xd4, 0x29, 0x4d, 0x4c, 0xa8, 0xe5, 0xff, 0xf9,
	0xf5, 0xfb, 0xe9, 0x83, 0xdf, 0x5d, 0xdb, 0x12, 0x5d, 0x1b, 0x5d, 0x81, 0x8d, 0x24, 0x26, 0x51,
	0xcf, 0xb5, 0x15, 0xf1, 0x8c, 0xe6, 0xd9, 0xc3, 0x1f, 0xb9, 0x26, 0x97, 0x46, 0xbb, 0x36, 0xb2,
	0xa0, 0x94, 0x95, 0xd2, 0xf3, 0xc8, 0x09, 0xf1, 0x14, 0xa9, 0x22, 0x54, 0xcb, 0x4d, 0x55, 0x9f,
	0xd5, 0x7b, 0x26, 0xa7, 0xfe, 0x76, 0xaa, 0xe8, 0x5c, 0x9c, 0x98, 0xe5, 0x5a, 0x91, 0x67, 0x7b,
	0xf6, 0xc7, 0x0f, 0x7f, 0x3f, 0xf8, 0xc9, 0x2a, 0x66, 0x6a, 0x26, 0x40, 0xaf, 0xc1, 0x46, 0x1f,
	0x7b, 0xd8, 0x1f, 0x10, 0x65, 0xad, 0x22, 0x54, 0xc5, 0x4e, 0x61, 0x62, 0xe6, 0x6b, 0xb9, 0xa7,
	0x7f, 0x3e, 0x7a, 0xfe, 0xdd, 0x63, 0x6b, 0x16, 0x43, 0xc7, 0x50, 0x9e, 0xa5, 0x8e, 0x29, 0xa6,
	0x49, 0xac, 0xac, 0xb3, 0xdc, 0xdb, 0x2b, 0x73, 0x1f, 0x31, 0xc9, 0x72, 0xf2, 0x2f, 0x7f, 0x9b,
	0x7e, 0xfa, 0x99, 0x35, 0x2b, 0x80, 0x2b, 0xb4, 0xeb, 0xb0, 0xce, 0x8f, 0x21, 0x43, 0x91, 0x2d,
	0xde, 0x24, 0x77, 0x71, 0xe2, 0x51, 0xf9, 0x02, 0x2a, 0x42, 0x9e, 0x31, 0x87, 0x3e, 0x91, 0x85,
	0x39, 0xba, 0x73, 0x2f, 0x90, 0x45, 0x54, 0x06, 0xe0, 0x68, 0x18, 0x11, 0x22, 0x4b, 0xda, 0x1b,
	0x90, 0xe3, 0x86, 0xe8, 0x22, 0x94, 0xf8, 0x6a, 0x61, 0x24, 0x43, 0x91, 0x53, 0xb7, 0x82, 0x68,
	0x84, 0x3d, 0x59, 0x58, 0x30, 0x6f, 0x45, 0xc1, 0x29, 0xf1, 0x65, 0x51, 0xbb, 0x0a, 0x9b, 0xd9,
	0xf9, 0xbb, 0x36, 0xda, 0x3d, 0x33, 0xad, 0xad, 0x89, 0x59, 0x5c, 0x4c, 0xeb, 0x40, 0x48, 0xe7,
	0xa5, 0xdd, 0x04, 0x39, 0x53, 0x1f, 0x87, 0x36, 0xa6, 0xc4, 0x22, 0x1f, 0xfd, 0xef, 0x26, 0xa4,
	0x2c, 0x1a, 0x9d, 0x0e, 0x59, 0x9c, 0xf7, 0x56, 0x7b, 0x17, 0x20, 0xb3, 0x33, 0x3d, 0x0f, 0x5d,
	0x01, 0x09, 0x7b, 0x9e, 0x22, 0x54, 0xa4, 0x6a, 0xa1, 0x79, 0x79, 0x55, 0x7b, 0xad, 0x54, 0x80,
	0x5e, 0x85, 0xb5, 0x10, 0x3b, 0xdc, 0xac, 0xd0, 0x2c, 0xcd, 0x85, 0xb7, 0xb1, 0x43, 0x2c, 0x16,
	0xd2, 0x72, 0xb0, 0x76, 0x2b, 0xf1, 0x3c, 0xed, 0x2b, 0x01, 0xd6, 0x52, 0x1a, 0xed, 0x00, 0xa4,
	0x81, 0x9e, 0xeb, 0xdb, 0xe4, 0x3e, 0x3f, 0xac, 0xb5, 0x99, 0x32, 0xdd, 0x94, 0x40, 0xdb, 0xc0,
	0x40, 0x2f, 0x76, 0x4f, 0xb9, 0xaf, 0x64, 0xe5, 0x53, 0xe2, 0xc8, 0x3d, 0x5d, 0xec, 0xa5, 0x01,
	0xc5, 0xfc, 0xe6, 0x65, 0x7b, 0xef, 0xa4, 0x04, 0xba, 0x0c, 0xeb, 0x2c, 0x3f, 0xbb, 0x45, 0x92,
	0xc5, 0x41, 0xca, 0x72, 0xfd, 0x3a, 0x67, 0x19, 0x68, 0x3e, 0x14, 0xa1, 0x9c, 0xd5, 0x73, 0x44,
	0xa2, 0x13, 0x77, 0x40, 0xd0, 0xcd, 0x45, 0x0f, 0x6c, 0x1b, 0xad, 0x2c, 0x5b, 0x5d, 0xc9, 0x6a,
	0x97, 0x3e, 0x7f, 0xf2, 0xf4, 0x1b, 0xb1, 0xa4, 0xe5, 0x67, 0xaf, 0xbe, 0x2d, 0xd4, 0xd0, 0x07,
	0x50, 0x5a, 0x9a, 0x10, 0x7a, 0xf9, 0xc5, 0xbd, 0xf3, 0xc9, 0xa9, 0x2f, 0xe9, 0xfc, 0xe1, 0xeb,
	0xb3, 0x87, 0xaf, 0x5f, 0x4f, 0x1f, 0xbe, 0xa6, 0x30, 0x63, 0xa4, 0x96, 0xe6, 0x3f, 0x97, 0x8f,
	0x5d, 0xfb, 0x93, 0xd4, 0xfd, 0x18, 0xb6, 0x32, 0x97, 0x77, 0x12, 0x12, 0x8d, 0x0f, 0x7d, 0x82,
	0xd0, 0x8b, 0xfe, 0x5d, 0xfb, 0x3f, 0xce, 0x9b, 0xd9, 0xa2, 0x65, 0xdb, 0xbe, 0x50, 0xeb, 0x3c,
	0x11, 0x26, 0xe6, 0x5f, 0x02, 0x3a, 0x84, 0x1c, 0xbf, 0x3a, 0xda, 0x01, 0xc0, 0x7d, 0x17, 0x07,
	0x23, 0xe2, 0x3b, 0xfb, 0x2d, 0xb4, 0x33, 0xa4, 0x34, 0x8c, 0xdb, 0x86, 0xe1, 0xb8, 0x74, 0x98,
	0xf4, 0xf5, 0x41, 0x30, 0x32, 0x16, 0x61, 0xf5, 0x52, 0xfa, 0xe5, 0xeb, 0x83, 0x31, 0xc1, 0x43,
	0xdd, 0x27, 0xb4, 0x29, 0xed, 0xe9, 0x0d, 0xb5, 0xbc, 0xd7, 0xdc, 0xd7, 0x1b, 0x7a, 0x43, 0xdf,
	0x6b, 0xb7, 0x5a, 0xad, 0xfd, 0x9a, 0x20, 0x34, 0x65, 0x1c, 0x86, 0x9e, 0x3b, 0x60, 0x3f, 0x2d,
	0xe3, 0xc3, 0x38, 0xf0, 0xdb, 0xe7, 0x18, 0x0b, 0x81, 0xf4, 0x7a, 0xe3, 0x1a, 0x2a, 0xc0, 0xe6,
	0xf4, 0x97, 0x6f, 0xa7, 0x8f, 0xbf, 0x78, 0xfe, 0xe8, 0xeb, 0xf7, 0x77, 0x61, 0x07, 0xc0, 0x0c,
	0xdd, 0x1b, 0x64, 0x6c, 0x26, 0x74, 0x88, 0xb6, 0xf2, 0xa2, 0xba, 0xf9, 0x5e, 0xdd, 0xbc, 0xdd,
	0xad, 0xdf, 0x20, 0xe3, 0x8a, 0xd8, 0x97, 0xa1, 0xbc, 0x24, 0xb8, 0xd0, 0xcf, 0xb1, 0xb6, 0x5e,
	0xfb, 0x37, 0x00, 0x00, 0xff, 0xff, 0x41, 0xcf, 0xed, 0x03, 0xa7, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountServiceClient interface {
	// 账户添加
	AccountAdd(ctx context.Context, in *AccountBase, opts ...grpc.CallOption) (*AccountBase, error)
	// 账户更新
	AccountUpdate(ctx context.Context, in *AccountUpdateReq, opts ...grpc.CallOption) (*empty.Empty, error)
	// 账户查询一个
	AccountQueryOne(ctx context.Context, in *AccountId, opts ...grpc.CallOption) (*AccountBase, error)
}

type accountServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountServiceClient(cc *grpc.ClientConn) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) AccountAdd(ctx context.Context, in *AccountBase, opts ...grpc.CallOption) (*AccountBase, error) {
	out := new(AccountBase)
	err := c.cc.Invoke(ctx, "/account.AccountService/AccountAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) AccountUpdate(ctx context.Context, in *AccountUpdateReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/account.AccountService/AccountUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) AccountQueryOne(ctx context.Context, in *AccountId, opts ...grpc.CallOption) (*AccountBase, error) {
	out := new(AccountBase)
	err := c.cc.Invoke(ctx, "/account.AccountService/AccountQueryOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
type AccountServiceServer interface {
	// 账户添加
	AccountAdd(context.Context, *AccountBase) (*AccountBase, error)
	// 账户更新
	AccountUpdate(context.Context, *AccountUpdateReq) (*empty.Empty, error)
	// 账户查询一个
	AccountQueryOne(context.Context, *AccountId) (*AccountBase, error)
}

// UnimplementedAccountServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (*UnimplementedAccountServiceServer) AccountAdd(ctx context.Context, req *AccountBase) (*AccountBase, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountAdd not implemented")
}
func (*UnimplementedAccountServiceServer) AccountUpdate(ctx context.Context, req *AccountUpdateReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountUpdate not implemented")
}
func (*UnimplementedAccountServiceServer) AccountQueryOne(ctx context.Context, req *AccountId) (*AccountBase, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountQueryOne not implemented")
}

func RegisterAccountServiceServer(s *grpc.Server, srv AccountServiceServer) {
	s.RegisterService(&_AccountService_serviceDesc, srv)
}

func _AccountService_AccountAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountBase)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).AccountAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/AccountAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).AccountAdd(ctx, req.(*AccountBase))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_AccountUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).AccountUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/AccountUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).AccountUpdate(ctx, req.(*AccountUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_AccountQueryOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).AccountQueryOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/AccountQueryOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).AccountQueryOne(ctx, req.(*AccountId))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AccountAdd",
			Handler:    _AccountService_AccountAdd_Handler,
		},
		{
			MethodName: "AccountUpdate",
			Handler:    _AccountService_AccountUpdate_Handler,
		},
		{
			MethodName: "AccountQueryOne",
			Handler:    _AccountService_AccountQueryOne_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "srv/account/proto/account.proto",
}
