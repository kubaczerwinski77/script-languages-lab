package main

import (
	"fmt"
	"os"
)

func main() {
	envString := ""
	for _, value := range os.Environ() {
		envString += value + " "
	}
	fmt.Println(envString)
	fmt.Println(os.Args[1:])
}