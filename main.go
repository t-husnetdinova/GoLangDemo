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

	color := "green"
	color = "blue"

	fmt.Println(color)
	// ---------------------------------------------------------

}