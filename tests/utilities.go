package tests

import (
	"io"
	"os"
)

// Helper function to capture output for testing.
func captureOutput(f func()) string {
	oldOutput := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldOutput

	return string(out)
}
