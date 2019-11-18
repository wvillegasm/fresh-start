package main

import "fmt"

func main() {
	x := []int8{20, 30, 40, 50}

	fmt.Printf("%T\n", x)
	fmt.Println(x)

	for i, v := range x[:2] {
		fmt.Println("Index:", i, "Value:", v)
	}

	x = append(x, 5,6,7,8,9)
	fmt.Println(x)

	y := []int8{51,52,53,54}

	x = append(x, y...)

	fmt.Println(x)
}
