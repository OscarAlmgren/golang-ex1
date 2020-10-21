package main

import "fmt"

func main() {
	foo()
	bar("Oscar")
}

// func (r receiver) identifier(parameter) (return(s)) { code ...}

func foo() {
	fmt.Println("hello from foo")
}

// everything in Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("hello,", s)
}
