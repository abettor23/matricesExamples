package main

import (
	"fmt"
)

func multiplyMatrices(matrix1 [3][5]float64, matrix2 [5][4]float64) [3][4]float64 {
	var result [3][4]float64
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 5; k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}
	return result
}

func main() {

	matrix1 := [3][5]float64{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
	}
	matrix2 := [5][4]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
		{17, 18, 19, 20},
	}

	result := multiplyMatrices(matrix1, matrix2)

	fmt.Println("Результат умножения матриц:")
	for _, row := range result {
		fmt.Println(row)
	}
}
