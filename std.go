package glnassistant

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/fatih/color"
)

func Stderr(text string) {
	fmt.Fprintf(os.Stderr, "[%s] %s\n", color.HiRedString("error"), text)
}

func Stdout(inBracket, text string) {

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Define a list of terminal colors from the fatih/color package
	colors := []*color.Color{
		// color.New(color.FgRed),
		color.New(color.FgGreen),
		color.New(color.FgYellow),
		color.New(color.FgBlue),
		color.New(color.FgMagenta),
		color.New(color.FgCyan),
		color.New(color.FgWhite),
	}

	// Select a random color
	randomColor := colors[rand.Intn(len(colors))]

	if len(inBracket) > 0 {
		fmt.Printf("[%s] %s\n", randomColor.Sprint(inBracket), text)
	} else {
		randomColor.Println(text)
	}

}
