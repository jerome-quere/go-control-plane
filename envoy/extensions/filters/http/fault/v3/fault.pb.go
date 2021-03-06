// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/fault/v3/fault.proto

package envoy_extensions_filters_http_fault_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v32 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	v31 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/common/fault/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type FaultAbort struct {
	Percentage *v3.FractionalPercent `protobuf:"bytes,3,opt,name=percentage,proto3" json:"percentage,omitempty"`
	// Types that are valid to be assigned to ErrorType:
	//	*FaultAbort_HttpStatus
	//	*FaultAbort_HeaderAbort_
	ErrorType            isFaultAbort_ErrorType `protobuf_oneof:"error_type"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *FaultAbort) Reset()         { *m = FaultAbort{} }
func (m *FaultAbort) String() string { return proto.CompactTextString(m) }
func (*FaultAbort) ProtoMessage()    {}
func (*FaultAbort) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c8de848a52219f2, []int{0}
}

func (m *FaultAbort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultAbort.Unmarshal(m, b)
}
func (m *FaultAbort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultAbort.Marshal(b, m, deterministic)
}
func (m *FaultAbort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultAbort.Merge(m, src)
}
func (m *FaultAbort) XXX_Size() int {
	return xxx_messageInfo_FaultAbort.Size(m)
}
func (m *FaultAbort) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultAbort.DiscardUnknown(m)
}

var xxx_messageInfo_FaultAbort proto.InternalMessageInfo

func (m *FaultAbort) GetPercentage() *v3.FractionalPercent {
	if m != nil {
		return m.Percentage
	}
	return nil
}

type isFaultAbort_ErrorType interface {
	isFaultAbort_ErrorType()
}

type FaultAbort_HttpStatus struct {
	HttpStatus uint32 `protobuf:"varint,2,opt,name=http_status,json=httpStatus,proto3,oneof"`
}

type FaultAbort_HeaderAbort_ struct {
	HeaderAbort *FaultAbort_HeaderAbort `protobuf:"bytes,4,opt,name=header_abort,json=headerAbort,proto3,oneof"`
}

func (*FaultAbort_HttpStatus) isFaultAbort_ErrorType() {}

func (*FaultAbort_HeaderAbort_) isFaultAbort_ErrorType() {}

func (m *FaultAbort) GetErrorType() isFaultAbort_ErrorType {
	if m != nil {
		return m.ErrorType
	}
	return nil
}

func (m *FaultAbort) GetHttpStatus() uint32 {
	if x, ok := m.GetErrorType().(*FaultAbort_HttpStatus); ok {
		return x.HttpStatus
	}
	return 0
}

func (m *FaultAbort) GetHeaderAbort() *FaultAbort_HeaderAbort {
	if x, ok := m.GetErrorType().(*FaultAbort_HeaderAbort_); ok {
		return x.HeaderAbort
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FaultAbort) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FaultAbort_HttpStatus)(nil),
		(*FaultAbort_HeaderAbort_)(nil),
	}
}

type FaultAbort_HeaderAbort struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FaultAbort_HeaderAbort) Reset()         { *m = FaultAbort_HeaderAbort{} }
func (m *FaultAbort_HeaderAbort) String() string { return proto.CompactTextString(m) }
func (*FaultAbort_HeaderAbort) ProtoMessage()    {}
func (*FaultAbort_HeaderAbort) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c8de848a52219f2, []int{0, 0}
}

func (m *FaultAbort_HeaderAbort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultAbort_HeaderAbort.Unmarshal(m, b)
}
func (m *FaultAbort_HeaderAbort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultAbort_HeaderAbort.Marshal(b, m, deterministic)
}
func (m *FaultAbort_HeaderAbort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultAbort_HeaderAbort.Merge(m, src)
}
func (m *FaultAbort_HeaderAbort) XXX_Size() int {
	return xxx_messageInfo_FaultAbort_HeaderAbort.Size(m)
}
func (m *FaultAbort_HeaderAbort) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultAbort_HeaderAbort.DiscardUnknown(m)
}

var xxx_messageInfo_FaultAbort_HeaderAbort proto.InternalMessageInfo

type HTTPFault struct {
	Delay                           *v31.FaultDelay       `protobuf:"bytes,1,opt,name=delay,proto3" json:"delay,omitempty"`
	Abort                           *FaultAbort           `protobuf:"bytes,2,opt,name=abort,proto3" json:"abort,omitempty"`
	UpstreamCluster                 string                `protobuf:"bytes,3,opt,name=upstream_cluster,json=upstreamCluster,proto3" json:"upstream_cluster,omitempty"`
	Headers                         []*v32.HeaderMatcher  `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty"`
	DownstreamNodes                 []string              `protobuf:"bytes,5,rep,name=downstream_nodes,json=downstreamNodes,proto3" json:"downstream_nodes,omitempty"`
	MaxActiveFaults                 *wrappers.UInt32Value `protobuf:"bytes,6,opt,name=max_active_faults,json=maxActiveFaults,proto3" json:"max_active_faults,omitempty"`
	ResponseRateLimit               *v31.FaultRateLimit   `protobuf:"bytes,7,opt,name=response_rate_limit,json=responseRateLimit,proto3" json:"response_rate_limit,omitempty"`
	DelayPercentRuntime             string                `protobuf:"bytes,8,opt,name=delay_percent_runtime,json=delayPercentRuntime,proto3" json:"delay_percent_runtime,omitempty"`
	AbortPercentRuntime             string                `protobuf:"bytes,9,opt,name=abort_percent_runtime,json=abortPercentRuntime,proto3" json:"abort_percent_runtime,omitempty"`
	DelayDurationRuntime            string                `protobuf:"bytes,10,opt,name=delay_duration_runtime,json=delayDurationRuntime,proto3" json:"delay_duration_runtime,omitempty"`
	AbortHttpStatusRuntime          string                `protobuf:"bytes,11,opt,name=abort_http_status_runtime,json=abortHttpStatusRuntime,proto3" json:"abort_http_status_runtime,omitempty"`
	MaxActiveFaultsRuntime          string                `protobuf:"bytes,12,opt,name=max_active_faults_runtime,json=maxActiveFaultsRuntime,proto3" json:"max_active_faults_runtime,omitempty"`
	ResponseRateLimitPercentRuntime string                `protobuf:"bytes,13,opt,name=response_rate_limit_percent_runtime,json=responseRateLimitPercentRuntime,proto3" json:"response_rate_limit_percent_runtime,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}              `json:"-"`
	XXX_unrecognized                []byte                `json:"-"`
	XXX_sizecache                   int32                 `json:"-"`
}

func (m *HTTPFault) Reset()         { *m = HTTPFault{} }
func (m *HTTPFault) String() string { return proto.CompactTextString(m) }
func (*HTTPFault) ProtoMessage()    {}
func (*HTTPFault) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c8de848a52219f2, []int{1}
}

func (m *HTTPFault) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTTPFault.Unmarshal(m, b)
}
func (m *HTTPFault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTTPFault.Marshal(b, m, deterministic)
}
func (m *HTTPFault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPFault.Merge(m, src)
}
func (m *HTTPFault) XXX_Size() int {
	return xxx_messageInfo_HTTPFault.Size(m)
}
func (m *HTTPFault) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPFault.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPFault proto.InternalMessageInfo

func (m *HTTPFault) GetDelay() *v31.FaultDelay {
	if m != nil {
		return m.Delay
	}
	return nil
}

func (m *HTTPFault) GetAbort() *FaultAbort {
	if m != nil {
		return m.Abort
	}
	return nil
}

func (m *HTTPFault) GetUpstreamCluster() string {
	if m != nil {
		return m.UpstreamCluster
	}
	return ""
}

func (m *HTTPFault) GetHeaders() []*v32.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HTTPFault) GetDownstreamNodes() []string {
	if m != nil {
		return m.DownstreamNodes
	}
	return nil
}

func (m *HTTPFault) GetMaxActiveFaults() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxActiveFaults
	}
	return nil
}

func (m *HTTPFault) GetResponseRateLimit() *v31.FaultRateLimit {
	if m != nil {
		return m.ResponseRateLimit
	}
	return nil
}

func (m *HTTPFault) GetDelayPercentRuntime() string {
	if m != nil {
		return m.DelayPercentRuntime
	}
	return ""
}

func (m *HTTPFault) GetAbortPercentRuntime() string {
	if m != nil {
		return m.AbortPercentRuntime
	}
	return ""
}

func (m *HTTPFault) GetDelayDurationRuntime() string {
	if m != nil {
		return m.DelayDurationRuntime
	}
	return ""
}

func (m *HTTPFault) GetAbortHttpStatusRuntime() string {
	if m != nil {
		return m.AbortHttpStatusRuntime
	}
	return ""
}

func (m *HTTPFault) GetMaxActiveFaultsRuntime() string {
	if m != nil {
		return m.MaxActiveFaultsRuntime
	}
	return ""
}

func (m *HTTPFault) GetResponseRateLimitPercentRuntime() string {
	if m != nil {
		return m.ResponseRateLimitPercentRuntime
	}
	return ""
}

func init() {
	proto.RegisterType((*FaultAbort)(nil), "envoy.extensions.filters.http.fault.v3.FaultAbort")
	proto.RegisterType((*FaultAbort_HeaderAbort)(nil), "envoy.extensions.filters.http.fault.v3.FaultAbort.HeaderAbort")
	proto.RegisterType((*HTTPFault)(nil), "envoy.extensions.filters.http.fault.v3.HTTPFault")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/fault/v3/fault.proto", fileDescriptor_1c8de848a52219f2)
}

var fileDescriptor_1c8de848a52219f2 = []byte{
	// 760 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcf, 0x4f, 0xe3, 0x46,
	0x14, 0xc7, 0xc9, 0x0f, 0x02, 0x4c, 0x40, 0x04, 0xd3, 0x52, 0x37, 0x6d, 0x69, 0x4a, 0xab, 0x2a,
	0x6d, 0xe9, 0x58, 0x72, 0x72, 0xa0, 0x39, 0xd0, 0x92, 0x22, 0x94, 0x45, 0x2c, 0x8a, 0xbc, 0xec,
	0x5e, 0xad, 0xc1, 0x99, 0x24, 0x96, 0xec, 0x19, 0x6b, 0x3c, 0x0e, 0xc9, 0x6d, 0xb9, 0xed, 0xdf,
	0xb0, 0x7f, 0xca, 0xde, 0x57, 0xe2, 0xba, 0xb7, 0xfd, 0x53, 0x56, 0x7b, 0x5a, 0xcd, 0x1b, 0x3b,
	0x09, 0x04, 0xb4, 0x70, 0x1b, 0xcf, 0x9b, 0xcf, 0xf7, 0xbd, 0xf9, 0xbe, 0x79, 0x46, 0x36, 0x65,
	0x23, 0x3e, 0xb1, 0xe8, 0x58, 0x52, 0x16, 0xfb, 0x9c, 0xc5, 0x56, 0xdf, 0x0f, 0x24, 0x15, 0xb1,
	0x35, 0x94, 0x32, 0xb2, 0xfa, 0x24, 0x09, 0xa4, 0x35, 0x6a, 0xe8, 0x05, 0x8e, 0x04, 0x97, 0xdc,
	0xf8, 0x1d, 0x18, 0x3c, 0x63, 0x70, 0xca, 0x60, 0xc5, 0x60, 0x7d, 0x74, 0xd4, 0xa8, 0xee, 0x6b,
	0x6d, 0x8f, 0xb3, 0xbe, 0x3f, 0xb0, 0x04, 0x4f, 0x24, 0x55, 0x52, 0xb0, 0x70, 0x3d, 0x1e, 0x46,
	0x9c, 0x51, 0x26, 0x63, 0xad, 0x5a, 0x6d, 0x3e, 0x58, 0x89, 0xc7, 0xc3, 0x90, 0xb3, 0x7b, 0x6b,
	0xa9, 0xfe, 0xa0, 0x29, 0x39, 0x89, 0x40, 0x3b, 0xa2, 0xc2, 0xa3, 0x2c, 0x0b, 0xee, 0x0e, 0x38,
	0x1f, 0x04, 0xd4, 0x82, 0xaf, 0xcb, 0xa4, 0x6f, 0x5d, 0x09, 0x12, 0x45, 0xaa, 0x50, 0x1d, 0xff,
	0x29, 0xe9, 0x45, 0xc4, 0x22, 0x8c, 0x71, 0x49, 0x24, 0xa4, 0x8c, 0x25, 0x91, 0x49, 0x16, 0xfe,
	0x65, 0x21, 0x3c, 0xa2, 0x42, 0x95, 0xe6, 0xb3, 0x41, 0x7a, 0xe4, 0xbb, 0x11, 0x09, 0xfc, 0x1e,
	0x51, 0x17, 0x4b, 0x17, 0x3a, 0xb0, 0x77, 0x5d, 0x40, 0xe8, 0x44, 0xd5, 0x79, 0x74, 0xc9, 0x85,
	0x34, 0xfe, 0x43, 0x28, 0x2d, 0x8d, 0x0c, 0xa8, 0x59, 0xa8, 0xe5, 0xea, 0x65, 0xbb, 0x86, 0xb5,
	0x8f, 0xaa, 0x76, 0x3c, 0x6a, 0xe0, 0x13, 0x41, 0x3c, 0x95, 0x86, 0x04, 0x5d, 0x7d, 0xd4, 0x99,
	0x63, 0x0c, 0x8c, 0xca, 0xca, 0x5d, 0x57, 0x57, 0x68, 0xe6, 0x6b, 0xb9, 0xfa, 0x46, 0xbb, 0xfc,
	0xb9, 0xbd, 0xfa, 0x67, 0xa9, 0xf2, 0xb1, 0x58, 0xbf, 0xc9, 0x75, 0x96, 0x1c, 0xa4, 0x4e, 0xbc,
	0x80, 0x03, 0x86, 0x87, 0xd6, 0x87, 0x94, 0xf4, 0xa8, 0x70, 0x89, 0xaa, 0xc0, 0x2c, 0x42, 0xce,
	0x43, 0xfc, 0xb8, 0xde, 0xe1, 0x59, 0xed, 0xb8, 0x03, 0x32, 0xb0, 0xee, 0x2c, 0x39, 0xe5, 0xe1,
	0xec, 0xb3, 0x7a, 0x8e, 0xca, 0x73, 0xd1, 0xd6, 0xbf, 0x6f, 0xdf, 0xbf, 0xd9, 0x6d, 0xa1, 0x03,
	0x9d, 0x43, 0xf7, 0x3d, 0xd5, 0xbf, 0x25, 0x6f, 0x3f, 0x20, 0xdf, 0x6a, 0x28, 0x01, 0x8c, 0xf6,
	0x9f, 0x22, 0xd0, 0xde, 0x42, 0x88, 0x0a, 0xc1, 0x85, 0xab, 0x8c, 0x34, 0x0a, 0x9f, 0xda, 0xb9,
	0xd3, 0xe2, 0x6a, 0xae, 0x92, 0xdf, 0xbb, 0x5e, 0x41, 0x6b, 0x9d, 0x8b, 0x8b, 0x2e, 0x9c, 0x35,
	0x4e, 0xd1, 0x72, 0x8f, 0x06, 0x64, 0x62, 0xe6, 0xc0, 0x89, 0xe6, 0xc3, 0x4e, 0xe8, 0xf7, 0x76,
	0xc7, 0x8b, 0x63, 0xc5, 0x3a, 0x5a, 0xc2, 0xe8, 0xa0, 0x65, 0xed, 0x6a, 0x1e, 0xb4, 0xec, 0xa7,
	0xbb, 0xea, 0x68, 0x01, 0xe3, 0x0f, 0x54, 0x49, 0xa2, 0x58, 0x0a, 0x4a, 0x42, 0xd7, 0x0b, 0x92,
	0x58, 0x52, 0x01, 0xcf, 0x63, 0xcd, 0xd9, 0xcc, 0xf6, 0xff, 0xd7, 0xdb, 0xc6, 0x21, 0x5a, 0xd1,
	0xde, 0xc7, 0x66, 0xb1, 0x56, 0xa8, 0x97, 0xed, 0xdf, 0xf0, 0x2d, 0x9f, 0x60, 0xae, 0x54, 0x16,
	0xed, 0xe8, 0x73, 0x22, 0xbd, 0x21, 0x15, 0x4e, 0x06, 0xa9, 0x54, 0x3d, 0x7e, 0xc5, 0xd2, 0x64,
	0x8c, 0xf7, 0x68, 0x6c, 0x2e, 0xd7, 0x0a, 0x2a, 0xd5, 0x6c, 0xff, 0x5c, 0x6d, 0x1b, 0x1d, 0xb4,
	0x15, 0x92, 0xb1, 0xab, 0xde, 0xe3, 0x88, 0xba, 0x50, 0x7e, 0x6c, 0x96, 0xe0, 0xae, 0x3f, 0x62,
	0x3d, 0x54, 0x38, 0x1b, 0x2a, 0xfc, 0xf2, 0x19, 0x93, 0x0d, 0xfb, 0x15, 0x09, 0x12, 0xea, 0x6c,
	0x86, 0x64, 0x7c, 0x04, 0x14, 0x5c, 0x35, 0x36, 0x86, 0x68, 0x5b, 0xd0, 0x38, 0xe2, 0x2c, 0xa6,
	0xae, 0x20, 0x92, 0xba, 0x81, 0x1f, 0xfa, 0xd2, 0x5c, 0x01, 0xad, 0x83, 0x27, 0xf6, 0xc0, 0x21,
	0x92, 0x9e, 0x29, 0xde, 0xd9, 0xca, 0x44, 0xa7, 0x5b, 0x86, 0x8d, 0xbe, 0x85, 0xe6, 0xb8, 0xe9,
	0xd0, 0xb8, 0x22, 0x61, 0xd2, 0x0f, 0xa9, 0xb9, 0x0a, 0x76, 0x6e, 0x43, 0x30, 0x9b, 0x2c, 0x1d,
	0x52, 0x0c, 0xb4, 0x61, 0x81, 0x59, 0xd3, 0x0c, 0x04, 0xef, 0x30, 0x4d, 0xb4, 0xa3, 0xf3, 0xf4,
	0x12, 0x01, 0xbf, 0x85, 0x29, 0x84, 0x00, 0xfa, 0x06, 0xa2, 0xc7, 0x69, 0x30, 0xa3, 0xfe, 0x41,
	0xdf, 0xeb, 0x4c, 0x73, 0x43, 0x3c, 0x05, 0xcb, 0x00, 0xee, 0xc0, 0x81, 0xce, 0x74, 0x84, 0xe7,
	0xd0, 0x85, 0x66, 0x4c, 0xd1, 0x75, 0x8d, 0xde, 0xb1, 0x3d, 0x43, 0xcf, 0xd0, 0xaf, 0xf7, 0xb8,
	0xbf, 0x70, 0xdb, 0x0d, 0x10, 0xf9, 0x79, 0xc1, 0xd3, 0xdb, 0x37, 0x6f, 0xd9, 0x6a, 0x3a, 0xff,
	0x46, 0x7f, 0x7d, 0x7d, 0x3a, 0xa7, 0x53, 0xd7, 0x3e, 0x7b, 0xf7, 0xfa, 0xe6, 0x43, 0x29, 0x5f,
	0xc9, 0xa3, 0xa6, 0xcf, 0x75, 0xbb, 0x23, 0xc1, 0xc7, 0x93, 0x47, 0x4e, 0x4c, 0x5b, 0xff, 0x44,
	0xbb, 0xea, 0xad, 0x75, 0x73, 0x97, 0x25, 0x78, 0x74, 0x8d, 0x2f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x4a, 0x72, 0x9e, 0x50, 0xb6, 0x06, 0x00, 0x00,
}
