package main

import (
	"encoding/json"
	"fmt"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

// ColorGroup is a struct
type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

// Animal is a struct
type Animal struct {
	Name  string // fields must start upper case for JSON to work
	Order string
}

func main() {
	rawBlob := `[{"Name": "Platypus", "Order": "Monotremata"},{"Name": "Quoll",    "Order": "Dasyuromorphia"}]`
	jsonBlob := []byte(rawBlob)
	fmt.Printf("%T", rawBlob)
	fmt.Printf("\n%T", jsonBlob)

	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	// os.Stdout.Write(b)
	fmt.Printf("\n%T", b)
	var colors ColorGroup
	err = json.Unmarshal(b, &colors) // Address. This needs to be a pointer address to where it's stored.
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("\n%+v", colors)
	// animals := []Animal{}
	var animals []Animal

	err = json.Unmarshal(jsonBlob, &animals) //  If v is nil or not a pointer, Unmarshal returns an InvalidUnmarshalError.
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("\n%+v\n", animals)
	for i, v := range animals {
		fmt.Println("I:", i, " V:", v)
	}

	s := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(s) // sort Ints
	fmt.Println(s)
	sort.Strings(group.Colors) // sort Strings
	fmt.Println(group.Colors)

	password := "mittlösenord"
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // bs = byte slice of hashed password
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println(string(bs)) // convert bytestream to string and print
	password = "mittnyalösenord"

	compare := bcrypt.CompareHashAndPassword(bs, []byte(password))
	if compare != nil {
		fmt.Println("Err:", compare)
		fmt.Println("You can't login")
		return
	}
	fmt.Println("You're logged in")
}
