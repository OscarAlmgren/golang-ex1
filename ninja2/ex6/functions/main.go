package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	fmt.Println("Opening file x")
	// defer to close file right before exit/return of function
	defer fmt.Println("Closeing file x")

	// identifier
	foo()
	// identifier(argument)
	bar("Oscar")

	s1 := woo("Fredrik")
	fmt.Println(s1)

	// multi-return arguments
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x, y)

	// variadic arguments
	addInts(1, 2, 3, 4, 5)

	sa1 := secretAgent{
		person: person{"James", "Bond"},
		ltk:    true,
	}

	// method in Go.
	sa1.speak()
}

// receiver attaches function to any value of type in receiver. So each secretAgent gets function speak(). It's a method
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

// func (r receiver) identifier(parameter) (return(s)) { code ...}
// func identifier
func foo() {
	fmt.Println("hello from foo")
}

// everything in Go is PASS BY VALUE
// func identifier(parameter)
func bar(s string) {
	fmt.Println("hello,", s)
}

// func identifier(parameter) return
func woo(s string) string {
	return fmt.Sprint("Hello from woo,", s)
}

// func identifier(parameter(s)) (return(s))
func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, ` , says "Hello"`)
	b := false
	return a, b
}

// variadic parameters
func addInts(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		fmt.Println("for index", i, "we add", v)
		sum += v
	}
	fmt.Println("sum:", sum)
}
