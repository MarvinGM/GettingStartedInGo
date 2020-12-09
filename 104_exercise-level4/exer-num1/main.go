package main

import "fmt"

func main() {

	zxc := [5]int{3, 4, 5, 6, 42}
	for i, v := range zxc {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", zxc)
}
