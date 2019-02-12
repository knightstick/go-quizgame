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
	questionFile, removeFile := createTempFile(t, "")
	defer removeFile()

	in := strings.NewReader("\n")
	out := &bytes.Buffer{}
	loader := &StubQuestionLoader{}
	game := &StubGame{}

	// New CLI with filename
	cli := &quizgame.CLI{QuestionLoader: loader, In: in, Out: out, Game: game}
	cli.Run(questionFile.Name())

	// Loaded the questions
	assertLoadedFile(t, loader, questionFile.Name())

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

func createTempFile(t *testing.T, body string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "questions")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(body))

	removeFile := func() {
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
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
