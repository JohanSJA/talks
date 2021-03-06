// Code generated by protoc-gen-go.
// source: player2.proto
// DO NOT EDIT!

/*
Package player is a generated protocol buffer package.

It is generated from these files:
	player2.proto

It has these top-level messages:
	GetTimeRequest
	GetTimeResponse
	SetTimeRequest
	SetTimeResponse
	GetLyricRequest
	GetLyricResponse
*/
package player

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// discarding unused import google_api1 "google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GetTimeRequest struct {
}

func (m *GetTimeRequest) Reset()         { *m = GetTimeRequest{} }
func (m *GetTimeRequest) String() string { return proto.CompactTextString(m) }
func (*GetTimeRequest) ProtoMessage()    {}

type GetTimeResponse struct {
	Time int64 `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
}

func (m *GetTimeResponse) Reset()         { *m = GetTimeResponse{} }
func (m *GetTimeResponse) String() string { return proto.CompactTextString(m) }
func (*GetTimeResponse) ProtoMessage()    {}

type SetTimeRequest struct {
	Time int64 `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
}

func (m *SetTimeRequest) Reset()         { *m = SetTimeRequest{} }
func (m *SetTimeRequest) String() string { return proto.CompactTextString(m) }
func (*SetTimeRequest) ProtoMessage()    {}

type SetTimeResponse struct {
}

func (m *SetTimeResponse) Reset()         { *m = SetTimeResponse{} }
func (m *SetTimeResponse) String() string { return proto.CompactTextString(m) }
func (*SetTimeResponse) ProtoMessage()    {}

type GetLyricRequest struct {
}

func (m *GetLyricRequest) Reset()         { *m = GetLyricRequest{} }
func (m *GetLyricRequest) String() string { return proto.CompactTextString(m) }
func (*GetLyricRequest) ProtoMessage()    {}

type GetLyricResponse struct {
	Lyric string `protobuf:"bytes,1,opt,name=lyric" json:"lyric,omitempty"`
}

func (m *GetLyricResponse) Reset()         { *m = GetLyricResponse{} }
func (m *GetLyricResponse) String() string { return proto.CompactTextString(m) }
func (*GetLyricResponse) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Player service

type PlayerClient interface {
	// Get the current playing time of the player.
	GetTime(ctx context.Context, in *GetTimeRequest, opts ...grpc.CallOption) (*GetTimeResponse, error)
	// Set the current playing time of the player.
	SetTime(ctx context.Context, in *SetTimeRequest, opts ...grpc.CallOption) (*SetTimeResponse, error)
	// Get the current lyric
	GetLyric(ctx context.Context, in *GetLyricRequest, opts ...grpc.CallOption) (Player_GetLyricClient, error)
}

type playerClient struct {
	cc *grpc.ClientConn
}

func NewPlayerClient(cc *grpc.ClientConn) PlayerClient {
	return &playerClient{cc}
}

func (c *playerClient) GetTime(ctx context.Context, in *GetTimeRequest, opts ...grpc.CallOption) (*GetTimeResponse, error) {
	out := new(GetTimeResponse)
	err := grpc.Invoke(ctx, "/player.Player/GetTime", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerClient) SetTime(ctx context.Context, in *SetTimeRequest, opts ...grpc.CallOption) (*SetTimeResponse, error) {
	out := new(SetTimeResponse)
	err := grpc.Invoke(ctx, "/player.Player/SetTime", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerClient) GetLyric(ctx context.Context, in *GetLyricRequest, opts ...grpc.CallOption) (Player_GetLyricClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Player_serviceDesc.Streams[0], c.cc, "/player.Player/GetLyric", opts...)
	if err != nil {
		return nil, err
	}
	x := &playerGetLyricClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Player_GetLyricClient interface {
	Recv() (*GetLyricResponse, error)
	grpc.ClientStream
}

type playerGetLyricClient struct {
	grpc.ClientStream
}

func (x *playerGetLyricClient) Recv() (*GetLyricResponse, error) {
	m := new(GetLyricResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Player service

type PlayerServer interface {
	// Get the current playing time of the player.
	GetTime(context.Context, *GetTimeRequest) (*GetTimeResponse, error)
	// Set the current playing time of the player.
	SetTime(context.Context, *SetTimeRequest) (*SetTimeResponse, error)
	// Get the current lyric
	GetLyric(*GetLyricRequest, Player_GetLyricServer) error
}

func RegisterPlayerServer(s *grpc.Server, srv PlayerServer) {
	s.RegisterService(&_Player_serviceDesc, srv)
}

func _Player_GetTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PlayerServer).GetTime(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Player_SetTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SetTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PlayerServer).SetTime(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Player_GetLyric_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetLyricRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PlayerServer).GetLyric(m, &playerGetLyricServer{stream})
}

type Player_GetLyricServer interface {
	Send(*GetLyricResponse) error
	grpc.ServerStream
}

type playerGetLyricServer struct {
	grpc.ServerStream
}

func (x *playerGetLyricServer) Send(m *GetLyricResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Player_serviceDesc = grpc.ServiceDesc{
	ServiceName: "player.Player",
	HandlerType: (*PlayerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTime",
			Handler:    _Player_GetTime_Handler,
		},
		{
			MethodName: "SetTime",
			Handler:    _Player_SetTime_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetLyric",
			Handler:       _Player_GetLyric_Handler,
			ServerStreams: true,
		},
	},
}
