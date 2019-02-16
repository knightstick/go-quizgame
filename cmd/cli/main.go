package main

import (
	"flag"
	"os"
	"time"

	quizgame "github.com/knightstick/quizgame"
)

const defaultProblemsFile = "problems.csv"
const defaultGameLengthSeconds = 30

func main() {
	filenamePtr := flag.String("file", defaultProblemsFile, "path for the csv file with the questions and answers")
	gameLengthSeconds := flag.Float64("timelimit", defaultGameLengthSeconds, "time limit, in seconds, for the game")
	flag.Parse()

	gameLength := time.Duration(*gameLengthSeconds) * time.Second

	timer := quizgame.WaitQuizTimer{SleepTime: gameLength}
	cli := quizgame.NewCLI(os.Stdin, os.Stdout, &quizgame.FileSystemQuestionLoader{}, timer, *filenamePtr)
	cli.Run()
}
