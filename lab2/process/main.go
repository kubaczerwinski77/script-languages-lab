package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	// flags
	ignFirstPtr := flag.Int("ignoreFirst", 0, "an int")
	ignLastPtr := flag.Int("ignoreLast", 0, "an int")
	delimiterPtr := flag.String("delimiter", "", "a string")
	separatorPtr := flag.String("separator", "", "a string")	
	projectPtr := flag.String("project", "", "a string")
	selectPtr := flag.String("select", "", "a string")
	flag.Parse()

	// variables
	separator := "\t"
	delimiter := ","
	isEmpty := true
	
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// change flag if the file is not empty
		isEmpty = false

		// read file from stdin libe by line
		line := scanner.Text()

		// ignore first N characters in each line
		if *ignFirstPtr > 0 {
			if *ignFirstPtr >= len(line) {
				fmt.Fprintf(os.Stderr, "You cannot skip same or more characters than there are in the line! Line characters: %d. Given: %d\n", len(line), *ignFirstPtr)
				os.Exit(1)
			}
			line = line[*ignFirstPtr:]
		}

		// ignore last N characters in each line
		if *ignLastPtr > 0 {
			if *ignLastPtr >= len(line) {
				fmt.Fprintf(os.Stderr, "You cannot skip same or more characters than there are in the line! Line characters: %d. Given: %d\n", len(line), *ignLastPtr)
				os.Exit(1)
			}
			line = line[:len(line) - *ignLastPtr]
		}

		// replace delimiter default string ("\t") with given in separator flag
		if *separatorPtr != "" {
			separator = *separatorPtr
		}
		
		// replaceAll all delimiter values with given separator ()
		if *delimiterPtr != "" {
			delimiter = *delimiterPtr
		}

		// split lines into columns
		if *projectPtr != "" {
			line = strings.ReplaceAll(line, delimiter, separator)
			row := strings.Split(line, separator)

			// checks if line is selected to print
			hasSelected := false
			if strings.Contains(line, *selectPtr) {
				hasSelected = true
			}

			// split columns passed in project flag arguments
			columnsString := strings.Split(*projectPtr, ",")

			for i, v := range columnsString {

				// convert flags type (string) to int
				columnInt, _ := strconv.Atoi(v)
				if columnInt < len(row) {	

					// checks if line is selected to print
					if hasSelected {
						if i == len(columnsString) - 1 {
							fmt.Print(row[columnInt] + "\n")
						} else {
							fmt.Print(row[columnInt] + separator)
						}
					}
				} else {
					os.Exit(1)
				}
			}

		// otherwise just print lines lines formatted before
		} else {
			line = strings.ReplaceAll(line, delimiter, separator)
			if strings.Contains(line, *selectPtr) {
				fmt.Println(line)
			}
		}
	}

	// sets exit code due to empty file
	if isEmpty {
		os.Exit(2)
	}
}