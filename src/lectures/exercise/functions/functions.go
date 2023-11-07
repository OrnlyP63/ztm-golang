//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func myName(name string) {
	fmt.Println("Your name is", name)
}

func twoTwos() (int, int) {
	return 2, 2
}

func addThreeNumer(a, b, c int) int {
	return a + b + c
}

func five() int {
	return 5
}

func hiThere() string {
	return "Hi there!"
}
func main() {
	myName("Phiphat")
	fmt.Println(hiThere())

	a, b := twoTwos()

	answer := addThreeNumer(five(), a, b)
	fmt.Println(answer)
}
