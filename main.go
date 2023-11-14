package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	switch {
	case a/10000 >= 1:
		fmt.Println(a / 10000)
	case a/1000 >= 1:
		fmt.Println(a / 1000)
	case a/100 >= 1:
		fmt.Println(a / 100)
	case a/10 >= 1:
		fmt.Println(a / 10)
	default:
		fmt.Println(a)
	}
}
