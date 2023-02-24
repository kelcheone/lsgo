package main

import (
	"fmt"
)

func All(dir string, t string) {
	files, err := ReadDir(dir)
	handleErr(err)

	if t == "" {
		for _, file := range files {
			if !hasFlag(file.Name(), ".") {
				fmt.Printf("%v\t", file.Name())
			}
		}
	} else {

		for _, file := range files {
			fmt.Printf("%v\t", file.Name())
		}
	}
	println()
}

func ListLongAll(dir string) {
	// get file permisions, size, date,
	fmt.Printf("%v:\n", dir)
	files, err := ReadDir(dir)
	handleErr(err)

	for _, file := range files {
		if dir == "." {
			all, _ := createLong(file.Name())
			fmt.Println(all)
		} else {
			tPath := fmt.Sprintf("%v/%v", dir, file.Name())
			all, _ := createLong(tPath)
			fmt.Println(all)
		}
	}

}
func ListLong(dir string)         {}
func ListLongFileSize(dir string) {}
func Reverse(dir string)          {}
func ListFileSize(dir string)     {}
func SortFileSize(dir string)     {}
