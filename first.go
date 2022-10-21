package main

import "fmt"

func main() {

	a := 2
	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		// 在此打印"two"并跳出
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4, 5:
		fmt.Println("four or five")
	default:
		fmt.Println("other")
	}