package main

import (
	"bufio"
	"mirea-crossplatform/exec"
	"os"
)

func main() {
	exec.StartGame()
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\v')
}
