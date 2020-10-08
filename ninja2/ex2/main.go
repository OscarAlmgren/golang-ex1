package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(6)
	fmt.Println(i)
	switch i {
	case 1, 3, 5:
		fmt.Println("This was odd")
		//fallthrough
	case 2, 4, 6:
		fmt.Println("This was even")
	default:
		fmt.Println("Default")
	}

}
