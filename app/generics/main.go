package main

import (
	"fmt"
)

type Number interface {
	int64 | float64
}

func main() {
	 // int64のmap初期化
	 ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// float64のmap初期化
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
    SumIntsOrFloats[string, int64](ints),
    SumIntsOrFloats[string, float64](floats))

	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
    SumIntsOrFloats(ints),
    SumIntsOrFloats(floats))

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
    SumNumbers(ints),
    SumNumbers(floats))
}

// mapの値 (int64) の合計値を返却.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
			s += v
	}
	return s
}

// mapの値 (float64) の合計値を返却.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
			s += v
	}
	return s
}

// マップmの値 (int64またはfloat64)の合計値を返却
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
			s += v
	}
	return s
}

// マップmの値 (int64またはfloat64)の合計値を返却
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
			s += v
	}
	return s
}