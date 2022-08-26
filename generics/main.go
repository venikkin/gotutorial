package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non generic sums %v and %v\n",
		SumInts(ints),
		SumFloats(floats))
	fmt.Printf("Generic sums %v and %v\n",
		SumIntsOrFloats(ints), SumIntsOrFloats(floats))
}

func SumInts(m map[string]int64) int64 {
	var sum int64
	for _, n := range m {
		sum += n
	}
	return sum
}

func SumFloats(m map[string]float64) float64 {
	var sum float64
	for _, n := range m {
		sum += n
	}
	return sum
}

func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}
