package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func nakedReturn_Add(x int, y int) (result int) {
	result = add(x, y)
	return
}

func multiReturn_add_sub(x int, y int) (addRes int, subRes int) {
	addRes = add(x, y)
	subRes = x - y
	return
}

func using_multiReturn_res() {
	var a, b = multiReturn_add_sub(10, 20)
	var _, d = multiReturn_add_sub(20, 30)

	fmt.Println(a,b,d)
}