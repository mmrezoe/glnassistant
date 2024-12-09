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
	fmt.Fprintf(os.Stderr, "[%s] %s\n", color.New(color.FgRed, color.Bold).Sprint("error"), text)
}

func Stdout(text string) {

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Define a list of terminal colors from the fatih/color package
	colors := []*color.Color{
		color.New(color.FgRed),
		color.New(color.FgGreen),
		color.New(color.FgYellow),
		color.New(color.FgBlue),
		color.New(color.FgMagenta),
		color.New(color.FgCyan),
		color.New(color.FgWhite),
	}

	// Select a random color
	randomColor := colors[rand.Intn(len(colors))]

	// Print a message with the random color
	randomColor.Println(text)

	// Print multiple messages with random colors
	for i := 0; i < 5; i++ {
		randomColor := colors[rand.Intn(len(colors))]
		randomColor.Printf("Message %d with random color\n", i+1)
	}

}
