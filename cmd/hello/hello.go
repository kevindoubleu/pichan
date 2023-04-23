package main

import (
	"fmt"

	"github.com/kevindoubleu/pichan/internal/hello"
)

func main() {
	greeting := hello.BuildGreeting("pichan")
	fmt.Println(greeting)
}
