package main

import "fmt"

func workingWithArrays()  {
	//Declaration
	var arr = [3]int{1,2,3}
	var arr2 = [...]string{"H","E","L","L","O"}
	var defInt = [2]int{}
	var defPartialInt = [3]int{1}
	var specificInitializationInt = [5]int{1:10,4:20}

	fmt.Println("arr Before: ",arr)
	arr[1]=2
	fmt.Println("arr After: ",arr)

	fmt.Println(arr2)
	fmt.Println(defInt)
	fmt.Println(defPartialInt)
	fmt.Println(specificInitializationInt)
	fmt.Println(len(arr))
}