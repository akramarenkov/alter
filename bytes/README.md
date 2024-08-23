# Bytes

## Split

### Benchmark

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/alter/bytes
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkSplitStd
BenchmarkSplitStd-12              8448201           139.4 ns/op         128 B/op           1 allocs/op
BenchmarkSplit
BenchmarkSplit-12                 9370696           129.5 ns/op         128 B/op           1 allocs/op
BenchmarkSplitPreparer
BenchmarkSplitPreparer-12        32791620            37.25 ns/op          0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/alter/bytes    4.889s
```

### Usage

Example:

```go
package main

import (
    "fmt"

    "github.com/akramarenkov/alter/bytes"
)

func main() {
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
```
