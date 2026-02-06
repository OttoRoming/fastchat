package fcprotocol

import (
	"net"
)

func writeBytes(conn net.Conn, buffer []byte) error {
	bytesWritten := 0

	for bytesWritten < len(buffer) {
		n, err := conn.Write(buffer[bytesWritten:])
		if err != nil {
			return err
		}

		bytesWritten += n
	}

	return nil
}

func readBytes(conn net.Conn, count int) ([]byte, error) {
	buffer := make([]byte, count)
	bytesRead := 0

	for bytesRead != count {
		header := make([]byte, count)
		n, err := conn.Read(header[bytesRead:])
		if err != nil {
			return header, err
		}

		bytesRead += n
	}

	return buffer, nil
}
