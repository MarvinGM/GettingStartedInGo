package main

import "fmt"

func main() {

	//anonymous struct
	s := struct {
		first     string         //basic field inside the struct
		friends   map[string]int //map inside the struct
		favDrinks []string       //slice inside the struct
	}{
		first: "James", //putting value to the basic field inside the struct

		friends: map[string]int{ //putting values to the map inside the struct
			"Marvin": 111,
			"Q":      222,
			"M":      333,
		},

		favDrinks: []string{ //putting values to the slice inside the struct
			"Martini",
			"Water",
		},
	}

	fmt.Println(s.first)     //display the value of basic field
	fmt.Println(s.friends)   //display the values of map
	fmt.Println(s.favDrinks) //display the values of slice

	for k, v := range s.friends { //breaking down the values inside the map
		fmt.Println(k, v)
	}

	for i, val := range s.favDrinks { //breaking down the values inside the slice
		fmt.Println(i, val)
	}

}
