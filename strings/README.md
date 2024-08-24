# Strings

## Split

### Benchmark

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/alter/strings
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkSplitStd-12              8205315           152.7 ns/op          80 B/op           1 allocs/op
BenchmarkSplit-12                 8196849           142.9 ns/op          80 B/op           1 allocs/op
BenchmarkSplitPreparer-12        34443756            34.30 ns/op          0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/alter/strings    4.849s
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
