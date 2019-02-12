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

func (loader *StubQuestionLoader) Load(filename string) {
	loader.loadCalls = loader.loadCalls + 1
	loader.loadedFile = filename
}

func TestCLIIntegration(t *testing.T) {
	// Create a new CSV
	questionFile, _ := ioutil.TempFile("", "questions")
	defer os.Remove(questionFile.Name())

	in := strings.NewReader("\n")
	out := &bytes.Buffer{}
	loader := &StubQuestionLoader{}

	// New CLI with filename
	cli := &quizgame.CLI{QuestionLoader: loader, In: in, Out: out}
	cli.Run(questionFile.Name())

	// Loaded the questions
	if loader.loadCalls != 1 {
		t.Errorf("did not load the questions")
	}

	if loader.loadedFile != questionFile.Name() {
		t.Errorf("expected to load file %s, but loaded %s", questionFile.Name(), loader.loadedFile)
	}

	// Answered the questions
	// Outputted the results
}
