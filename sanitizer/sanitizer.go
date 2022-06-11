package sanitizer

import (
	"bufio"
	"fmt"
	"log"
	"os"

	markdown "github.com/MichaelMure/go-term-markdown"
	"golang.org/x/term"
)

func Init(debugger, logger *log.Logger, dgb *bool, body string) {
	w, _, _ := term.GetSize(int(os.Stdin.Fd()))

	fmt.Print(string(markdown.Render(body, w, 0)))
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
