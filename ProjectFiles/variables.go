package main

import "fmt"

// func variables() {
// 	var name string = "GoLang" 
// 	age :=10
// 	// Constant
// 	const pi = 3.142

// 	// Used Printf instaed of Println becuase Printf is for pritning formatted texts.
// 	fmt.Printf("Language: %s, Age: %d, Pi: %.3f\n", name, age, pi)
// }


// Section 2
// func variables() {
// 	var FavNumber int = 7
// 	var ChoiceDec float64 = 3.21
// 	Golang := false

// 	fmt.Println("Favorite Number: ", FavNumber)
// 	fmt.Printf("Favorite Decimal: %.1f\n", ChoiceDec)
// 	fmt.Println("Do I like GoLang?: ", Golang)
// }

// Using if else Conditionals
// func variables() {
// 	age := 66

// 	if age >= 18 && age <= 65  {
// 		fmt.Println("You are an Adult")
// 	} else if age < 14 && age < 12 {
// 		fmt.Println("You are a child")
// 	} else {
// 		fmt.Println("You are a senior")
// 	}
// }


// Using Switch Conditionals
// func variables() {
// 	age := 20

// 	switch age {
// 	case 10: 
// 	fmt.Println("You are a Baby")
// 	case 16: 
// 	fmt.Println("You are a Toddler")
// 	default:
// 		fmt.Println("You are an Adult")
// 	}
// }

// Write a program to check if a number is positive, negative, or zero.
func variables() {
	num := -1

	if num > 0 {
		fmt.Println("Number is Positive")
	} else if num < 0 {
		fmt.Println("Number is Negative")
	} else {
		fmt.Println("Number is zero")
	}
}