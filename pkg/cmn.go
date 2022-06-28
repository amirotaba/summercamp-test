package pkg

import "strconv"

func Unique(intSlice []int) []int {
	allKeys := make(map[int]bool)
	ulist := []int{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			ulist = append(ulist, item)
		}
	}
	return ulist
}

func Parse(str string) int {
	if i, err := strconv.Atoi(str); err == nil {
		return i
	} else {
		return -10000
	}
}