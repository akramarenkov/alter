package strings_test

import (
	"fmt"

	"github.com/akramarenkov/alter/strings"
)

func ExampleSplit() {
	buffer := make([]string, 10)

	preparer := func(length int) []string {
		if length > len(buffer) {
			return nil
		}

		return buffer[:length]
	}

	fmt.Println(strings.Split("1 2 3 4 5", " "))
	fmt.Println(strings.Split("1 2 3 4 5", " ", preparer))

	// Output:
	// [1 2 3 4 5]
	// [1 2 3 4 5]
}
