// Code generated by Kitex v0.8.0. DO NOT EDIT.

package greet

import (
	"context"
	team "generic-kitex-test/grpc_test/kitex_gen/hello/cloudwego/team"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return greetServiceInfo
}

var greetServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "greet"
	handlerType := (*team.Greet)(nil)
	methods := map[string]kitex.MethodInfo{
		"Hello": kitex.NewMethodInfo(helloHandler, newHelloArgs, newHelloResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "hello",
		"ServiceFilePath": ``,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.8.0",
		Extra:           extra,
	}
	return svcInfo
}

func helloHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(team.MyReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(team.Greet).Hello(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *HelloArgs:
		success, err := handler.(team.Greet).Hello(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*HelloResult)
		realResult.Success = success
	}
	return nil
}
func newHelloArgs() interface{} {
	return &HelloArgs{}
}

func newHelloResult() interface{} {
	return &HelloResult{}
}

type HelloArgs struct {
	Req *team.MyReq
}

func (p *HelloArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(team.MyReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *HelloArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *HelloArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *HelloArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *HelloArgs) Unmarshal(in []byte) error {
	msg := new(team.MyReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var HelloArgs_Req_DEFAULT *team.MyReq

func (p *HelloArgs) GetReq() *team.MyReq {
	if !p.IsSetReq() {
		return HelloArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *HelloArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *HelloArgs) GetFirstArgument() interface{} {
	return p.Req
}

type HelloResult struct {
	Success *team.MyResp
}

var HelloResult_Success_DEFAULT *team.MyResp

func (p *HelloResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(team.MyResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *HelloResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *HelloResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *HelloResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *HelloResult) Unmarshal(in []byte) error {
	msg := new(team.MyResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelloResult) GetSuccess() *team.MyResp {
	if !p.IsSetSuccess() {
		return HelloResult_Success_DEFAULT
	}
	return p.Success
}

func (p *HelloResult) SetSuccess(x interface{}) {
	p.Success = x.(*team.MyResp)
}

func (p *HelloResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *HelloResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Hello(ctx context.Context, Req *team.MyReq) (r *team.MyResp, err error) {
	var _args HelloArgs
	_args.Req = Req
	var _result HelloResult
	if err = p.c.Call(ctx, "Hello", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
