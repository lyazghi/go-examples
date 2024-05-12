package main

import (
	"fmt"
	"go-examples/examples"
)

// To run the program, put the code in hello-world.go and use go run.
// Sometimes we’ll want to build our programs into binaries. We can do this using go build.
// We can then execute the built binary directly.

func main() {
	fmt.Println("\n\n" + "----------------------------------------------------------")
	fmt.Println(" 		G O L A N G  E X A M P L E S")
	fmt.Println("----------------------------------------------------------")

	// 2 - VALUES
	examples.Values()
	// 3 - Variables
	examples.Variables()
	// 4 - CONSTANTS
	examples.Constants()
	// 5 - FOR
	examples.For()
	// 6 - IF ELSE
	examples.IfElse()
	// 7 - SWITCH
	examples.Switch()
	// 8 - ARRAYS
	examples.Arrays()
	// 9 - SLICES
	examples.Slices()
	// 10 - MAPS
	examples.Maps()
	// 11 - RANGE
	examples.Range()
	// 12 - FUNCTIONS
	examples.Functions()
	// 13 - MULTIPLE RETURN VALUES
	examples.MultipleReturnValues()
	// 14 - VARIADIC FUNCTIONS
	examples.VariadicFnctions()
	// 15 - CLOSURES
	examples.Closures()
	// 16 - Recursion
	examples.Recursion()
	// 17 - Pointers
	examples.Pointers()
	// 18 - String & Runes
	examples.StringsAndRunes()
}
