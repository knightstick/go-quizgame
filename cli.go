package quizgame

import "io"

// QuestionLoader takes a filename and can load a list of questions
type QuestionLoader interface {
	Load(filename string)
}

// CLI is the command line interface to the quizgame
type CLI struct {
	QuestionLoader QuestionLoader
	In             io.Reader
	Out            io.Writer
}

// NewCLI creates a new CLI to play the quiz game
func NewCLI(in io.Reader, out io.Writer) *CLI {
	loader := &FileSystemQuestionLoader{}
	return &CLI{QuestionLoader: loader}
}

// Run runs the whole game
func (cli CLI) Run(filename string) {
	cli.QuestionLoader.Load(filename)
}
