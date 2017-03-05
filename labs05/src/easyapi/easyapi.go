package easyapi

import (
	"log"
	"net"
	"strings"

	"time"

	"github.com/golang/protobuf/proto"
)

const packetHeadSize = 4 + 2

type EasyAPI struct {
	services []IService
	handle   IHandle
	headBuf  []byte
	headData [packetHeadSize]byte
	listener net.Listener
}

type IService interface {
	ServiceID() byte
	NewRequest(byte) proto.Message
	HandleRequest(byte, proto.Message) proto.Message
}

type IHandle interface {
	Handle(conn net.Conn, easyAPI *EasyAPI)
}

func NewEasyAPI() *EasyAPI {
	return &EasyAPI{
		services: make([]IService, 256),
		handle:   &DefaultTcpHandle{},
	}
}

func (easyAPI *EasyAPI) RegisterService(s IService) {
	easyAPI.services[s.ServiceID()] = s
}

func (easyAPI *EasyAPI) Listen(network, address string) (net.Listener, error) {
	listener, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}

	easyAPI.listener = listener
	return listener, nil
}

func (easyAPI *EasyAPI) Close() {
	if easyAPI.listener != nil {
		easyAPI.listener.Close()
	}
}

func (easyAPI *EasyAPI) Serve(listener net.Listener) {
	log.Printf("easyapi serve on %s\n", listener.Addr().String())

	var tempDelay time.Duration

	for {
		conn, err := listener.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				time.Sleep(tempDelay)
				continue
			}
			if strings.Contains(err.Error(), "use of closed network connection") {
				break
			}
			return
		}
		go easyAPI.handle.Handle(conn, easyAPI)
	}
}

func (easyAPI *EasyAPI) Dial(network, address string) (net.Conn, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
