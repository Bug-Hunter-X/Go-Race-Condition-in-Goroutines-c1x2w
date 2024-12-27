```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var m sync.Mutex
        var count int

        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        m.Lock()
                        count++
                        m.Unlock()
                }(i) // Pass i to the goroutine
        }

        wg.Wait()
        fmt.Println("Count:", count)
}
```