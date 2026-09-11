package calc

import (
	"math"
)

func Add(x, y float64) float64 {
	return x + y
}

func Sub(x, y float64) float64 {
	return x - y
}

func Mult(x, y float64) float64 {
	return x * y
}

func Dev(x, y float64) float64 {
	return x / y
}

func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}
