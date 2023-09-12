//go:build !(js && wasm)

package exec

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

func Run() {
	go game.Start(proceedCliOutput, InputChannel, ExitGame)
	reader := bufio.NewReader(os.Stdin)
	go func() {
		for {
			readString, _ := reader.ReadString('\n')
			readString = strings.TrimRight(readString, "\n")
			readString = strings.TrimRight(readString, "\r")
			InputChannel <- readString
		}
	}()
}
