/*
Hands-on exercise #47
For this exercise, do the following:
	● Create a slice to store the names of all of the states in the United States of America.
		○ Use make and append to do this.
		○ Goal: do not have the array that underlies the slice created more than once.
	● Print out
		○ the len
		○ the cap
		○ the values, along with their index position, without using the range clause.
*/

package main

import "fmt"

func main() {

	usa := make([]string, 0, 50)

	fmt.Printf("len = %v \t cap = %v\n", len(usa), cap(usa))

	usa = append(usa, "Alabama", "Alaska", "Arizona", "Arkansas", "California",
		"Colorado", "Connecticut", "Delaware", "Florida", "Georgia",
		"Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas",
		"Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts",
		"Michigan", "Minnesota", "Mississippi", "Missouri", "Montana",
		"Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico",
		"New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma",
		"Oregon", "Pennsylvania", "Rhode Island", "South Carolina", "South Dakota",
		"Tennessee", "Texas", "Utah", "Vermont", "Virginia",
		"Washington", "West Virginia", "Wisconsin", "Wyoming")

	fmt.Printf("len = %v \t cap = %v\n", len(usa), cap(usa))

	for idx := 0; idx < len(usa); idx++ {
		fmt.Printf("%v - \t %v\n", idx, usa[idx])
	}
}
