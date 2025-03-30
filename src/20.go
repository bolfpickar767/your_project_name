package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Generate a random number between 1 and 100
    var num int = rand.Intn(100) + 1

    fmt.Printf("The random number is: %d\n", num)
    time.Sleep(time.Second * 2)
}
