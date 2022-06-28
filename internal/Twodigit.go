package internal

import (
	"divar/pkg"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func parse(str string) int {
	if i, err := strconv.Atoi(str); err == nil {
		return i
	} else {
		return -1
	}
}

func Twodigit() {
	var m string
	fmt.Scan(&m)
	numbers := recnumber(m)
	fmt.Println("_______________________________")
	rule(numbers)
}

func recnumber(str string) []int {
	var a, b int
	list := strings.Split(str, "")
	//fmt.Println("list: ",list)
	numbers := make([]int, 0)
	temp := make([]string, 0)
	for i := 0; i < len(list)-1; i++ {
		a, b = 0, 0
		a = parse(list[i])
		b = parse(list[i+1])
		if a > 0 {
			if b >= 0 {
				temp = nil
				temp = append(temp, list[i], list[i+1])
				two := parse(strings.Join(temp, ""))
				numbers = append(numbers, two)
				//fmt.Println(a, b, temp)
			}
		}
	}
	//fmt.Println("numbers: ",numbers)
	return numbers
}

func rule(all []int) {
	key := make([]int, 0)
	//fmt.Println("dict all: ", dict_all)
	for y := range all {
		sqrt := int(math.Sqrt(float64(all[y])))
		for t := 2; t <= sqrt; t++ {
			if all[y]%t == 0 {
				key = append(key, all[y])
			}
		}
	}
	key = pkg.Unique(key)
	//fmt.Println("keys:", key, "alls: ", all)
	for r := range all {
		for e := range key {
			if all[r] == key[e] {
				all[r] = 0
			}
		}
	}
	for h := range all {
		if all[h] != 0 {
			fmt.Println(all[h])
		}
	}
	return
}