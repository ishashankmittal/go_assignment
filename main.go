package main

import (
	"savsch/go_assignment/tasks"
	"fmt"
	"time"
	"sync"
)

func main() {
	tasks.PrintHelloWorld()
	tasks.PrintInteger(100)
	tasks.PrintSum(10,101)

	// goroutines task:
	fmt.Println("\nGoroutines task:")
	var wg sync.WaitGroup
	wg.Add(2)
        go func() {
            defer wg.Done()
            printFirstHalf()
        }()
        go func() {
            defer wg.Done()
            printSecondHalf()
        }()
	wg.Wait()
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
