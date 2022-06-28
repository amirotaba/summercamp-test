package internal

import (
	"bufio"
	"divar/pkg"
	"fmt"
	"os"
	"strings"
)

func Peace() {
	var n int
	times := make([][]int, 0)
	fmt.Scan(&n)
	Free(Meeting(n, times), FindDur())
}

func Meeting(n int, times [][]int) [][]int {
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++{
		time := make([]int, 0)
		text ,_ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		subtext := strings.Split(text, " ")
		for o := 0; o < len(subtext); o++{
			time = append(time, pkg.Parse(subtext[o]))
		}
		times = append(times, time)
	}
	return times
}

func FindDur() int{
	reader := bufio.NewReader(os.Stdin)
	text ,_ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	subtext := strings.Split(text, " ")
	duration := pkg.Parse(subtext[1])
	return duration
}

func Free(times [][]int, duration int) {
	output := make([]int, 0)
	if times[0][1] - duration >= 0 {
		output = append(output, 0, duration)
	}else{
		for i := range times {
			if i < len(times)-1 {
				if times[i+1][1] - times[i][2] >= duration {
					start := times[i][2]
					output = append(output, start, start+duration)
				}
			}else {
				start := times[len(times)-1][2]
				output = append(output, start, start+duration)
			}
		}
	}
	fmt.Println("_______________________________")
	fmt.Println(output[0], output[1])
	return
}