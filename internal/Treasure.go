package internal

import (
	"bufio"
	"divar/pkg"
	"fmt"
	"os"
	"strings"
)

func Max(slice []int) int {
	var max = slice[0]
	for _, value := range slice {
		if max < value {
			max = value
		}
	}
	return max
}

func Treasure() {
	var m int
	fmt.Scan(&m)
	matrix := make([][]int, 0)
	matrix = Read(m, matrix)
	output := Rules(matrix)
	fmt.Println("_______________________________")
	fmt.Println(output)
}

func Read(m int, matrix [][]int) [][]int {
	reader := bufio.NewReader(os.Stdin)
	for u := 0; u < m; u++{
		line := make([]int, 0)
		text ,_ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		subtext := strings.Split(text, " ")
		for o := 0; o < len(subtext); o++{
			line = append(line, pkg.Parse(subtext[o]))
		}
		matrix = append(matrix, line)
	}
	return matrix
}

func Rules(mat [][]int) int {
	var sum int
	var sumList []int
	for x := 0; x < len(mat); x++{
		if x == 0 {
			for y := range mat {
				if y == 0 {
					sum = mat[x][y] + mat[x][y+1] + mat[x+1][y]
					sumList = append(sumList, sum)
				}else if y == len(mat) - 1{
					sum = mat[x][y] + mat[x][y-1] + mat[x+1][y]
					sumList = append(sumList, sum)
				}else{
					sum = mat[x][y] + mat[x][y-1] + mat[x][y+1] + mat[x+1][y]
					sumList = append(sumList, sum)
				}
			}
		}else if x == len(mat) - 1{
			for y := range mat {
				if y == 0 {
					sum = mat[x][y] + mat[x][y+1] + mat[x-1][y]
					sumList = append(sumList, sum)
				}else if y == len(mat) - 1{
					sum = mat[x][y] + mat[x][y-1] + mat[x-1][y]
					sumList = append(sumList, sum)
				}else{
					sum = mat[x][y] + mat[x][y-1] + mat[x][y+1] + mat[x-1][y]
					sumList = append(sumList, sum)
				}
			}
		}else{
			for y := range mat {
				if y == 0{
					sum = mat[x][y] + mat[x-1][y] + mat[x+1][y] + mat[x][y+1]
					sumList = append(sumList, sum)
				}else if y == len(mat) - 1{
					sum = mat[x][y] + mat[x-1][y] + mat[x+1][y] + mat[x][y-1]
					sumList = append(sumList, sum)
				}else{
					sum = mat[x][y] + mat[x+1][y] + mat[x-1][y] + mat[x][y-1] + mat[x][y+1]
					sumList = append(sumList, sum)
				}
			}
		}
	}
	return Max(sumList)
}