package main

import "fmt"

func loops()  {
	for i := 0; i < 5; i++ {
		if i==3 {
			continue
		}
		if i==4{
			break
		}
		fmt.Println(i)
	}

	var arr = [5]int{1,2,3,4,5}
	for idx, val := range arr{
		fmt.Println(idx, " - ",val)
	}

	for _,val := range arr{
		fmt.Println(val)
	}
	for idx,_ := range arr{
		fmt.Println(idx)
	}

}