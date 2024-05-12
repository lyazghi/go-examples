package examples

import "fmt"

// This fact function calls itself until it reaches the base case of fact(0).
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func Recursion() {
	fmt.Println("\n\n" + "--------------------------------------------------------------")
	fmt.Println("	16 -  R E C U R S I O N ")
	fmt.Println("--------------------------------------------------------------")

	// Go supports recursive functions. Here’s a classic example.
	fmt.Println(fact(7))

	// Closures can also be recursive, but this requires the closure to be declared with a typed var explicitly before it’s defined.
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		// Since fib was previously declared in main, Go knows which function to call with fib here.
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}
