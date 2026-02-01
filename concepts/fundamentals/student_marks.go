package fundamentals

import "fmt"

func StudentMarks() {
	marks := make(map[string]int)
	marks["A"] = 100
	marks["B"] = 90
	marks["C"] = 50

	var toCheck string = "A"

	_, exist := marks[toCheck]

	if !exist {
		fmt.Println("User Does not exist")
	} else {
		fmt.Printf("Marks of %v is %v", toCheck, marks[toCheck])
	}
}
