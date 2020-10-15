package main

import "fmt"

// create a new type (struct). complex data type
type person struct {
	// Field names may be...

	first string // explicitly (IdentifierList)
	last  string
	age   int
}

// combine structs to new types
type secretAgent struct {
	// Field names may be...
	person //  implicitly (EmbeddedField)
	ltk    bool
}

// structs
func main() {
	// create a _value_ of type "person" (object)
	p1 := person{
		first: "Oscar",
		last:  "Almgren",
		age:   38,
	}
	p2 := person{
		first: "Ida",
		last:  "Almgren",
		age:   29,
	}
	sa := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   47,
		},
		ltk: true,
	}
	// anonymous struct
	diplomat := struct {
		first   string
		last    string
		country string
	}{
		first:   "Andreas",
		last:    "Bergman",
		country: "Sverige",
	}
	fmt.Println(p1, "\n", p2, "\n", sa, "\n", diplomat)
}
