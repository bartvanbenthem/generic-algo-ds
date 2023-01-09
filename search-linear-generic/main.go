package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 500_000_000

type DataType interface {
	~float64 | ~int | ~string
}

func linearSearch[T DataType](slice []T, target T) bool {
	// Return true if T is in the slice
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return true
		}
	}
	return false
}
func main() {
	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = 100.0 * rand.Float64()
	}

	start := time.Now()
	result := linearSearch[float64](data, 54.0)
	elapsed := time.Since(start)
	fmt.Println("Time to search slice of 100_000_000 floats using linearSearch = ", elapsed)
	fmt.Println("Result of search is ", result)

	start = time.Now()
	result = linearSearch[float64](data, data[size/2])
	elapsed = time.Since(start)
	fmt.Println("Time to search slice of 100_000_000 floats using linearSearch = ", elapsed)
	fmt.Println("Result of search is ", result)
}
