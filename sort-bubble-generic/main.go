package main

import (
	"fmt"
)

type Ordered interface {
	~float64 | ~int | ~string
}

func bubblesort[T Ordered](data []T) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func main() {
	numbers := []int{99, 55, 3, 4, 5, 111, 2, 1, 65, 88}
	numbers2 := []float64{3.5, -2.4, 12.8, 9.1}
	names := []string{"Zachary", "John", "Moe", "Jim", "Robert"}

	bubblesort[int](numbers)
	fmt.Println(numbers)

	bubblesort[float64](numbers2)
	fmt.Println(numbers2)

	bubblesort[string](names)
	fmt.Println(names)
}
