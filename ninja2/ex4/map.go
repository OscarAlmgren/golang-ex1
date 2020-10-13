package main

import "fmt"

func main() {

	fmt.Println("\n== map command ==")
	m := map[int]string{
		131: "Oscar Almgren",
		93:  "Andreas Bergman",
		1:   "Mårten Sundmark",
	}
	fmt.Println("m:", m)
	fmt.Println("\n== comma ok idiom ==")
	_, ok := m[100]
	fmt.Println("found:", ok)
	fmt.Println("\n== add to map ==")
	m[100] = "Marcus Molander"
	fmt.Println("\n== range over m ==")
	for k, v := range m {
		fmt.Println("key:", k, "value:", v)
	}
	fmt.Println("\n== delete from map ==")
	delete(m, 1)
	fmt.Println(m)

	fmt.Println("\n== Make map command ==")
	elements := make(map[string]string) // map[key-type]value-type
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Flourine"
	elements["Ne"] = "Neon"

	fmt.Println(elements)
	fmt.Println(elements["Li"])

	name, ok := elements["Un"] // hitta name, bool svar om det var OK eller EJ OK.
	fmt.Println(name, ok)
}
