package echo_service

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"

	pb "github.com/nayotta/metathings-component-echo/proto"
)

type EchoServiceOption struct{}

type EchoService struct {
	opt *EchoServiceOption
}

func (self *EchoService) HANDLE_GRPC_Echo(ctx context.Context, in *any.Any) (*any.Any, error) {
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

func (self *EchoService) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Text: req.GetText().GetValue()}, nil
}

func (self *EchoService) StreamingEcho(stm pb.EchoService_StreamingEchoServer) error {
	var req *pb.EchoRequest
	var err error

	for {
		if req, err = stm.Recv(); err != nil {
			return err
		}

		res := &pb.EchoResponse{Text: req.GetText().GetValue()}
		if err = stm.Send(res); err != nil {
			return err
		}
	}
}

func NewEchoService(opt *EchoServiceOption) (*EchoService, error) {
	return &EchoService{
		opt: opt,
	}, nil
}
