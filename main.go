/*
Hands-on exercise #42
	● Using a COMPOSITE LITERAL:
		○ create an ARRAY which holds 5 VALUES of TYPE int
		○ assign VALUES to each index position.
	● Range over the array and print the values out.
		○ print out the VALUE and the TYPE
*/

package main

import "fmt"

func main() {

	ai := [5]int{}

	ai[0] = 100
	ai[1] = 101
	ai[2] = 102
	ai[3] = 103
	ai[4] = 104

	fmt.Println(ai)
	fmt.Printf("%T\n", ai)
	for idx, val := range ai {
		fmt.Printf("index = %v \t type = %T \t value = %v\n", idx, val, val)
	}

}
