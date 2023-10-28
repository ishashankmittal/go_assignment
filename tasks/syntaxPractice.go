package tasks

import (
	"fmt"
)
// Task 1: Write a function that prints "Hello World" to the console.

func PrintHelloWorld() {
	fmt.Print("Hello World")
	// Implement the code to print "Hello World"
}

// Task 2: Write a function that takes an integer and prints it to the console using string formatting.

func PrintInteger(n int) {
	fmt.Printf("%d", n)
	// Implement the code to print the integer n.
}

// Task 3: Write a function that takes two integers, adds them, and prints the result.

func PrintSum(a, b int) {
	sum := a+b
	fmt.Printf("%d",sum)
	// Implement the code to print the sum of a and b.
}

// Task 4: Write a function that prints a message with placeholders for a name and age.

func PrintMessageWithPlaceholders(name string, age int) {
	fmt.Printf("Hello, my name is %s and I am %d years old.",name,age)
	// Implement the code to print a message like "Hello, my name is [name] and I am [age] years old."
}

// Task 5: Write a function that prints a formatted date in the "YYYY-MM-DD" format.

func PrintFormattedDate(year, month, day int) {
	fmt.Printf("%d-%02d-%02d", year, month, day)
	// Implement the code to print the date in "YYYY-MM-DD" format.
}
