package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	HOST                  = "localhost"
	PORT                  = "9979"
	TYPE                  = "tcp"
	ShutdownServerCommand = "Q"
)

func Start() {
	fmt.Println("Server Running...")

	server, error := net.Listen(TYPE, HOST+":"+PORT)
	// there are no exceptions we handle error with if (hay panic tambien investigar)
	fmt.Println("Listening on " + HOST + ":" + PORT)
	fmt.Println("Waiting for client...")

	if error != nil {
		fmt.Println("Error listening:", error.Error())
		os.Exit(1)
	}
	// delay the execution of the function or method or an anonymous method until the nearby functions returns.
	acceptor := Acceptor{Listener: server}
	go stop(acceptor)
	start_receiver(acceptor)

}

func stop(acceptor Acceptor) {

	var quit bool
	for !quit {
		/* process info until someone enters exit */
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		value := strings.TrimSpace(input)
		if value == ShutdownServerCommand {
			/* close loop */
			quit = true
			stop_receiver(acceptor)
		}
	}
}
