//go:build !(js && wasm)

package main

import (
	"bufio"
	"fmt"
	"mirea-crossplatform/game"
	"os"
	"strings"
)

func proceedCliOutput(text interface{}) {
	fmt.Print(text)
}

func run() {
	go game.Start(proceedCliOutput, inputChannel, exitGame)
	reader := bufio.NewReader(os.Stdin)
	go func() {
		for {
			readString, _ := reader.ReadString('\n')
			readString = strings.TrimRight(readString, "\n")
			readString = strings.TrimRight(readString, "\r")
			inputChannel <- readString
		}
	}()
}
