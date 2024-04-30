package main

import "fmt"
import "github.com/t-husnetdinova/GoLangDemo/hello"
import "rsc.io/quote"

func main() {
	// call HelloWorld function
	fmt.Println(hello.HelloWorld())

	
	// test out quote function
	fmt.Println(quote.Go())


	// type conversion example
	speed := 100 // int
	force := 2.5 // float64
	// first convert speed to float64, then convert it back to int to assing to itself
	speed = int(float64(speed) * force)
	fmt.Println(speed) 


	//
}