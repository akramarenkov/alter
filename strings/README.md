# Strings

## Split

### Benchmark

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/alter/strings
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkSplitStd
BenchmarkSplitStd-12              9856176           140.9 ns/op          80 B/op           1 allocs/op
BenchmarkSplit
BenchmarkSplit-12                 8776906           138.9 ns/op          80 B/op           1 allocs/op
BenchmarkSplitPreparer
BenchmarkSplitPreparer-12        32606250            34.81 ns/op          0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/alter/strings    4.054s
```

### Usage

Example:

```go
package main

import (
    "fmt"

    "github.com/akramarenkov/alter/strings"
)

func main() {
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
```
