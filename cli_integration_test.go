package quizgame_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/knightstick/quizgame"
)

type StubQuestionLoader struct {
	loadCalls  int
	loadedFile string
}

func (loader *StubQuestionLoader) Load(filename string) []quizgame.Question {
	loader.loadCalls = loader.loadCalls + 1
	loader.loadedFile = filename

	return []quizgame.Question{}
}

type StubGame struct {
	playCalls int
}

func (game *StubGame) Play([]quizgame.Question) {
	game.playCalls = game.playCalls + 1
}

func TestCLIIntegration(t *testing.T) {
	// Create a new CSV
	questionFile, _ := ioutil.TempFile("", "questions")
	defer os.Remove(questionFile.Name())

	in := strings.NewReader("\n")
	out := &bytes.Buffer{}
	loader := &StubQuestionLoader{}
	game := &StubGame{}

	// New CLI with filename
	cli := &quizgame.CLI{QuestionLoader: loader, In: in, Out: out, Game: game}
	cli.Run(questionFile.Name())

	// Loaded the questions
	if loader.loadCalls != 1 {
		t.Errorf("did not load the questions")
	}

	if loader.loadedFile != questionFile.Name() {
		t.Errorf("expected to load file %s, but loaded %s", questionFile.Name(), loader.loadedFile)
	}

	// Played the game
	if game.playCalls != 1 {
		t.Error("did not play the game")
	}

	// Outputted the results
	expectedOutput := "You scored 0 out of 0\n"
	actualOutput := out.String()

	if actualOutput != expectedOutput {
		t.Errorf("expected output %s, got %s", expectedOutput, out.String())
	}
}
