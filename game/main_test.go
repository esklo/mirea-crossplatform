package game

import (
	"fmt"
	"io"
	"math/rand"
	"os/exec"
	"testing"
)

var output string

func proceedOutput(text interface{}) {
	output += text.(string)
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func TestBasicCode(t *testing.T) {
	var input []string
	numStrings := rand.Intn(10) + 5
	for i := 0; i < numStrings; i++ {
		stringLength := rand.Intn(20) + 1
		randomString := generateRandomString(stringLength)
		input = append(input, randomString)
	}
	input = append(input, "NO")
	cmd := exec.Command("./vintbas", "tvplot.bas")
	pipe, err := cmd.StdinPipe()
	if err != nil {
		t.Fatal(err)
	}
	for _, str := range input {
		_, err = io.WriteString(pipe, fmt.Sprintf("%s\n", str))
		if err != nil {
			t.Fatal(err)
		}
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
	go Start(proceedOutput, inputChannel, exitChannel, func() int {
		return 1
	})
	for _, str := range input {
		inputChannel <- str
	}
	<-exitChannel

	if string(expectedOutput) != output {
		t.Errorf("The output of the Basic program does not match the expected result.\nExpected: %s\nActually: %s", expectedOutput, output)
	}
}
