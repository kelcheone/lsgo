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
	sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })

	handleErr(err)

	var all []Long

	all = AddToLong(dir, files, all, false)

	if len(all) != 0 {
		sort.Slice(all, func(i, j int) bool { return all[i].name < all[j].name })
		for _, file := range all {
			fmt.Printf("%v \t %v \t%v\t %v \n", file.perm, file.size, file.time, file.name)
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
			fmt.Printf("%v \t %v \t%v\t %v \n", file.perm, file.size, file.time, file.name)

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
				fmt.Printf("%v \t", file.name)
			}
		}

	}
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
		sort.Slice(all, func(i, j int) bool { return all[i].sizeint > all[j].sizeint })
		for _, file := range all {
			if !hasFlag(file.name, ".") {
				fmt.Printf("%v \t", file.name)
			}
		}

	}
}
