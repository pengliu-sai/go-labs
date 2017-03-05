package service1

import (
	"easyapi"
	"easyapi/easyapi_toy/services"

	"github.com/golang/protobuf/proto"
)

type Service1 struct {
	msgs map[byte][2]proto.Message
}

func New() easyapi.IService {
	service := &Service1{
		msgs: make(map[byte][2]proto.Message),
	}
	service.registerMsgs()
	return service
}

func (this *Service1) ServiceID() byte {
	return byte(services.ServiceID_SERVICE1)
}

func (this *Service1) NewRequest(msgID byte) proto.Message {
	switch MsgID(msgID) {
	case MsgID_ADD:
		return &AddIn{}
	}
	return nil
}

func (this *Service1) registerMsgs() {
	this.msgs[byte(MsgID_ADD)] = [2]proto.Message{&AddIn{}, &AddOut{}}
}

func (this *Service1) HandleRequest(msgID byte, req proto.Message) (rsp proto.Message) {
	switch MsgID(msgID) {
	case MsgID_ADD:
		return this.Add(req.(*AddIn))
	}
	return nil
}

func (this *Service1) Add(in *AddIn) *AddOut {
	return &AddOut{in.A + in.B}
}
