package tests

import (
	"criticic/cops_go_assignment/tasks"
	"fmt"
	"testing"
)

var Score int = 0
var Total int = 40

func TestMain(m *testing.M) {
	m.Run()
	fmt.Printf("\nScore: %d out of %d\n", Score*2, Total) // Assuming 2 points per correct test.
}

func TestPrintHelloWorld(t *testing.T) {
	// Test Task 1: Print "Hello, World!" and check the output.
	// Redirect standard output for testing.
	capturedOutput := captureOutput(func() {
		tasks.PrintHelloWorld()
	})

	expectedOutput := "Hello World"
	if capturedOutput != expectedOutput {
		errorMessage := fmt.Sprintf("\nExpected output: %s, Got: %s\n", expectedOutput, capturedOutput)
		t.Error(errorMessage)
	} else {
		Score++
		passMessage := fmt.Sprint("\nTestPrintHelloWorld passed successfully!\n\n")
		fmt.Println(passMessage)
	}
}

func TestPrintInteger(t *testing.T) {
	// Test Task 2: Print an integer and check the output.
	capturedOutput := captureOutput(func() {
		tasks.PrintInteger(42)
	})

	expectedOutput := "42"
	if capturedOutput != expectedOutput {
		errorMessage := fmt.Sprintf("\nExpected output: %s, Got: %s\n", expectedOutput, capturedOutput)
		t.Error(errorMessage)
	} else {
		Score++
		passMessage := fmt.Sprint("\nTestPrintInteger passed successfully!\n\n")
		fmt.Println(passMessage)
	}
}

func TestPrintSum(t *testing.T) {
	// Test Task 3: Print the sum of two integers and check the output.
	capturedOutput := captureOutput(func() {
		tasks.PrintSum(10, 20)
	})

	expectedOutput := "30"
	if capturedOutput != expectedOutput {
		errorMessage := fmt.Sprintf("\nExpected output: %s, Got: %s\n", expectedOutput, capturedOutput)
		t.Error(errorMessage)
	} else {
		Score++
		passMessage := fmt.Sprint("\nTestPrintSum passed successfully!\n")
		fmt.Println(passMessage)
	}
}

func TestPrintMessageWithPlaceholders(t *testing.T) {
	// Test Task 4: Write a function that prints a message with placeholders for a name and age.
	capturedOutput := captureOutput(func() {
		tasks.PrintMessageWithPlaceholders("Shashank Mittal", 20)
	})

	expectedOutput := "Hello, my name is Shashank Mittal and I am 20 years old."
	if capturedOutput != expectedOutput {
		errorMessage := fmt.Sprintf("\nExpected output: %s, Got: %s\n", expectedOutput, capturedOutput)
		t.Error(errorMessage)
	} else {
		Score++
		passMessage := fmt.Sprint("\nTestPrintMessageWithPlaceholders passed successfully!\n")
		fmt.Println(passMessage)
	}
}

func TestPrintFormattedDate(t *testing.T) {
	// Test Task 5: Write a function that prints a formatted date in the "YYYY-MM-DD" format.
	capturedOutput := captureOutput(func() {
		tasks.PrintFormattedDate(2023, 10, 18)
	})

	expectedOutput := "2023-10-18"
	if capturedOutput != expectedOutput {
		errorMessage := fmt.Sprintf("\nExpected output: %s, Got: %s\n", expectedOutput, capturedOutput)
		t.Error(errorMessage)
	} else {
		Score++
		passMessage := fmt.Sprint("\nTestPrintFormattedDate passed successfully!\n")
		fmt.Println(passMessage)
	}
}

func TestPrintMessageBasedOnValue(t *testing.T) {
	// Test case for a positive number
	capturedOutput := captureOutput(func() {
		tasks.PrintMessageBasedOnValue(5)
	})

	expectedOutput := "Positive\n"
	if capturedOutput != expectedOutput {
		t.Errorf("For 5, Expected output: %s, Got: %s", expectedOutput, capturedOutput)
	} else {
		Score++
		passMessage := fmt.Sprint("\nFor 5,TestPrintMessageBasedOnValue passed successfully!\n")
		fmt.Println(passMessage)
	}

	// Test case for a negative number
	capturedOutput = captureOutput(func() {
		tasks.PrintMessageBasedOnValue(-5)
	})

	expectedOutput = "Negative\n"
	if capturedOutput != expectedOutput {
		t.Errorf("For -5, Expected output: %s, Got: %s", expectedOutput, capturedOutput)
	} else {
		Score++
		passMessage := fmt.Sprint("\nFor -5,TestPrintMessageBasedOnValue passed successfully!\n")
		fmt.Println(passMessage)
	}

	// Test case for zero
	capturedOutput = captureOutput(func() {
		tasks.PrintMessageBasedOnValue(0)
	})

	expectedOutput = "Zero\n"
	if capturedOutput != expectedOutput {
		t.Errorf("For 0, Expected output: %s, Got: %s", expectedOutput, capturedOutput)
	} else {
		Score++
		passMessage := fmt.Sprint("\nFor 0,TestPrintMessageBasedOnValue passed successfully!\n")
		fmt.Println(passMessage)
	}
}

func TestCheckDayOfWeek(t *testing.T) {
	// Test case for a weekday
	capturedOutput := captureOutput(func() {
		tasks.CheckDayOfWeek("Monday")
	})

	expectedOutput := "Weekday\n"
	if capturedOutput != expectedOutput {
		errorMessage := fmt.Sprintf("\nExpected output: %s, Got: %s\n", expectedOutput, capturedOutput)
		t.Error(errorMessage)
	} else {
		Score++
		passMessage := fmt.Sprint("\nFor Weekday,TestCheckDayOfWeek passed successfully!\n")
		fmt.Println(passMessage)
	}

	// Test case for a weekend day
	capturedOutput = captureOutput(func() {
		tasks.CheckDayOfWeek("Saturday")
	})

	expectedOutput = "Weekend\n"
	if capturedOutput != expectedOutput {
		errorMessage := fmt.Sprintf("\nExpected output: %s, Got: %s\n", expectedOutput, capturedOutput)
		t.Error(errorMessage)
	} else {
		Score++
		passMessage := fmt.Sprint("\nFor Weekend,TestCheckDayOfWeek passed successfully!\n")
		fmt.Println(passMessage)
	}

	// Test case for an invalid day
	capturedOutput = captureOutput(func() {
		tasks.CheckDayOfWeek("InvalidDay")
	})

	expectedOutput = "Invalid day\n"
	if capturedOutput != expectedOutput {
		errorMessage := fmt.Sprintf("\nExpected output: %s, Got: %s\n", expectedOutput, capturedOutput)
		t.Error(errorMessage)
	} else {
		Score++
		passMessage := fmt.Sprint("\nFor Invalid day,TestCheckDayOfWeek passed successfully!\n")
		fmt.Println(passMessage)
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
	} else {
		Score++
		passMessage := fmt.Sprint("\nFor Sum greater that 10, TestCheckIfSumIsLessThan10 passed successfully!\n")
		fmt.Println(passMessage)
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
	} else {
		Score++
		passMessage := fmt.Sprint("\nFor Sum greater that 10, TestCheckIfSumIsLessThan10 passed successfully!\n")
		fmt.Println(passMessage)
	}
}

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
	} else {
		Score++
		passMessage := fmt.Sprint("\nTestDefineAndPrintStruct passed successfully!\n")
		fmt.Println(passMessage)
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
	} else {
		Score++
		passMessage := fmt.Sprint("\nTestAddMethodToStruct passed successfully!\n")
		fmt.Println(passMessage)
	}
}

func TestCreateAndPrintSlice(t *testing.T) {
	capturedOutput := captureOutput(func() {
		tasks.CreateAndPrintSlice()
	})

	expectedOutput := "[1 2 3 4 5]\n"
	if capturedOutput != expectedOutput {
		t.Errorf("Expected output: %s, Got: %s", expectedOutput, capturedOutput)
	} else {
		Score++
		passMessage := fmt.Sprint("\nTestCreateAndPrintSlice passed successfully!\n")
		fmt.Println(passMessage)
	}
}

func TestCalculateAndPrintSum(t *testing.T) {
	capturedOutput := captureOutput(func() {
		tasks.CalculateAndPrintSum([]int{1, 2, 3, 4, 5})
	})

	expectedOutput := "Sum: 15\n"
	if capturedOutput != expectedOutput {
		t.Errorf("Expected output: %s, Got: %s", expectedOutput, capturedOutput)
	} else {
		Score++
		passMessage := fmt.Sprint("\nTestCalculateAndPrintSum passed successfully!\n")
		fmt.Println(passMessage)
	}
}

func TestCreateAndPrintDaysOfWeek(t *testing.T) {
	capturedOutput := captureOutput(func() {
		tasks.CreateAndPrintDaysOfWeek()
	})

	expectedOutput := "Monday\nTuesday\nWednesday\nThursday\nFriday\nSaturday\nSunday\n"
	if capturedOutput != expectedOutput {
		t.Errorf("Expected output: %s, Got: %s", expectedOutput, capturedOutput)
	} else {
		Score++
		passMessage := fmt.Sprint("\nTestCreateAndPrintDaysOfWeek passed successfully!\n")
		fmt.Println(passMessage)
	}
}

func TestFilterAndPrintOddNumbers(t *testing.T) {
	capturedOutput := captureOutput(func() {
		tasks.FilterAndPrintOddNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	})

	expectedOutput := "Odd Numbers: [1 3 5 7 9]\n"
	if capturedOutput != expectedOutput {
		t.Errorf("Expected output: %s, Got: %s", expectedOutput, capturedOutput)
	} else {
		Score++
		passMessage := fmt.Sprint("\nTestFilterAndPrintOddNumbers passed successfully!\n")
		fmt.Println(passMessage)
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
	} else {
		Score++
		passMessage := fmt.Sprint("\nTestPrintMapKeysAndValues passed successfully!\n")
		fmt.Println(passMessage)
	}
}
