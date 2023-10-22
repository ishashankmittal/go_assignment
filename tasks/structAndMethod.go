package tasks

import "fmt"


// Define a struct named Person with fields Name, Age, and City
type Person struct {
	Name string
	Age  int
	City string
}

// Task 1: Define a Struct and Print Its Fields

// Write a function named DefineAndPrintStruct that defines a struct named Person with the following fields:
// Name (string)
// Age (int)
// City (string)
// Then print its fields.

func DefineAndPrintStruct() {
	// Define a struct named Person with fields Name, Age, and City , Again.

	// Comment out this instance of the Person struct with sample values.
	// person := Person{
	// 	Name: "John",
	// 	Age:  30,
	// 	City: "New York",
	// }
	// Print the fields of the struct instance.
	person := Person{
		Name: "John",
		Age:  30,
		City: "New York",
	}
	// Print the fields of the struct instance.
	fmt.Printf("Name: %s, Age: %d, City: %s\n", person.Name, person.Age, person.City)
}

// Task 2: Add a Method to a Struct
// Extend the Person struct from Task 6 with a method named GetInfo that takes no arguments and returns a string. This method should return a formatted string that includes the person's name, age, and city.

// Construct the GetInfo method such that it takes a struct of type Person
// and returns a string.

// For example, if the person's name is "Alice," age is 25, and city is "New York,"
// the GetInfo method should return "Name: Alice, Age: 25, City: New York."

// here
func (p Person) GetInfo() string {
	return "Name: " + p.Name + ", Age: " + fmt.Sprint(p.Age) + ", City: " + p.City
}


// Comment out function TestAddMethodToStruct in /tests/structAndMethod.go to test the Task 2
