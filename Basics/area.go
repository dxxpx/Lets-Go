package main

import "fmt"

func findArea()  {
	//typed Constant
	const PI float32 = 3.14

	//Untyped Constant
	const unTypedPI =3.14

	// const(
	// 	PI = 3.14
	// 	unTypedPI float32 = 3.14
	// )

	var rad float32 = 3
	var area = PI * (rad *rad)
	fmt.Println(area)
}