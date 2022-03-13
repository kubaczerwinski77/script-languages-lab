package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	params := os.Args[1:]
	// loop through each environmental variable
	for _, item := range os.Environ() {
		// split to pair {key, value}
		pair := strings.SplitN(item, "=", 2)
		// loop through each params and check if param is substring of key
		for _, param := range params {
			if strings.Contains(pair[0], strings.ToUpper(param)) {
				fmt.Println(item)
			}
		}
	}
}