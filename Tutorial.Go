package main

import (
	"fmt"
)

//To run a Go program, go into your terminal and type "go run <the name of your file.go>" example: for this Tutorial.Go, we would type: "go run Tutorial.Go"
// This is a global variable, make sure to specify the variable type.
var num1 int = 5

var num2 int = 8

func main() {

	fmt.Println("saikyonanda")

	//basically it is going to add the amount and the := will elminate the need
	//to specify the variable type
	//u have to put "print" as "Print" cuz its gay like that

	sum := num1 + num2
	fmt.Println(sum)

	if num1 < num2 {
		fmt.Println("working comparison function using global variables")
	}

	fmt.Println("This is an array in GO")

	//initialize an array
	a := []int{8, 8, 8}

	//append to a list, it produces a new list btw dosen"t alter og list
	a = append(a, 18)

	fmt.Println(a)

	//this is a map in go
	// a map in go is similar to a dictionary in python
	fmt.Println("This is a map in Go")
	fruits := make(map[string]int)

	fruits["apple"] = 2
	fruits["mango"] = 4
	fruits["kiwi"] = 8

	//deletes the apple item in the dict of "fruits"
	delete(fruits, "apple")

	fmt.Println(fruits)

	//only type of looop in go is for loop
	//this loop goes up to 8

	for i := 0; i < 8; i++ {
		fmt.Println(i)
	}

	//can also make a while loop like this

	abg := 0
	bods := 6

	for abg < bods {
		fmt.Println(abg)
		abg++

	}

	//this is how we iterate over each elemnt of an array
	//bro this language wants to be "used" soo badly bro i swear

	arr := []string{"x", "y", "z"}

	for index, value := range arr {
		fmt.Println("index", index, "value:", value)
	}

	//iterating over each elemnt of a map; u get the key instead of the index
	bob := make(map[string]string)
	bob["a"] = "alpha"
	bob["b"] = "beta"

	for key, value := range bob {
		fmt.Println("key", key, "value:", value)
	}

	result := sob(6, 3)
	fmt.Println(result)

	print()
}

func sob(x int, y int) int {
	return x + y
}

// function to access a global variable
//also you have to use double quotes for print statements

/*func nexus() {

	if num1 < num2 {
		fmt.Print("working comparison function using global variables")
	}

}*/

//questions to think about:
//why do files in Golang affect each other?
//How to validate user input in GoLang? How do I check if their input is actually a number or a string?
//modulo operator not built in?
//how to run without typing go run.file name?
