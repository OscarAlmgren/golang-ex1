package main

import "fmt"

type person struct {
	first string
	last  string
}

// structs
func main() {
	p1 := person{
		first: "Oscar",
		last:  "Almgren",
	}
	p2 := person{
		first: "Ida",
		last:  "Almgren",
	}
	fmt.Println(p1, "\n", p2)
}
