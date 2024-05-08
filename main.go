package main

import (
	"fmt"
	"os"
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

	// EXERCISES AND SOLUTIONS HERE: https://github.com/inancgumus/learngo/tree/master

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
	// var (
	// 	lang    string
	// 	version int
	// )
	// lang, version = "go", 2
	// fmt.Println(lang, "version", version)
	// ---------------------------------------------------------

	// ---------------------------------------------------------
	// EXERCISE: Multi Assign #2
	//
	//  1. Assign the correct values to the variables
	//     to match to the EXPECTED OUTPUT below
	//  2. Print the variables
	// HINT
	//  Use multiple Println calls to print each sentence.
	//
	// EXPECTED OUTPUT
	//  Air is good on Mars
	//  It's true
	//  It is 19.5 degrees

	// SOLUTION
	// var (
	// 	planet string
	// 	isTrue bool
	// 	temp   float64
	// )
	// planet, isTrue, temp = "Mars", true, 19.5
	// fmt.Println("Air is good on ", planet)
	// fmt.Println("It's ", isTrue)
	// fmt.Println("It is ", temp, " degrees")
	// ---------------------------------------------------------

	// ---------------------------------------------------------
	// EXERCISE: Multi Short Func
	//
	// 	1. Declare two variables using multiple short declaration syntax
	//  2. Initialize the variables using `multi` function below
	//  3. Discard the 1st variable's value in the declaration
	//  4. Print only the 2nd variable
	// NOTE
	//  You should use `multi` function
	//  while initializing the variables
	//
	// EXPECTED OUTPUT
	//  4

	// SOLUTION
	// ADD YOUR DECLARATIONS HERE
	// _, b := multi()
	// fmt.Println(b)
	// ---------------------------------------------------------

	// ---------------------------------------------------------
	// EXERCISE: Swapper
	//  1. Change `color` to "orange"
	//     and `color2` to "green" at the same time
	//     (use multiple-assignment)
	//  2. Print the variables
	//
	// EXPECTED OUTPUT
	//  orange green

	// SOLUTION
	// color, color2 := "red", "blue"
	// color, color2 = "orange", "green"
	// fmt.Println(color, color2)
	// ---------------------------------------------------------

	// ---------------------------------------------------------
	// EXERCISE: Swapper #2
	//  1. Swap the values of `red` and `blue` variables
	//  2. Print them
	//
	// EXPECTED OUTPUT
	//  blue red

	// SOLUTION
	// red, blue := "red", "blue"
	// red, blue = "blue", "red"
	// fmt.Println(red, blue)
	// ---------------------------------------------------------

	// ---------------------------------------------------------
	// EXERCISE: Discard The File
	//  1. Print only the directory using `path.Split`
	//  2. Discard the file part
	//
	// RESTRICTION
	//  Use short declaration
	//
	// EXPECTED OUTPUT
	//  secret/

	// SOLUTION
	// dir, _ := path.Split("secret/file.txt")
	// fmt.Println(dir)
	// ---------------------------------------------------------

	// ---------------------------------------------------------
	// EXERCISE: Convert and Fix
	//  Fix the code by using the conversion expression.
	//
	// EXPECTED OUTPUT
	//  15.5

	// SOLUTION
	// a, b := 10, 5.5
	// fmt.Println(float64(a) + b)
	// ---------------------------------------------------------

	// ---------------------------------------------------------
	// EXERCISE: Convert and Fix #2
	//  Fix the code by using the conversion expression.
	//
	// EXPECTED OUTPUT
	//  10.5

	// SOLUTION
	// a, b := 10, 5.5
	// a = int(b)
	// fmt.Println(float64(a) + b)
	// ---------------------------------------------------------

	// ---------------------------------------------------------
	// EXERCISE: Convert and Fix #3
	//  Fix the code.
	//
	// EXPECTED OUTPUT
	//  5.5

	// SOLUTION
	// fmt.Println(5.5)
	// ---------------------------------------------------------

	// ---------------------------------------------------------
	// EXERCISE: Convert and Fix #4
	//  Fix the code.
	//
	// EXPECTED OUTPUT
	//  9.5

	// SOLUTION
	// age := 2
	// fmt.Println(7.5 + float64(age))
	// ---------------------------------------------------------

	// ---------------------------------------------------------
	// EXERCISE: Convert and Fix #5
	//  Fix the code.
	//
	// HINTS
	//   maximum of int8  can be 127
	//   maximum of int16 can be 32767
	//
	// EXPECTED OUTPUT
	//  1127

	// SOLUTION
	// min := int8(127)
	// max := int16(1000)

	// FIX THE CODE HERE
	// fmt.Println(int8(max) + min) // this prints 103
	// fmt.Println(max + int16(min))
	// ---------------------------------------------------------

	// ---------------------------------------------------------
	// EXERCISE: Count the Arguments
	//  Print the count of the command-line arguments
	// INPUT
	//  bilbo balbo bungo
	//
	// EXPECTED OUTPUT
	//  There are 3 names.

	// UNCOMMENT & FIX THIS CODE
	// count := len(os.Args) - 1

	// // UNCOMMENT IT & THEN DO NOT TOUCH THIS CODE
	// fmt.Printf("There are %d names.\n", count)

	// run like this:
	// $ go run main.go bilbo balbo bungo
	// ---------------------------------------------------------

	// ---------------------------------------------------------
	// EXERCISE: Print the Path
	//
	//  Print the path of the running program by getting it
	//  from `os.Args` variable.
	//
	// HINT
	//  Use `go build` to build your program.
	//  Then run it using the compiled executable program file.
	//
	// EXPECTED OUTPUT SHOULD INCLUDE THIS
	//  myprogram

	// SOLUTION
	// fmt.Println(os.Args[0])
	// ---------------------------------------------------------


	// ---------------------------------------------------------
	// EXERCISE: Print Your Name
	//
	//  Get it from the first command-line argument
	//
	// INPUT
	//  Call the program using your name
	// EXPECTED OUTPUT
	//  It should print your name
	// EXAMPLE
	//  go run main.go inanc
	//    inanc
	// BONUS: Make the output like this:
	//
	//  go run main.go inanc
	//    Hi inanc
	//    How are you?

	// SOLUTION
	fmt.Println(os.Args[1])

	// bonus
	fmt.Println("Hello", os.Args[1])
	fmt.Println("How are you?")
	// ---------------------------------------------------------
}

// THIS IS FOR EXCERCISE: Multi Short Func (must be called outside of main)
// commenting this now because I might forget after I add a million other little exercises to this and comment this out
// Hi, Future Tammy :) Hope you are having a good day!
// func multi() (int, int) {
// 	return 5, 4
// }