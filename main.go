/*
Hands-on exercise #44
	Using the code from the previous example, use SLICING to create the following new slices
	which are then printed:
		● [42 43 44 45 46]
		● [47 48 49 50 51]
		● [44 45 46 47 48]
		● [43 44 45 46 47]
*/

package main

import "fmt"

func main() {

	ai := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	s1 := ai[:5]
	fmt.Printf("%#v\n", s1)

	s2 := ai[5:]
	fmt.Printf("%#v\n", s2)

	s3 := ai[2:7]
	fmt.Printf("%#v\n", s3)

	s4 := ai[1:6]
	fmt.Printf("%#v\n", s4)

}
