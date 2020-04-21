package main

import "fmt"

type kfvbroder int

var x kfvbroder
var y int

func main() {
	x = 1982
	fmt.Printf("%T\n", x)
	fmt.Println("The value x:", x)

	fmt.Println("The value y:", y)
	y = int(x)
	fmt.Printf("%T\t%v\n", y, y)

}
