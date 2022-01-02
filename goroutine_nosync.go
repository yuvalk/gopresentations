package main

import (
	"fmt"
	"sync"
)

//START OMIT
func main() {
    v := 0
    var wg sync.WaitGroup
    wg.Add(2)
    go func() {
        v = 1
        wg.Done()
    }()
    go func() {
        fmt.Println(v)
        wg.Done()
    }()
    wg.Wait()
}
//END OMIT
