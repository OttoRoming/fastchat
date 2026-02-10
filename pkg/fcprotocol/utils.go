package fcprotocol

import (
	"net"
)

// writeBytes writes the entire buffer to the connection.
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

// readBytes reads the specified number of bytes from the connection.
func readBytes(conn net.Conn, count int) ([]byte, error) {
	buffer := make([]byte, count)
	bytesRead := 0

	for bytesRead < count {
		n, err := conn.Read(buffer[bytesRead:])
		if err != nil {
			return buffer, err
		}

		bytesRead += n
	}

	return buffer, nil
}
