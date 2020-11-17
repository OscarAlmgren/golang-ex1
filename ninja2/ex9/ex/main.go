/* Ninja level 9 */

package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Spoken from speak for", p.first)
}

func saySomething(h human) {
	h.speak()
}

func main() {

	p1 := person{
		first: "Oscar",
		last:  "Almgren",
		age:   38,
	}

	saySomething(&p1)
	// var wg sync.WaitGroup
	// wg.Add(2)

	// go func() {
	// 	fmt.Println("hello from 1")
	// 	wg.Done()
	// }()

	// go func() {
	// 	fmt.Println("hello from 2")
	// 	wg.Done()
	// }()

	// fmt.Println("Mid go runs:", runtime.NumGoroutine())

	// wg.Wait()
	// fmt.Println("CPU:", runtime.NumCPU())
	// fmt.Println("GOMAX:", runtime.GOMAXPROCS(0))
	// fmt.Println("Go runs:", runtime.NumGoroutine())
	// fmt.Println("About to exit")
}
