package main

import "fmt"

func main() {
	foo()
	bar("Oscar")
	s1 := woo("Fredrik")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x, y)
}

// func (r receiver) identifier(parameter) (return(s)) { code ...}

func foo() {
	fmt.Println("hello from foo")
}

// everything in Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("hello,", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo,", s)
}

func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, ` , says "Hello"`)
	b := false
	return a, b
}
