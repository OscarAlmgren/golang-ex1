package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

type broder struct {
	Number    int
	FirstName string
	LastName  string
	Rank      int
}

const (
	fris = iota
	lkfv
	rkfv
	kkfv
	k1klkfv
	kmskkfv
)

func main() {
	br131 := broder{131, "Oscar Fredrik", "Almgren", rkfv}
	br1 := broder{1, "Mårten", "Sundmark", kmskkfv}
	br93 := broder{93, "Andreas", "Bergman", k1klkfv}

	fmt.Println("Br131:", unsafe.Sizeof(br131))

	brothers := []broder{br131, br1, br93}
	fmt.Println("Brothers:", unsafe.Sizeof(brothers), "len:", len(brothers), "cap:", cap(brothers))
	fmt.Println(brothers)

	bs, err := json.Marshal(brothers)
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println(string(bs))

	var jsonBrothers []broder

	jsonError := json.Unmarshal(bs, &jsonBrothers)
	if jsonError != nil {
		fmt.Println("Err:", jsonError)
	}
	fmt.Println(jsonBrothers)
}
