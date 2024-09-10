package mymath

import "math"

func Abs(x float64) float64 {
    if x < 0 {
        return -x
    } else {
        return x
    }
}

func Max(a, b float64) float64 {
    if a > b {
        return a
    } else {
        return b
    }
}

func Sqrt(x float64) float64 {
    return math.Sqrt(x)
}

func Yn(m, n int) float64 {
    result := m * n
    return float64(result)

}
