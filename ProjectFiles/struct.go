package main

import "fmt"


type Person struct {
	name string
	age int
	occupation string
}

func structMain() {
	p := Person {
		name: "Ivie",
		age: 25,
		occupation: "Software Engineer",
	}

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("Occupation:", p.occupation)

	// Modify fields
	p.age = 30
	fmt.Println("Modiefied Field:", p.age)


	// Classwork
	type BookTitle struct {
		Title string
		Author string
		Pages int
	}

	book := BookTitle {
		Title: "Lord of the Rings",
		Author: "J.R.R. Tolkien",
        Pages: 11789,
	}

	fmt.Println("Title Of Book:", book.Title)
	fmt.Println("Author Of Book:", book.Author)
	fmt.Println("Pages Of Book:", book.Pages)
}