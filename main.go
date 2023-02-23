// This app will be added to path and will be invocked by calling `gols`
// `gols` will take all the commands required by the ls command
// It will use filepath package to access the directories
package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	var dir string
	var t string
	if len(args) == 0 {
		dir = "."
	} else if len(args) == 1 {
		dir = args[0]
	} else if len(args) == 2 {
		dir = args[1]
		t = args[0]
	} else {
		if len(args) >= 2 {
			dir = args[1]
			t = args[0]
		}
		fmt.Printf("lsgo: Cannot access '%v': No such file in directory.\n", args[2:])
	}

	if t == "-a" {
		All(dir)
	} else if t == "-la" {
		//list long format including hidden files
	} else if t == "-l" {
		//list with long format - show permissions
	} else if t == "ls" {
		//list with long format with file size
	} else if t == "-r" {
		// list files in reverse order
	} else if t == "-s" {
		// List file size
	} else if t == "-S" {
		// Sort by file size
	}

}
