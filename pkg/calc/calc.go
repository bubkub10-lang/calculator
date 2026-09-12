package calc

import (
	"math"
)

// def math
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

// trigo
func Sin(x float64) float64 {
	return math.Sin(x)
}
func Cos(x float64) float64 {
	return math.Cos(x)
}
func Tan(x float64) float64 {
	return math.Tan(x)
}
func Asin(x float64) float64 {
	return math.Asin(x)
}
func Acos(x float64) float64 {
	return math.Acos(x)
}
func Atan(x float64) float64 {
	return math.Atan(x)
}
func DegtoRad(deg float64) float64 {
	return deg * (math.Pi / 180)
}
func RadtoDeg(rad float64) float64 {
	return rad * (180 / math.Pi)
}

// logarithms and exp
func Exp(x float64) float64 {
	return math.Exp(x)
}
func Log(x float64) float64 {
	return math.Log(x)
}
func Log2(x float64) float64 {
	return math.Log2(x)
}
func Log10(x float64) float64 {
	return math.Log10(x)
}

// Roots
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}
func Cbrt(x float64) float64 {
	return math.Cbrt(x)
}
