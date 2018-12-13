package echo_service

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	component_pb "github.com/nayotta/metathings/pkg/proto/component"

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

type StreamingEchoServerWrapper struct {
	component_pb.ModuleService_StreamCallServer
}

func (self *StreamingEchoServerWrapper) Send(res *pb.EchoResponse) error {
	res_val, err := ptypes.MarshalAny(res)
	if err != nil {
		return err
	}

	out := &component_pb.StreamCallResponse{
		Response: &component_pb.StreamCallResponse_Data{
			Data: &component_pb.StreamCallDataResponse{
				Value: res_val,
			},
		},
	}

	err = self.ModuleService_StreamCallServer.Send(out)

	if err != nil {
		return err
	}

	return nil
}

func (self *StreamingEchoServerWrapper) Recv() (*pb.EchoRequest, error) {
	var in pb.EchoRequest

	req, err := self.ModuleService_StreamCallServer.Recv()
	if err != nil {
		return nil, err
	}

	req_val := req.GetData().GetValue()

	err = ptypes.UnmarshalAny(req_val, &in)
	if err != nil {
		return nil, err
	}

	return &in, nil
}

func NewStreamingEchoServerWrapper(upstm component_pb.ModuleService_StreamCallServer) pb.EchoService_StreamingEchoServer {
	return &StreamingEchoServerWrapper{upstm}
}

func (self *EchoService) HANDLE_GRPC_StreamingEcho(upstm component_pb.ModuleService_StreamCallServer) error {
	return self.StreamingEcho(NewStreamingEchoServerWrapper(upstm))
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
