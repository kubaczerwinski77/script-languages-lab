package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"time"
)

func main() {
	// flags
	recursivePtr := flag.Bool("R", false, "a bool")
	onlyDirPtr := flag.Bool("d", false, "a bool")
	sizePtr := flag.Bool("s", false, "a bool")
	sortPtr := flag.String("sort", "alpha", "a string")
	flag.Parse()

	// result slice with File objects
	result := []File{}

	// depends on R flag, choose the way of traversing
	if (*recursivePtr) {
		err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			if err != nil {
					return err
			}

			if path != "." {
				result = append(result, File{path, info.Size(), info.IsDir(), info.ModTime()})
			}

			return nil
		})
		if err != nil {
			fmt.Println(err)
		}
	} else {
		items, err := ioutil.ReadDir(".")
		if err != nil {
			fmt.Println(err)
		}
		for _, item := range items {
			file := File{item.Name(), item.Size(), item.IsDir(), item.ModTime()}
			result = append(result, file)
		}
	}

	display(*onlyDirPtr, *sizePtr, *sortPtr, result)
}

func display(onlyDir, size bool, dateSort string, result []File) {
	if dateSort == "date" {
		sort.SliceStable(result, func(i, j int) bool {
			return result[i].modTime.After(result[j].modTime)
		})
	}

	for _, item := range result {
		// skip if we want only directories
		if onlyDir && !item.isDir {
			continue
		}
		if size {
			fmt.Printf("%-15s %-6d %-20s\n", item.name, item.size, parseTime(item.modTime))
		} else {
			fmt.Printf("%-15s %-20s\n", item.name, parseTime(item.modTime))
		}
	}
}

func parseTime(t time.Time) string {
	return t.Format("2/1/2006 15:04:05")
}

type File struct {
	name string
	size int64
	isDir bool
	modTime time.Time
}
