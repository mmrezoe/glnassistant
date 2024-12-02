package go_assistant

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func Stderr(mode string, text string) {
	if mode == "error" || mode == "err" {
		fmt.Fprintf(os.Stderr, "[%s] %s\n", color.HiRedString(mode), text)
	} else {
		fmt.Fprintf(os.Stderr, "[%s] %s\n", color.HiCyanString(mode), text)
	}
}
