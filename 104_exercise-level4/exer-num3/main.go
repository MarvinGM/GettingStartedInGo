package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	//slicing from index 42 to 45
	fmt.Println(x[:5])

	//slicing from index 47 to 51
	fmt.Println(x[5:])

	//slicing from index 44 to 48
	fmt.Println(x[2:7])

	//slicing from index 43 to 47
	fmt.Println(x[1:6])

}
