package main

import "fmt"

func variables() {
	var name string = "GoLang" 
	age :=10
	// Constant
	const pi = 3.142

	// fmt.Println("Language:", name)
	// fmt.Println("Age:", age)
	// fmt.Println("Pi:", pi)

	// Used Printf instaed of Println becuase Printf is for pritning formatted texts.
	fmt.Printf("Language: %s, Age: %d, Pi: %.3f\n", name, age, pi)
}