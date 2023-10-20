package main

import (
	"fmt"
	"time"
)
func printFirstHalf(){
	for i:=1; i<6; i++ {
		fmt.Println(i)
		time.Sleep(1*time.Second)
	}
}	
func printSecondHalf(){
	for i:=6; i<=10; i++ {
		fmt.Println(i)
		time.Sleep(1*time.Second)	
}
}
func main() {
	go printSecondHalf()
	printFirstHalf()
		
}
