package main

import (
	"fmt"
	"sort"
)

func All(dir string, t string) {
	files, err := ReadDir(dir)
	handleErr(err)
	sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })

	if t == "" {
		for _, file := range files {
			if !hasFlag(file.Name(), ".") {
				lsPrint(file)
			}
		}
	} else {
		for _, file := range files {
			lsPrint(file)
		}
	}

	println()
}

func ListLongAll(dir string) {
	fmt.Printf("%v:\n", dir)
	files, err := ReadDir(dir)
	sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })

	handleErr(err)

	var all []Long

	all = AddToLong(dir, files, all, false)

	if len(all) != 0 {
		sort.Slice(all, func(i, j int) bool { return all[i].name < all[j].name })
		for _, file := range all {
			lsPrintLong(file.file)
		}
	}

}
func ListLong(dir string) {
	fmt.Printf("%v:\n", dir)
	files, err := ReadDir(dir)
	sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })
	handleErr(err)
	var all []Long

	all = AddToLong(dir, files, all, true)
	if len(all) != 0 {
		sort.Slice(all, func(i, j int) bool { return all[i].name < all[j].name })
		for _, file := range all {
			lsPrintLong(file.file)
		}
	}
}

func Reverse(dir string) {
	files, err := ReadDir(dir)
	handleErr(err)
	sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })

	var all []Long
	all = AddToLong(dir, files, all, false)
	if len(all) != 0 {
		sort.Slice(all, func(i, j int) bool { return all[i].name > all[j].name })
		for _, file := range all {
			if !hasFlag(file.name, ".") {
				lsPrint(file.file)
			}
		}

	}
	fmt.Println()
}
func ListFileSize(dir string) {
	files, err := ReadDir(dir)
	handleErr(err)
	sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })

	var all []Long
	all = AddToLong(dir, files, all, false)
	if len(all) != 0 {
		for _, file := range all {
			if !hasFlag(file.name, ".") {
				fmt.Printf("%v %v\t", file.size, file.name)
			}
		}

	}

}
func SortFileSize(dir string) {
	files, err := ReadDir(dir)
	handleErr(err)
	sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })

	var all []Long
	all = AddToLong(dir, files, all, false)
	if len(all) != 0 {
		sort.Slice(all, func(i, j int) bool { return all[i].sizeInt > all[j].sizeInt })
		for _, file := range all {
			if !hasFlag(file.name, ".") {
				lsPrint(file.file)
			}
		}

	}
	fmt.Println()
}
