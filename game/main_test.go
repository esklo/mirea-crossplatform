package game

import (
	"io"
	"os/exec"
	"testing"
)

var output string

func proceedOutput(text interface{}) {
	output += text.(string)
}

func TestBasicCode(t *testing.T) {

	t.Cleanup(func() {
		exec.Command("sudo", "systemsetup", "-setusingnetworktime", "on").Run()
	})
	exec.Command("sudo", "systemsetup", "-setusingnetworktime", "off").Run()
	exec.Command("sudo", "date", "0415110022").Run()

	cmd := exec.Command("./vintbas", "tvplot.bas")
	pipe, err := cmd.StdinPipe()
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.WriteString(pipe, "YES\n")
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.WriteString(pipe, "NO\n")
	if err != nil {
		t.Fatal(err)
	}

	err = pipe.Close()
	if err != nil {
		t.Fatal(err)
	}
	expectedOutput, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Basic error: %v", err)
	}
	inputChannel := make(chan string)
	exitChannel := make(chan struct{})
	// values are hardcoded to match basic output for date 0415110022
	go Start(proceedOutput, inputChannel, exitChannel, []int{4, 7, 8, 3, 5, 10, 4, 5, 2, 1, 7, 1})
	inputChannel <- "YES"
	inputChannel <- "NO"
	<-exitChannel

	if string(expectedOutput) != output {
		t.Errorf("The output of the Basic program does not match the expected result.\nExpected: %s\nActually: %s", expectedOutput, output)
	}
}
