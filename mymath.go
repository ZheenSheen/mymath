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

func Sqrt(x float64) int {
    return int(math.Sqrt(x))
}

func Yn(m, n int) int {
    result := 1
    for i := m; i <= n; i++ {
        result *= i
    }
    return result
}
