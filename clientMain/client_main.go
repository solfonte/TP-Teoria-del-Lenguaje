package main

import (
	"truco/app/client"
	"truco/app/common"
)

func main() {
	common.CallClearScreen()
	client.Start()
}
