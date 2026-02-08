package main

import (
	"fmt"
	"math"
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
	value := ""
	// fcmul.Unmarshal(`"test"`, &value)
	fmt.Printf("value: %v\n", value)

	fmt.Printf("math.MaxInt64: %s\n", fmt.Sprint(math.MaxInt64))

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
