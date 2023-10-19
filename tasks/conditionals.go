package tasks

// Task 1: Write a function that takes an integer and prints a message based on the value:
// - If the integer is positive, print "Positive."
// - If the integer is negative, print "Negative."
// - If the integer is zero, print "Zero."

func PrintMessageBasedOnValue(num int) {
	// Implement the code to print the appropriate message based on the value of num.
	if num>0{
		fmt.Println("Positive.")
	}else if num<0{
		fmt.Println("Negative.")
	}else{
		fmt.Println("Zero.")
	}
}

// Task 2: Write a function that takes a string and checks if it is a valid day of the week.
// - If the input is "Monday" or "Tuesday" or "Wednesday" or "Thursday" or "Friday",
//   print "Weekday."
// - If the input is "Saturday" or "Sunday", print "Weekend."
// - If the input is anything else, print "Invalid day."

func CheckDayOfWeek(input string) {
	// Implement the code to check the validity of the day and print the corresponding message.
}

// Task 3: Write a function that takes a map of string keys and integer values and calculates the sum of the values.
// If the sum of values is greater than 10, print "Sum is greater than 10."
// If the sum of values is less than or equal to 10, print "Sum is 10 or less."

func CheckIfSumIsLessThan10(inputMap map[string]int) {
	// Implement the code.
}
