package main

import "fmt"

// To run the program, put the code in hello-world.go and use go run.
// Sometimes we’ll want to build our programs into binaries. We can do this using go build.
// We can then execute the built binary directly.

func main() {
	fmt.Println("\n\n" + "----------------------------------------------------------")
	fmt.Println(" 		G O L A N G  E X A M P L E S")
	fmt.Println("----------------------------------------------------------")

	// 2 - Values
	Values()
	// 3 - Variables
	Variables()
	// 4 - Constants
	Constants()
	// 5 - For
	For()
	// 6 - If Else
	IfElse()
	// 7 - Switch
	Switch()
	// 8 - Arrays
	Arrays()
	// 9 - Slices
	Slices()
}
