package main

const N int = 10000
type Number interface {
	int64 | float64
}

func sumNonGenerics(a float64, b float64) float64 {
	var result float64 = 0
	for i := 0; i < N; i++ {
		result += a + b
	}
	return result
}

func sumGenerics[T Number](a T, b T) T {
	var result T = 0
	for i := 0; i < N; i++ {
		result += a + b
	}
	return result
}