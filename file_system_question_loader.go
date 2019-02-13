package quizgame

import (
	"encoding/csv"
	"os"
)

// FileSystemQuestionLoader loads the questions from a file, using the filename
type FileSystemQuestionLoader struct{}

// Load takes a filename and loads the questions from the file
func (loader FileSystemQuestionLoader) Load(filename string) []Question {
	file, _ := os.Open(filename)
	defer file.Close()

	reader := csv.NewReader(file)
	questionsStrings, _ := reader.ReadAll()

	questions := make([]Question, 0)
	for _, questionString := range questionsStrings {
		questions = append(questions, Question{
			Question: questionString[0],
			Answer:   questionString[1],
		})
	}

	return questions
}
