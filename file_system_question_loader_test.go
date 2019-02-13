package quizgame_test

import (
	"reflect"
	"testing"

	"github.com/knightstick/quizgame"
)

func TestFileSystemQuestionLoader(t *testing.T) {
	t.Run("loads 0 questions from empty file", func(t *testing.T) {
		questionFile, removeFile := quizgame.CreateTempFile(t, "")
		defer removeFile()

		loader := &quizgame.FileSystemQuestionLoader{}
		questions := loader.Load(questionFile.Name())

		expected := []quizgame.Question{}
		actual := questions

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected to load questions %v, got %v", expected, actual)
		}
	})

	t.Run("loads 3 questions from question file", func(t *testing.T) {
		questionFile, removeFile := quizgame.CreateTempFile(t, "1+1,2\n1+2,3\n1+3,4\n")
		defer removeFile()

		loader := &quizgame.FileSystemQuestionLoader{}
		questions := loader.Load(questionFile.Name())

		expected := []quizgame.Question{
			{Question: "1+1", Answer: "2"},
			{Question: "1+2", Answer: "3"},
			{Question: "1+3", Answer: "4"},
		}
		actual := questions

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected to load questions %v, got %v", expected, actual)
		}
	})
}
