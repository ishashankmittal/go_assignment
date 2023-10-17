package tests

import (
	"ishashankmittal/go_assignment/tasks"
	"testing"
)

type TestPerson struct {
	tasks.Person
}

func TestDefineAndPrintStruct(t *testing.T) {
	// Test case for defining and printing the Person struct
	capturedOutput := captureOutput(func() {
		tasks.DefineAndPrintStruct()
	})

	expectedOutput := "Name: John, Age: 30, City: New York\n"
	if capturedOutput != expectedOutput {
		t.Errorf("Expected output: %s, Got: %s", expectedOutput, capturedOutput)
	}
}

func TestAddMethodToStruct(t *testing.T) {
	// Create a Person instance with sample values
	testperson := TestPerson{
		Person: tasks.Person{Name: "Alice",
			Age:  25,
			City: "New York"},
	}

	// Test the GetInfo method
	info := testperson.Person.GetInfo()
	expectedInfo := "Name: Alice, Age: 25, City: New York"
	if info != expectedInfo {
		t.Errorf("Expected info: %s, Got: %s", expectedInfo, info)
	}
}
