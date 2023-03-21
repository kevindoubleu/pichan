// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: habits.proto

package habits

import (
	context "context"

	empty "github.com/golang/protobuf/ptypes/empty"
	pichan "github.com/kevindoubleu/pichan/proto/pichan"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ScorecardsClient is the client API for Scorecards service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScorecardsClient interface {
	Describe(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*pichan.Description, error)
	Insert(ctx context.Context, in *Scorecard, opts ...grpc.CallOption) (*Scorecard, error)
	List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ScorecardList, error)
}

type scorecardsClient struct {
	cc grpc.ClientConnInterface
}

func NewScorecardsClient(cc grpc.ClientConnInterface) ScorecardsClient {
	return &scorecardsClient{cc}
}

func (c *scorecardsClient) Describe(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*pichan.Description, error) {
	out := new(pichan.Description)
	err := c.cc.Invoke(ctx, "/habits.Scorecards/Describe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scorecardsClient) Insert(ctx context.Context, in *Scorecard, opts ...grpc.CallOption) (*Scorecard, error) {
	out := new(Scorecard)
	err := c.cc.Invoke(ctx, "/habits.Scorecards/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scorecardsClient) List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ScorecardList, error) {
	out := new(ScorecardList)
	err := c.cc.Invoke(ctx, "/habits.Scorecards/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScorecardsServer is the server API for Scorecards service.
// All implementations must embed UnimplementedScorecardsServer
// for forward compatibility
type ScorecardsServer interface {
	Describe(context.Context, *empty.Empty) (*pichan.Description, error)
	Insert(context.Context, *Scorecard) (*Scorecard, error)
	List(context.Context, *empty.Empty) (*ScorecardList, error)
	mustEmbedUnimplementedScorecardsServer()
}

// UnimplementedScorecardsServer must be embedded to have forward compatible implementations.
type UnimplementedScorecardsServer struct {
}

func (UnimplementedScorecardsServer) Describe(context.Context, *empty.Empty) (*pichan.Description, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Describe not implemented")
}
func (UnimplementedScorecardsServer) Insert(context.Context, *Scorecard) (*Scorecard, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedScorecardsServer) List(context.Context, *empty.Empty) (*ScorecardList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedScorecardsServer) mustEmbedUnimplementedScorecardsServer() {}

// UnsafeScorecardsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScorecardsServer will
// result in compilation errors.
type UnsafeScorecardsServer interface {
	mustEmbedUnimplementedScorecardsServer()
}

func RegisterScorecardsServer(s grpc.ServiceRegistrar, srv ScorecardsServer) {
	s.RegisterService(&Scorecards_ServiceDesc, srv)
}

func _Scorecards_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScorecardsServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/habits.Scorecards/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScorecardsServer).Describe(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scorecards_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Scorecard)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScorecardsServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/habits.Scorecards/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScorecardsServer).Insert(ctx, req.(*Scorecard))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scorecards_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScorecardsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/habits.Scorecards/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScorecardsServer).List(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Scorecards_ServiceDesc is the grpc.ServiceDesc for Scorecards service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Scorecards_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "habits.Scorecards",
	HandlerType: (*ScorecardsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Describe",
			Handler:    _Scorecards_Describe_Handler,
		},
		{
			MethodName: "Insert",
			Handler:    _Scorecards_Insert_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Scorecards_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "habits.proto",
}
