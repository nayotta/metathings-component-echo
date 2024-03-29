package echo_service

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	log "github.com/sirupsen/logrus"

	pb "github.com/nayotta/metathings-component-echo/proto"
	component "github.com/nayotta/metathings/pkg/component"
	component_ext_firmware_service "github.com/nayotta/metathings/pkg/component_ext/firmware/service"
	component_pb "github.com/nayotta/metathings/proto/component"
)

type EchoService struct {
	*component_ext_firmware_service.ComponentExtFirmwareService
	logger log.FieldLogger
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
	text := req.GetText().GetValue()
	self.logger.WithField("#method", "Echo").Infof(text)
	return &pb.EchoResponse{Text: text}, nil
}

type StreamingEchoServerWrapper struct {
	component_pb.ModuleService_StreamCallServer
}

func (self *StreamingEchoServerWrapper) Send(res *pb.StreamingEchoResponse) error {
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

func (self *StreamingEchoServerWrapper) Recv() (*pb.StreamingEchoRequest, error) {
	var in pb.StreamingEchoRequest

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
	var req *pb.StreamingEchoRequest
	var err error

	self.logger.Infof("streaming echo started")
	defer self.logger.Infof("streaming echo closed")
	for {
		if req, err = stm.Recv(); err != nil {
			return err
		}

		text := req.GetText().GetValue()
		self.logger.WithField("#method", "StreamingEcho").Infof(text)

		res := &pb.StreamingEchoResponse{Text: text}
		if err = stm.Send(res); err != nil {
			return err
		}
	}
}

func (self *EchoService) InitModuleService(m *component.Module) error {
	var err error

	self.logger = m.Logger()
	self.ComponentExtFirmwareService, err = component_ext_firmware_service.NewComponentExtFirmwareService(m)
	if err != nil {
		return err
	}

	return nil
}
