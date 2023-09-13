package main

import (
	"bufio"
	"mirea-crossplatform/exec"
	"os"
)

func main() {
	exec.StartGame()

	// infinite waiting for input (to avoid kernel panic)
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\v')
}
