package main

import (
	"fmt"
	"math"
	"math/rand"
)

var x, x1, x2, c, a float64

func f(x float64) float64 {
	return 4*x + 1
}

func Round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

func testingStraightLine(x1 float64, x2 float64, c float64, a float64) {

	testResult := false
	for i := 0; i < 100; i++ {
		x := x1 + rand.Float64()*(x2-x1)
		result := Round(5, (f(x)-f(c))/(x-c))
		if result != a {
			testResult = true
			fmt.Printf("The requirement is not met for point X = %f, f(x) = %f: %.1f\n", x, f(x), result)
		}
	}
	if testResult == true {
		fmt.Printf("The requirement is not met for at least one point.\n")
	} else {
		fmt.Printf("The result was met for all points!\n")
	}
}

func main() {
	testingStraightLine(0, 10, 2, 4)
}
