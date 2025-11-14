package main

import "fmt"

func plusMinus(ar []int32, arCount uint32) (float32, float32, float32) {
	var countPositive uint32 = 0
	var countNegative uint32 = 0
	var countZero uint32 = 0
	
	for _, value := range ar {
		if value > 0 {
			countPositive++
		} else if value < 0 {
			countNegative++
		} else {
			countZero++
		}
	}

	return float32(countPositive) / float32(arCount), float32(countNegative) / float32(arCount), float32(countZero) / float32(arCount)
}

func main() {
	var arCount uint32
	fmt.Scan(&arCount)

	var ar = make([]int32, arCount)

	for i := 0; i < int(arCount); i++ {
		fmt.Scan(&ar[i])
	}

	resultPositive, resultNegative, resultZero := plusMinus(ar,arCount)
	fmt.Printf("%.6f\n%.6f\n%.6f", resultPositive, resultNegative, resultZero)
}