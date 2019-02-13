package main

import (
	"flag"
	"os"

	quizgame "github.com/knightstick/quizgame"
)

func main() {
	filenamePtr := flag.String("file", "problems.csv", "path for the csv file with the questions and answers")
	flag.Parse()

	cli := quizgame.NewCLI(os.Stdin, os.Stdout, &quizgame.FileSystemQuestionLoader{}, *filenamePtr)
	cli.Run()
}
