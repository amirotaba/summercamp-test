package internal

import (
	"bufio"
	"divar/pkg"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Limbo() {
	var n, m int
	times := make(map[int]int, n+m)
	fmt.Scan(&n)
	times = Counter(n, times)
	fmt.Scan(&m)
	times = Counter(m, times)
	commons := common(times)
	sum := Sum(commons)
	fmt.Println("_______________________________")
	fmt.Println(sum)
}

func Counter(x int, times map[int]int) map[int]int {
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < x; i++{
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		stext := strings.Split(text, " ")
		times[pkg.Parse(stext[1])] = pkg.Parse(stext[2])
	}
	return times
}

func common(times map[int]int) []int {
	var common int
	starts := make([]int, 0)
	commons := make([]int, 0)
	for i := range times {
		starts = append(starts, i)
	}
	sort.Ints(starts)
	for u := 0; u < len(starts)-1; u++{
		if times[starts[u]] > starts[u+1] {
			if times[starts[u+1]] < times[starts[u]] {
				common = times[starts[u+1]] - starts[u+1]
				commons = append(commons, common)
			}else {
				common = times[starts[u]] - starts[u+1]
				commons = append(commons, common)
			}
		}
	}
	return commons
}

func Sum(list []int) int {
	sum := 0
	for u := range list {
		sum += list[u]
	}
	return sum
}