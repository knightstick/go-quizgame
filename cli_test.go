package quizgame_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/knightstick/quizgame"
)

func TestCLI(t *testing.T) {
	t.Run("Scores 0 of 0 when no questions", func(t *testing.T) {
		questionFile, removeFile := createTempFile(t, "")
		defer removeFile()

		in := strings.NewReader("\n")
		out := &bytes.Buffer{}
		loader := &StubQuestionLoader{questions: []quizgame.Question{}}
		game := &StubGame{}

		cli := &quizgame.CLI{QuestionLoader: loader, In: in, Out: out, Game: game}
		cli.Run(questionFile.Name())

		assertLoadedFile(t, loader, questionFile.Name())
		assertGamePlayed(t, game, []quizgame.Question{})
		assertOutput(t, out, "You scored 0 out of 0\n")
	})

	t.Run("Can score 2 out of 3", func(t *testing.T) {
		questionFile, removeFile := createTempFile(t, "question,answer\n1+1,2\n1+2,3\n1+3,4\n")
		defer removeFile()

		in := strings.NewReader("2\n3\n666\n")
		out := &bytes.Buffer{}

		questions := []quizgame.Question{
			quizgame.Question{
				Question: "1+1",
				Answer:   "2",
			},
			quizgame.Question{
				Question: "1+2",
				Answer:   "3",
			},
			quizgame.Question{
				Question: "1+3",
				Answer:   "4",
			},
		}

		loader := &StubQuestionLoader{questions: questions}
		game := &StubGame{
			score: 2,
			total: 3,
		}

		cli := &quizgame.CLI{QuestionLoader: loader, In: in, Out: out, Game: game}
		cli.Run(questionFile.Name())

		assertLoadedFile(t, loader, questionFile.Name())
		assertGamePlayed(t, game, questions)
		assertOutput(t, out, "You scored 2 out of 3\n")
	})
}
