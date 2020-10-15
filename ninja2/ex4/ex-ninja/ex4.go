package main

import "fmt"

func main() {
	fmt.Println("Hello")
	names := map[string][]string{
		"Oscar_Almgren": []string{"Design Thinking", "Old Fashioned", "Ida"},
	}

	names["James_Bond"] = []string{"Shaken, not stirred", "Martinis", "Women"}
	names["Miss_Money"] = []string{"James Bond", "Literature", "Computer Science"}
	names["Dr_No"] = []string{"Being evil", "Ice cream", "Sunsets"}

	for i, v := range names {
		fmt.Println("Index:", i)
		for _, interest := range v {
			fmt.Println("Interest:", interest)
		}
	}
}
