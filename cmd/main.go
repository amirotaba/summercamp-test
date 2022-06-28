package main

import (
	"divar/internal"
	"fmt"
)

func main() {
	var q int
	fmt.Println("Enter question number: ")
	fmt.Println("1 or 2(software engineering")
	fmt.Scan(&q)
	Load(q)
}

func Load(n int) {
	var t, y int
	switch n {
	case 1:
		fmt.Println("enter test number: ")
		fmt.Println("1, 2, 3, 4")
		fmt.Scan(&t)
		fmt.Println("_______________________________")
		switch t {
		case 1:
			internal.Diameter()
		case 2:
			internal.Encode()
		case 3:
			internal.Twodigit()
		case 4:
			internal.Treasure()
		}
	case 2:
		fmt.Println("enter test number: ")
		fmt.Println("1, 2, 3")
		fmt.Scan(&y)
		fmt.Println("_______________________________")
		switch y {
		case 1:
			internal.Peace()
		case 2:
			internal.Limbo()
		case 3:
			internal.Calendar()
		}
		}
}