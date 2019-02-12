package quizgame

import "io"

// Question is the basic quiz question the user has to answer
type Question struct{}

// QuestionLoader takes a filename and can load a list of questions
type QuestionLoader interface {
	Load(filename string) []Question
}

// Game allows the user to play and answer the questions
type Game interface {
	Play([]Question)
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
	cli.QuestionLoader.Load(filename)
	cli.Game.Play([]Question{})
	cli.Out.Write([]byte("You scored 0 out of 0\n"))
}
