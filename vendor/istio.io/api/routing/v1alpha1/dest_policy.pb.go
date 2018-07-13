// Code generated by protoc-gen-go. DO NOT EDIT.
// source: routing/v1alpha1/dest_policy.proto

package v1alpha1 // import "istio.io/api/routing/v1alpha1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import duration "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Load balancing algorithms supported by Envoy.
type LoadBalancing_SimpleLBPolicy int32

const (
	// Simple round robin policy.
	LoadBalancing_ROUND_ROBIN LoadBalancing_SimpleLBPolicy = 0
	// The least request load balancer uses an O(1) algorithm which selects
	// two random healthy hosts and picks the host which has fewer active
	// requests.
	LoadBalancing_LEAST_CONN LoadBalancing_SimpleLBPolicy = 1
	// The random load balancer selects a random healthy host. The random
	// load balancer generally performs better than round robin if no health
	// checking policy is configured.
	LoadBalancing_RANDOM LoadBalancing_SimpleLBPolicy = 2
)

var LoadBalancing_SimpleLBPolicy_name = map[int32]string{
	0: "ROUND_ROBIN",
	1: "LEAST_CONN",
	2: "RANDOM",
}
var LoadBalancing_SimpleLBPolicy_value = map[string]int32{
	"ROUND_ROBIN": 0,
	"LEAST_CONN":  1,
	"RANDOM":      2,
}

func (x LoadBalancing_SimpleLBPolicy) String() string {
	return proto.EnumName(LoadBalancing_SimpleLBPolicy_name, int32(x))
}
func (LoadBalancing_SimpleLBPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dest_policy_101b2e64e366bbad, []int{1, 0}
}

// DestinationPolicy defines client/caller-side policies that determine how
// to handle traffic bound to a particular destination service. The policy
// specifies configuration for load balancing and circuit breakers. For
// example, a simple load balancing policy for the ratings service would
// look as follows:
//
//     metadata:
//       name: ratings-lb-policy
//       namespace: default # optional (default is "default")
//     spec:
//       destination:
//         name: ratings
//       loadBalancing:
//         name: ROUND_ROBIN
//
// The FQDN of the destination service is composed from the destination name and meta namespace fields, along with
// a platform-specific domain suffix
// (e.g. on Kubernetes, "reviews" + "default" + "svc.cluster.local" -> "reviews.default.svc.cluster.local").
//
// A destination policy can be restricted to a particular version of a
// service or applied to all versions. It can also be restricted to calls from
// a particular source. For example, the following load balancing policy
// applies to version v1 of the ratings service running in the prod
// environment but only when called from version v2 of the reviews service:
//
//
//     metadata:
//       name: ratings-lb-policy
//       namespace: default
//     spec:
//       source:
//         name: reviews
//         labels:
//           version: v2
//       destination:
//         name: ratings
//         labels:
//           env: prod
//           version: v1
//       loadBalancing:
//         name: ROUND_ROBIN
//
// *Note:* Destination policies will be applied only if the corresponding
// tagged instances are explicitly routed to. In other words, for every
// destination policy defined, at least one route rule must refer to the
// service version indicated in the destination policy.
type DestinationPolicy struct {
	// Optional: Destination uniquely identifies the destination service associated
	// with this policy.
	Destination *IstioService `protobuf:"bytes,1,opt,name=destination,proto3" json:"destination,omitempty"`
	// Optional: Source uniquely identifies the source service associated
	// with this policy.
	Source *IstioService `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	// Load balancing policy.
	LoadBalancing *LoadBalancing `protobuf:"bytes,3,opt,name=load_balancing,json=loadBalancing,proto3" json:"load_balancing,omitempty"`
	// Circuit breaker policy.
	CircuitBreaker *CircuitBreaker `protobuf:"bytes,4,opt,name=circuit_breaker,json=circuitBreaker,proto3" json:"circuit_breaker,omitempty"`
	// (-- Other custom policy implementations --)
	Custom               *any.Any `protobuf:"bytes,100,opt,name=custom,proto3" json:"custom,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DestinationPolicy) Reset()         { *m = DestinationPolicy{} }
func (m *DestinationPolicy) String() string { return proto.CompactTextString(m) }
func (*DestinationPolicy) ProtoMessage()    {}
func (*DestinationPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_dest_policy_101b2e64e366bbad, []int{0}
}
func (m *DestinationPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestinationPolicy.Unmarshal(m, b)
}
func (m *DestinationPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestinationPolicy.Marshal(b, m, deterministic)
}
func (dst *DestinationPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestinationPolicy.Merge(dst, src)
}
func (m *DestinationPolicy) XXX_Size() int {
	return xxx_messageInfo_DestinationPolicy.Size(m)
}
func (m *DestinationPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_DestinationPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_DestinationPolicy proto.InternalMessageInfo

func (m *DestinationPolicy) GetDestination() *IstioService {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *DestinationPolicy) GetSource() *IstioService {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *DestinationPolicy) GetLoadBalancing() *LoadBalancing {
	if m != nil {
		return m.LoadBalancing
	}
	return nil
}

func (m *DestinationPolicy) GetCircuitBreaker() *CircuitBreaker {
	if m != nil {
		return m.CircuitBreaker
	}
	return nil
}

func (m *DestinationPolicy) GetCustom() *any.Any {
	if m != nil {
		return m.Custom
	}
	return nil
}

// Load balancing policy to use when forwarding traffic. These policies
// directly correlate to [load balancer
// types](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/load_balancing)
// supported by Envoy. Example,
//
// ```yaml
// metadata:
//   name: reviews-lb-policy
//   namespace: default
// spec:
//   destination:
//     name: reviews
//   loadBalancing:
//     name: RANDOM
// ```
type LoadBalancing struct {
	// Types that are valid to be assigned to LbPolicy:
	//	*LoadBalancing_Name
	//	*LoadBalancing_Custom
	LbPolicy             isLoadBalancing_LbPolicy `protobuf_oneof:"lb_policy"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *LoadBalancing) Reset()         { *m = LoadBalancing{} }
func (m *LoadBalancing) String() string { return proto.CompactTextString(m) }
func (*LoadBalancing) ProtoMessage()    {}
func (*LoadBalancing) Descriptor() ([]byte, []int) {
	return fileDescriptor_dest_policy_101b2e64e366bbad, []int{1}
}
func (m *LoadBalancing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancing.Unmarshal(m, b)
}
func (m *LoadBalancing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancing.Marshal(b, m, deterministic)
}
func (dst *LoadBalancing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancing.Merge(dst, src)
}
func (m *LoadBalancing) XXX_Size() int {
	return xxx_messageInfo_LoadBalancing.Size(m)
}
func (m *LoadBalancing) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancing.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancing proto.InternalMessageInfo

type isLoadBalancing_LbPolicy interface {
	isLoadBalancing_LbPolicy()
}

type LoadBalancing_Name struct {
	Name LoadBalancing_SimpleLBPolicy `protobuf:"varint,1,opt,name=name,proto3,enum=istio.routing.v1alpha1.LoadBalancing_SimpleLBPolicy,oneof"`
}
type LoadBalancing_Custom struct {
	Custom *any.Any `protobuf:"bytes,2,opt,name=custom,proto3,oneof"`
}

func (*LoadBalancing_Name) isLoadBalancing_LbPolicy()   {}
func (*LoadBalancing_Custom) isLoadBalancing_LbPolicy() {}

func (m *LoadBalancing) GetLbPolicy() isLoadBalancing_LbPolicy {
	if m != nil {
		return m.LbPolicy
	}
	return nil
}

func (m *LoadBalancing) GetName() LoadBalancing_SimpleLBPolicy {
	if x, ok := m.GetLbPolicy().(*LoadBalancing_Name); ok {
		return x.Name
	}
	return LoadBalancing_ROUND_ROBIN
}

func (m *LoadBalancing) GetCustom() *any.Any {
	if x, ok := m.GetLbPolicy().(*LoadBalancing_Custom); ok {
		return x.Custom
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LoadBalancing) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LoadBalancing_OneofMarshaler, _LoadBalancing_OneofUnmarshaler, _LoadBalancing_OneofSizer, []interface{}{
		(*LoadBalancing_Name)(nil),
		(*LoadBalancing_Custom)(nil),
	}
}

func _LoadBalancing_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LoadBalancing)
	// lb_policy
	switch x := m.LbPolicy.(type) {
	case *LoadBalancing_Name:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Name))
	case *LoadBalancing_Custom:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Custom); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LoadBalancing.LbPolicy has unexpected type %T", x)
	}
	return nil
}

func _LoadBalancing_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LoadBalancing)
	switch tag {
	case 1: // lb_policy.name
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.LbPolicy = &LoadBalancing_Name{LoadBalancing_SimpleLBPolicy(x)}
		return true, err
	case 2: // lb_policy.custom
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(any.Any)
		err := b.DecodeMessage(msg)
		m.LbPolicy = &LoadBalancing_Custom{msg}
		return true, err
	default:
		return false, nil
	}
}

func _LoadBalancing_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LoadBalancing)
	// lb_policy
	switch x := m.LbPolicy.(type) {
	case *LoadBalancing_Name:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Name))
	case *LoadBalancing_Custom:
		s := proto.Size(x.Custom)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Circuit breaker configuration for Envoy. The circuit breaker
// implementation is fine-grained in that it tracks the success/failure
// rates of individual hosts in the load balancing pool. Hosts that
// continually return errors for API calls are ejected from the pool for a
// pre-defined period of time.
// See Envoy's
// [circuit breaker](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/circuit_breaking)
// and [outlier detection](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/outlier)
// for more details.
type CircuitBreaker struct {
	// Types that are valid to be assigned to CbPolicy:
	//	*CircuitBreaker_SimpleCb
	//	*CircuitBreaker_Custom
	CbPolicy             isCircuitBreaker_CbPolicy `protobuf_oneof:"cb_policy"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *CircuitBreaker) Reset()         { *m = CircuitBreaker{} }
func (m *CircuitBreaker) String() string { return proto.CompactTextString(m) }
func (*CircuitBreaker) ProtoMessage()    {}
func (*CircuitBreaker) Descriptor() ([]byte, []int) {
	return fileDescriptor_dest_policy_101b2e64e366bbad, []int{2}
}
func (m *CircuitBreaker) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CircuitBreaker.Unmarshal(m, b)
}
func (m *CircuitBreaker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CircuitBreaker.Marshal(b, m, deterministic)
}
func (dst *CircuitBreaker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreaker.Merge(dst, src)
}
func (m *CircuitBreaker) XXX_Size() int {
	return xxx_messageInfo_CircuitBreaker.Size(m)
}
func (m *CircuitBreaker) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreaker.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreaker proto.InternalMessageInfo

type isCircuitBreaker_CbPolicy interface {
	isCircuitBreaker_CbPolicy()
}

type CircuitBreaker_SimpleCb struct {
	SimpleCb *CircuitBreaker_SimpleCircuitBreakerPolicy `protobuf:"bytes,1,opt,name=simple_cb,json=simpleCb,proto3,oneof"`
}
type CircuitBreaker_Custom struct {
	Custom *any.Any `protobuf:"bytes,2,opt,name=custom,proto3,oneof"`
}

func (*CircuitBreaker_SimpleCb) isCircuitBreaker_CbPolicy() {}
func (*CircuitBreaker_Custom) isCircuitBreaker_CbPolicy()   {}

func (m *CircuitBreaker) GetCbPolicy() isCircuitBreaker_CbPolicy {
	if m != nil {
		return m.CbPolicy
	}
	return nil
}

func (m *CircuitBreaker) GetSimpleCb() *CircuitBreaker_SimpleCircuitBreakerPolicy {
	if x, ok := m.GetCbPolicy().(*CircuitBreaker_SimpleCb); ok {
		return x.SimpleCb
	}
	return nil
}

func (m *CircuitBreaker) GetCustom() *any.Any {
	if x, ok := m.GetCbPolicy().(*CircuitBreaker_Custom); ok {
		return x.Custom
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CircuitBreaker) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CircuitBreaker_OneofMarshaler, _CircuitBreaker_OneofUnmarshaler, _CircuitBreaker_OneofSizer, []interface{}{
		(*CircuitBreaker_SimpleCb)(nil),
		(*CircuitBreaker_Custom)(nil),
	}
}

func _CircuitBreaker_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CircuitBreaker)
	// cb_policy
	switch x := m.CbPolicy.(type) {
	case *CircuitBreaker_SimpleCb:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SimpleCb); err != nil {
			return err
		}
	case *CircuitBreaker_Custom:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Custom); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CircuitBreaker.CbPolicy has unexpected type %T", x)
	}
	return nil
}

func _CircuitBreaker_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CircuitBreaker)
	switch tag {
	case 1: // cb_policy.simple_cb
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CircuitBreaker_SimpleCircuitBreakerPolicy)
		err := b.DecodeMessage(msg)
		m.CbPolicy = &CircuitBreaker_SimpleCb{msg}
		return true, err
	case 2: // cb_policy.custom
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(any.Any)
		err := b.DecodeMessage(msg)
		m.CbPolicy = &CircuitBreaker_Custom{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CircuitBreaker_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CircuitBreaker)
	// cb_policy
	switch x := m.CbPolicy.(type) {
	case *CircuitBreaker_SimpleCb:
		s := proto.Size(x.SimpleCb)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CircuitBreaker_Custom:
		s := proto.Size(x.Custom)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A simple circuit breaker can be set based on a number of criteria such as
// connection and request limits. For example, the following destination
// policy sets a limit of 100 connections to "reviews" service version
// "v1" backends.
//
// ```yaml
// metadata:
//   name: reviews-cb-policy
//   namespace: default
// spec:
//   destination:
//     name: reviews
//     labels:
//       version: v1
//   circuitBreaker:
//     simpleCb:
//       maxConnections: 100
// ```
//
// The following destination policy sets a limit of 100 connections and
// 1000 concurrent requests, with no more than 10 req/connection to
// "reviews" service version "v1" backends. In addition, it configures
// hosts to be scanned every 5 mins, such that any host that fails 7
// consecutive times with 5XX error code will be ejected for 15 minutes.
//
// ```yaml
// metadata:
//   name: reviews-cb-policy
//   namespace: default
// spec:
//   destination:
//     name: reviews
//     labels:
//       version: v1
//   circuitBreaker:
//     simpleCb:
//       maxConnections: 100
//       httpMaxRequests: 1000
//       httpMaxRequestsPerConnection: 10
//       httpConsecutiveErrors: 7
//       sleepWindow: 15m
//       httpDetectionInterval: 5m
// ```
type CircuitBreaker_SimpleCircuitBreakerPolicy struct {
	// Maximum number of connections to a backend.
	MaxConnections int32 `protobuf:"varint,1,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections,omitempty"`
	// Maximum number of pending requests to a backend. Default 1024
	HttpMaxPendingRequests int32 `protobuf:"varint,2,opt,name=http_max_pending_requests,json=httpMaxPendingRequests,proto3" json:"http_max_pending_requests,omitempty"`
	// Maximum number of requests to a backend. Default 1024
	HttpMaxRequests int32 `protobuf:"varint,3,opt,name=http_max_requests,json=httpMaxRequests,proto3" json:"http_max_requests,omitempty"`
	// Minimum time the circuit will be open. format: 1h/1m/1s/1ms. MUST
	// BE >=1ms. Default is 30s.
	SleepWindow *duration.Duration `protobuf:"bytes,4,opt,name=sleep_window,json=sleepWindow,proto3" json:"sleep_window,omitempty"`
	// Number of 5XX errors before circuit is opened. Defaults to 5.
	HttpConsecutiveErrors int32 `protobuf:"varint,5,opt,name=http_consecutive_errors,json=httpConsecutiveErrors,proto3" json:"http_consecutive_errors,omitempty"`
	// Time interval between ejection sweep analysis. format:
	// 1h/1m/1s/1ms. MUST BE >=1ms. Default is 10s.
	HttpDetectionInterval *duration.Duration `protobuf:"bytes,6,opt,name=http_detection_interval,json=httpDetectionInterval,proto3" json:"http_detection_interval,omitempty"`
	// Maximum number of requests per connection to a backend. Setting this
	// parameter to 1 disables keep alive.
	HttpMaxRequestsPerConnection int32 `protobuf:"varint,7,opt,name=http_max_requests_per_connection,json=httpMaxRequestsPerConnection,proto3" json:"http_max_requests_per_connection,omitempty"`
	// Maximum % of hosts in the load balancing pool for the destination
	// service that can be ejected by the circuit breaker. Defaults to
	// 10%.
	HttpMaxEjectionPercent int32 `protobuf:"varint,8,opt,name=http_max_ejection_percent,json=httpMaxEjectionPercent,proto3" json:"http_max_ejection_percent,omitempty"`
	// Maximum number of retries that can be outstanding to all hosts in a
	// cluster at a given time. Defaults to 3.
	HttpMaxRetries       int32    `protobuf:"varint,9,opt,name=http_max_retries,json=httpMaxRetries,proto3" json:"http_max_retries,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) Reset() {
	*m = CircuitBreaker_SimpleCircuitBreakerPolicy{}
}
func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) String() string { return proto.CompactTextString(m) }
func (*CircuitBreaker_SimpleCircuitBreakerPolicy) ProtoMessage()    {}
func (*CircuitBreaker_SimpleCircuitBreakerPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_dest_policy_101b2e64e366bbad, []int{2, 0}
}
func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CircuitBreaker_SimpleCircuitBreakerPolicy.Unmarshal(m, b)
}
func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CircuitBreaker_SimpleCircuitBreakerPolicy.Marshal(b, m, deterministic)
}
func (dst *CircuitBreaker_SimpleCircuitBreakerPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreaker_SimpleCircuitBreakerPolicy.Merge(dst, src)
}
func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) XXX_Size() int {
	return xxx_messageInfo_CircuitBreaker_SimpleCircuitBreakerPolicy.Size(m)
}
func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreaker_SimpleCircuitBreakerPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreaker_SimpleCircuitBreakerPolicy proto.InternalMessageInfo

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) GetMaxConnections() int32 {
	if m != nil {
		return m.MaxConnections
	}
	return 0
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) GetHttpMaxPendingRequests() int32 {
	if m != nil {
		return m.HttpMaxPendingRequests
	}
	return 0
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) GetHttpMaxRequests() int32 {
	if m != nil {
		return m.HttpMaxRequests
	}
	return 0
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) GetSleepWindow() *duration.Duration {
	if m != nil {
		return m.SleepWindow
	}
	return nil
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) GetHttpConsecutiveErrors() int32 {
	if m != nil {
		return m.HttpConsecutiveErrors
	}
	return 0
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) GetHttpDetectionInterval() *duration.Duration {
	if m != nil {
		return m.HttpDetectionInterval
	}
	return nil
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) GetHttpMaxRequestsPerConnection() int32 {
	if m != nil {
		return m.HttpMaxRequestsPerConnection
	}
	return 0
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) GetHttpMaxEjectionPercent() int32 {
	if m != nil {
		return m.HttpMaxEjectionPercent
	}
	return 0
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) GetHttpMaxRetries() int32 {
	if m != nil {
		return m.HttpMaxRetries
	}
	return 0
}

func init() {
	proto.RegisterType((*DestinationPolicy)(nil), "istio.routing.v1alpha1.DestinationPolicy")
	proto.RegisterType((*LoadBalancing)(nil), "istio.routing.v1alpha1.LoadBalancing")
	proto.RegisterType((*CircuitBreaker)(nil), "istio.routing.v1alpha1.CircuitBreaker")
	proto.RegisterType((*CircuitBreaker_SimpleCircuitBreakerPolicy)(nil), "istio.routing.v1alpha1.CircuitBreaker.SimpleCircuitBreakerPolicy")
	proto.RegisterEnum("istio.routing.v1alpha1.LoadBalancing_SimpleLBPolicy", LoadBalancing_SimpleLBPolicy_name, LoadBalancing_SimpleLBPolicy_value)
}

func init() {
	proto.RegisterFile("routing/v1alpha1/dest_policy.proto", fileDescriptor_dest_policy_101b2e64e366bbad)
}

var fileDescriptor_dest_policy_101b2e64e366bbad = []byte{
	// 667 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x6d, 0x4f, 0x13, 0x4b,
	0x14, 0xa6, 0x85, 0xf6, 0xc2, 0xf4, 0xd2, 0x96, 0xc9, 0xbd, 0xdc, 0xa5, 0xb9, 0x2a, 0x36, 0xbe,
	0x10, 0x63, 0xb6, 0x01, 0x8d, 0x89, 0x09, 0x7e, 0xe8, 0x0b, 0x04, 0x4c, 0x69, 0xeb, 0xa2, 0x31,
	0xf1, 0xcb, 0x38, 0x9d, 0x3d, 0x96, 0xd1, 0xed, 0xcc, 0x3a, 0x3b, 0x5b, 0xe0, 0xef, 0xf8, 0x9b,
	0x4c, 0xfc, 0x07, 0xfe, 0x0e, 0xd3, 0xd9, 0xe9, 0xcb, 0x82, 0x8d, 0xf0, 0x71, 0xcf, 0x79, 0x9e,
	0x67, 0x9e, 0xf3, 0x9c, 0xd3, 0xa2, 0xaa, 0x92, 0xb1, 0xe6, 0x62, 0x50, 0x1b, 0xed, 0xd2, 0x20,
	0x3c, 0xa3, 0xbb, 0x35, 0x1f, 0x22, 0x4d, 0x42, 0x19, 0x70, 0x76, 0xe9, 0x86, 0x4a, 0x6a, 0x89,
	0x37, 0x79, 0xa4, 0xb9, 0x74, 0x2d, 0xd2, 0x9d, 0x20, 0x2b, 0x5b, 0x03, 0x29, 0x07, 0x01, 0xd4,
	0x0c, 0xaa, 0x1f, 0x7f, 0xaa, 0x51, 0x61, 0x29, 0x95, 0xbb, 0x57, 0x5b, 0x7e, 0xac, 0xa8, 0xe6,
	0x52, 0xd8, 0xfe, 0xfd, 0x6b, 0xcf, 0x8e, 0x0b, 0x40, 0x54, 0x1c, 0x40, 0x02, 0xa9, 0xfe, 0xcc,
	0xa2, 0x8d, 0x16, 0x44, 0x9a, 0x0b, 0x43, 0xec, 0x19, 0x47, 0xf8, 0x10, 0x15, 0xfc, 0x59, 0xd1,
	0xc9, 0x6c, 0x67, 0x76, 0x0a, 0x7b, 0x0f, 0xdc, 0xdf, 0x3b, 0x74, 0x8f, 0xc7, 0xe5, 0x53, 0x50,
	0x23, 0xce, 0xc0, 0x9b, 0x27, 0xe2, 0x7d, 0x94, 0x8f, 0x64, 0xac, 0x18, 0x38, 0xd9, 0x5b, 0x48,
	0x58, 0x0e, 0x6e, 0xa3, 0x62, 0x20, 0xa9, 0x4f, 0xfa, 0x34, 0xa0, 0x82, 0x71, 0x31, 0x70, 0x96,
	0x8d, 0xca, 0xc3, 0x45, 0x2a, 0x6d, 0x49, 0xfd, 0xc6, 0x04, 0xec, 0xad, 0x07, 0xf3, 0x9f, 0xb8,
	0x8b, 0x4a, 0x8c, 0x2b, 0x16, 0x73, 0x4d, 0xfa, 0x0a, 0xe8, 0x17, 0x50, 0xce, 0x8a, 0x91, 0x7b,
	0xb4, 0x48, 0xae, 0x99, 0xc0, 0x1b, 0x09, 0xda, 0x2b, 0xb2, 0xd4, 0x37, 0x7e, 0x8a, 0xf2, 0x2c,
	0x8e, 0xb4, 0x1c, 0x3a, 0xbe, 0xd1, 0xf9, 0xc7, 0x4d, 0xd6, 0xe1, 0x4e, 0xd6, 0xe1, 0xd6, 0xc5,
	0xa5, 0x67, 0x31, 0xd5, 0x1f, 0x19, 0xb4, 0x9e, 0xf2, 0x87, 0x5f, 0xa3, 0x15, 0x41, 0x87, 0x60,
	0xd2, 0x2d, 0xee, 0x3d, 0xbf, 0xd1, 0x50, 0xee, 0x29, 0x1f, 0x86, 0x01, 0xb4, 0x1b, 0xc9, 0xa2,
	0x8e, 0x96, 0x3c, 0xa3, 0x81, 0xdd, 0xa9, 0x97, 0xec, 0x62, 0x2f, 0x47, 0x4b, 0x53, 0x37, 0xaf,
	0x50, 0x31, 0xad, 0x84, 0x4b, 0xa8, 0xe0, 0x75, 0xdf, 0x75, 0x5a, 0xc4, 0xeb, 0x36, 0x8e, 0x3b,
	0xe5, 0x25, 0x5c, 0x44, 0xa8, 0x7d, 0x50, 0x3f, 0x7d, 0x4b, 0x9a, 0xdd, 0x4e, 0xa7, 0x9c, 0xc1,
	0x08, 0xe5, 0xbd, 0x7a, 0xa7, 0xd5, 0x3d, 0x29, 0x67, 0x1b, 0x05, 0xb4, 0x16, 0xf4, 0xed, 0xf9,
	0x56, 0xbf, 0xe7, 0x50, 0x31, 0x1d, 0x15, 0xfe, 0x88, 0xd6, 0x22, 0x23, 0x4f, 0x58, 0xdf, 0x5e,
	0x4f, 0xfd, 0x66, 0x29, 0xdb, 0x01, 0xd3, 0xc5, 0xe9, 0xb0, 0xab, 0x89, 0x6a, 0xb3, 0x7f, 0xdb,
	0x81, 0x2b, 0xdf, 0x56, 0x50, 0x65, 0xb1, 0x34, 0x7e, 0x8c, 0x4a, 0x43, 0x7a, 0x41, 0x98, 0x14,
	0x02, 0xd8, 0xf8, 0x74, 0x23, 0x63, 0x3b, 0xe7, 0x15, 0x87, 0xf4, 0xa2, 0x39, 0xab, 0xe2, 0x97,
	0x68, 0xeb, 0x4c, 0xeb, 0x90, 0x8c, 0xd1, 0x21, 0x08, 0x9f, 0x8b, 0x01, 0x51, 0xf0, 0x35, 0x86,
	0x48, 0x47, 0xc6, 0x4a, 0xce, 0xdb, 0x1c, 0x03, 0x4e, 0xe8, 0x45, 0x2f, 0x69, 0x7b, 0xb6, 0x8b,
	0x9f, 0xa0, 0x8d, 0x29, 0x75, 0x4a, 0x59, 0x36, 0x94, 0x92, 0xa5, 0x4c, 0xb1, 0xfb, 0xe8, 0xef,
	0x28, 0x00, 0x08, 0xc9, 0x39, 0x17, 0xbe, 0x3c, 0xb7, 0x97, 0xba, 0x75, 0x6d, 0xc8, 0x96, 0xfd,
	0xc1, 0x7b, 0x05, 0x03, 0x7f, 0x6f, 0xd0, 0xf8, 0x05, 0xfa, 0xcf, 0xbc, 0xc4, 0xa4, 0x88, 0x80,
	0xc5, 0x9a, 0x8f, 0x80, 0x80, 0x52, 0x52, 0x45, 0x4e, 0xce, 0xbc, 0xf7, 0xef, 0xb8, 0xdd, 0x9c,
	0x75, 0x0f, 0x4c, 0x13, 0xbf, 0xb1, 0x3c, 0x1f, 0x74, 0x32, 0x2f, 0xe1, 0x42, 0x83, 0x1a, 0xd1,
	0xc0, 0xc9, 0xff, 0xc9, 0x80, 0x91, 0x6c, 0x4d, 0x88, 0xc7, 0x96, 0x87, 0x0f, 0xd1, 0xf6, 0xb5,
	0xa1, 0x49, 0x08, 0x6a, 0x2e, 0x6a, 0xe7, 0x2f, 0xe3, 0xe9, 0xff, 0x2b, 0x19, 0xf4, 0x40, 0xcd,
	0x82, 0x4f, 0xe5, 0x0e, 0x9f, 0xad, 0xbb, 0x10, 0x14, 0x03, 0xa1, 0x9d, 0xd5, 0x54, 0xee, 0x07,
	0xb6, 0xdd, 0x4b, 0xba, 0x78, 0x07, 0x95, 0xe7, 0x2c, 0x68, 0xc5, 0x21, 0x72, 0xd6, 0x92, 0xe5,
	0x4e, 0x9f, 0x34, 0xd5, 0xf1, 0x59, 0xb3, 0xc9, 0x59, 0x37, 0xee, 0x7d, 0xb8, 0x93, 0x5c, 0x2c,
	0x97, 0x35, 0x1a, 0xf2, 0xda, 0xd5, 0xff, 0xd2, 0x7e, 0xde, 0x84, 0xf0, 0xec, 0x57, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x42, 0x2f, 0xdb, 0xec, 0xdd, 0x05, 0x00, 0x00,
}