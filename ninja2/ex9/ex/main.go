package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Go runs:", runtime.NumGoroutine())
	fmt.Printf("GOMAXPROCS:%d\n", runtime.GOMAXPROCS(0))

	counter := 0
	var atomicCounter int64
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&atomicCounter, 1)
			fmt.Println("Atomic:", atomic.LoadInt64(&atomicCounter))
			mu.Lock()
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched() // hey start another thread
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Go runs:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Go runs:", runtime.NumGoroutine())
	fmt.Println("Count:", counter)
}
