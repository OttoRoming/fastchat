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
	defer conn.Close()

	msg, err := fcprotocol.ReadMessage(conn)
	if err != nil {
		log.Warn("failed to read message", "err", err)
		return
	}

	if !msg.Confidential() {
		log.Info("message received", "msg", msg)
	}

	switch msg.(type) {
	case fcprotocol.ReqUptime:
		response := fcprotocol.AckUptime{
			Uptime: getUptimeFormatted(),
		}

		fcprotocol.SendMessage(response, conn)
	}
}

func main() {
	loadDB()
	defer db.Close()

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(startTime)
	log.Info("Server started in", "microseconds", elapsed.Microseconds())
	log.Info("Listening for fcprotocol messages on", "address", address)

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
