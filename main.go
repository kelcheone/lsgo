// This app will be added to path and will be invocked by calling `gols`
// `gols` will take all the commands required by the ls command
// It will use filepath package to access the directories
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	var dir string
	var t string

	if len(args) > 0 {
		if listFlag(args) {
			t = args[0]
		}
	}

	if len(args) == 0 {
		dir = "."
	} else if len(args) == 1 {
		if listFlag(args) {
			dir = "."
		} else {
			dir = args[0]
		}

	} else if len(args) == 2 {
		if listFlag(args) {
			dir = args[1]
			t = args[0]
		} else {
			dir = args[0]
			fmt.Printf("lsgo: Cannot access '%v': No such file in directory.\n", strings.Join(args[1:], ", "))
		}
	} else if len(args) >= 3 {
		if listFlag(args) {
			t = args[0]
			dir = args[1]
			fmt.Printf("lsgo: Cannot access '%v': No such file in directory.\n", strings.Join(args[2:], " "))
		} else {
			dir = args[0]
		}
	}

	if t == "" {
		All(dir, t)
	} else if t == "-a" {
		All(dir, t)

	} else if t == "-la" {
		//list long format including hidden files
		ListLongAll(dir)
	} else if t == "-l" {
		//list with long format - show permissions
		ListLong(dir)
	} else if t == "-ls" {
		//list with long format with file size
		ListLongFileSize(dir)
	} else if t == "-r" {
		// list files in reverse order
		Reverse(dir)
	} else if t == "-s" {
		// List file size
		ListFileSize(dir)
	} else if t == "-S" {
		// Sort by file size
		SortFileSize(dir)
	}

}
