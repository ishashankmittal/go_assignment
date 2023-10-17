package tests

import (
	"ishashankmittal/go_assignment/tasks"
	"testing"
)

func TestPrintHelloWorld(t *testing.T) {
	// Test Task 1: Print "Hello, World!" and check the output.
	// Redirect standard output for testing.
	capturedOutput := captureOutput(func() {
		tasks.PrintHelloWorld()
	})

	expectedOutput := "Hello World"
	if capturedOutput != expectedOutput {
		t.Errorf("Expected output: %s, Got: %s", expectedOutput, capturedOutput)
	}
}

func TestPrintInteger(t *testing.T) {
	// Test Task 2: Print an integer and check the output.
	capturedOutput := captureOutput(func() {
		tasks.PrintInteger(42)
	})

	expectedOutput := "42"
	if capturedOutput != expectedOutput {
		t.Errorf("Expected output: %s, Got: %s", expectedOutput, capturedOutput)
	}
}

func TestPrintSum(t *testing.T) {
	// Test Task 3: Print the sum of two integers and check the output.
	capturedOutput := captureOutput(func() {
		tasks.PrintSum(10, 20)
	})

	expectedOutput := "30"
	if capturedOutput != expectedOutput {
		t.Errorf("Expected output: %s, Got: %s", expectedOutput, capturedOutput)
	}
}

func TestPrintMessageWithPlaceholders(t *testing.T) {
	// Test Task 4: Write a function that prints a message with placeholders for a name and age.
	capturedOutput := captureOutput(func() {
		tasks.PrintMessageWithPlaceholders("Shashank Mittal", 20)
	})

	expectedOutput := "Hello, my name is Shashank Mittal and I am 20 years old."
	if capturedOutput != expectedOutput {
		t.Errorf("Expected output: %s, Got: %s", expectedOutput, capturedOutput)
	}
}

func TestPrintFormattedDate(t *testing.T) {
	// Test Task 5: Write a function that prints a formatted date in the "YYYY-MM-DD" format.
	capturedOutput := captureOutput(func() {
		tasks.PrintFormattedDate(2023, 10, 18)
	})

	expectedOutput := "2023-10-18"
	if capturedOutput != expectedOutput {
		t.Errorf("Expected output: %s, Got: %s", expectedOutput, capturedOutput)
	}
}
