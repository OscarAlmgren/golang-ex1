package main

import "fmt"

type kfvbroder int

func main() {
	var x kfvbroder

	x = 42

	fmt.Printf("%T\n", x)
	fmt.Println(x)
}
