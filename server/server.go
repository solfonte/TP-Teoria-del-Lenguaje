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
	PORT                  = "9988"
	TYPE                  = "tcp"
	ShutdownServerCommand = "Q"
)

func start() {
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
	receiver := Receiver{listenerSocket: server}
	go stop()
	go start_receiver(receiver)

}

func stop() {
	var quit bool
	for !quit {
		/* process info until someone enters exit */
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		value := strings.TrimSpace(input)
		if value == ShutdownServerCommand {
			/* close loop */
			fmt.Println("entre a cerrar")
			quit = true
		}
	}
}
