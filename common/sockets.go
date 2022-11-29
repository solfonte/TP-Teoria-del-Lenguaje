package common

import (
	"net"
	"strings"
	"time"
)

func Set_deadline(connection net.Conn) {
	connection.SetReadDeadline(time.Now().Add(1 * time.Second))
}

func Receive(connection net.Conn) (string, error) {

	buffer := make([]byte, 4096)
	mLen, err := connection.Read(buffer)

	msg := ""
	if mLen > 0 && err == nil {
		msg = string(buffer[:mLen])
	}

	return msg, err
}

func Send(connection net.Conn, message string) error {
	_, err := connection.Write([]byte(strings.TrimRight(message, "\n")))
	return err
}
