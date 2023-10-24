package main

import ("shinigami-777/go_assignment/tasks"
	     "fmt"
	     "time"
)

func main() {
	tasks.PrintHelloWorld()
	tasks.PrintInteger(100)
	tasks.PrintSum(10,101)
	//goroutines task implementation
	go first_half()
	go second_half()
	time.Sleep(time.Second*10)
}

//task 3 functions
func first_half(){
    for i:=1;i<=5;i++{
    	fmt.Println(i)
    	time.Sleep(time.Second*1)
    }
}
func second_half(){
	for i:=6;i<=10;i++{
		fmt.Println(i)
		time.Sleep(time.Second*1)
	}
}
