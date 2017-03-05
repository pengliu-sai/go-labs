package main

import (
	"fmt"
	"os"

	"easyapi"
	"easyapi/easyapi_toy/services/service1"
)

func main() {
	easyapi := easyapi.NewEasyAPI()
	easyapi.RegisterService(service1.New())

	l, err := easyapi.Listen("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}

	defer l.Close()
	easyapi.Serve(l)
}
