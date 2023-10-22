package tasks

import "fmt"


// Task 1: Write a function that creates and prints a slice of integers with values 1 to 5.

func CreateAndPrintSlice() {
	// Implement the code to create and print a slice of integers.
	array := []int{1, 2, 3, 4, 5, 6}
	slice := array[0:5]
	fmt.Println(slice)
}

// Task 2: Write a function that takes a slice of integers and calculates and prints their sum.

func CalculateAndPrintSum(numbers []int) {
	// Implement the code to calculate and print the sum of the numbers in the slice.
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	fmt.Println("Sum:", sum)
}

// Task 3: Write a function that creates and prints an array of strings containing days of the week.

func CreateAndPrintDaysOfWeek() {
	// Implement the code to create and print an array of strings containing days of the week.
	days := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
}

// Task 4: Write a function that takes a slice of integers, filters out even numbers, and prints the remaining odd numbers.

func FilterAndPrintOddNumbers(numbers []int) {
	// Implement the code that takes a slice of integers, filters out even numbers, and prints the remaining odd numbers.
	oddNumbers := []int{}
	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 != 0 {
			oddNumbers = append(oddNumbers, numbers[i])
		}
	}
	fmt.Println("Odd Numbers:", oddNumbers)
}

// Task 5: Write a function that takes a map of string keys and integer values and prints the keys and values.

func PrintMapKeysAndValues(data map[string]int) {
	// Implement the code to print the keys and values in the map.
	for key, value := range data {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
