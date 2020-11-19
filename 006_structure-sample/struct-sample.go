package main

import "fmt"

//structure encapsulation
type Car struct {
	Name    string
	Age     int
	ModelNo int
}

//method a1 print
func (a1 Car) Print() {
	fmt.Println(a1)
}

//method a2 drive
func (a2 Car) Drive() {
	fmt.Println("Driving...")
}

//method a1 getname
func (a1 Car) getName() string {
	return a1.Name
}

func main() {
	//initialization
	a1 := Car{"Chevy", 1, 2}
	fmt.Println("Printing normally a1: ", a1)

	//another initialization which is better
	a2 := Car{
		Name:    "Toyota",
		Age:     3,
		ModelNo: 4,
	}
	fmt.Println("Printing normally a2: ", a2)

	//calling in another methods
	//printing by calling the method a1 print
	a1.Print()

	//printing by calling the method a2 print
	a2.Drive()

	//getting specific variable inside the structure
	fmt.Println("getName of a1.Name: ", a1.getName())

}
