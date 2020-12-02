package main

import "fmt"

func main() {
	// c := make(chan int) // non-buffered channel
	// go func() {
	// 	c <- 42
	// }()

	// d := make(chan int, 1) // buffered channel (try to stay away from them according to tips)
	// d <- 131

	// fmt.Println(<-c)
	// fmt.Println(<-d)

	c := make(chan int)

	go foo(c)

	bar(c)

	fmt.Println("Exit")

}

// send
func foo(c chan<- int) {
	c <- 42
}

// receive
func bar(c <-chan int) {
	fmt.Println("Value:", <-c)
}
