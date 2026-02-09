/*
Fast Chat Protocol

Provides utilities for creating, sending, and receiving messages
*/
package fcprotocol

import (
	"encoding/binary"
	"fmt"
	"math"
	"net"

	"github.com/OttoRoming/fastchat/pkg/fcmul"
)

const (
	headerLength   = 8
	correctVersion = 1
)

type packet struct {
	Method uint16
	Body   string
}

func packageMessage(msg Message) (packet, error) {
	body, err := fcmul.Marshal(msg)
	if err != nil {
		return packet{}, err
	}

	packet := packet{
		Method: msg.method(),
		Body:   body,
	}

	err = packet.validate()
	if err != nil {
		return packet, err
	}

	return packet, nil
}

func (msg *packet) validate() error {
	if len(msg.Body) > math.MaxUint32 {
		return fmt.Errorf("body can not be larger than %d", math.MaxUint32)
	}

	return nil
}

func (msg *packet) send(conn net.Conn) error {
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

func readPacket(conn net.Conn) (packet, error) {
	var result packet

	header, err := readBytes(conn, headerLength)
	if err != nil {
		return result, err
	}

	version := binary.BigEndian.Uint16(header[:2])
	if version != correctVersion {
		return result, fmt.Errorf("unsupported version %d", version)
	}

	result.Method = binary.BigEndian.Uint16(header[2:4])
	if result.Method >= methodLimit {
		return result, fmt.Errorf("invalid method %d", result.Method)
	}

	bodyLength := binary.BigEndian.Uint32(header[4:8])
	bodyBytes, err := readBytes(conn, int(bodyLength))
	if err != nil {
		return result, err
	}

	result.Body = string(bodyBytes)

	return result, nil
}

func SendMessage(msg Message, conn net.Conn) error {
	packet, err := packageMessage(msg)
	if err != nil {
		return err
	}

	err = packet.send(conn)
	if err != nil {
		return err
	}

	return nil
}
