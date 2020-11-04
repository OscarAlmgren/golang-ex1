package main

import "fmt"

func main() {
	a := 1982
	fmt.Println("value:", a)
	fmt.Println("adr:", &a) // gives you the address

	fmt.Printf("type a: %T\n", a)   // what type is a? -> int
	fmt.Printf("type &a: %T\n", &a) // what type is &a? -> *int

	b := &a // gives us an address.
	fmt.Println("mem adr:", b)
	fmt.Println("mem value:", *b) // gives you the value stored at this address (dereferencing)

	*b = 1991
	fmt.Println(a)
}
