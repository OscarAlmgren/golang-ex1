package main

import "fmt"

func main() {
	// array
	var x [5]int
	x[1] = 42
	fmt.Println(x)

	// slice (allows you to GROUP together VALUES of the same TYPE)
	xx := []int{11, 22, 33, 44, 55} // slice []type{} <- slice format
	fmt.Println(xx)
	fmt.Println(len(xx))
	//fmt.Println(cap(xx))
	fmt.Println("index 2:", xx[2])
	for i, v := range xx {
		fmt.Println(i, v)
	}

	// : operator for slicing slices
	fmt.Println(xx[1:])  // start at this position
	fmt.Println(xx[1:3]) // up to but not including 3rd position

	for i := 0; i <= 4; i++ {
		fmt.Println(i, xx[i])
	}
}
