package exec

var InputChannel chan string
var ExitGame chan struct{}

func StartGame() {
	InputChannel = make(chan string)
	ExitGame = make(chan struct{})
	Run()
	//Ожидаем выхода из игры
	<-ExitGame
}
