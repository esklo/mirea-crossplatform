package exec

var InputChannel chan string
var ExitGame chan struct{}

func StartGame() {
	// this channel is used to transmit user input
	InputChannel = make(chan string)
	// this channel will close as soon as the game is over
	ExitGame = make(chan struct{})

	// either cli.go or web.go (defined during build)
	Run()

	// waiting for the end of the game
	<-ExitGame
}
