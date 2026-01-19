package main

import "fmt"

func increment_count(x int) int {
	if x == 10 {
		return 0
	}
	fmt.Println(x)
	return increment_count(x + 1)
}