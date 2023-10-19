package tasks

import "fmt"

// Task 1: Write a function that creates and prints a slice of integers with values 1 to 5.

func CreateAndPrintSlice() {
	// Implement the code to create and print a slice of integers.
	demo := []int{5, 7}
	fmt.Println(demo)
}

// Task 2: Write a function that takes a slice of integers and calculates and prints their sum.

func CalculateAndPrintSum(numbers []int) {
	// Implement the code to calculate and print the sum of the numbers in the slice.
	number := []int{6, 3}
	var sum int = 0
	for i := 0; i < 2; i++ {
		sum += number[i]
	}
	fmt.Println(sum)

}

// Task 3: Write a function that creates and prints an array of strings containing days of the week.

func CreateAndPrintDaysOfWeek() {
	// Implement the code to create and print an array of strings containing days of the week.
	arr := [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(arr)
}

func FilterAndPrintOddNumbers(numbers []int) {
	// Implement the code that takes a slice of integers, filters out even numbers, and prints the remaining odd numbers.
	for i := 0; i < 4; i++ {
		if numbers[i]%2 == 0 {
			fmt.Println(numbers[i])
		}
	}
}

// Task 5: Write a function that takes a map of string keys and integer values and prints the keys and values.

func PrintMapKeysAndValues(data map[string]int) {
	// Implement the code to print the keys and values in the map.
	fmt.Println(data)
}
