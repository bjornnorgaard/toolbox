package main

import (
	"fmt"

	"github.com/bjornnorgaard/toolbox/oda"
)

func main() {
	converted, err := oda.ReplaceAllCharsWith("./main.go", "ReplaceAllCharsWith", 'X')
	if err != nil {
		panic(err)
	}

	fmt.Println(converted)
}
