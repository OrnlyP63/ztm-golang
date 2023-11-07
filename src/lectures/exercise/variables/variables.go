//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var myFavoriteColor = "black"
	fmt.Println("my favorite color is", myFavoriteColor)

	birthYear, age := 1995, 28
	fmt.Println("Born in", birthYear, "aged", age)

	var (
		first = "P"
		last  = "C"
	)
	fmt.Println("Intitials", first, last)

	var myAgeInDays int
	myAgeInDays = 365 * age
	fmt.Println(myAgeInDays)

}
