// Code generated by protoc-gen-go.
// source: bid.proto
// DO NOT EDIT!

/*
Package bid is a generated protocol buffer package.

It is generated from these files:
	bid.proto

It has these top-level messages:
	BidRequest
	BidResponse
*/
package bid

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The request containing a user's identifier
type BidRequest struct {
	UserIdentifier string `protobuf:"bytes,1,opt,name=user_identifier" json:"user_identifier,omitempty"`
}

func (m *BidRequest) Reset()                    { *m = BidRequest{} }
func (m *BidRequest) String() string            { return proto.CompactTextString(m) }
func (*BidRequest) ProtoMessage()               {}
func (*BidRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The response containing a bid or no-bid
type BidResponse struct {
	Bid bool `protobuf:"varint,1,opt,name=bid" json:"bid,omitempty"`
}

func (m *BidResponse) Reset()                    { *m = BidResponse{} }
func (m *BidResponse) String() string            { return proto.CompactTextString(m) }
func (*BidResponse) ProtoMessage()               {}
func (*BidResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*BidRequest)(nil), "bid.BidRequest")
	proto.RegisterType((*BidResponse)(nil), "bid.BidResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Bidder service

type BidderClient interface {
	// Sends a response
	Bid(ctx context.Context, in *BidRequest, opts ...grpc.CallOption) (*BidResponse, error)
}

type bidderClient struct {
	cc *grpc.ClientConn
}

func NewBidderClient(cc *grpc.ClientConn) BidderClient {
	return &bidderClient{cc}
}

func (c *bidderClient) Bid(ctx context.Context, in *BidRequest, opts ...grpc.CallOption) (*BidResponse, error) {
	out := new(BidResponse)
	err := grpc.Invoke(ctx, "/bid.Bidder/Bid", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Bidder service

type BidderServer interface {
	// Sends a response
	Bid(context.Context, *BidRequest) (*BidResponse, error)
}

func RegisterBidderServer(s *grpc.Server, srv BidderServer) {
	s.RegisterService(&_Bidder_serviceDesc, srv)
}

func _Bidder_Bid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(BidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BidderServer).Bid(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Bidder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bid.Bidder",
	HandlerType: (*BidderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Bid",
			Handler:    _Bidder_Bid_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0xca, 0x4c, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x06, 0x32, 0x95, 0x54, 0xb9, 0xb8, 0x9c, 0x32, 0x53,
	0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xc4, 0xb9, 0xf8, 0x4b, 0x8b, 0x53, 0x8b, 0xe2,
	0x33, 0x53, 0x52, 0xf3, 0x4a, 0x32, 0xd3, 0x32, 0x53, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x95, 0xa4, 0xb8, 0xb8, 0xc1, 0xca, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xb8, 0xb9, 0x40,
	0x9a, 0xc1, 0x72, 0x1c, 0x46, 0x26, 0x5c, 0x6c, 0x40, 0xb9, 0x94, 0xd4, 0x22, 0x21, 0x2d, 0x2e,
	0x66, 0x20, 0x4b, 0x88, 0x5f, 0x0f, 0x64, 0x09, 0xc2, 0x58, 0x29, 0x01, 0x84, 0x00, 0xc4, 0x00,
	0x25, 0x86, 0x24, 0x36, 0xb0, 0x23, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x16, 0xf7, 0x1d,
	0x9f, 0x91, 0x00, 0x00, 0x00,
}
