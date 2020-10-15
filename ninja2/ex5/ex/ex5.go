package main

import "fmt"

type person struct {
	first    string
	last     string
	flavours []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	p1 := person{
		first:    "Oscar",
		last:     "Almgren",
		flavours: []string{"Vanilla", "Chocolate"},
	}
	fmt.Println(p1)
	for _, v := range p1.flavours {
		fmt.Println(v)
	}

	truck := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}

	sedan := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: false,
	}

	fmt.Println(truck.doors)
	fmt.Println(sedan.color)

	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Daniel":   12345,
			"Jonathan": 98765,
		},
		favDrinks: []string{
			"Martini",
			"Whisky",
		},
	}
	for i, v := range s.friends {
		fmt.Println("Namn:", i, "\tnr:", v)
	}
	for _, v := range s.favDrinks {
		fmt.Println(v)
	}
}
