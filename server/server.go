package server

import (
	//"bufio"
	"fmt"
	"net"
	"os"
	//"strings"
	"os/signal"
    "syscall"
)

const (
	HOST                  = "localhost"
	PORT                  = "9988"
	TYPE                  = "tcp"
)

func Start() {
	
	server, error := net.Listen(TYPE, HOST+":"+PORT)
	// there are no exceptions we handle error with if (hay panic tambien investigar)
	fmt.Println("Listening on " + HOST + ":" + PORT)
	fmt.Println("Waiting for client...")

	if error != nil {
		fmt.Println("Error listening:", error.Error())
		os.Exit(1)
	}
	// delay the execution of the function or method or an anonymous method until the nearby functions returns.
	acceptor := Acceptor{listenerSocket: server}
	go stop(acceptor)
	start_receiver(acceptor)

	
}

func stop(acceptor Acceptor) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	stop_receiver(acceptor)
}
