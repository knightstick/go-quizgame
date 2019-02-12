package quizgame

// FileSystemQuestionLoader loads the questions from a file, using the filename
type FileSystemQuestionLoader struct{}

// Load takes a filename and loads the questions from the file
func (loader FileSystemQuestionLoader) Load(filename string) []Question {
	return make([]Question, 0)
}
