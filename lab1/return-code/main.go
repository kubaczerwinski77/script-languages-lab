package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// read from stdin
	scanner := bufio.NewScanner(os.Stdin)
	words := []string{}
	for scanner.Scan() {
		// split the line of input text
		line := strings.Split(scanner.Text(), " ")
		for _, word := range line {
			lower := strings.ToLower(word)
			words = append(words, strings.Trim(lower, "!?:;.,"))
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	// handle program arguments
	args := []string{}
	for _, arg := range os.Args[1:] {
		args = append(args, strings.ToLower(arg))
	}

	// make a dictionary to sum up words repetitions
	dict := make(map[string]int)
	for _, word := range words {
		for _, arg := range args {
			if word == arg {
				dict[word] += 1
			}
		}
	}

	// find one word with most repetitions
	if len(dict) == 0 {
		fmt.Println(0)
	}

	var maxKey string
	maxValue := 0
	for key, value := range dict {
		if value > maxValue {
			maxValue = value
			maxKey = key
		}
	}

	// return proper value depends on args
	for index, arg := range args {
		if arg == maxKey {
			os.Exit(index + 1)
		}
	}
}