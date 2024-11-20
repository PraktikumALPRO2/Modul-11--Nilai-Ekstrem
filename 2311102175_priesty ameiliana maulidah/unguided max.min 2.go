package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	weights := make([]float64, x)
	for i := 0; i < x; i++ {
		fmt.Scan(&weights[i])
	}

	totalWeights := make([]float64, (x+y-1)/y)
	for i := 0; i < x; i++ {
		totalWeights[i/y] += weights[i]
	}

	for _, total := range totalWeights {
		fmt.Printf("%.2f ", total)
	}
	fmt.Println()

	averageWeights := make([]float64, len(totalWeights))
	for i, total := range totalWeights {
		averageWeights[i] = total / float64(y)
		if i == len(totalWeights)-1 && x%y != 0 {
			averageWeights[i] = total / float64(x%y)
		}
	}

	for _, avg := range averageWeights {
		fmt.Printf("%.2f ", avg)
	}
	fmt.Println()
}
