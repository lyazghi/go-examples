package examples

import "fmt"

// Functions are central in Go. We’ll learn about functions with a few different examples.
// Here’s a function that takes two ints and returns their sum as an int.
func plus(a int, b int) int {
	// Go requires explicit returns, i.e. it won’t automatically return the value of the last expression.
	return a + b
}

// When you have multiple consecutive parameters of the same type, you may omit the type name for the like-typed parameters up to the final parameter that declares the type.
func plusPlus(a, b, c int) int {
	return a + b + c
}

func Functions() {
	fmt.Println("\n\n" + "----------------------------------")
	fmt.Println("	12 - F U N C T I O N S ")
	fmt.Println("----------------------------------")

	// Call a function just as you’d expect, with name(args).
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	// There are several other features to Go functions. One is multiple return values, which we’ll look at next.
}
