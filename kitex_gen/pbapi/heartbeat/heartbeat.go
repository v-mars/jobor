// Code generated by Kitex v0.7.0. DO NOT EDIT.

package heartbeat

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	pbapi "jobor/kitex_gen/pbapi"
)

func serviceInfo() *kitex.ServiceInfo {
	return heartbeatServiceInfo
}

var heartbeatServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "Heartbeat"
	handlerType := (*pbapi.Heartbeat)(nil)
	methods := map[string]kitex.MethodInfo{
		"RegistryWorker": kitex.NewMethodInfo(registryWorkerHandler, newRegistryWorkerArgs, newRegistryWorkerResult, false),
		"SendHeartbeat":  kitex.NewMethodInfo(sendHeartbeatHandler, newSendHeartbeatArgs, newSendHeartbeatResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "srv_rpc",
		"ServiceFilePath": "",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.7.0",
		Extra:           extra,
	}
	return svcInfo
}

func registryWorkerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(pbapi.RegistryReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(pbapi.Heartbeat).RegistryWorker(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *RegistryWorkerArgs:
		success, err := handler.(pbapi.Heartbeat).RegistryWorker(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RegistryWorkerResult)
		realResult.Success = success
	}
	return nil
}
func newRegistryWorkerArgs() interface{} {
	return &RegistryWorkerArgs{}
}

func newRegistryWorkerResult() interface{} {
	return &RegistryWorkerResult{}
}

type RegistryWorkerArgs struct {
	Req *pbapi.RegistryReq
}

func (p *RegistryWorkerArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *RegistryWorkerArgs) Unmarshal(in []byte) error {
	if len(in) == 0 {
		return nil
	}
	msg := new(pbapi.RegistryReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RegistryWorkerArgs_Req_DEFAULT *pbapi.RegistryReq

func (p *RegistryWorkerArgs) GetReq() *pbapi.RegistryReq {
	if !p.IsSetReq() {
		return RegistryWorkerArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RegistryWorkerArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RegistryWorkerArgs) GetFirstArgument() interface{} {
	return p.Req
}

type RegistryWorkerResult struct {
	Success *pbapi.Empty
}

var RegistryWorkerResult_Success_DEFAULT *pbapi.Empty

func (p *RegistryWorkerResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *RegistryWorkerResult) Unmarshal(in []byte) error {
	if len(in) == 0 {
		return nil
	}
	msg := new(pbapi.Empty)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RegistryWorkerResult) GetSuccess() *pbapi.Empty {
	if !p.IsSetSuccess() {
		return RegistryWorkerResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RegistryWorkerResult) SetSuccess(x interface{}) {
	p.Success = x.(*pbapi.Empty)
}

func (p *RegistryWorkerResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RegistryWorkerResult) GetResult() interface{} {
	return p.Success
}

func sendHeartbeatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(pbapi.HeartbeatReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(pbapi.Heartbeat).SendHeartbeat(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *SendHeartbeatArgs:
		success, err := handler.(pbapi.Heartbeat).SendHeartbeat(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SendHeartbeatResult)
		realResult.Success = success
	}
	return nil
}
func newSendHeartbeatArgs() interface{} {
	return &SendHeartbeatArgs{}
}

func newSendHeartbeatResult() interface{} {
	return &SendHeartbeatResult{}
}

type SendHeartbeatArgs struct {
	Req *pbapi.HeartbeatReq
}

func (p *SendHeartbeatArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *SendHeartbeatArgs) Unmarshal(in []byte) error {
	if len(in) == 0 {
		return nil
	}
	msg := new(pbapi.HeartbeatReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SendHeartbeatArgs_Req_DEFAULT *pbapi.HeartbeatReq

func (p *SendHeartbeatArgs) GetReq() *pbapi.HeartbeatReq {
	if !p.IsSetReq() {
		return SendHeartbeatArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SendHeartbeatArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SendHeartbeatArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SendHeartbeatResult struct {
	Success *pbapi.Empty
}

var SendHeartbeatResult_Success_DEFAULT *pbapi.Empty

func (p *SendHeartbeatResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *SendHeartbeatResult) Unmarshal(in []byte) error {
	if len(in) == 0 {
		return nil
	}
	msg := new(pbapi.Empty)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SendHeartbeatResult) GetSuccess() *pbapi.Empty {
	if !p.IsSetSuccess() {
		return SendHeartbeatResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SendHeartbeatResult) SetSuccess(x interface{}) {
	p.Success = x.(*pbapi.Empty)
}

func (p *SendHeartbeatResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SendHeartbeatResult) GetResult() interface{} {
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

func (p *kClient) RegistryWorker(ctx context.Context, Req *pbapi.RegistryReq) (r *pbapi.Empty, err error) {
	var _args RegistryWorkerArgs
	_args.Req = Req
	var _result RegistryWorkerResult
	if err = p.c.Call(ctx, "RegistryWorker", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SendHeartbeat(ctx context.Context, Req *pbapi.HeartbeatReq) (r *pbapi.Empty, err error) {
	var _args SendHeartbeatArgs
	_args.Req = Req
	var _result SendHeartbeatResult
	if err = p.c.Call(ctx, "SendHeartbeat", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
