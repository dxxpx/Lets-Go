package main

import "fmt"

func conditions() {
	x := 20
	y := 18
	if x > y {
		fmt.Println("X is greater")
	}else if x < y{
		fmt.Println("Y is greater")
	} else {
		fmt.Println("Both are equal")
	}
}