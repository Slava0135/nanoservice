package main

import (
	"fmt"
	"slava0135/nanoservice/generate"
)

func main() {
    fmt.Println("Hello, World!")
    l := generate.NewLayout()
    for _, v := range l {
        for _, v := range v {
            if v {
                fmt.Print("#")
            } else {
                fmt.Print("_")
            }
        } 
        fmt.Println()
    }
}