package quizgame

import "io"

// QuestionLoader takes a filename and can load a list of questions
type QuestionLoader interface {
	Load(filename string)
}

// CLI is the command line interface to the quizgame
type CLI struct {
	filename       string
	questionLoader QuestionLoader
}

// NewCLI creates a new CLI to play the quiz game
func NewCLI(filename string, in io.Reader, out io.Writer) *CLI {
	loader := &FileSystemQuestionLoader{}
	return NewCLIWithLoader(filename, in, out, loader)
}

// NewCLIWithLoader allows creating a CLI with a specific implementation of
// QuestionLoader
func NewCLIWithLoader(filename string, in io.Reader, out io.Writer, loader QuestionLoader) *CLI {
	return &CLI{
		filename:       filename,
		questionLoader: loader,
	}
}

// Play runs the whole game
func (cli CLI) Play() {
	cli.questionLoader.Load(cli.filename)
}
