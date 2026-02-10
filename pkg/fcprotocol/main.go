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

// packageMessage packages a message into a packet
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

// validate validates a pacekt as being correct
func (msg *packet) validate() error {
	if len(msg.Body) > math.MaxUint32 {
		return fmt.Errorf("body can not be larger than %d", math.MaxUint32)
	}

	return nil
}

// send sends the packet on the connection
func (msg *packet) send(conn net.Conn) error {
	header := make([]byte, headerLength)
	binary.BigEndian.PutUint16(header[0:2], correctVersion)
	binary.BigEndian.PutUint16(header[2:4], msg.Method)
	binary.BigEndian.PutUint32(header[4:8], uint32(len(msg.Body)))

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

// readPacket reads a packet from the connection
func readPacket(conn net.Conn) (packet, error) {
	var result packet

	header, err := readBytes(conn, headerLength)
	if err != nil {
		return result, err
	}

	version := binary.BigEndian.Uint16(header[0:2])
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

	err = result.validate()
	if err != nil {
		return result, err
	}

	return result, nil
}

// ReadMessage reads a message from the connection
func ReadMessage(conn net.Conn) (Message, error) {
	packet, err := readPacket(conn)
	if err != nil {
		return nil, err
	}

	var result Message

	switch packet.Method {
	case methodReqUptime:
		result = &ReqUptime{}
	case methodAckUptime:
		result = &AckUptime{}
	case methodReqSignUp:
		result = &ReqSignUp{}
	case methodReqLogIn:
		result = &ReqLogin{}
	case methodAckSignedin:
		result = &AckSignedIn{}
	case methodErrUsernameInUse:
		result = ErrUsernameInUse{}
	case methodReqSendChat:
		result = &ReqSendMessage{}
	case methodAckChatSent:
		result = &AckMessageSent{}
	case methodReqGetHistory:
		result = &ReqGetHistory{}
	case methodAckHistory:
		result = &AckHistory{}
	case methodErrAccountNotInUse:
		result = ErrAccountNotInUse{}
	case methodErrFailedRead:
		result = &ErrFailedRead{}
	default:
		return nil, fmt.Errorf("unsupported method: %d", packet.Method)
	}

	if packet.Body == "" {
		return nil, fmt.Errorf("empty body")
	}

	err = fcmul.Unmarshal(packet.Body, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// SendMessage sends a message on the connection
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

func
kkkkkkkkkkk
