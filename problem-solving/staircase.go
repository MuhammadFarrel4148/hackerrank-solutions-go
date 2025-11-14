package main

import "fmt"

func staircase(n uint32) {
	for i := 1; i <= int(n); i++ {
		for j := 0; j < int(n) - i; j++ {
			fmt.Printf(" ")
		}

		for k := int(n) - i; k <= int(n) - 1; k++ {
			fmt.Printf("#")
		}

		fmt.Println()
	}
}

func main() {
	var n uint32
	fmt.Scan(&n)

	staircase(n)
}