package quizgame_test

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/knightstick/quizgame"
)

type StubGame struct {
	playCalls int
	questions []quizgame.Question
	score     int
	total     int
}

func (game *StubGame) Play() {
	game.playCalls = game.playCalls + 1
}

func (game *StubGame) Score() int {
	return game.score
}

func (game *StubGame) NumberOfQuestions() int {
	return game.total
}

func TestCLI(t *testing.T) {
	t.Run("Scores 0 of 0 when no questions", func(t *testing.T) {
		in := strings.NewReader("\n")
		out := &bytes.Buffer{}

		game := &StubGame{questions: []quizgame.Question{}}

		cli := &quizgame.CLI{In: in, Out: out, Game: game}
		cli.Run()

		assertGamePlayed(t, game, []quizgame.Question{})
		quizgame.AssertOutput(t, out, "You scored 0 out of 0\n")
	})

	t.Run("Can score 2 out of 3", func(t *testing.T) {
		in := strings.NewReader("2\n3\n666\n")
		out := &bytes.Buffer{}

		questions := []quizgame.Question{
			quizgame.Question{Question: "1+1", Answer: "2"},
			quizgame.Question{Question: "1+2", Answer: "3"},
			quizgame.Question{Question: "1+3", Answer: "4"},
		}

		game := &StubGame{score: 2, total: 3, questions: questions}

		cli := &quizgame.CLI{In: in, Out: out, Game: game}
		cli.Run()

		assertGamePlayed(t, game, questions)
		quizgame.AssertOutput(t, out, "You scored 2 out of 3\n")
	})
}

func assertGamePlayed(t *testing.T, game *StubGame, questions []quizgame.Question) {
	t.Helper()

	if game.playCalls != 1 {
		t.Error("did not play the game")
	}

	if !reflect.DeepEqual(game.questions, questions) {
		t.Errorf("expected questions %v, got %v", questions, game.questions)
	}
}
