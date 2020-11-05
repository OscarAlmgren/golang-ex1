package main

import "fmt"

/*
All values are stored in memory. Every location in memory has an address. A pointer is a memory address.
When to use them?

Pointers allow you to share a value stored in some memory location. Use pointers when
1. you don’t want to pass around a lot of data (enterprise load)
2. you want to change the data at a location (http client?)
*/

func main() {
	a := 1982
	fmt.Println("value:", a) // a is int 1982
	fmt.Println("adr:", &a)  // gives you the address 0xc00001a090

	fmt.Printf("type a: %T\n", a)   // what type is a? -> int
	fmt.Printf("type &a: %T\n", &a) // what type is &a? -> *int

	b := &a                       // b is now the same address as *a, 0xc00001a090
	fmt.Println("mem adr:", b)    // 0xc00001a090
	fmt.Println("mem value:", *b) // gives you the value stored at this address (dereferencing) -> 1982

	*b = 1991        // the pointer *b is set to 1991, being the same address as a. 0xc00001a090
	fmt.Println(*&a) // the value stored at 0xc00001a090 (a) was therefor updated to 1991 when dereferenced
}
