package main

import (
	"fmt"
	"math"
)

func main() {
	n := foo()
	i, s := bar()
	fmt.Println(n, i, s)

	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(varFoo(xi...))
	fmt.Println(varBar(xi))

	p1 := person{"Oscar", "Almgren", 38}
	p1.speak()

	circle := CIRCLE{radius: 3.0}
	INFO(circle)

	square := SQUARE{length: 2.0, witdh: 4.0}
	INFO(square)

	thirteen := foo()
	fmt.Printf("%T\t%v\n", thirteen, thirteen)

	str := echoEcho("Oscar")
	fmt.Println(str())

	fmt.Println(echoEcho("Rango")())
}

// ex1
func foo() int {
	return 13
}
func bar() (int, string) {
	return 42, "Oscar"
}

// ex2
func varFoo(xi ...int) int {
	fmt.Printf("%T\t%v\t", xi, xi)
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func varBar(xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

// ex3 not important

// ex4
type person struct {
	first, last string
	age         int
}

func (p person) speak() {
	fmt.Println(p.first, p.last, p.age)
}

// ex5

// SQUARE type
type SQUARE struct {
	length, witdh float64
}

func (s SQUARE) area() float64 {
	area := s.length * s.witdh
	return area
}

// CIRCLE type
type CIRCLE struct {
	radius float64
}

func (c CIRCLE) area() float64 {
	area := c.radius * 2 * math.Pi
	return area
}

// SHAPE interface
type SHAPE interface {
	area() float64
}

// INFO func
func INFO(s SHAPE) {
	fmt.Println(s.area())
}

// ex8
// stuff
func echoEcho(s string) func() string {
	fmt.Println("Ropa", s, "?")
	return func() string {
		return "Svara " + s + " !"
	}
}
