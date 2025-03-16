/*
Hands-on exercise #43
	● Using a COMPOSITE LITERAL:
	● create a SLICE of TYPE int
	● assign these 10 VALUES
		42, 43, 44, 45, 46, 47, 48, 49, 50, 51
	● Range over the slice and print the values out.
		○ print out the VALUE and the TYPE
*/

package main

import "fmt"

func main() {

	ai := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Printf("%#v\n", ai)
	for idx, val := range ai {
		fmt.Printf("index = %v \t type = %T \t value = %v\n", idx, val, val)
	}

}
