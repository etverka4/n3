package main
import (
    "fmt"
    "sync"
)
func main() {
    var swg sync.WaitGroup
    nezn := make(chan int)
    swg.Add(2)
    go_1 := func(){
        defer swg.Done()
        text := <- nezn
        fmt.Print(text)
        fmt.Print("\n")
    }
    go_2 := func(){
        defer swg.Done()
        nezn <- 1
        fmt.Print("done")
    }
    go go_1()
    go go_2()
    swg.Wait()
    close(nezn)
}
# n3
