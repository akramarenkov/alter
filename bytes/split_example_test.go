package bytes_test

import (
	"fmt"

	"github.com/akramarenkov/alter/bytes"
)

func ExampleSplit() {
	buffer := make([][]byte, 10)

	preparer := func(length int) [][]byte {
		if length > len(buffer) {
			return nil
		}

		return buffer[:length]
	}

	fmt.Println(bytes.Split([]byte("1 2 3 4 5"), []byte(" ")))
	fmt.Println(bytes.Split([]byte("1 2 3 4 5"), []byte(" "), preparer))

	// Output:
	// [[49] [50] [51] [52] [53]]
	// [[49] [50] [51] [52] [53]]
}
