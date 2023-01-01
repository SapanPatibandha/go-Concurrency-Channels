package main

import (
	"fmt"
)

func timeThread(number []int, ch chan int) {
	// result := number * 3
	// fmt.Println(number * 3)
	// ch <- result

	for _, elem := range number {
		ch <- elem * 3
	}
}

func main() {
	fmt.Println("we are starting program!")

	arr := []int{2, 3, 4}

	ch := make(chan int, 3)

	go timeThread(arr, ch)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Result is %v \n", <-ch)
	}

	// time.Sleep(time.Second)
}
