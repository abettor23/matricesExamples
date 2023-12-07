package main

import (
	"fmt"
)

// для удобного представления кода присвоил переменные на строки матрицы
func determinant(matrix [3][3]float64) float64 {
	a, b, c := matrix[0][0], matrix[0][1], matrix[0][2]
	d, e, f := matrix[1][0], matrix[1][1], matrix[1][2]
	g, h, i := matrix[2][0], matrix[2][1], matrix[2][2]

	return a*(e*i-f*h) - b*(d*i-f*g) + c*(d*h-e*g)
}

func main() {

	matrix := [3][3]float64{
		{3, 9, 1},
		{5, 1, 7},
		{6, 5, 9}, //числа матрицы из головы
	}

	fmt.Println("Определитель равен:", determinant(matrix))

}
