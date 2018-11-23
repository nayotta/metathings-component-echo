package echo_service

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"

	pb "github.com/nayotta/metathings-component-echo/proto"
)

type EchoModule struct{}

func (self *EchoModule) PROCESS_GRPC_Echo(ctx context.Context, in *any.Any) (*any.Any, error) {
	var err error
	req := &pb.EchoRequest{}

	if err = ptypes.UnmarshalAny(in, req); err != nil {
		return nil, err
	}

	res, err := self.Echo(ctx, req)
	if err != nil {
		return nil, err
	}

	out, err := ptypes.MarshalAny(res)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (self *EchoModule) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Text: req.GetText().GetValue()}, nil
}
