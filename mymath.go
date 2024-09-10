package mymath

import "math"

func Sqrt(x float64) float64 {
    return math.Sqrt(x)
}

func Add(a, b float64) float64 {
    return a + b
}

func Subtract(a, b float64) float64 {
    return a - b
}

func Multiply(a, b float64) float64 {
    return a * b
}

func Divide(a, b float64) float64 {
    if b != 0 {
        return a / b
    }
    return math.NaN()
}
