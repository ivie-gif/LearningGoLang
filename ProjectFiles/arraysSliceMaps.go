package main

import "fmt"

// func arraysSliceMaps() {
// 	var numbers [5] int
// 	numbers[0] = 2
// 	numbers[1] = 4
// 	numbers[2] = 6
// 	numbers[3] = 8
// 	numbers[4] = 10

// 	fruits := [5]string{"Apple", "Banana", "Orange", "Grapes", "Mango"}
	
// 	fmt.Println("Numbers:", numbers[2])
// 	fmt.Println("Fruits:", fruits[3])

// }


// Create an array to store your top 3 favorite books.
// Print each book using a loop.
func arraysSliceMaps() {
	favouriteBooks := [3]string{"Bible", "Novels", "Magazines"}

	for i, book := range favouriteBooks {
		fmt.Println("Favourite Book", i+1, ":", book)
	}

}