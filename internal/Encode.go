package internal

import "fmt"

//
//import (
//	"divar/pkg"
//	"fmt"
//	"strconv"
//	"strings"
//)
//
func Encode()  {
	fmt.Println("Sorry this program is under maintenance. :(")
	//var m int
	//fmt.Scan(&m)
	//list := splitToGigits(m)
	//ulist := pkg.Unique(list)
	//dict := countlist(list)
	//fmt.Println(Output(Code(ulist, dict)))
	//fmt.Println(Output(key))
}
//
//func reverseInt(l []int) {
//	for i, j := 0, len(l)-1; i < j; i, j = i+1, j-1 {
//		l[i], l[j] = l[j], l[i]
//		//fmt.Println("for i:", i,"for j:", j)
//	}
//}
//
//func splitToGigits(n int) []int {
//	var list []int
//
//	for n != 0 {
//		list = append(list, n % 10)
//		n /= 10
//	}
//
//	reverseInt(list)
//
//	fmt.Println("list:", list)
//	return list
//}
//
//func countlist( arr []int) map[int]int {
//	dict := make(map[int]int)
//	for _, num := range arr {
//		dict[num] = dict[num] + 1
//	}
//	fmt.Println("dict:", dict)
//	return dict
//}
//
//func Code(ulist []int, dict map[int]int) []int {
//	key := make([]int, 0)
//	fmt.Println("before for:", key)
//	for u := 0; u < len(ulist); u++ {
//		key = append(key, dict[ulist[u]], ulist[u])
//		fmt.Println(ulist[u], u , key)
//	}
//	fmt.Println("key:", key)
//	return key
//}
//
//func Output(key []int) string {
//	var keytext []string
//	for y := range key {
//		number := key[y]
//		text := strconv.Itoa(number)
//		keytext = append(keytext, text)
//
//	}
//	integer := strings.Join(keytext, "")
//	return integer
//}