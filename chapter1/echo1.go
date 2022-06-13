package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("name is " + os.Args[0])

	for index, arg := range os.Args[1:] {
		fmt.Printf("index = %d, arg = %v\n", index, arg)
	}
}
