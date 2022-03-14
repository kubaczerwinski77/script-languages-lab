package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("SUPER;COMPLICATED;NAME", "my/environmental/variable")
	args := os.Args[1:]
	result := []string{}
	for _, arg := range args {
		exists := false
		for _, env := range os.Environ() {
			pair := strings.SplitN(env, "=", 2)
			if strings.Contains(pair[0], arg) {
				result = append(result, env)
				exists = true
			}
		}

		if !exists {
			result = append(result, arg + "=NONE")
		}
	}

	for _, v := range result {
		displayVariable(v)
	}
}

func displayVariable(variable string) {
	if strings.Contains(variable, ";") {
		slice := strings.Split(variable, ";")
		complexName := ""
		for _, v := range slice {
			complexName += "\t" + v + "\n"
		}
		fmt.Println(complexName)
	} else {
		fmt.Println(variable)
	}
}