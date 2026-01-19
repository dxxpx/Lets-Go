package main

import "fmt"

type Animal struct{
	name string
	habitat string
	legs int
	tail bool
}

func structures()  {
	var dog Animal
	dog.habitat = "Cities"
	dog.legs = 4
	dog.name = "leo"
	dog.tail = true

	var spider Animal

	fmt.Println(dog)
	init_Spider(spider)
}

func init_Spider(animal Animal)  {
	animal.habitat = "Musty areas"
	animal.legs = 8
	animal.name = "Brook"
	animal.tail = false

	fmt.Println(animal)
}