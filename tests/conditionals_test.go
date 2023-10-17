package tests

import (
	"ishashankmittal/go_assignment/tasks"
	"testing"
)

func TestPrintMessageBasedOnValue(t *testing.T) {
	// Test case for a positive number
	capturedOutput := captureOutput(func() {
		tasks.PrintMessageBasedOnValue(5)
	})

	expectedOutput := "Positive\n"
	if capturedOutput != expectedOutput {
		t.Errorf("For 5, Expected output: %s, Got: %s", expectedOutput, capturedOutput)
	}

	// Test case for a negative number
	capturedOutput = captureOutput(func() {
		tasks.PrintMessageBasedOnValue(-5)
	})

	expectedOutput = "Negative\n"
	if capturedOutput != expectedOutput {
		t.Errorf("For -5, Expected output: %s, Got: %s", expectedOutput, capturedOutput)
	}

	// Test case for zero
	capturedOutput = captureOutput(func() {
		tasks.PrintMessageBasedOnValue(0)
	})

	expectedOutput = "Zero\n"
	if capturedOutput != expectedOutput {
		t.Errorf("For 0, Expected output: %s, Got: %s", expectedOutput, capturedOutput)
	}
}

func TestCheckDayOfWeek(t *testing.T) {
	// Test case for a weekday
	capturedOutput := captureOutput(func() {
		tasks.CheckDayOfWeek("Monday")
	})

	expectedOutput := "Weekday\n"
	if capturedOutput != expectedOutput {
		t.Errorf("For Monday, Expected output: %s, Got: %s", expectedOutput, capturedOutput)
	}

	// Test case for a weekend day
	capturedOutput = captureOutput(func() {
		tasks.CheckDayOfWeek("Saturday")
	})

	expectedOutput = "Weekend\n"
	if capturedOutput != expectedOutput {
		t.Errorf("For Saturday, Expected output: %s, Got: %s", expectedOutput, capturedOutput)
	}

	// Test case for an invalid day
	capturedOutput = captureOutput(func() {
		tasks.CheckDayOfWeek("InvalidDay")
	})

	expectedOutput = "Invalid day\n"
	if capturedOutput != expectedOutput {
		t.Errorf("For InvalidDay, Expected output: %s, Got: %s", expectedOutput, capturedOutput)
	}
}

func TestCheckIfSumIsLessThan10(t *testing.T) {
	// Test case for a sum greater than 10
	inputMap := map[string]int{
		"A": 5,
		"B": 6,
	}

	capturedOutput := captureOutput(func() {
		tasks.CheckIfSumIsLessThan10(inputMap)
	})

	expectedOutput := "Sum is greater than 10\n"
	if capturedOutput != expectedOutput {
		t.Errorf("Expected output: %s, Got: %s", expectedOutput, capturedOutput)
	}

	// Test case for a sum equal to 10
	inputMap = map[string]int{
		"A": 5,
		"B": 5,
	}

	capturedOutput = captureOutput(func() {
		tasks.CheckIfSumIsLessThan10(inputMap)
	})

	expectedOutput = "Sum is 10 or less\n"
	if capturedOutput != expectedOutput {
		t.Errorf("Expected output: %s, Got: %s", expectedOutput, capturedOutput)
	}
}
