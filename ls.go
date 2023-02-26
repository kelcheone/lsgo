package main

import (
	"fmt"
	"sort"
)

func All(dir string, t string) {
	files, err := ReadDir(dir)
	sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })
	handleErr(err)

	var all []Long

	if t == "" {
		for _, file := range files {
			if !hasFlag(file.Name(), ".") {
				fmt.Printf("%v\t", file.Name())
			}
		}
	} else if t == "-r" {
		for i := len(files) - 1; i >= 0; i-- {
			if !hasFlag(files[i].Name(), ".") {
				fmt.Printf("%v\t", files[i].Name())
			}
		}
	} else if t == "-s" {
		for _, file := range files {
			if dir == "." {
				long, _ := createLong(file.Name())
				all = append(all, long)
			} else {
				tPath := fmt.Sprintf("%v/%v", dir, file.Name())
				long, _ := createLong(tPath)
				all = append(all, long)
			}
		}
		if len(all) != 0 {
			// fmt.Print(all)
			for _, file := range all {
				if !hasFlag(file.name, ".") {
					fmt.Printf("%v %v\t", file.size, file.name)
				}
			}
		}
	} else if t == "-S" {
		for _, file := range files {
			if dir == "." {
				long, _ := createLong(file.Name())
				all = append(all, long)
			} else {
				tPath := fmt.Sprintf("%v/%v", dir, file.Name())
				long, _ := createLong(tPath)
				all = append(all, long)
			}
		}
		if len(all) != 0 {
			sort.Slice(all, func(i, j int) bool { return all[i].sizeint > all[j].sizeint })
			for _, file := range all {
				if !hasFlag(file.name, ".") {
					fmt.Printf("%v \t", file.name)
				}
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

	for _, file := range files {
		if dir == "." {
			long, _ := createLong(file.Name())
			all = append(all, long)
		} else {
			tPath := fmt.Sprintf("%v/%v", dir, file.Name())
			long, _ := createLong(tPath)
			all = append(all, long)
		}
	}

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

	for _, file := range files {
		if dir == "." {
			long, _ := createLong(file.Name())

			if !hasFlag(long.name, ".") {
				all = append(all, long)
			}

		} else {
			tPath := fmt.Sprintf("%v/%v", dir, file.Name())
			long, _ := createLong(tPath)
			if !hasFlag(long.name, ".") {
				all = append(all, long)
			}
		}
	}
	if len(all) != 0 {
		sort.Slice(all, func(i, j int) bool { return all[i].name < all[j].name })
		for _, file := range all {
			fmt.Printf("%v \t %v \t%v\t %v \n", file.perm, file.size, file.time, file.name)

		}
	}
}

// func ListLongFileSize(dir string) {}
// func Reverse(dir string)      { All(dir, "-r") }
func ListFileSize(dir string) {

}
func SortFileSize(dir string) {}
