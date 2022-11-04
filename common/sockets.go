package common

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func Receive(connection net.Conn) (string, error) {
	str, error := bufio.NewReader(connection).ReadString('\n')
	return strings.TrimSpace(str), error
}

func Send(connection net.Conn, message string) error {
	_, error := fmt.Fprintf(connection, message+"\n")
	return error
}
