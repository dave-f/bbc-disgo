package main

import (
    "fmt"
    "io/ioutil"
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

	data, err := ioutil.ReadAll(f)
	fmt.Println("Number of bytes read:", len(data))

    defer f.Close()
}
