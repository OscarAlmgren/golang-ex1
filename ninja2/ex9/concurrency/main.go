package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
	Intel released the first commercially available multi-core CPUs in 2006
	Google built and released Go lang in 2007 to natively take advantage of multi-core CPUs.

	Concurrency vs Parallellism

	Concurrency - a design pattern to write code that can run in parallel on multi-core, but does not guarantee parallelism

	Parallellism - run code at the same time in parallell on multiple CPUs

	Steps:
	wg.Add(n)
	wg.Wait()
	wg.Done() - in other method
*/

var wg sync.WaitGroup

func main() {
	fmt.Println("Go rout:", runtime.NumGoroutine())
	fmt.Println("OS:\t", runtime.GOOS)
	fmt.Println("ARCH:\t", runtime.GOARCH)
	fmt.Println("CPUs:\t", runtime.NumCPU())

	wg.Add(1)
	go foo()
	bar()
	fmt.Println("Go rout:", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar:", i)
	}

}
