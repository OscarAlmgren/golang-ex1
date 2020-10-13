package main

import "fmt"

func main() {

	// array
	fmt.Println("== regular array ==")
	var x [5]int
	x[1] = 42
	fmt.Println(x)

	// slice (allows you to GROUP together VALUES of the same TYPE)
	// slices are everywhere in GO. https://blog.golang.org/slices-intro They build on arrays to provide great power and convenience
	fmt.Println("== slice []int ==")
	xx := []int{11, 22, 33, 44, 55} // slice []type{} <- slice format []T
	fmt.Println(xx)
	fmt.Println(len(xx))
	//fmt.Println(cap(xx))
	fmt.Println("index 2:", xx[2])
	for i, v := range xx {
		fmt.Println(i, v)
	}

	// : operator for slicing slices
	fmt.Println("== : operator command ==")
	fmt.Println(xx[1:])  // start at this position
	fmt.Println(xx[1:3]) // up to but not including 3rd position

	for i := 0; i <= 4; i++ {
		fmt.Println(i, xx[i])
	}

	// appending to a slice
	fmt.Println("== append command ==")
	xx = append(xx, 77, 88, 99)
	fmt.Println(xx)

	y := []int{1, 2, 3, 4, 5}
	xx = append(xx, y...)
	fmt.Println(xx)

	// deleting from a slice
	xx = append(xx[:2], xx[4:]...) // use the append but split up the slice in two parts.
	fmt.Println(xx)

	// make command
	fmt.Println("== Make command ==")
	s := make([]string, 3, 3)
	s[0] = "Oscar"
	s[1] = "Fredrik"
	s[2] = "Almgren"
	//s[3] = "d.y." // ger panic: runtime error: index out of range [3] with length 3
	fmt.Println(s, "length:", len(s), "capacity:", cap(s))

	// multi-dimension slice (tänk Excel-blad som har 2 dimensioner eller 2+ flera!)
	// vanliga slices
	fmt.Println("== Multi dimension slices ==")
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println("jb slice", jb)
	mp := []string{"Miss", "Penny", "Strawberry", "Champagne"}
	fmt.Println("mp slice", mp)

	// multi-dimension
	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
