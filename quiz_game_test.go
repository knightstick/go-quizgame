package quizgame_test

import (
	"strings"
	"testing"

	"github.com/knightstick/quizgame"
)

func TestQuizGameScore(t *testing.T) {
	t.Run("base case game with no questions", func(t *testing.T) {
		questions := []quizgame.Question{}

		assertScore(t, questions, "\n", 0)
	})

	t.Run("perfect score on 3 question", func(t *testing.T) {
		questions := []quizgame.Question{
			{Question: "1+1", Answer: "2"},
			{Question: "1+1", Answer: "2"},
			{Question: "1+1", Answer: "2"},
		}
		answers := "2\n2\n2\n"

		assertScore(t, questions, answers, 3)
	})

	t.Run("one correct answer in 3", func(t *testing.T) {
		questions := []quizgame.Question{
			{Question: "1+1", Answer: "2"},
			{Question: "Meaning of life", Answer: "42"},
			{Question: "1+2", Answer: "3"},
		}
		answers := "hello\n42\ni don't know\n"

		assertScore(t, questions, answers, 1)
	})
}

func TestQuizGameNumberOfQuestions(t *testing.T) {
	t.Run("no questions have total 0", func(t *testing.T) {
		questions := []quizgame.Question{}

		assertNumberOfQuestions(t, questions, 0)
	})

	t.Run("3 questions have total 3", func(t *testing.T) {
		questions := []quizgame.Question{
			{Question: "1+1", Answer: "2"},
			{Question: "Meaning of life", Answer: "42"},
			{Question: "1+2", Answer: "3"},
		}

		assertNumberOfQuestions(t, questions, 3)
	})
}

func assertScore(t *testing.T, questions []quizgame.Question, answers string, score int) {
	t.Helper()

	game := quizgame.NewQuizGame(strings.NewReader(answers), questions)
	game.Play()

	actualScore := game.Score()

	if actualScore != score {
		t.Errorf("expected score of %d, got %d", score, actualScore)
	}
}

func assertNumberOfQuestions(t *testing.T, questions []quizgame.Question, total int) {
	t.Helper()

	game := quizgame.NewQuizGame(strings.NewReader(""), questions)

	actualNumberOfQuestions := game.NumberOfQuestions()

	if actualNumberOfQuestions != total {
		t.Errorf("expected number of questions to be %d, got %d", actualNumberOfQuestions, total)
	}
}
