package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"sync"

	"github.com/funny/binary"
	"github.com/golang/protobuf/proto"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
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
	for i := 10; i > 0; i-- {
		addIn := &AddIn{
			A: 1,
			B: 2,
		}

		data, err := proto.Marshal(addIn)
		if err != nil {
			fmt.Println("Error to marshal message because of ", err.Error())
		}

		var buff = binary.Buffer{
			Data: make([]byte, len(data)+4),
		}
		buff.WriteUint32LE(uint32(len(data)))
		buff.WriteBytes(data)

		_, err = conn.Write(buff.Data)

		if err != nil {
			fmt.Println("Error to send message because of ", err.Error())
			break
		}
	}
}

func handleRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		var headBuf = make([]byte, 4)
		if _, err := io.ReadFull(conn, headBuf); err != nil {
			return
		}

		packageSize := int(binary.GetUint32LE(headBuf))
		packet := make([]byte, packageSize)

		fmt.Println("read message len: ", packageSize)

		if _, err := io.ReadFull(conn, packet); err != nil {
			fmt.Println("Error to read message because of ", err.Error())
		}

		addIn := &AddIn{}
		err := proto.Unmarshal(packet, addIn)
		if err != nil {
			fmt.Println("Error to unmarshal message because of ", err.Error())
		}
		fmt.Printf("addIn A: %d, B: %d\n", addIn.A, addIn.B)
	}
}
