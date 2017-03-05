package easyapi

import (
	"fmt"
	"io"
	"net"

	"github.com/funny/binary"
	"github.com/golang/protobuf/proto"
)

type DefaultTcpHandle struct {
}

func (*DefaultTcpHandle) Handle(conn net.Conn, easyAPI *EasyAPI) {
	easyAPI.headBuf = easyAPI.headData[:]

	if _, err := io.ReadFull(conn, easyAPI.headBuf); err != nil {
		return
	}

	packetSize := int(binary.GetUint32LE(easyAPI.headBuf[0:4]))
	packet := make([]byte, packetSize)

	if _, err := io.ReadFull(conn, packet); err != nil {
		return
	}

	serviceID := easyAPI.headBuf[4]
	msgID := easyAPI.headBuf[5]

	service := easyAPI.services[serviceID]
	newReq := service.NewRequest(msgID)

	err := proto.Unmarshal(packet, newReq)

	if err != nil {
		fmt.Println("Error to unmarshal request message because of ", err.Error())
		return
	}

	newRsp := service.HandleRequest(msgID, newReq)
	if newReq != nil {
		newRsp, err := proto.Marshal(newRsp)

		if err != nil {
			fmt.Println("Error to marshal response message because of ", err.Error())
			return
		}

		var buff = binary.Buffer{Data: make([]byte, len(newRsp)+packetHeadSize)}
		buff.WriteUint32LE(uint32(len(newRsp)))
		buff.WriteUint8(serviceID)
		buff.WriteUint8(msgID)
		buff.WriteBytes(newRsp)
		conn.Write(buff.Data)
	}
}
