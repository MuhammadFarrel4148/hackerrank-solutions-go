package main

import "fmt"

func sumArray(ar []uint64) uint64 {
	var result uint64
	result = 0
	
	for _, value := range ar {
		result += value
	}

	return result
}

func main() {
	var arCount uint32
	fmt.Scan(&arCount);

	var ar = make([]uint64, arCount)

	for i := 0; i < int(arCount); i++ {
		fmt.Scan(&ar[i])
	}

	result := sumArray(ar)
	fmt.Println(result)
}