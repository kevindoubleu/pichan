package main

import (
	"fmt"

	"github.com/kevindoubleu/pichan/pkg/hello"
)

func main() {
	greeting := hello.BuildGreeting()
	fmt.Println(greeting)
}
