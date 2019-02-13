package quizgame_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/knightstick/quizgame"
)

func TestCLIIntegration(t *testing.T) {
	t.Run("Scores 0 out of 0 when no questions", func(t *testing.T) {
		questionFile, removeFile := quizgame.CreateTempFile(t, "")
		defer removeFile()

		in := strings.NewReader("\n")
		out := &bytes.Buffer{}

		cli := quizgame.NewCLI(in, out, &quizgame.FileSystemQuestionLoader{}, questionFile.Name())
		cli.Run()

		quizgame.AssertOutput(t, out, "You scored 0 out of 0\n")
	})

	t.Run("Can score 2 out of 3", func(t *testing.T) {
		questionFile, removeFile := quizgame.CreateTempFile(t, "question,answer\n1+1,2\n1+2,3\n1+3,4\n")
		defer removeFile()

		in := strings.NewReader("2\n3\n666\n")
		out := &bytes.Buffer{}

		cli := quizgame.NewCLI(in, out, &quizgame.FileSystemQuestionLoader{}, questionFile.Name())
		cli.Run()

		t.Skip("skipping until all pieces integrate together")
		quizgame.AssertOutput(t, out, "You scored 2 out of 3\n")
	})
}
