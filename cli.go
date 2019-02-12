package quizgame

import (
	"fmt"
	"io"
)

// Question is the basic quiz question the user has to answer
type Question struct {
	Question string
	Answer   string
}

// QuestionLoader takes a filename and can load a list of questions
type QuestionLoader interface {
	Load(filename string) []Question
}

// Game allows the user to play and answer the questions
type Game interface {
	Play([]Question)
	Score() int
	NumberOfQuestions() int
}

// CLI is the command line interface to the quizgame
type CLI struct {
	QuestionLoader QuestionLoader
	In             io.Reader
	Out            io.Writer
	Game           Game
}

// NewCLI creates a new CLI to play the quiz game
func NewCLI(in io.Reader, out io.Writer) *CLI {
	loader := &FileSystemQuestionLoader{}
	game := &QuizGame{}
	return &CLI{
		QuestionLoader: loader,
		Game:           game,
		In:             in,
		Out:            out,
	}
}

// Run runs the whole game
func (cli CLI) Run(filename string) {
	questions := cli.QuestionLoader.Load(filename)
	cli.Game.Play(questions)

	score := cli.Game.Score()
	total := cli.Game.NumberOfQuestions()
	cli.Out.Write([]byte(fmt.Sprintf("You scored %d out of %d\n", score, total)))
}
