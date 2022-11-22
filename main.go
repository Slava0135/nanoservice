package main

import (
	"fmt"
	"math/rand"
	"slava0135/nanoservice/generate"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(generate.NewGameLayout())
}
