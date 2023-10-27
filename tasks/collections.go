package tasks

import (
	"fmt"
    "sort"
)

// Task 1: Write a function that creates and prints a slice of integers with values 1 to 5.

func CreateAndPrintSlice() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	// Implement the code to create and print a slice of integers.
}

// Task 2: Write a function that takes a slice of integers and calculates and prints their sum.

func CalculateAndPrintSum(numbers []int) {
	sum := 0
    for i := 0; i < len(numbers); i++ {
        sum += numbers[i]
    }

    fmt.Println("Sum:", sum)
	// Implement the code to calculate and print the sum of the numbers in the slice.
}

// Task 3: Write a function that creates and prints an array of strings containing days of the week.

func CreateAndPrintDaysOfWeek() {
	daysOfWeek := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

    for i := 0; i < len(daysOfWeek); i++ {
        fmt.Println(daysOfWeek[i])
    }

	// Implement the code to create and print an array of strings containing days of the week.
}

// Task 4: Write a function that takes a slice of integers, filters out even numbers, and prints the remaining odd numbers.

func FilterAndPrintOddNumbers(numbers []int) {
    var oddNumbers []int

    for i := 0; i < len(numbers); i++ {
        if numbers[i]%2 != 0 {
            oddNumbers = append(oddNumbers, numbers[i])
        }
    }

    fmt.Println("Odd Numbers:", oddNumbers)
	// Implement the code that takes a slice of integers, filters out even numbers, and prints the remaining odd numbers.
}

// Task 5: Write a function that takes a map of string keys and integer values and prints the keys and values.

func PrintMapKeysAndValues(data map[string]int) {
	keys := make([]string, 0, len(data)) 

    for key := range data {
        keys = append(keys, key)
    }

	sort.Strings(keys) // Sort the keys

    for i := 0; i < len(keys); i++ {
        key := keys[i]
        value := data[key]
        fmt.Printf("Key: %s, Value: %d\n", key, value)
    }
	// Implement the code to print the keys and values in the map.
}
