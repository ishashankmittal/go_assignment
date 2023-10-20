package tasks
import "fmt"
// Task 1: Write a function that creates and prints a slice of integers with values 1 to 5.

func CreateAndPrintSlice() {
	var sl []int
	sl=make([]int,5)
	sl=[]int{1,2,3,4,5}
	fmt.Println(sl)
	// Implement the code to create and print a slice of integers.
}

// Task 2: Write a function that takes a slice of integers and calculates and prints their sum.

func CalculateAndPrintSum(numbers []int) {
	sum:=0
	for i:=0; i<len(numbers); i++{
		ele:=numbers[i]
		sum+=ele
	}
	/*================another method===========================
	for _,v:=range numbers{
		sum+=v
	}*/
	fmt.Println("Sum:",sum)
	// Implement the code to calculate and print the sum of the numbers in the slice.
}

// Task 3: Write a function that creates and prints an array of strings containing days of the week.

func CreateAndPrintDaysOfWeek() {
	// Implement the code to create and print an array of strings containing days of the week.
	var arr [7]string
	arr=[7]string{"Monday","Tuesday","Wednesday","Thursday","Friday","Saturday","Sunday"}
	for _,v:= range arr{
		fmt.Println(v)
	}
}

// Task 4: Write a function that takes a slice of integers, filters out even numbers, and prints the remaining odd numbers.

func FilterAndPrintOddNumbers(numbers []int) {
	var p []int
	p=make([]int,0)
	// Implement the code that takes a slice of integers, filters out even numbers, and prints the remaining odd numbers.
	for _,v:= range numbers{
		if v%2!=0{
			p=append(p,v)//appending the odd value in a new slice
		}
	}
	fmt.Println("Odd Numbers:",p)
}

// Task 5: Write a function that takes a map of string keys and integer values and prints the keys and values.

func PrintMapKeysAndValues(data map[string]int) {
	// Implement the code to print the keys and values in the map.
	for k,v := range data {
		fmt.Print("Key: ",k)
		fmt.Println(", Value:",v)
	}
	}

