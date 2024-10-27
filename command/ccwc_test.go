package command

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestCountingWordsSuccessFully(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	NewWcFromFile("c", "../test.txt").Exec()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout

	assert.Equal(t, "../test.txt\n# Bytes: 104 \t\n", string(out))
}

func TestCountingBytesSuccessFully(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	NewWcFromFile("w", "../test.txt").Exec()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout

	assert.Equal(t, "../test.txt\n# Words: 12 \t\n", string(out))
}

func TestCountingLinesSuccessFully(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	NewWcFromFile("l", "../test.txt").Exec()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout

	assert.Equal(t, "../test.txt\n# Lines: 9 \t\n", string(out))
}

func TestCountingCharactersSuccessFully(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	NewWcFromFile("m", "../test.txt").Exec()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout

	assert.Equal(t, "../test.txt\n# Characters: 104 \t\n", string(out))
}

func TestCountingCharactersAndLinesSuccessFully(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	NewWcFromFile("ml", "../test.txt").Exec()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout

	assert.Equal(t, "../test.txt\n# Lines: 9 \t# Characters: 104 \t\n", string(out))
}
