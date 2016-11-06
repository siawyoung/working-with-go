/**
 * control-structures.go
 * If, Else, Switch and Conditionals
 */

// standard stuff
package main

import "fmt"

func main() {

	// define an int variable
	num := 1

	// Go is not picky, conditionals don't require parentheses
	if num > 3 {
		fmt.Println("Many")
	}

	// Go is picky, "else" must be on the same line as closing if bracket
	if num == 1 {
		fmt.Println("One")
	} else if num == 2 {
		fmt.Println("Two")
	} else {
		fmt.Println("Many")
	}

	// Switch statement takes conditionals and auto breaks
	// it compares each case against the expression provided
	// e.g. num == 1
	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Way too many")
	}

	// If the expression is omitted, go assumes it as `true`
	// i.e. it compares each case against  `true`
	// e.g. true == (num == 1)
	switch {
	case num == 1:
		fmt.Println("One")
	case num == 2:
		fmt.Println("Two")
	case num > 2:
		fmt.Println("Many")
	default:
		fmt.Println("Thrown over boat")
	}

}
