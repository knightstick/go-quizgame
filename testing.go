package quizgame

import (
	"bytes"
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

// AssertOutput checks that the output you pass in matches the string you expect
func AssertOutput(t *testing.T, out *bytes.Buffer, expected string) {
	t.Helper()

	actual := out.String()

	if actual != expected {
		t.Errorf("expected output %s, got %s", expected, actual)
	}
}
