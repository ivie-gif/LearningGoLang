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
// func arraysSliceMaps() {
// 	favouriteBooks := [3]string{"Bible", "Novels", "Magazines"}

// 	for i, book := range favouriteBooks {
// 		fmt.Println("Favourite Book", i+1, ":", book)
// 	}

// }

// Learning Slice
// func arraysSliceMaps() {
// 	arr :=[9]int {1,2,3,4,5,6,7,8,9}
// 	slice := arr[3:8]

// 	fmt.Println("Slice:", slice)

// 	numb :=[]int{10,20,30,40,50}
// 	numb = append(numb, 60, 70, 80)

// 	fmt.Println("Numbers:", numb)


// 	// Classwork
// 	arrMovie :=[]string{"Avengers", "Spiderman", "Batman", "Superman", "Ironman"}
// 	arrMovie = append(arrMovie, "Hulk", "Thor")
	
// 	MovieSlice := arrMovie[1:3]
	
// 	fmt.Println("Movies:", arrMovie)
// 	fmt.Println("Slice:", MovieSlice)
// }


func arraysSliceMaps() { 
	capitals := map[string] string{
		"Nigeria": "Abuja",
		"Ghana": "Accra",
		"Togo": "Lome",
		"Benin": "Cotonou",
		"South Africa": "Cape Town",
	}
	capitals["Egypt"] = "Cairo"

	fmt.Println("Capitals:", capitals)
	fmt.Println("Capital of Nigeria:", capitals["Nigeria"])


	if capital, found := capitals["Zambia"]; found {
		fmt.Println("Capital of Zambia:", capital)
	} else {
		fmt.Println("Capital of Zambia not found")
	}


	// Classwork
	countries := map[string]string {
		"Canada": "Ottawa",
		"United States": "Washington D.C",
		"Mexico": "Mexico City",
	}

	countries["Brazil"] = "Brasilia"
	
	fmt.Println("Countries:", countries)


	if country, exists := countries["Mexico"]; exists {
		fmt.Println("Capital of Mexico:", country)
	} else {
		fmt.Println(`Capital of Mexico not found`)
	}


}