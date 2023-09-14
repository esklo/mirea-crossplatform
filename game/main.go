package game

import (
	"fmt"
	"math/rand"
	"strings"
)

var outputHandler func(interface{})
var randomFunc func() int

func Start(proceedOutput func(interface{}), inputChannel chan string, exitGame chan struct{}, rndFunc func() int) {
	outputHandler = proceedOutput
	if rndFunc != nil {
		randomFunc = rndFunc
	} else {
		randomFunc = func() int {
			return rand.Intn(10) + 1
		}
	}

	printHeader()

	for {
		a, b, c, d, e, f := generatePlot()
		outputHandler(fmt.Sprintf("\nTHE %s IS ABOUT A %s %s WHO IS %s AT\n", a, b, c, d))
		outputHandler(fmt.Sprintf("%s AND WHO %s.\n\n", e, f))

		outputHandler("\nANOTHER (YES OR NO)? ")
		if <-inputChannel == "NO" {
			break
		}
	}

	outputHandler("\nO.K.  HOPE YOU HAVE A SUCCESSFUL TV SHOW!!\n")
	close(exitGame)
}

func printHeader() {
	outputHandler(strings.Repeat(" ", 26) + "TVPLOT\n")
	outputHandler(strings.Repeat(" ", 20) + "CREATIVE COMPUTING\n")
	outputHandler(strings.Repeat(" ", 18) + "MORRISTOWN, NEW JERSEY\n\n\n\n")
	outputHandler("THIS PROGRAM AUTOMATICALLY COMES UP WITH TELEVISION\n")
	outputHandler("SHOWS GUARANTEED TO APPEAL TO THE MASSES AND WIN\n")
	outputHandler("HIGH NEILSEN RATINGS.\n\n")
	outputHandler("HERE IS THE FIRST PLOT:\n")
}

func generatePlot() (string, string, string, string, string, string) {
	plots := [][10]string{
		{"PROGRAM", "REPORT", "SPECIAL", "SERIES", "STORY", "PROGRAM", "REPORT", "SPECIAL", "SERIES", "STORY"},
		{"SWINGING", "BRILLIANT", "SALTY", "HILARIOUS", "SENSITIVE", "DODDERING", "HENPECKED", "DEDICATED", "THOUGHTFUL", "HEAVY"},
		{"GIRL COWHAND", "LITTLE BOY", "SCIENTEST", "LAWYER", "TOWN MARSHALL", "DENTIST", "BUS DRIVER", "JUNGLE MAN", "SECRET AGENT", "COLLIE"},
		{"A WHIZ", "A FLOP", "MEDIOCRE", "A SUCCESS", "A DISASTER", "A WHIZ", "A FLOP", "MEDIOCRE", "A SUCCESS", "A DISASTER"},
		{"SOLVING CRIMES", "ROPING COWS", "COOKING HEALTH FOOD", "PITCHING WOO", "PROTECTING ECOLOGY", "HELPING CHILDREN", "TWO-FISTED DRINKING", "FIGHTING FIRES", "HERDING ELEPHANTS", "WINNING RACES"},
		{"RECOVERS THE JEWELS", "FOILS THE SPIES", "DESTROYS THE CITY", "FINDS LOVE", "SAVES THE ANIMALS", "CONFESSES", "DISCOVERS THE SECRET", "STOPS THE FLOOD", "HELPS THE DOG", "MAKES THE SACRIFICE"},
	}

	values := []string{}
	for _, plot := range plots {
		values = append(values, plot[randomFunc()-1])
	}

	return values[0], values[1], values[2], values[3], values[4], values[5]
}
