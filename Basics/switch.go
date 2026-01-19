package main

import "fmt"

func switchCase()  {
	x:=1

	switch x {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	
	case 3:
		fmt.Println("Wednesday")
	
	case 4:
		fmt.Println("Thursday")
	
	case 5:
		fmt.Println("Friday")
	
	case 6:
		fmt.Println("Saturday")
	
	case 7:
		fmt.Println("Sunday")
	
	default:
		fmt.Println("Not valid")
	}

	switch x {
	case 1,3,5,7:
		fmt.Println("Odd day")
	case 2,4,6:
		fmt.Println("Even day")
	default:
		fmt.Println("Not a valid Day")
	}
}