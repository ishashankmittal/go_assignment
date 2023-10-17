package tasks

import "fmt"

// Task 1: Write a function that prints "Hello, World!" to the console.

func PrintHelloWorld() {
	// Implement the code to print "Hello, World!"
	fmt.Print("Hello World")
}

// Task 2: Write a function that takes an integer and prints it to the console using string formatting.

func PrintInteger(n int) {
	// Implement the code to print the integer n.
	fmt.Printf("%v", n)
}

// Task 3: Write a function that takes two integers, adds them, and prints the result.

func PrintSum(a, b int) {
	// Implement the code to print the sum of a and b.
	sum := a * b
	fmt.Printf("%v", sum)
}

// Task 4: Write a function that prints a message with placeholders for a name and age.

func PrintMessageWithPlaceholders(name string, age int) {
	// Implement the code to print a message like "Hello, my name is [name] and I am [age] years old."
	fmt.Printf("Hello, my name is %s and I am %d years old.", name, age)
}

// Task 5: Write a function that prints a formatted date in the "YYYY-MM-DD" format.

func PrintFormattedDate(year, month, day int) {
	// Implement the code to print the date in "YYYY-MM-DD" format.
	fmt.Printf("%d-%d-%d", year, month, day)
}
