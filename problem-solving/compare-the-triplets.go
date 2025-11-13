package main

import "fmt"

func compareTriplets(a []uint32, b []uint32) (uint32, uint32) {
	var scoreA uint32 = 0
	var scoreB uint32 = 0
	
	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			scoreA++
		} else if a[i] < b[i] {
			scoreB++
		}
	}

	return scoreA, scoreB
}

func main() {
    const size = 3
    var a = make([]uint32, size)
    var b = make([]uint32, size)
    
    for i := 0; i < size; i++ {
        fmt.Scan(&a[i]) 
    }

    for i := 0; i < size; i++ {
        fmt.Scan(&b[i]) 
    }
    
	scoreA, scoreB := compareTriplets(a, b);
	fmt.Println(scoreA, scoreB)
}