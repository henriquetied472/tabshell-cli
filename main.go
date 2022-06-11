package main

import (
	"flag"
	"log"
	"os"

	"github.com/henriquetied472/tabshell-cli/questions"
)

var debugger log.Logger
var logger log.Logger

func main() {
	dgb := flag.Bool("debug", false, "Append the log file")
	flag.Parse()

	out, _ := os.OpenFile("./logs/latest.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer out.Close()
	debugger = *log.Default()
	logger = *log.Default()
	debugger.SetOutput(out)

	debugger.SetPrefix("DEBUG: ")
	debugger.Printf("TabShell Initialized!!")
	logger.Printf("Tabshell Initialized!!")

	defer logger.Printf("TabShell Finished!!")

	questions.Init(&debugger, &logger, dgb)
}
