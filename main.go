package main

import ("fmt"
		// "github.com/t-husnetdinova/GoLangDemo/hello"
		// "rsc.io/quote"
		// "os"
)

func main() {
	// call HelloWorld function
	// fmt.Println(hello.HelloWorld())

	// ---------------------------------------------------------

	// test out quote function
	// fmt.Println(quote.Go())

	// ---------------------------------------------------------

	// type conversion example
	// speed := 100 // int
	// force := 2.5 // float64
	// // first convert speed to float64, then convert it back to int to assing to itself
	// speed = int(float64(speed) * force)
	// fmt.Println(speed) 

	// ---------------------------------------------------------
	
	// Args
	// args's type is a slice of a string (a slice can store multiple string values, they are unnamed)
	// to see the difference between printf and println
	// fmt.Printf("%#v\n", os.Args)
	// fmt.Println(os.Args)

	// fmt.Println("Path:", os.Args[0])
	// fmt.Println("1st argument:", os.Args[1])
	// fmt.Println("2nd argument:", os.Args[2])
	// fmt.Println("3rd argument:", os.Args[3])
	// fmt.Println("Number of items inside os.Args", len(os.Args))


	// ---------------------------------------------------------
	// EXERCISE: Make It Blue
	//
	//  1. Change `color` variable's value to "blue"
	//  2. Print it
	//
	// EXPECTED OUTPUT
	//  blue

	// SOLUTION
	// color := "green"
	// color = "blue"
	// fmt.Println(color)
	// ---------------------------------------------------------


	// ---------------------------------------------------------
	// EXERCISE: Variables To Variables
	//
	//  1. Change the value of `color` variable to "dark green"
	//  2. Do not assign "dark green" to `color` directly
	//     Instead, use the `color` variable again while assigning to `color`
	//  3. Print it
	//
	// RESTRICTIONS
	//  WRONG ANSWER, DO NOT DO THIS:
	//  `color = "dark green"`
	//
	// HINT
	//  + operator can concatenate string values
	//
	//  Example:
	//  "h" + "e" + "y" returns "hey"
	//
	// EXPECTED OUTPUT
	//  dark green

	// SOLUTION
	// color := "green"
	// fmt.Println("dark " + color)
	// ---------------------------------------------------------


	// ---------------------------------------------------------
	// EXERCISE: Assign With Expressions
	//
	//  1. Multiply 3.14 with 2 and assign it to `n` variable
	//  2. Print the `n` variable
	//
	// HINT
	//  Example: 3 * 2 = 6
	//  * is the product operator (it multiplies numbers)
	//
	// EXPECTED OUTPUT
	//  6.28

	// SOLUTION
	// n := 0.

	// answer needs to be a float64 assigned to variable n
	// n = 3.14*2
	// fmt.Println(n)
	// --------------------------------------------------------


	// --------------------------------------------------------
	// EXERCISE: Find the Rectangle's Perimeter
	//
	//  1. Find the perimeter of a rectangle
	//     Its width is  5
	//     Its height is 6
	//  2. Assign the result to the `perimeter` variable
	//  3. Print the `perimeter` variable
	// HINT
	//  Formula = 2 times the width and height
	// EXPECTED OUTPUT
	//  22
	
	// SOLUTION
	// var (
	// 	perimeter int
	// 	width, height = 5, 6
	// )

	// perimeter = 2 * (width + height)
	// fmt.Println(perimeter)	
	// ---------------------------------------------------------


	// ---------------------------------------------------------
	// EXERCISE: Multi Assign
	//
	//  1. Assign "go" to `lang` variable
	//     and assign 2 to `version` variable
	//     using a multiple assignment statement
	//  2. Print the variables
	// EXPECTED OUTPUT
	//  go version 2

	// SOLUTION
	var (
		lang    string
		version int
	)
	lang, version = "go", 2
	fmt.Println(lang, "version", version)
	// ---------------------------------------------------------
}