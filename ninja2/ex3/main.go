package main

import "fmt"

func main() {
	for i := 0; i < 10001; i++ {
		if i%1000 == 0 {
			fmt.Println(i)
		}
	}
}
