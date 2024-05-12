package examples

import "fmt"

// Here’s a function that will take an arbitrary number of ints as arguments.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	// Within the function, the type of nums is equivalent to []int. We can call len(nums), iterate over it with range, etc.
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func VariadicFnctions() {
	fmt.Println("\n\n" + "--------------------------------------------------------------")
	fmt.Println("	14 - V A R I A D I C   F U N C T I O N S ")
	fmt.Println("--------------------------------------------------------------")

	// Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.
	// Variadic functions can be called in the usual way with individual arguments.
	sum(1, 2)
	sum(1, 2, 3)

	// If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	// Another key aspect of functions in Go is their ability to form closures, which we’ll look at next.
}
