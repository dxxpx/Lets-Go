package fundamentals

import "fmt"

func ShoppingCart() {
	items := make([]string, 0)

	items = append(items, "book", "Grocery")
	items = append(items, "Milkshake", "Vinyl")

	fmt.Println("\nLength of slice : ", len(items))
	fmt.Println("Capacity of slice : ", cap(items))

	fmt.Println("Before")
	for index, value := range items {
		fmt.Printf("Item%d: %s || ", index, value)
	}
	fmt.Println()

	deleteIndex := 2
	items = append(items[:deleteIndex], items[deleteIndex+1:]...)

	fmt.Println("After")
	for index, value := range items {
		fmt.Printf("Item%d: %s || ", index, value)
	}
}
