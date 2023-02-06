package main

import (
	"fmt"

	"github.com/kevindoubleu/pichan/pkg/hello"
)

func main() {
	greeting := hello.BuildGreeting("pichan")
	fmt.Println(greeting)
}
