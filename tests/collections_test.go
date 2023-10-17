package tests

import (
	"ishashankmittal/go_assignment/tasks"
	"testing"
)

func TestCreateAndPrintSlice(t *testing.T) {
	capturedOutput := captureOutput(func() {
		tasks.CreateAndPrintSlice()
	})

	expectedOutput := "[1 2 3 4 5]\n"
	if capturedOutput != expectedOutput {
		t.Errorf("Expected output: %s, Got: %s", expectedOutput, capturedOutput)
	}
}

func TestCalculateAndPrintSum(t *testing.T) {
	capturedOutput := captureOutput(func() {
		tasks.CalculateAndPrintSum([]int{1, 2, 3, 4, 5})
	})

	expectedOutput := "Sum: 15\n"
	if capturedOutput != expectedOutput {
		t.Errorf("Expected output: %s, Got: %s", expectedOutput, capturedOutput)
	}
}

func TestCreateAndPrintDaysOfWeek(t *testing.T) {
	capturedOutput := captureOutput(func() {
		tasks.CreateAndPrintDaysOfWeek()
	})

	expectedOutput := "Monday\nTuesday\nWednesday\nThursday\nFriday\nSaturday\nSunday\n"
	if capturedOutput != expectedOutput {
		t.Errorf("Expected output: %s, Got: %s", expectedOutput, capturedOutput)
	}
}

func TestFilterAndPrintOddNumbers(t *testing.T) {
	capturedOutput := captureOutput(func() {
		tasks.FilterAndPrintOddNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	})

	expectedOutput := "Odd Numbers: [1 3 5 7 9]\n"
	if capturedOutput != expectedOutput {
		t.Errorf("Expected output: %s, Got: %s", expectedOutput, capturedOutput)
	}
}

func TestPrintMapKeysAndValues(t *testing.T) {
	data := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	capturedOutput := captureOutput(func() {
		tasks.PrintMapKeysAndValues(data)
	})

	expectedOutput := "Key: A, Value: 1\nKey: B, Value: 2\nKey: C, Value: 3\n"
	if capturedOutput != expectedOutput {
		t.Errorf("Expected output: %s, Got: %s", expectedOutput, capturedOutput)
	}
}
