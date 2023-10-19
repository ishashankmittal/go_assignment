package tasks

import "fmt"

// Task 1: Write a function that prints "Hello World" to the console.

func PrintHelloWorld() {
	fmt.Println("Hello World!")

	// Implement the code to print "Hello World"
}

// Task 2: Write a function that takes an integer and prints it to the console using string formatting.
func PrintInteger(n int) {
	var num int
	fmt.Scanln(&num)
	fmt.Println(num)
	// Implement the code to print the integer n.
}

// Task 3: Write a function that takes two integers, adds them, and prints the result.

func PrintSum(a, b int) {
	// Implement the code to print the sum of a and b.
	var num1, num2 int
	fmt.Println(num1 + num2)
}

// Task 4: Write a function that prints a message with placeholders for a name and age.

func PrintMessageWithPlaceholders(name string, age int) {
	// Implement the code to print a message like "Hello, my name is [name] and I am [age] years old."
	fmt.Print("Hello, my name is ")
	fmt.Print(name)
	fmt.Print(" and I am ")
	fmt.Print(age)
	fmt.Print(" years old.")
	fmt.Println()

}

// Task 5: Write a function that prints a formatted date in the "YYYY-MM-DD" format.

func PrintFormattedDate(year, month, day int) {
	// Implement the code to print the date in "YYYY-MM-DD" format.
	fmt.Println("The Date is ")
	fmt.Print(year)
	fmt.Print("-")
	fmt.Print(month)
	fmt.Print("-")
	fmt.Print(day)
}
