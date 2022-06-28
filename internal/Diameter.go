package internal

import (
	"fmt"
)

func Diameter() {
	var m int
	fmt.Scan(&m)
	numbers := makeRange(m*m)
	matrix := MakeMatrix(m, numbers)
	diameters := Find(m, matrix)
	output := Rule1(diameters, numbers)
	fmt.Println("_______________________________")
	fmt.Println(output)
}

func makeRange(max int) []int {
	a := make([]int, max)
	for i := range a {
		a[i] = 1 + i
	}
	return a
}

func MakeMatrix(m int, numbers []int) [][]int {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = numbers[m*i : (i+1)*m]
	}
	return matrix
}

func Find(m int, matrix [][]int) []int {
	list := make([]int, m*m)
	for u := range matrix {
		list = append(list, matrix[u][u])
		t := -(u - (m - 1))
		if t != u {
			list = append(list, matrix[u][t])
		}
	}
	return list
}

func Rule1(list []int, numbers []int) int {
	key := make([]int, 0)
	for r := range list {
		if list[r] == 1 {
			key = append(key, list[r])
		} else{
			for e := range numbers {
				if list[r] == 3*numbers[e]+1 {
					key = append(key, list[r])
				}
			}
		}
	}
	sum := 0
	for w := range key {
		sum += key[w]
	}
	return sum
}