package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	// define error flag
	errorPtr := flag.Bool("e", false, "a bool")
	linesPtr := flag.Int("lines", 1, "an int")
	flag.Parse()

	// easy queue implementation on slice
	queue := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if len(queue) >= *linesPtr {
			queue = queue[1:]
		}
		queue = append(queue, scanner.Text())
	}

	for _, v := range queue {
		fmt.Println(v)
	}

	// if e flag is set, exit program without errors
	if *errorPtr {
		os.Exit(0)
	}

	// otherwise output message on stderr
	if len(queue) < *linesPtr {
		fmt.Fprintf(os.Stderr, "Missing lines: %d\n", *linesPtr - len(queue))
		os.Exit(2)
	}

}