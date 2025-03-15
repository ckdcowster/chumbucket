/*
Hands-on exercise #41
	● Use a slice literal to store these elements in a slice, then also print out the slice and the
		length of the slice.
	● if you can, try to get a "for range" loop working on the slice
		○ you can reference the documentation for a "for range" loop here:
			■ https://go.dev/doc/effective_go#for
*/

package main

import "fmt"

func main() {

	ice_creams := []string{
		"Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie",
		"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)",
		"Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)",
		"Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)",
		"Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)",
		"Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different",
		"Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)",
	}

	fmt.Println(ice_creams)
	fmt.Printf("size of ice_creams is %v\n", len(ice_creams))
	fmt.Printf("type of ice_creams is %T\n", ice_creams)

	for idx, value := range ice_creams {
		fmt.Printf("%v - %v\n", idx, value)
	}
}
