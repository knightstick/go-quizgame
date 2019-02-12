package main

import (
	"flag"
	"os"

	quizgame "github.com/knightstick/quizgame"
)

func main() {
	filenamePtr := flag.String("file", "problems.csv", "path for the csv file with the questions and answers")

	cli := quizgame.NewCLI(os.Stdin, os.Stdout)
	cli.Run(*filenamePtr)
}
