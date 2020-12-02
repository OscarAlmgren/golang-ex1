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

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	// go foo(c)

	for y := range c {
		fmt.Println(y)
	}

	fmt.Println("Exit")

}

// send
func foo(c chan<- int) {

	for i := 0; i < 100; i++ {
		c <- i
	}

	close(c)
}

// // receive
// func bar(c <-chan int) {
// 	fmt.Println("Value:", <-c)
// }
