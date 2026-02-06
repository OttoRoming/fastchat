/*
Fast Chat Protocol

Provides utilities for creating, sending, and receiving messages
*/
package fcprotocol

import (
	"encoding/binary"
	"fmt"
	"net"
	"math"
)

const (
	MethodRequestSignUp uint16 = iota
	methodLimit
)

const (
	headerLength = 8
	correctVersion = 1
)

type Message struct {
	Method uint16
	Body string
}

func NewMessage(method uint16, body string) (Message, error) {
	msg := Message{
		Method: method,
		Body: body,
	}

	err := msg.validate()
	if err != nil {
		return msg, err
	}

	return msg, nil
}

func (msg *Message)validate() error {
	if (len(msg.Body) > math.MaxUint32) {
		return fmt.Errorf("body can not be larger than %d", math.MaxUint32)
	}

	return nil
}

func (msg *Message)Send(conn net.Conn) error {
	header := make([]byte, headerLength)
	binary.BigEndian.PutUint16(header, correctVersion)
	binary.BigEndian.PutUint16(header[2:], msg.Method)
	binary.BigEndian.PutUint32(header[4:], uint32(len(msg.Body)))

	err := writeBytes(conn, header)
	if err != nil {
		return err
	}

	body := []byte(msg.Body)
	err = writeBytes(conn, body)
	if err != nil {
		return err
	}

	return nil
}


func Parse(conn net.Conn) (Message, error) {
	var result Message

	header, err := readBytes(conn, headerLength)
	if err != nil {
		return result, err
	}

	version := binary.BigEndian.Uint16(header[:2])
	if (version != correctVersion) {
		return result, fmt.Errorf("unsupported version %d", version)
	}

	result.Method = binary.BigEndian.Uint16(header[2:4])
	if result.Method >= methodLimit {
		return result, fmt.Errorf("invalid method %d", result.Method)
	}

	bodyLength := binary.BigEndian.Uint32(header[4:8])
	bodyBytes, err := readBytes(conn, int(bodyLength))
	result.Body = string(bodyBytes)

	return result, nil
}
