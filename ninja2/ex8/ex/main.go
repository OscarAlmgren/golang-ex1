package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type user struct {
	First string
	Age   int
}

type person []struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

// SortByAge for person
type SortByAge person

func (a SortByAge) Len() int           { return len(a) }
func (a SortByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// SortByLast for person
type SortByLast person

func (a SortByLast) Len() int           { return len(a) }
func (a SortByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	jsonBlob, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Printf("%T\n", jsonBlob)
	fmt.Println(string(jsonBlob))

	str := []byte(`[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`)
	// fmt.Println(str)

	var persons person

	err = json.Unmarshal(str, &persons)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%+v\n", persons)
	for _, v := range persons {
		fmt.Println("Name:", v.First, v.Last, " Age:", v.Age)
		for _, saying := range v.Sayings {
			fmt.Println("Sayings:", saying)
		}
	}

	// Write to std out
	json.NewEncoder(os.Stdout).Encode(persons)

	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
