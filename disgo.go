package main

import (
    "fmt"
//    "io"
    "os"
)

func main() {

    fmt.Println("Hello, World!")

    if len(os.Args) != 2 {
		fmt.Println("Usage: disgo <filename>")
        return
    }

    f, err := os.Open(os.Args[1])

    if err != nil {
        fmt.Println(err)
        return
    }

    defer f.Close()
}
