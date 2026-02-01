package fundamentals

import "fmt"

// SimpleCounter demonstrates the use of continue and break statements in a loop.
// It iterates from 1 to 9, skipping the iteration when i equals 5 (using continue),
// and terminates the loop when i equals 8 (using break). The loop prints each
// iteration number to standard output, except when skipped or after breaking.
func SimpleCounter() {
	for i := 1; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
		if i == 8 {
			break
		}
	}
}
