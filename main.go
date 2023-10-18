package main

import (
	"fmt"
	"time"
)

// implement task in main.go

// Task: Print Numbers 1 to 10 Using Goroutines

// Write a Go program that uses goroutines to print numbers from 1 to 10. The program should do the following:

// Create two goroutines, printFirstHalf and printSecondHalf, that print numbers from 1 to 5 and numbers from 6 to 10, respectively.

// Use goroutines to ensure that the numbers are printed concurrently.

// Print the numbers in ascending order with a 1-second delay between each number.

func main() {
	go printFirstHalf()
	go printSecondHalf()
	time.Sleep(10 * time.Second)
}

func printFirstHalf() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func printSecondHalf() {
	for i := 6; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}
