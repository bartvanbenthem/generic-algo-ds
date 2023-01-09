package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 50_000_000

type DataType interface {
	~float64 | ~int | ~string
}

func IsSorted[T DataType](data []T) bool {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}

func quicksort[T DataType](data []T, low, high int) {
	if low < high {
		var pivot = partition(data, low, high)
		quicksort(data, low, pivot)
		quicksort(data, pivot+1, high)
	}
}

func partition[T DataType](data []T, low, high int) int {
	// Pick a lowest bound element as a pivot value
	var pivot = data[low]
	var i = low
	var j = high
	for i < j {
		for data[i] <= pivot && i < high {
			i++
		}
		for data[j] > pivot && j > low {
			j--
		}

		if i < j {
			data[i], data[j] = data[j], data[i]
		}
	}
	data[low] = data[j]
	data[j] = pivot
	return j
}

func main() {

	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = 100.0 * rand.Float64()
	}

	start := time.Now()
	quicksort[float64](data, 0, len(data)-1)
	elapsed := time.Since(start)
	fmt.Println("Elapsed time for concurrent quicksort = ", elapsed)
	fmt.Println("Is sorted: ", IsSorted(data))

}
