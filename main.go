package main

var inputChannel chan string
var exitGame chan struct{}

func main() {
	// Каналы для ввода и выхода из игры
	inputChannel = make(chan string)
	exitGame = make(chan struct{})
	run()
	//Ожидаем выхода из игры
	<-exitGame
}
