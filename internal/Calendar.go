package internal

import (
	"bufio"
	"divar/pkg"
	"fmt"
	"os"
	"strings"
)

func Calendar() {
	var id, nom, n, duration, busyDay, nothing int
	var bad string
	busy := make([]int, 0)
	nomlist := make([]int, 0)
	times := make([][][]int, 0)
	fmt.Scan(&n)
	for i := 0; i < n; i++{
		fmt.Scan(&id)
		fmt.Scan(&bad)
		busyDay = Convert(bad)
		busy = append(busy, busyDay)
		fmt.Scan(&nom)
		nomlist = append(nomlist, nom)
		times = append(times, Matrix(nom))
	}

	fmt.Scan(&nothing, &duration)
	key := Order(times, nomlist, busy)
	fmt.Println("_______________________________")
	fmt.Println(FreeTimes(key, busy, duration))
}

func Convert(day string) int {
	output := 0
	var week = []string{"MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"}
	for i := range week {
		if week[i] == day {
			output = i
		}
	}
	return output
}

func Reconvert(day int) string {
	var output string
	var week = []string{"MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"}
	output = week[day]
	return output
}

func Matrix(nom int) [][]int {
	reader := bufio.NewReader(os.Stdin)
	matrix := make([][]int, 0)
	for i := 0; i < nom; i++{
		list := make([]int, 0)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		subtext := strings.Split(text, " ")
		list = append(list, Convert(subtext[1]))
		for y := 2; y < len(subtext); y++{
			list = append(list, pkg.Parse(subtext[y]))
		}
		matrix = append(matrix, list)
	}
	return matrix
}

func Order(list [][][]int, nom []int, busy []int) [][]int {
	times := make([][]int, 0)
	for i := range list {
		for u := 0; u < nom[i]; u++{
			times = append(times, list[i][u])
		}
	}
	pass := make([]int, 0)
	key := make([][]int, 0)
	for y := range times{
		pass = nil
		for t := range busy {
			if times[y][0] != busy[t] {
				pass = append(pass, 0)
				if len(pass) == len(busy) {
					key = append(key, times[y])
				}
			}
		}
	}
	org := make([][]int, 0)
	for t := 0; t < 7; t++{
		for r := range key {
			if key[r][0] == t {
				org = append(org, key[r])
			}
		}
	}
	return org
}


func FreeTimes(key [][]int, busy []int, dur int) (string, int, int) {
	m := make(map[int]int)
	for i := 0; i < 7; i++ {
		for u := range key {
			if key[u][0] == i {
				m[i] += 1
			}
		}
	}
	list := make([][]int, 0)
	start := 0
	end := 0
	output := make([][]int, 0)
	for y := range m {
		list = nil
		end += m[y]
		start = end - m[y]
		for t := start; t < end; t++{
			list = append(list, key[t])
		}
		if list[0][1] >= dur {
			temp := make([]int, 0)
			temp = append(temp, list[0][0], 0, dur)
			output = append(output, temp)
			break
		}
		if 32400000 - list[len(list)-1][2] >= dur {
			output = append(output, list[len(list)-1])
		}
		for r := range list {
			if r < len(list)-1{
				if list[r+1][1] - list[r][2] >= dur {
					output = append(output, list[r])
				}
			}
		}
	}
	freeday := make([]int, 0)
	for t := 0; t < 7; t++{
		found := false
		for r := range busy {
			if t == r {
				found = true
			}
		}
		if found == false {
			freeday = append(freeday, t)
			break
		}
	}
	var day string
	answer := make([]int, 0)
	if freeday[0] < output[0][0] {
		day = Reconvert(freeday[0])
		answer = append(answer, 0, dur)
	}else{
		day = Reconvert(output[0][0])
		for w := 1; w < 3; w++ {
			answer = append(answer, output[0][w])
		}
	}
	return day, answer[0], answer[1]
}