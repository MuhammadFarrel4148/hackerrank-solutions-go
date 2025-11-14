package main

import (
	"fmt"
)

func diagonalDifference(ar [][]int32, arCount uint32) int32 {
	var leftSum int32 = 0
	var rightSum int32 = 0
	var i uint32

	for i = 0; i < arCount; i++ {
		leftSum += ar[i][i]
		rightSum += ar[i][arCount-1-i]
	}

	if leftSum > rightSum {
		return leftSum - rightSum
	}
	return rightSum - leftSum
}

func main() {
	var arCount uint32
	fmt.Scan(&arCount)

	var ar = make([][]int32, arCount)

	for i := 0; i < int(arCount); i++ {
		ar[i] = make([]int32, arCount)

		for y := 0; y < int(arCount); y++ {
			fmt.Scan(&ar[i][y])
		}
	}

	result := diagonalDifference(ar, arCount)
	fmt.Println(result)
}
