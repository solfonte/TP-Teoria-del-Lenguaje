package common

import (
	"net"
	"strings"
)

func Receive(connection net.Conn) (string, error) {
	buffer := make([]byte, 4096)
	mLen, err := connection.Read(buffer)
	msg := ""
	if mLen > 0 {
		msg = string(buffer[:mLen])
	}
	return msg, err
}

func Send(connection net.Conn, message string) error {
	_, error := connection.Write([]byte(strings.TrimRight(message, "\n")))
	return error
}
