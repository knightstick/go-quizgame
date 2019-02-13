package quizgame_test

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/knightstick/quizgame"
)

type StubQuestionLoader struct {
	loadCalls  int
	loadedFile string
	questions  []quizgame.Question
}

func (loader *StubQuestionLoader) Load(filename string) []quizgame.Question {
	loader.loadCalls = loader.loadCalls + 1
	loader.loadedFile = filename

	return loader.questions
}

type StubGame struct {
	playCalls int
	questions []quizgame.Question
	score     int
	total     int
}

func (game *StubGame) Play(questions []quizgame.Question) {
	game.playCalls = game.playCalls + 1
	game.questions = questions
}

func (game *StubGame) Score() int {
	return game.score
}

func (game *StubGame) NumberOfQuestions() int {
	return game.total
}

func TestCLIIntegration(t *testing.T) {
	t.Run("Scores 0 out of 0 when no questions", func(t *testing.T) {
		questionFile, removeFile := quizgame.CreateTempFile(t, "")
		defer removeFile()

		in := strings.NewReader("\n")
		out := &bytes.Buffer{}

		cli := quizgame.NewCLI(in, out)
		cli.Run(questionFile.Name())

		assertOutput(t, out, "You scored 0 out of 0\n")
	})

	t.Run("Can score 2 out of 3", func(t *testing.T) {
		questionFile, removeFile := quizgame.CreateTempFile(t, "question,answer\n1+1,2\n1+2,3\n1+3,4\n")
		defer removeFile()

		in := strings.NewReader("2\n3\n666\n")
		out := &bytes.Buffer{}

		cli := quizgame.NewCLI(in, out)
		cli.Run(questionFile.Name())

		t.Skip("skipping until all pieces integrate together")
		assertOutput(t, out, "You scored 2 out of 3\n")
	})
}

func assertLoadedFile(t *testing.T, loader *StubQuestionLoader, filename string) {
	t.Helper()

	if loader.loadCalls != 1 {
		t.Errorf("did not load the questions")
	}

	if loader.loadedFile != filename {
		t.Errorf("expected to load file %s, but loaded %s", filename, loader.loadedFile)
	}
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

func assertOutput(t *testing.T, out *bytes.Buffer, expected string) {
	actual := out.String()

	if actual != expected {
		t.Errorf("expected output %s, got %s", expected, actual)
	}
}
