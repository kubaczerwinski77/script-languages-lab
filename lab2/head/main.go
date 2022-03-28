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
	linesPtr := flag.Int("lines", 0, "an int")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if *linesPtr > 0 {
			fmt.Println(scanner.Text())
			*linesPtr -= 1
		}		
	}

	// if e flag is set, exit program without errors
	if *errorPtr {
		os.Exit(0)
	}

	// otherwise output message on stderr
	if *linesPtr > 0 {
		fmt.Fprintf(os.Stderr, "Missing lines: %d\n", *linesPtr)
		os.Exit(2)
	}

}