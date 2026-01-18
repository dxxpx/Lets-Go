package main

import "fmt"

func slices()  {
	var arr = [6]int{0,1,2,3,4,5}
	
	//Creating slices 
	//1)
	var myslice = []int{2,3,4,5,6,7}

	//2)
	var arrSlice = arr[2:4]

	//3)
	var makeSlice = make([]int, 5,10) //(type,len,cap)

	fmt.Println("Length: ", len(myslice))
	fmt.Println("Capacity: ", cap(myslice))
	fmt.Println(makeSlice)

	fmt.Println(myslice[3])
	myslice[3] = 9
	fmt.Println(myslice[3])
	
	myslice = append(myslice, 10,11,12)
	fmt.Println("After appending by items",myslice)

	myslice = append(myslice, arrSlice...)
	fmt.Println("After appending another slice", myslice)
}