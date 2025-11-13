package main

import "fmt"

func sumArray(ar []uint32) uint32 {
	var result uint32 = 0

	for _, value := range ar {
		result += value
	}

	return result
}

func main() {
	var arCount uint32
	var ar []uint32
	var tempValue uint32

	fmt.Scanf("%d", &arCount)

	for i := 0; i < int(arCount); i++ {
		fmt.Scanf("%d", &tempValue)
		ar = append(ar, tempValue)
	}

	result := sumArray(ar)
	fmt.Println(result)
}
