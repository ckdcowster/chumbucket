/*
Hands-on exercise #48
	Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
	slice:
		"James", "Bond", "Shaken, not stirred"
		"Miss", "Moneypenny", "I'm 008."
*/

package main

import "fmt"

func main() {

	x1 := []string{"James", "Bond", "Shaken, not stirred"}
	x2 := []string{"Miss", "Moneypenny", "I'm 008."}
	xss := [][]string{x1, x2}

	for x, item := range xss {
		fmt.Printf("%v - %v\n", x, item)
		for y, value := range item {
			fmt.Printf("%v -\t %v\n", y, value)
		}
	}
}
