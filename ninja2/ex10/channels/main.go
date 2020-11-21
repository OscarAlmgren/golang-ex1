package main

import "fmt"

func main() {
	c := make(chan int) // non-buffered channel
	go func() {
		c <- 42
	}()

	d := make(chan int, 1) // buffered channel
	d <- 131

	fmt.Println(<-c)
	fmt.Println(<-d)
}
