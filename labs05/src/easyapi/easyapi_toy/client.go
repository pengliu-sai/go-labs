package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"sync"

	"easyapi"
	"easyapi/easyapi_toy/services"
	"easyapi/easyapi_toy/services/service1"

	"github.com/funny/binary"
	"github.com/golang/protobuf/proto"
)

func main() {
	easyapi := easyapi.NewEasyAPI()

	conn, err := easyapi.Dial("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}

	defer conn.Close()

	fmt.Println("Connecting to " + conn.RemoteAddr().String())

	var wg sync.WaitGroup
	wg.Add(2)

	go handleWrite(conn, &wg)
	go handleRead(conn, &wg)

	wg.Wait()
}

func handleWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	addIn := &service1.AddIn{A: 1, B: 2}
	addReq, err := proto.Marshal(addIn)
	if err != nil {
		fmt.Println("Error to marshal req message because of ", err.Error())
		os.Exit(1)
	}

	var buff = binary.Buffer{Data: make([]byte, len(addReq)+6)}
	buff.WriteUint32LE(uint32(len(addReq)))
	buff.WriteUint8(uint8(services.ServiceID_SERVICE1))
	buff.WriteUint8(uint8(service1.MsgID_ADD))
	buff.WriteBytes(addReq)

	_, err = conn.Write(buff.Data)

	if err != nil {
		fmt.Println("Error to send message because of ", err.Error())
	}
}

func handleRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	var headBuf = make([]byte, 6)

	if _, err := io.ReadFull(conn, headBuf); err != nil {
		return
	}

	packageSize := int(binary.GetUint32LE(headBuf[0:4]))
	packet := make([]byte, packageSize)

	fmt.Println("read message len: ", packageSize)

	if _, err := io.ReadFull(conn, packet); err != nil {
		fmt.Println("Error to read message because of ", err.Error())
	}

	addOut := &service1.AddOut{}
	err := proto.Unmarshal(packet, addOut)
	if err != nil {
		fmt.Println("Error to unmarshal message because of ", err.Error())
	}

	fmt.Printf("addOut C: %d\n", addOut.C)
}
