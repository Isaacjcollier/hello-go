package main

import (
	"fmt"
	"strings"

	"github.com/isaacjcollier/add"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(add.Add(1, 2))

	fmt.Println(strings.Join([]string{"hi", "duhde"}, " "))
}
