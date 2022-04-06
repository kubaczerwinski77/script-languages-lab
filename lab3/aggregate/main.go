package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	// define what aggregate function user what to use
	functionPtr := flag.String("using", "", "a string")
	labelPtr := flag.String("label", "", "a string")
	sepPtr := flag.String("separator", ";", "a string")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	var result float64
	switch *functionPtr {
	case "min":
		min := math.MaxFloat64
		for scanner.Scan() {
			num, empty := parseNumber(scanner.Text())
			if empty {
				continue
			}

			if (num < min) {
				min = num
			}
		}
		result = min
	case "max":
		max := 2.2E-308 // <- min float64 number in Go
		for scanner.Scan() {
			num, empty := parseNumber(scanner.Text())
			if empty {
				continue
			}

			if (num > max) {
				max = num
			}
		}
		result = max
	case "sum":
		sum := 0.0
		for scanner.Scan() {
			num, empty := parseNumber(scanner.Text())
			if empty {
				continue
			}

			sum += num
		}
		result = sum
	case "avg":
		sum := 0.0
		count := 0
		for scanner.Scan() {
			num, empty := parseNumber(scanner.Text())
			if empty {
				continue
			}

			sum += num
			count += 1
		}
		result = sum / float64(count)
	case "count":
		count := 0
		for scanner.Scan() {
			_, empty := parseNumber(scanner.Text())
			if empty {
				continue
			}

			count += 1
		}
		result = float64(count)
	case "countAll":
		count := 0
		for scanner.Scan() {
			parseNumber(scanner.Text())
			count += 1
		}
		result = float64(count)
	case "median":
		slice := []float64{}
		for scanner.Scan() {
			num, empty := parseNumber(scanner.Text())
			if empty {
				continue
			}

			slice = append(slice, num)
		}

		sort.Float64s(slice)

		length := len(slice)
		// cant calculate median if stdin was empty
		if length == 0 {
			fmt.Fprintf(os.Stderr, "You must aggregate at least one number")
			os.Exit(1)
		}

		var res float64
		if (length % 2 == 0) {
			res = (slice[length / 2 - 1] + slice[length / 2]) / 2.0
		} else {
			res = slice[length / 2]
		}
		result = res
	default:
		fmt.Println("there is no such option for this flag")
	}
	
	fmt.Printf("%s%s%f\n", *labelPtr, *sepPtr, result)
}

func parseNumber(num string) (float64, bool) {
	if len(num) != 0 {
		n, err := strconv.ParseFloat(num, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Number \"%s\" cannot be converted into Float64.\n", num)
			os.Exit(1)
		}
		return n, false
	} else {
		return 0.0, true
	}
}