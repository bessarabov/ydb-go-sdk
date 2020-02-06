// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ydb_experimental.proto

package Ydb_Experimental

import (
	Ydb "github.com/yandex-cloud/ydb-go-sdk/api/protos/Ydb"
	Ydb_Issue "github.com/yandex-cloud/ydb-go-sdk/api/protos/Ydb_Issue"
	Ydb_Operations "github.com/yandex-cloud/ydb-go-sdk/api/protos/Ydb_Operations"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ExecuteStreamQueryRequest_ProfileMode int32

const (
	ExecuteStreamQueryRequest_PROFILE_MODE_UNSPECIFIED ExecuteStreamQueryRequest_ProfileMode = 0
	ExecuteStreamQueryRequest_NONE                     ExecuteStreamQueryRequest_ProfileMode = 1
	ExecuteStreamQueryRequest_BASIC                    ExecuteStreamQueryRequest_ProfileMode = 2
)

var ExecuteStreamQueryRequest_ProfileMode_name = map[int32]string{
	0: "PROFILE_MODE_UNSPECIFIED",
	1: "NONE",
	2: "BASIC",
}

var ExecuteStreamQueryRequest_ProfileMode_value = map[string]int32{
	"PROFILE_MODE_UNSPECIFIED": 0,
	"NONE":                     1,
	"BASIC":                    2,
}

func (x ExecuteStreamQueryRequest_ProfileMode) String() string {
	return proto.EnumName(ExecuteStreamQueryRequest_ProfileMode_name, int32(x))
}

func (ExecuteStreamQueryRequest_ProfileMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{3, 0}
}

type UploadRowsRequest struct {
	Table                string                          `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Rows                 *Ydb.TypedValue                 `protobuf:"bytes,2,opt,name=rows,proto3" json:"rows,omitempty"`
	OperationParams      *Ydb_Operations.OperationParams `protobuf:"bytes,3,opt,name=operation_params,json=operationParams,proto3" json:"operation_params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *UploadRowsRequest) Reset()         { *m = UploadRowsRequest{} }
func (m *UploadRowsRequest) String() string { return proto.CompactTextString(m) }
func (*UploadRowsRequest) ProtoMessage()    {}
func (*UploadRowsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{0}
}

func (m *UploadRowsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadRowsRequest.Unmarshal(m, b)
}
func (m *UploadRowsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadRowsRequest.Marshal(b, m, deterministic)
}
func (m *UploadRowsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadRowsRequest.Merge(m, src)
}
func (m *UploadRowsRequest) XXX_Size() int {
	return xxx_messageInfo_UploadRowsRequest.Size(m)
}
func (m *UploadRowsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadRowsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadRowsRequest proto.InternalMessageInfo

func (m *UploadRowsRequest) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *UploadRowsRequest) GetRows() *Ydb.TypedValue {
	if m != nil {
		return m.Rows
	}
	return nil
}

func (m *UploadRowsRequest) GetOperationParams() *Ydb_Operations.OperationParams {
	if m != nil {
		return m.OperationParams
	}
	return nil
}

type UploadRowsResponse struct {
	Operation            *Ydb_Operations.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *UploadRowsResponse) Reset()         { *m = UploadRowsResponse{} }
func (m *UploadRowsResponse) String() string { return proto.CompactTextString(m) }
func (*UploadRowsResponse) ProtoMessage()    {}
func (*UploadRowsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{1}
}

func (m *UploadRowsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadRowsResponse.Unmarshal(m, b)
}
func (m *UploadRowsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadRowsResponse.Marshal(b, m, deterministic)
}
func (m *UploadRowsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadRowsResponse.Merge(m, src)
}
func (m *UploadRowsResponse) XXX_Size() int {
	return xxx_messageInfo_UploadRowsResponse.Size(m)
}
func (m *UploadRowsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadRowsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadRowsResponse proto.InternalMessageInfo

func (m *UploadRowsResponse) GetOperation() *Ydb_Operations.Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

type UploadRowsResult struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadRowsResult) Reset()         { *m = UploadRowsResult{} }
func (m *UploadRowsResult) String() string { return proto.CompactTextString(m) }
func (*UploadRowsResult) ProtoMessage()    {}
func (*UploadRowsResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{2}
}

func (m *UploadRowsResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadRowsResult.Unmarshal(m, b)
}
func (m *UploadRowsResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadRowsResult.Marshal(b, m, deterministic)
}
func (m *UploadRowsResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadRowsResult.Merge(m, src)
}
func (m *UploadRowsResult) XXX_Size() int {
	return xxx_messageInfo_UploadRowsResult.Size(m)
}
func (m *UploadRowsResult) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadRowsResult.DiscardUnknown(m)
}

var xxx_messageInfo_UploadRowsResult proto.InternalMessageInfo

type ExecuteStreamQueryRequest struct {
	YqlText              string                                `protobuf:"bytes,1,opt,name=yql_text,json=yqlText,proto3" json:"yql_text,omitempty"`
	Parameters           map[string]*Ydb.TypedValue            `protobuf:"bytes,2,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ProfileMode          ExecuteStreamQueryRequest_ProfileMode `protobuf:"varint,3,opt,name=profile_mode,json=profileMode,proto3,enum=Ydb.Experimental.ExecuteStreamQueryRequest_ProfileMode" json:"profile_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *ExecuteStreamQueryRequest) Reset()         { *m = ExecuteStreamQueryRequest{} }
func (m *ExecuteStreamQueryRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteStreamQueryRequest) ProtoMessage()    {}
func (*ExecuteStreamQueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{3}
}

func (m *ExecuteStreamQueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteStreamQueryRequest.Unmarshal(m, b)
}
func (m *ExecuteStreamQueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteStreamQueryRequest.Marshal(b, m, deterministic)
}
func (m *ExecuteStreamQueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteStreamQueryRequest.Merge(m, src)
}
func (m *ExecuteStreamQueryRequest) XXX_Size() int {
	return xxx_messageInfo_ExecuteStreamQueryRequest.Size(m)
}
func (m *ExecuteStreamQueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteStreamQueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteStreamQueryRequest proto.InternalMessageInfo

func (m *ExecuteStreamQueryRequest) GetYqlText() string {
	if m != nil {
		return m.YqlText
	}
	return ""
}

func (m *ExecuteStreamQueryRequest) GetParameters() map[string]*Ydb.TypedValue {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *ExecuteStreamQueryRequest) GetProfileMode() ExecuteStreamQueryRequest_ProfileMode {
	if m != nil {
		return m.ProfileMode
	}
	return ExecuteStreamQueryRequest_PROFILE_MODE_UNSPECIFIED
}

type ExecuteStreamQueryResponse struct {
	Status               Ydb.StatusIds_StatusCode  `protobuf:"varint,1,opt,name=status,proto3,enum=Ydb.StatusIds_StatusCode" json:"status,omitempty"`
	Issues               []*Ydb_Issue.IssueMessage `protobuf:"bytes,2,rep,name=issues,proto3" json:"issues,omitempty"`
	Result               *ExecuteStreamQueryResult `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ExecuteStreamQueryResponse) Reset()         { *m = ExecuteStreamQueryResponse{} }
func (m *ExecuteStreamQueryResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteStreamQueryResponse) ProtoMessage()    {}
func (*ExecuteStreamQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{4}
}

func (m *ExecuteStreamQueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteStreamQueryResponse.Unmarshal(m, b)
}
func (m *ExecuteStreamQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteStreamQueryResponse.Marshal(b, m, deterministic)
}
func (m *ExecuteStreamQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteStreamQueryResponse.Merge(m, src)
}
func (m *ExecuteStreamQueryResponse) XXX_Size() int {
	return xxx_messageInfo_ExecuteStreamQueryResponse.Size(m)
}
func (m *ExecuteStreamQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteStreamQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteStreamQueryResponse proto.InternalMessageInfo

func (m *ExecuteStreamQueryResponse) GetStatus() Ydb.StatusIds_StatusCode {
	if m != nil {
		return m.Status
	}
	return Ydb.StatusIds_STATUS_CODE_UNSPECIFIED
}

func (m *ExecuteStreamQueryResponse) GetIssues() []*Ydb_Issue.IssueMessage {
	if m != nil {
		return m.Issues
	}
	return nil
}

func (m *ExecuteStreamQueryResponse) GetResult() *ExecuteStreamQueryResult {
	if m != nil {
		return m.Result
	}
	return nil
}

type StreamQueryPing struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamQueryPing) Reset()         { *m = StreamQueryPing{} }
func (m *StreamQueryPing) String() string { return proto.CompactTextString(m) }
func (*StreamQueryPing) ProtoMessage()    {}
func (*StreamQueryPing) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{5}
}

func (m *StreamQueryPing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamQueryPing.Unmarshal(m, b)
}
func (m *StreamQueryPing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamQueryPing.Marshal(b, m, deterministic)
}
func (m *StreamQueryPing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamQueryPing.Merge(m, src)
}
func (m *StreamQueryPing) XXX_Size() int {
	return xxx_messageInfo_StreamQueryPing.Size(m)
}
func (m *StreamQueryPing) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamQueryPing.DiscardUnknown(m)
}

var xxx_messageInfo_StreamQueryPing proto.InternalMessageInfo

type StreamQueryProgress struct {
	// Types that are valid to be assigned to Kind:
	//	*StreamQueryProgress_Ping
	Kind                 isStreamQueryProgress_Kind `protobuf_oneof:"kind"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *StreamQueryProgress) Reset()         { *m = StreamQueryProgress{} }
func (m *StreamQueryProgress) String() string { return proto.CompactTextString(m) }
func (*StreamQueryProgress) ProtoMessage()    {}
func (*StreamQueryProgress) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{6}
}

func (m *StreamQueryProgress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamQueryProgress.Unmarshal(m, b)
}
func (m *StreamQueryProgress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamQueryProgress.Marshal(b, m, deterministic)
}
func (m *StreamQueryProgress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamQueryProgress.Merge(m, src)
}
func (m *StreamQueryProgress) XXX_Size() int {
	return xxx_messageInfo_StreamQueryProgress.Size(m)
}
func (m *StreamQueryProgress) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamQueryProgress.DiscardUnknown(m)
}

var xxx_messageInfo_StreamQueryProgress proto.InternalMessageInfo

type isStreamQueryProgress_Kind interface {
	isStreamQueryProgress_Kind()
}

type StreamQueryProgress_Ping struct {
	Ping *StreamQueryPing `protobuf:"bytes,1,opt,name=ping,proto3,oneof"`
}

func (*StreamQueryProgress_Ping) isStreamQueryProgress_Kind() {}

func (m *StreamQueryProgress) GetKind() isStreamQueryProgress_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *StreamQueryProgress) GetPing() *StreamQueryPing {
	if x, ok := m.GetKind().(*StreamQueryProgress_Ping); ok {
		return x.Ping
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StreamQueryProgress) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StreamQueryProgress_Ping)(nil),
	}
}

type ExecuteStreamQueryResult struct {
	// Types that are valid to be assigned to Result:
	//	*ExecuteStreamQueryResult_ResultSet
	//	*ExecuteStreamQueryResult_Profile
	//	*ExecuteStreamQueryResult_Progress
	Result               isExecuteStreamQueryResult_Result `protobuf_oneof:"result"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *ExecuteStreamQueryResult) Reset()         { *m = ExecuteStreamQueryResult{} }
func (m *ExecuteStreamQueryResult) String() string { return proto.CompactTextString(m) }
func (*ExecuteStreamQueryResult) ProtoMessage()    {}
func (*ExecuteStreamQueryResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{7}
}

func (m *ExecuteStreamQueryResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteStreamQueryResult.Unmarshal(m, b)
}
func (m *ExecuteStreamQueryResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteStreamQueryResult.Marshal(b, m, deterministic)
}
func (m *ExecuteStreamQueryResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteStreamQueryResult.Merge(m, src)
}
func (m *ExecuteStreamQueryResult) XXX_Size() int {
	return xxx_messageInfo_ExecuteStreamQueryResult.Size(m)
}
func (m *ExecuteStreamQueryResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteStreamQueryResult.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteStreamQueryResult proto.InternalMessageInfo

type isExecuteStreamQueryResult_Result interface {
	isExecuteStreamQueryResult_Result()
}

type ExecuteStreamQueryResult_ResultSet struct {
	ResultSet *Ydb.ResultSet `protobuf:"bytes,1,opt,name=result_set,json=resultSet,proto3,oneof"`
}

type ExecuteStreamQueryResult_Profile struct {
	Profile string `protobuf:"bytes,2,opt,name=profile,proto3,oneof"`
}

type ExecuteStreamQueryResult_Progress struct {
	Progress *StreamQueryProgress `protobuf:"bytes,3,opt,name=progress,proto3,oneof"`
}

func (*ExecuteStreamQueryResult_ResultSet) isExecuteStreamQueryResult_Result() {}

func (*ExecuteStreamQueryResult_Profile) isExecuteStreamQueryResult_Result() {}

func (*ExecuteStreamQueryResult_Progress) isExecuteStreamQueryResult_Result() {}

func (m *ExecuteStreamQueryResult) GetResult() isExecuteStreamQueryResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *ExecuteStreamQueryResult) GetResultSet() *Ydb.ResultSet {
	if x, ok := m.GetResult().(*ExecuteStreamQueryResult_ResultSet); ok {
		return x.ResultSet
	}
	return nil
}

func (m *ExecuteStreamQueryResult) GetProfile() string {
	if x, ok := m.GetResult().(*ExecuteStreamQueryResult_Profile); ok {
		return x.Profile
	}
	return ""
}

func (m *ExecuteStreamQueryResult) GetProgress() *StreamQueryProgress {
	if x, ok := m.GetResult().(*ExecuteStreamQueryResult_Progress); ok {
		return x.Progress
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExecuteStreamQueryResult) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExecuteStreamQueryResult_ResultSet)(nil),
		(*ExecuteStreamQueryResult_Profile)(nil),
		(*ExecuteStreamQueryResult_Progress)(nil),
	}
}

func init() {
	proto.RegisterEnum("Ydb.Experimental.ExecuteStreamQueryRequest_ProfileMode", ExecuteStreamQueryRequest_ProfileMode_name, ExecuteStreamQueryRequest_ProfileMode_value)
	proto.RegisterType((*UploadRowsRequest)(nil), "Ydb.Experimental.UploadRowsRequest")
	proto.RegisterType((*UploadRowsResponse)(nil), "Ydb.Experimental.UploadRowsResponse")
	proto.RegisterType((*UploadRowsResult)(nil), "Ydb.Experimental.UploadRowsResult")
	proto.RegisterType((*ExecuteStreamQueryRequest)(nil), "Ydb.Experimental.ExecuteStreamQueryRequest")
	proto.RegisterMapType((map[string]*Ydb.TypedValue)(nil), "Ydb.Experimental.ExecuteStreamQueryRequest.ParametersEntry")
	proto.RegisterType((*ExecuteStreamQueryResponse)(nil), "Ydb.Experimental.ExecuteStreamQueryResponse")
	proto.RegisterType((*StreamQueryPing)(nil), "Ydb.Experimental.StreamQueryPing")
	proto.RegisterType((*StreamQueryProgress)(nil), "Ydb.Experimental.StreamQueryProgress")
	proto.RegisterType((*ExecuteStreamQueryResult)(nil), "Ydb.Experimental.ExecuteStreamQueryResult")
}

func init() { proto.RegisterFile("ydb_experimental.proto", fileDescriptor_ac21a693e2c386a5) }

var fileDescriptor_ac21a693e2c386a5 = []byte{
	// 671 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdf, 0x6e, 0xd3, 0x3e,
	0x14, 0x6e, 0xd6, 0xae, 0x6b, 0x4f, 0x7f, 0x5a, 0x3b, 0xff, 0x10, 0x74, 0x05, 0x89, 0x11, 0x34,
	0xa9, 0x42, 0x28, 0x85, 0x82, 0x34, 0x04, 0x37, 0xd0, 0x2e, 0x53, 0x8b, 0xe8, 0x1f, 0xdc, 0x6d,
	0x12, 0x70, 0x11, 0xa5, 0x8d, 0xa9, 0xa2, 0xa6, 0x71, 0x66, 0x3b, 0xac, 0xb9, 0xe4, 0x35, 0x78,
	0x10, 0x9e, 0x80, 0x87, 0xe2, 0x12, 0xc5, 0x71, 0xbb, 0x6c, 0xd0, 0x0a, 0x6e, 0x22, 0xfb, 0xf8,
	0xfb, 0xbe, 0xf3, 0xf9, 0x9c, 0x13, 0xc3, 0xed, 0xc8, 0x19, 0x5b, 0x64, 0x11, 0x10, 0xe6, 0xce,
	0x89, 0x2f, 0x6c, 0xcf, 0x08, 0x18, 0x15, 0x14, 0x55, 0x3e, 0x38, 0x63, 0xc3, 0x4c, 0xc5, 0x6b,
	0x4f, 0x66, 0xee, 0xcc, 0x9d, 0xb3, 0x46, 0x10, 0x8e, 0x3d, 0x77, 0xd2, 0xb0, 0x03, 0xb7, 0x21,
	0xa1, 0xbc, 0x11, 0x4b, 0xb8, 0x9c, 0x87, 0xc4, 0x9a, 0x13, 0xce, 0xed, 0x29, 0x49, 0x34, 0x6a,
	0x8f, 0x37, 0x32, 0x68, 0x40, 0x98, 0x2d, 0x5c, 0xea, 0x2b, 0x74, 0x63, 0x23, 0x9a, 0x0b, 0x5b,
	0x84, 0xdc, 0x9a, 0x50, 0x87, 0x70, 0x45, 0xa8, 0x6f, 0x24, 0x7c, 0xb1, 0xbd, 0x50, 0x19, 0xd1,
	0xbf, 0x69, 0xb0, 0x77, 0x16, 0x78, 0xd4, 0x76, 0x30, 0xbd, 0xe4, 0x98, 0x5c, 0x84, 0x84, 0x0b,
	0x74, 0x0b, 0xb6, 0x85, 0x3d, 0xf6, 0x48, 0x55, 0x3b, 0xd0, 0xea, 0x45, 0x9c, 0x6c, 0xd0, 0x43,
	0xc8, 0x31, 0x7a, 0xc9, 0xab, 0x5b, 0x07, 0x5a, 0xbd, 0xd4, 0x2c, 0x1b, 0x71, 0x1d, 0x4e, 0xa3,
	0x80, 0x38, 0xe7, 0xb1, 0x20, 0x96, 0x87, 0xe8, 0x2d, 0x54, 0x56, 0xf6, 0xad, 0xc0, 0x66, 0xf6,
	0x9c, 0x57, 0xb3, 0x92, 0x70, 0x5f, 0x12, 0x06, 0xcb, 0x43, 0x7e, 0xb5, 0x1c, 0x4a, 0x18, 0x2e,
	0xd3, 0xeb, 0x01, 0xbd, 0x07, 0x28, 0xed, 0x8d, 0x07, 0xd4, 0xe7, 0x04, 0x1d, 0x41, 0x71, 0x05,
	0x94, 0x06, 0x4b, 0xcd, 0xfd, 0xb5, 0xd2, 0xf8, 0x0a, 0xab, 0x23, 0xa8, 0x5c, 0x93, 0x0b, 0x3d,
	0xa1, 0x7f, 0xcd, 0xc2, 0xbe, 0xb9, 0x20, 0x93, 0x50, 0x90, 0x91, 0x60, 0xc4, 0x9e, 0xbf, 0x0f,
	0x09, 0x8b, 0x96, 0x75, 0xd8, 0x87, 0x42, 0x74, 0xe1, 0x59, 0x82, 0x2c, 0x84, 0x2a, 0xc5, 0x4e,
	0x74, 0xe1, 0x9d, 0x92, 0x85, 0x40, 0x9f, 0x00, 0xe4, 0xed, 0x88, 0x20, 0x2c, 0x2e, 0x49, 0xb6,
	0x5e, 0x6a, 0xbe, 0x32, 0x6e, 0x8e, 0x86, 0xb1, 0x56, 0xdb, 0x18, 0xae, 0xd8, 0xa6, 0x2f, 0x58,
	0x84, 0x53, 0x72, 0xe8, 0x23, 0xfc, 0x17, 0x30, 0xfa, 0xd9, 0xf5, 0x88, 0x35, 0xa7, 0x0e, 0x91,
	0x05, 0xdc, 0x6d, 0x1e, 0xfd, 0x93, 0x7c, 0xc2, 0xef, 0x51, 0x87, 0xe0, 0x52, 0x70, 0xb5, 0xa9,
	0xf5, 0xa1, 0x7c, 0x23, 0x35, 0xaa, 0x40, 0x76, 0x46, 0x22, 0x75, 0xc3, 0x78, 0x89, 0x0e, 0x61,
	0x5b, 0x4e, 0xc9, 0xba, 0x5e, 0x27, 0xa7, 0x2f, 0xb7, 0x5e, 0x68, 0xfa, 0x6b, 0x28, 0xa5, 0x72,
	0xa1, 0x7b, 0x50, 0x1d, 0xe2, 0xc1, 0x49, 0xf7, 0x9d, 0x69, 0xf5, 0x06, 0xc7, 0xa6, 0x75, 0xd6,
	0x1f, 0x0d, 0xcd, 0x76, 0xf7, 0xa4, 0x6b, 0x1e, 0x57, 0x32, 0xa8, 0x00, 0xb9, 0xfe, 0xa0, 0x6f,
	0x56, 0x34, 0x54, 0x84, 0xed, 0xd6, 0x9b, 0x51, 0xb7, 0x5d, 0xd9, 0xd2, 0x7f, 0x68, 0x50, 0xfb,
	0xd3, 0x45, 0x54, 0xbf, 0x9f, 0x42, 0x3e, 0x19, 0x71, 0x69, 0x70, 0x57, 0x35, 0x7b, 0x24, 0x43,
	0x5d, 0x87, 0xab, 0x55, 0x3b, 0xbe, 0xa8, 0x02, 0xa2, 0x06, 0xe4, 0xe5, 0x5f, 0xb7, 0x6c, 0xcc,
	0x1d, 0x49, 0xe9, 0xc6, 0xa1, 0xe4, 0xdb, 0x4b, 0xfe, 0x46, 0xac, 0x60, 0xa8, 0x05, 0x79, 0x26,
	0x07, 0x42, 0xcd, 0xea, 0xa3, 0xbf, 0x2b, 0x75, 0xcc, 0xc0, 0x8a, 0xa9, 0xef, 0x41, 0x39, 0x75,
	0x38, 0x74, 0xfd, 0xa9, 0x7e, 0x0e, 0xff, 0xa7, 0x43, 0x8c, 0x4e, 0x19, 0xe1, 0x1c, 0x1d, 0x41,
	0x2e, 0x70, 0xfd, 0xa9, 0x1a, 0xde, 0x07, 0xbf, 0xe7, 0xba, 0xa1, 0xd3, 0xc9, 0x60, 0x49, 0x68,
	0xe5, 0x21, 0x37, 0x73, 0x7d, 0x47, 0xff, 0xae, 0x41, 0x75, 0x9d, 0x1f, 0xd4, 0x00, 0x48, 0x1c,
	0x59, 0x9c, 0x08, 0x95, 0x63, 0x57, 0xe6, 0x48, 0x00, 0x23, 0x22, 0x3a, 0x19, 0x5c, 0x64, 0xcb,
	0x0d, 0xaa, 0xc1, 0x8e, 0x1a, 0x10, 0xd9, 0xee, 0x62, 0x27, 0x83, 0x97, 0x01, 0xd4, 0x86, 0x42,
	0xa0, 0x6c, 0xab, 0xd2, 0x1c, 0x6e, 0xb6, 0xab, 0xc0, 0x9d, 0x0c, 0x5e, 0x11, 0x5b, 0x85, 0x65,
	0x75, 0x5b, 0xcf, 0xe1, 0xee, 0x84, 0xce, 0x8d, 0xc8, 0xf6, 0x1d, 0xb2, 0x30, 0x22, 0x67, 0x6c,
	0xa4, 0x1f, 0xd8, 0x16, 0x4a, 0xcb, 0x0e, 0xe5, 0x8b, 0xf5, 0x53, 0xd3, 0xc6, 0x79, 0xf9, 0x56,
	0x3d, 0xfb, 0x15, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x73, 0x18, 0x7c, 0x92, 0x05, 0x00, 0x00,
}

const ()

// SetOperationParams implements ydb generic interface for setting
// operation parameters inside driver implementation.
func (m *UploadRowsRequest) SetOperationParams(v *Ydb_Operations.OperationParams) {
	m.OperationParams = v
}
