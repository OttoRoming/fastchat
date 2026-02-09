package main

import (
	"fmt"

	"github.com/OttoRoming/fastchat/pkg/fcmul"
)

const (
// address = "localhost:4040"
)

// func handleConnection(conn net.Conn) {
// 	defer conn.Close()

// 	msg, err := fcprotocol.Parse(conn)
// 	if err != nil {
// 		return
// 	}

// 	fmt.Printf("msg: %v\n", msg)
// }

func main() {
	data, err := fcmul.Parse(`{"hello" -> "world"}`)
	if err != nil {
		panic(err)
	}

	fmt.Printf("data: %v\n", data)

	// listener, err := net.Listen("tcp", address)
	// if err != nil {
	// 	panic(err)
	// }

	// defer listener.Close()

	// fmt.Printf("Listening on %s\n", address)

	// for {
	// 	conn, err := listener.Accept()
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	go handleConnection(conn)
	// }
}
