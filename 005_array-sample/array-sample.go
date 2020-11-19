package main

import (
	"fmt"
)

func todo() {
	//var arr []int - declare an array
	//array initialization
	arr := []int{1, 2, 3, 4}
	arr2 := []string{"hi", "my", "name"}

	//append value or data in an array
	arr = append(arr, 5)
	arr2 = append(arr2, "is marvin")

	//printing the array
	fmt.Println(arr, arr2)
}

func main() {
	//calling the todo function at the top
	todo()

}
