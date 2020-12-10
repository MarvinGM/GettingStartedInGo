package main

import "fmt"

type person struct {
	fname  string
	lname  string
	flavor []string
}

func main() {
	p1 := person{
		fname: "James",
		lname: "Bond",
		flavor: []string{
			"Pandan",
			"Ube",
		},
	}

	p2 := person{
		fname: "Marvin",
		lname: "Mahilum",
		flavor: []string{
			"Chocolate",
			"Mango",
		},
	}

	m := map[string]person{
		p1.lname: p1,
		p2.lname: p2,
	}

	for k, v := range m {
		fmt.Println("Key in map: ", k)
		fmt.Println(v.fname)
		fmt.Println(v.lname)
		for i, val := range v.flavor {
			fmt.Println(i, val)
		}
		fmt.Println("----------")
	}
}
