package quizgame_test

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/knightstick/quizgame"
)

type stubQuizTimer struct {
	sleepCalls int
}

func (timer *stubQuizTimer) Sleep() {
	timer.sleepCalls++
}

type slowInput struct{}

func (in *slowInput) Read(p []byte) (int, error) {
	time.Sleep(1 * time.Second)

	return 8, nil
}

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

func TestQuizTimer(t *testing.T) {
	t.Run("times out if waiting too long", func(t *testing.T) {
		timer := &stubQuizTimer{}
		in := &slowInput{}
		questions := []quizgame.Question{
			quizgame.Question{Question: "1+1", Answer: "2"},
		}
		game := quizgame.NewQuizGame(in, &bytes.Buffer{}, questions, timer)
		game.Play()

		if timer.sleepCalls != 1 {
			t.Error("expected to call the timer")
		}
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

func TestQuizGameOutput(t *testing.T) {
	t.Run("3 questions get printed", func(t *testing.T) {
		questions := []quizgame.Question{
			{Question: "1+1", Answer: "2"},
			{Question: "Meaning of life", Answer: "42"},
			{Question: "1+2", Answer: "3"},
		}

		out := &bytes.Buffer{}

		game := quizgame.NewQuizGame(strings.NewReader("\n\n\n"), out, questions, QuickTimer)
		game.Play()

		expectedOutput := "Problem #1: 1+1 = Problem #2: Meaning of life = Problem #3: 1+2 = "
		quizgame.AssertOutput(t, out, expectedOutput)
	})
}

func assertScore(t *testing.T, questions []quizgame.Question, answers string, score int) {
	t.Helper()

	game := quizgame.NewQuizGame(strings.NewReader(answers), &bytes.Buffer{}, questions, QuickTimer)
	game.Play()

	actualScore := game.Score()

	if actualScore != score {
		t.Errorf("expected score of %d, got %d", score, actualScore)
	}
}

func assertNumberOfQuestions(t *testing.T, questions []quizgame.Question, total int) {
	t.Helper()

	game := quizgame.NewQuizGame(strings.NewReader(""), &bytes.Buffer{}, questions, QuickTimer)

	actualNumberOfQuestions := game.NumberOfQuestions()

	if actualNumberOfQuestions != total {
		t.Errorf("expected number of questions to be %d, got %d", actualNumberOfQuestions, total)
	}
}
