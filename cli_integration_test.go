package quizgame_test

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/knightstick/quizgame"
)

var QuickTimer = &quizgame.WaitQuizTimer{SleepTime: 1 * time.Microsecond}

func TestCLIIntegration(t *testing.T) {
	t.Run("Scores 0 out of 0 when no questions", func(t *testing.T) {
		questionFile, removeFile := quizgame.CreateTempFile(t, "")
		defer removeFile()

		in := strings.NewReader("\n")
		out := &bytes.Buffer{}

		cli := quizgame.NewCLI(in, out, &quizgame.FileSystemQuestionLoader{}, QuickTimer, questionFile.Name())
		cli.Run()

		quizgame.AssertOutput(t, out, "You scored 0 out of 0\n")
	})

	t.Run("Can score 2 out of 3", func(t *testing.T) {
		questionFile, removeFile := quizgame.CreateTempFile(t, "1+1,2\n1+2,3\n1+3,4\n")
		defer removeFile()

		in := strings.NewReader("2\n3\n666\n")
		out := &bytes.Buffer{}

		cli := quizgame.NewCLI(in, out, &quizgame.FileSystemQuestionLoader{}, QuickTimer, questionFile.Name())
		cli.Run()

		expectedOutput := "Problem #1: 1+1 = Problem #2: 1+2 = Problem #3: 1+3 = You scored 2 out of 3\n"
		quizgame.AssertOutput(t, out, expectedOutput)
	})
}
