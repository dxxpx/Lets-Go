package main

import "fmt"

func operators()  {

	//Arithemetic
	var a int = 10;
	var b int = a + 20;

	fmt.Println("a",a,"b",b)

	var addtn int = b+20
	var subtr int = b-20
	var mult int = b*20
	var div int  = b/10
	var mod int = b%10
	a++
	fmt.Printf("Addtn: %d Subtr: %d Mult: %d Div: %d Modulus: %d increment: %d\n", addtn,subtr,mult,div,mod,a)
	a--
	fmt.Printf("Decrement: %d \n",a)

	//Assignment
	var x=10
	x+=5
	fmt.Println("X: ",x)

	//Comparision
	fmt.Println(a>b)
	fmt.Println(a<b)
	fmt.Println(a<=b)
	fmt.Println(a>=b)
	fmt.Println(a==b)
	fmt.Println(a!=b)

	//Logical
	fmt.Println(a>5 && b>10)
	fmt.Println(a>5 || b>10)
	fmt.Println(!(a>5 || b>10))

}