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
	cmd := exec.Command("vintbas", "tvplot.bas")
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

	pipe.Close()
	expectedOutput, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Ошибка выполнения Basic программы: %v", err)
	}

	inputChannel := make(chan string)
	exitChannel := make(chan struct{})
	go Start(proceedOutput, inputChannel, exitChannel)
	inputChannel <- "YES"
	inputChannel <- "NO"
	<-exitChannel

	if string(expectedOutput) != output {
		t.Errorf("Вывод Basic программы не соответствует ожидаемому результату.\nОжидаемо: %s\nФактически: %s", expectedOutput, string(output))
	}
}
