//go:build js && wasm

package main

import (
	"mirea-crossplatform/game"
	"slices"
	"syscall/js"
)

var webOutputElement js.Value
var webInput string

var stopKeys = []string{"Meta", "Shift", "Tab", "CapsLock", "Alt", "Control"}

func proceedWebOutput(text interface{}) {
	currentText := webOutputElement.Get("textContent").String()
	newText := currentText + text.(string)
	webOutputElement.Set("textContent", newText)
}

func proceedWebBackspace() {
	currentText := webOutputElement.Get("textContent").String()
	newText := currentText
	if len(currentText) > 0 {
		newText = currentText[:len(currentText)-1]
	}
	webOutputElement.Set("textContent", newText)
}

func proceedWebInput(_ js.Value, args []js.Value) interface{} {
	key := args[0].String()
	if slices.Contains(stopKeys, key) {
		return nil
	}
	if key == "Backspace" {
		if len(webInput) > 0 {
			webInput = webInput[:len(webInput)-1]
			proceedWebBackspace()
		}
		return nil
	}

	if key == "Enter" {
		proceedWebOutput("\n")
		inputChannel <- webInput
		webInput = ""
		return nil
	}
	webInput += key
	proceedWebOutput(key)
	return nil
}

func run() {

	//Определяем куда будет выводиться webOutputElement
	webOutputElement = js.Global().Get("document").Call("getElementById", "root")

	//Регистрируем функцию, которую будем вызывать из JS для передачи данных в go
	js.Global().Set("sendInput", js.FuncOf(proceedWebInput))

	//Запускаем игру
	go game.Start(proceedWebOutput, inputChannel, exitGame)
}
