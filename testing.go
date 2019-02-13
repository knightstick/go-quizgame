package quizgame

import (
	"io/ioutil"
	"os"
	"testing"
)

// CreateTempFile is a helper to help testing with actual file reading
func CreateTempFile(t *testing.T, body string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "quizgame")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(body))

	removeFile := func() {
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
