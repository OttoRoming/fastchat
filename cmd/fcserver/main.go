package main

import (
	"fmt"
	"net"
	"time"

	"github.com/OttoRoming/fastchat/pkg/fcprotocol"
	"github.com/charmbracelet/log"
)

const (
	address = "localhost:4040"
)

var (
	startTime time.Time
	data      struct {
		Users []struct {
			Username string
			Possword string
		}

		Chats []struct {
			From    string
			To      string
			Content string
		}
	}
)

func init() {
	startTime = time.Now()
}

func getUptimeFormatted() string {
	elapsed := time.Since(startTime)
	days, hours, minutes := int(elapsed.Hours())/24, int(elapsed.Hours())%24, int(elapsed.Minutes())%60
	formattedString := fmt.Sprintf("up %d day, %d hour, %d min", days, hours, minutes)

	return formattedString
}

func handleConnection(conn net.Conn) {
	message, err := fcprotocol.ReadMessage(conn)
	if err != nil {
		log.Warn("failed to read message", "err", err)
		return
	}

	if !message.Confidential() {
		log.Info("message received", "message", message)
	}

	switch message.(type) {
	case *fcprotocol.ReqUptime:
		log.Info("message got requptime")
		response := fcprotocol.AckUptime{
			Uptime: getUptimeFormatted(),
		}

		err := fcprotocol.SendMessage(response, conn)
		if err != nil {
			log.Warn("failed to send message", "err", err)
		}
		log.Info("message sent", "message", response)
	}
}

func main() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(startTime)
	log.Info("Server started", "time (µs)", elapsed.Microseconds(), "address", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Warn("connection failed", "err", err)
			continue
		}

		go handleConnection(conn)
	}

	// err = listener.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
